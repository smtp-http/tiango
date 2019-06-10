/*
	Author        : tuxpy
	Email         : q8886888@qq.com.com
	Create time   : 2017-11-04 23:13:08
	Filename      : main.go
	Description   :
*/

package main

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"time"
	//"utils"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

func ToStruct(rows *sql.Rows, to interface{}) error {
	v := reflect.ValueOf(to)
	if v.Elem().Type().Kind() != reflect.Struct {
		return errors.New("Expect a struct")
	}

	scan_dest := []interface{}{}
	column_names, _ := rows.Columns()

	addr_by_column_name := map[string]interface{}{}

	for i := 0; i < v.Elem().NumField(); i++ {
		one_value := v.Elem().Field(i)
		column_name := v.Elem().Type().Field(i).Tag.Get("sql")
		if column_name == "" {
			column_name = one_value.Type().Name()
		}
		addr_by_column_name[column_name] = one_value.Addr().Interface()
	}

	for _, column_name := range column_names {
		scan_dest = append(scan_dest, addr_by_column_name[column_name])
	}

	return rows.Scan(scan_dest...)

}

func main() {
	db, err := sql.Open("sqlite3", "f:\\db\\firstlink.db")
	if err != nil {
		log.Fatal(err)
	}
	//utils.CheckErrorPanic(err)
	fmt.Println(db)
	db.Exec(`CREATE TABLE userinfo (
		work_id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(64) NULL,
		departname VARCHAR(64) NULL,
		created DATE NULL
	)`)

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		tx.Commit()
	}()
	//utils.CheckErrorPanic(err)

	//	stmt, err := tx.Prepare("INSERT INTO userinfo(name, departname, created) values(?, ?, ?);")
	//	utils.CheckErrorPanic(err)
	//
	//	_, err = stmt.Exec("我", "爬虫", "2015-12-15")
	//	_, err = stmt.Exec("你", "客服", "2015-12-16")
	//	utils.CheckErrorPanic(err)

	rows, err := tx.Query("SELECT work_id, name, created FROM userinfo")
	if err != nil {
		log.Fatal(err)
	}
	//utils.CheckErrorPanic(err)

	type Record struct {
		Name    string    `sql:"name"`
		WorkID  int       `sql:"work_id"`
		Created time.Time `sql:"created"`
	}

	for rows.Next() {
		record := Record{}
		err := ToStruct(rows, &record)
		if err != nil {
			log.Fatal(err)
		}
		//utils.CheckErrorPanic()
		fmt.Println(record.WorkID, record.Name, record.Created)
	}
}
