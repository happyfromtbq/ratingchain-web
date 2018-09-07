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

// insertGetId : insert data and get id
func InsertGetId(dba *gorose.Database) (int64, error) {
	_, err := dba.Insert()
	if err != nil {
		return 0, err
	}
	return dba.LastInsertId, nil
}
