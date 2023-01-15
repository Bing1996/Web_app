package main

import (
	"Web_App/repository/mysql"
	"fmt"
)

func main() {
	err := mysql.Init()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if (mysql.DBMulti["cloud"]).Migrator().HasTable("user") {
		fmt.Printf("table %s not found", "user")
	} else {
		fmt.Printf("table %s found", "user")
	}

}
