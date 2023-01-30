package mysql

import (
	"Web_App/asset/settings"
	"Web_App/model"
	"fmt"
	"net/url"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db      *gorm.DB
	DBMulti = make(map[string]*gorm.DB)
)

func Init() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=%s",
		settings.Conf.MysqlConfig.User,
		settings.Conf.MysqlConfig.Password,
		settings.Conf.MysqlConfig.Host,
		settings.Conf.MysqlConfig.Port,
		settings.Conf.MysqlConfig.DataBaseName,
		settings.Conf.MysqlConfig.Charset,
		url.QueryEscape("Local"),
	)

	// connect to Mysql
	dbConn, err := gorm.Open(mysql.New(mysql.Config{
		DSN:        dsn,
		DriverName: settings.Conf.MysqlConfig.DriverName,
	}))
	if err != nil {
		zap.L().Fatal(fmt.Sprintf("cannot connect to mysql: %s", err))
		return err
	}

	sqlDB, _ := dbConn.DB()
	sqlDB.SetMaxOpenConns(settings.Conf.MysqlConfig.MaxOpenConns)
	sqlDB.SetMaxIdleConns(settings.Conf.MysqlConfig.MaxIdleConns)

	db = dbConn

	DBMulti["cloud"] = dbConn
	if err = registerAllModelTable(); err != nil {
		return err
	}

	return nil
}

func Close() {
	s, err := db.DB()
	if err != nil {
		return
	}
	_ = s.Close()
}

// 注册所有基于Model的关键表
func registerAllModelTable() error {
	// 创建用户表
	if !(DBMulti["cloud"]).Migrator().HasTable("users") {
		fmt.Printf("table %s not found, create table...\n", "user")
		err := (DBMulti["cloud"]).Migrator().CreateTable(&model.User{})
		if err != nil {
			// 创建table失败
			return err
		}
	} else {
		fmt.Printf("table %s found\n", "user")
	}

	return nil
}
