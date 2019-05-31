package dbops

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err    error
)

//go中设计的sql适用于长连接的，一直创建链接可能导致关闭连接的时候出现很多close() 等待
func init() {
	dbConn, err = sql.Open("mysql", "root:123456@tcp(111.230.92.97:3306)/video_server?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
}
