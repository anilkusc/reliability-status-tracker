package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func Select() []Source {
	database, err := sql.Open("sqlite3", "./db/database.db")
	if err != nil {
		fmt.Println("Open database error")
	}
	defer database.Close()
	query := "SELECT * FROM status;"
	rows, err := database.Query(query)
	if err != nil {
		fmt.Println("Query error")
	}

	defer rows.Close()
	var source Source
	var sources []Source
	for rows.Next() {
		rows.Scan(&source.Host, &source.Desired, &source.Interval, &source.Method, &source.Proxy, &source.LastCode)
		sources = append(sources, source)

	}
	return sources
}

func Insert(source Source) string {
	database, err := sql.Open("sqlite3", "./db/database.db")
	if err != nil {
		fmt.Println("Open database error")
		return `{"status":"FAIL"}`
	}
	defer database.Close()
	statement, err := database.Prepare("INSERT INTO status (host,desired,interval,method,proxy,lastCode) VALUES (?,?,?,?,?,?)")
	if err != nil {
		fmt.Println("prepare statement error")
		return `{"status":"FAIL"}`
	}
	statement.Exec(source.Host, source.Desired, source.Interval, source.Method, source.Proxy, source.LastCode)

	return `{"status":"OK"}`
}

func Delete(source Source) string {
	database, err := sql.Open("sqlite3", "./db/database.db")
	if err != nil {
		fmt.Println("Open database error")
		return `{"status":"FAIL"}`
	}
	defer database.Close()
	statement, err := database.Prepare("DELETE FROM status where host=? AND desired=? AND interval=? AND method=?  AND proxy=? ;")
	if err != nil {
		fmt.Println("prepare statement error")
		return `{"status":"FAIL"}`
	}
	statement.Exec(source.Host, source.Desired, source.Interval, source.Method, source.Proxy)

	return `{"status":"OK"}`
}
func Update(source Source) string {
	database, err := sql.Open("sqlite3", "./db/database.db")
	if err != nil {
		fmt.Println("Open database error")
		return `{"status":"FAIL"}`
	}
	defer database.Close()
	statement, err := database.Prepare("UPDATE status SET lastCode = ? WHERE host=? AND desired=? AND interval=? AND method=?  AND proxy=? ;")
	if err != nil {
		fmt.Println("prepare statement error")
		return `{"status":"FAIL"}`
	}
	time.Sleep(time.Second)
	statement.Exec(source.LastCode, source.Host, source.Desired, source.Interval, source.Method, source.Proxy)

	return `{"status":"OK"}`
}
