package mysql

import (
	"fmt"
	"net/url"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() error {
	host := viper.GetString("mysql.host")
	port := viper.GetInt("mysql.port")
	database := viper.GetString("mysql.database_name")
	username := viper.GetString("mysql.user")
	password := viper.GetString("mysql.password")
	charset := viper.GetString("mysql.charset")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=%s",
		username,
		password,
		host,
		port,
		database,
		charset,
		url.QueryEscape("Local"),
	)

	// connect to Mysql
	dbConn, err := gorm.Open(mysql.New(mysql.Config{
		DSN:        dsn,
		DriverName: viper.GetString("datasource.driverName"),
	}))
	if err != nil {
		zap.L().Fatal(fmt.Sprintf("cannot connect to mysql: %s", err))
		return err
	}

	// AutoMigrate
	//err = db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.User{})
	//if err != nil {
	//	zap.L().Fatal(fmt.Sprintf("cannot create table: %s", err))
	//	return err
	//}

	db = dbConn
	return nil
}
