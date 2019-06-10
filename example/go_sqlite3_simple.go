package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

var Path string = "F:\\db\\foo.db"

func main() {
	os.Remove(Path)

	db, err := sql.Open("sqlite3", Path)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	create table foo (id integer not null primary key, name text);
	`
	// delete from foo;
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}


	sqlStr := `
	create table data1 (id integer not null primary key, name text,attributes text);
	`
	// delete from data1;
	_, err = db.Exec(sqlStr)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStr)
		return
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into foo(id, name) values(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	for i := 0; i < 100; i++ {
		_, err = stmt.Exec(i, fmt.Sprintf("hp%03d", i))
		if err != nil {
			log.Fatal(err)
		}
	}
	tx.Commit()

	rows, err := db.Query("select id, name from foo")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("================================")
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err = db.Prepare("select name from foo where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var name string
	err = stmt.QueryRow("3").Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("+++++++++++++++++++++++++++++++++++")
	fmt.Println(name)

	_, err = db.Exec("delete from foo where id = 1")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("delete from foo where id = 2")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("delete from foo where id = 3")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("insert into foo(id, name) values(1, 'foo'), (2, 'bar'), (3, 'baz')")
	if err != nil {
		log.Fatal(err)
	}

	rows, err = db.Query("select id, name from foo")
	if err != nil {
		log.Fatal(err)
	}
	//defer rows.Close()

	fmt.Println("-------------------------------------")

	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	rows.Close()

	rows,err = db.Query("SELECT name FROM foo")
	if err != nil {
		log.Fatal(err)
	}

	//defer rows.Close()

	fmt.Println("-------------=========----------------")

	for rows.Next() {
		//var id int
		var name string
		err = rows.Scan(&name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println( name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	rows.Close()

	err = stmt.QueryRow("30").Scan(&name)
	if err != nil {
		log.Fatal(err)
		fmt.Printf("+++++++ err: %v\n",err)
	} 
	fmt.Println("+++++***********************+++++++++++")
	fmt.Println(name)
}
