package repositories

import (
	"github.com/happyfromtbq/ratingchain-web/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gohouse/gorose"
)

var conn gorose.Connection


func InitDb() {
//*gorose.Database{
	var dbconfig = config.DbConfig
	connection, err := gorose.Open(dbconfig, "mysql_dev")
	if err != nil {
		fmt.Println(err)
		return
	}
	conn = connection

	// close DB
	//defer connection.Close()
	//
	//db := connection.NewDB()
	//return db
}
