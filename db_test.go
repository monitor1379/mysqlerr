package mysqlerr_test

/*
 * @Date: 2020-10-31 14:46:13
 * @LastEditors: aiden.deng (Zhenpeng Deng)
 * @LastEditTime: 2020-10-31 15:08:11
 */

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func TestMySQLErr(t *testing.T) {

	cfg := mysql.Config{
		User:   "monitor1379",
		Passwd: "deng123",
		Net:    "tcp",
		Addr:   "127.0.0.1:5500",
		DBName: "dev",
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	checkErr(err)

	fmt.Println(db)
	rows, err := db.Query("show databases")
	checkErr(err)

	for rows.Next() {
		var s string
		err := rows.Scan(&s)
		checkErr(err)
		fmt.Println(s)
	}
}
