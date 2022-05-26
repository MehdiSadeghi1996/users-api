package users_db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const (
	mysql_users_username = "root"
	mysql_users_password = "16101374"
	mysql_users_host     = "192.168.92.129:3306"
	mysql_users_schema   = "users_db"
	//they should get from os env
)

var (
	Client *sql.DB
)

func init() {

	datasourcename := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		mysql_users_username,
		mysql_users_password,
		mysql_users_host,
		mysql_users_schema)

	var err error

	Client, err = sql.Open("mysql", datasourcename)

	if err != nil {
		panic(err)
	}
	//mysql.SetLogger()
	log.Println("database successfully configured")

}
