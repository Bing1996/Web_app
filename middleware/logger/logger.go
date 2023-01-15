package logger

import (
	"Web_App/asset/settings"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

type LogConfig struct {
	Level      string `json:"level,omitempty"`
	Filename   string `json:"filename,omitempty"`
	MaxSize    int    `json:"max_size,omitempty"`
	MaxAge     int    `json:"max_age,omitempty"`
	MaxBackups int    `json:"max_backups,omitempty"`
}

// Init 创建日志实例
func Init() (err error) {
	writeSyncer := getLogWriter(
		settings.Conf.LogConfig.Filename,
		settings.Conf.LogConfig.MaxSize,
		settings.Conf.LogConfig.MaxBackups,
		settings.Conf.LogConfig.MaxAge,
	)

	encoder := getEncoder()

	var l = new(zapcore.Level)

	if err = l.UnmarshalText([]byte(settings.Conf.LogConfig.Level)); err != nil {
		return
	}

	var core zapcore.Core
	if settings.Conf.GinConfig.Mode == "dev" {
		// 开发模式，日志输出至终端
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee(
			// 文件输出编码
			zapcore.NewCore(encoder, writeSyncer, l),
			// 终端输出编码
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		)
	} else {
		core = zapcore.NewCore(encoder, writeSyncer, l)
	}

	logger = zap.New(core, zap.AddCaller())

	// 替换zap包中全局logger实例，后续在其它保重只需要只用zap.L()调用即可
	zap.ReplaceGlobals(logger)

	return nil
}

// GinLogger 接受gin框架默认的日志
func GinLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		path := ctx.Request.URL.Path
		query := ctx.Request.URL.RawQuery
		ctx.Next()

		cost := time.Since(start)
		logger.Info(path,
			zap.Int("status", ctx.Writer.Status()),
			zap.String("method", ctx.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", ctx.ClientIP()),
			zap.String("user-agent", ctx.Request.UserAgent()),
			zap.String("errors", ctx.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("code", cost),
		)
	}
}

// GinRecovery recover项目可能出现的panic, 并使用zap相关日志
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") ||
							strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(context.Request, false)
				if brokenPipe {
					logger.Error(context.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					_ = context.Error(err.(error)) // nolint: err check
					context.Abort()
					return
				}
				if stack {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				context.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		context.Next()
	}
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}

	return zapcore.AddSync(lumberJackLogger)
}
