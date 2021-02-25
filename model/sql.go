package sql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
var DB *sql.DB
func  Db() {
	db,_:=sql.Open("mysql","root:@tcp(127.0.0.1:3306)/test")
	err:=db.Ping()
	if err !=nil{
		fmt.Print("error")
		return
	}
	DB =db
}
