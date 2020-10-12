package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func Select() []Source {
	database, err := sql.Open("sqlite3", "./db/database.db")
	if err != nil {
		fmt.Println(err)
	}
	//defer database.Close()
	query := "SELECT * FROM status;"
	rows, err := database.Query(query)
	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()
	var source Source
	var sources []Source
	for rows.Next() {
		rows.Scan(&source.Host, &source.Interval, &source.Method, &source.Proxy, &source.LastCode)
		sources = append(sources, source)
	}
	return sources
}

func Insert(source Source) string {
	database, err := sql.Open("sqlite3", "./db/database.db")
	if err != nil {
		fmt.Println(err)
		return "failed to open sql"
	}
	defer database.Close()
	statement, _ := database.Prepare("INSERT INTO status (host, interval,method,proxy) VALUES (?,?,?,?)")
	statement.Exec(source.Host, source.Interval, source.Method, source.Proxy)
	return "OK"
}

func Delete(source Source) string {
	database, err := sql.Open("sqlite3", "./db/database.db")
	if err != nil {
		fmt.Println(err)
		return "failed to open sql"
	}
	defer database.Close()
	statement, _ := database.Prepare("DELETE FROM status where host=? AND interval=? AND method=?  AND proxy=? )")
	statement.Exec(source.Host, source.Interval, source.Method, source.Proxy)
	return "OK"
}
func Update(source Source) string {
	database, err := sql.Open("sqlite3", "./db/database.db")
	if err != nil {
		fmt.Println(err)
		return "failed to open sql"
	}
	defer database.Close()
	statement, _ := database.Prepare("UPDATE status SET lastCode = ? WHERE host=? AND interval=? AND method=?  AND proxy=? ;")
	statement.Exec(source.LastCode, source.Host, source.Interval, source.Method, source.Proxy)
	return "OK"
}
