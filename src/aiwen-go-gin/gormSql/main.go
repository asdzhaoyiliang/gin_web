package main

import (
	//go 提供sql 和mysql进行交互的api
	"database/sql"
	"fmt"
	"log"
	//导入go的驱动包
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

//func Init() {
//	ConnectDb()
//
//}
func main() {
	ConnectDb()
	//insert()
	//selectDb()
	updateDb()
	deleteDb()
	defer Db.Close()
}
func insert() {
	res, err := Db.Exec("insert into person(`name`,`age`) value(?,?);", "z1", 1)

	fmt.Println(res)
	fmt.Println(err)
	fmt.Println("insert success")
}
func selectDb() {
	res, err := Db.Query("select * from person")
	fmt.Println(res)
	fmt.Println(err)
	for res.Next() {
		var id, age int
		var name string
		res.Scan(&id, &name, &age)
		a, _ := res.Columns()
		fmt.Println(a)
		fmt.Printf("id:%d name:%s age:%d\n", id, name, age)
	}
}
func updateDb() {
	res, err := Db.Exec("update person set name = ? where id = ?", "z2", 3)
	if err != nil {
		log.Fatal(err)
	} else {
		lastInsertId, _ := res.LastInsertId()
		rowsEffected, _ := res.RowsAffected()
		fmt.Println(lastInsertId, rowsEffected)
	}
}
func deleteDb() {
	_, err := Db.Exec("delete from person where id = ?", "2")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(err)
	}

}
func ConnectDb() {

	db, _ := sql.Open("mysql", "wanart:wanart@tcp(127.0.0.1:3306)/ginsql")
	//defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("success")
	Db = db
}
