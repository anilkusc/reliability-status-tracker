package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func NewDbConn() *sql.DB {
	database, err := sql.Open("sqlite3", "./db/database.db")
	if err != nil {
		fmt.Println("Open database error")
	}
	//defer database.Close()
	return database
}

func Select(database *sql.DB) []Source {
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

func Insert(database *sql.DB, source Source) string {
	statement, err := database.Prepare("INSERT INTO status (host,desired,interval,method,proxy,lastCode) VALUES (?,?,?,?,?,?)")
	if err != nil {
		fmt.Println(err)
		fmt.Println("prepare statement error")
		return `{"status":"FAIL"}`
	}
	statement.Exec(source.Host, source.Desired, source.Interval, source.Method, source.Proxy, source.LastCode)
	restart = true
	time.Sleep(time.Second)
	restart = false
	return `{"status":"OK"}`
}

func Delete(database *sql.DB, source Source) string {

	statement, err := database.Prepare("DELETE FROM status where host=? AND desired=? AND interval=? AND method=?  AND proxy=? ;")
	if err != nil {
		fmt.Println("prepare statement error")
		return `{"status":"FAIL"}`
	}
	statement.Exec(source.Host, source.Desired, source.Interval, source.Method, source.Proxy)
	restart = true
	time.Sleep(time.Second)
	restart = false
	return `{"status":"OK"}`
}
func Update(database *sql.DB, source Source) string {
	if os.Getenv("DEBUG") == "true" {
		fmt.Println(source)
	}
	statement, err := database.Prepare("UPDATE status SET lastCode = ? WHERE host=? AND desired=? AND interval=? AND method=?  AND proxy=? ;")
	if err != nil {
		fmt.Println("prepare statement error")
		return `{"status":"FAIL"}`
	}
	time.Sleep(time.Second)
	statement.Exec(source.LastCode, source.Host, source.Desired, source.Interval, source.Method, source.Proxy)

	return `{"status":"OK"}`
}
