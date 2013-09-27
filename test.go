package main

import (
	"code.google.com/p/go-sqlite/go1/sqlite3"
	"fmt"
)

func main() {
	c, _ := sqlite3.Open("sqlite.db")
	c.Exec("CREATE TABLE x(a, b, c)")

	args := sqlite3.NamedArgs{"$a": 1, "$b": "demo"}
	c.Exec("INSERT INTO x VALUES($a, $b, $c)", args) // $c will be NULL

	sql := "SELECT rowid, * FROM x"
	row := make(sqlite3.RowMap)
	for s, err := c.Query(sql); err == nil; err = s.Next() {
		var rowid int64
		s.Scan(&rowid, row)     // Assigns 1st column to rowid, the rest to row
		fmt.Println(rowid, row) // Prints "1 map[a:1 b:demo c:<nil>]"
	}
	fmt.Printf("hello\n")
}
