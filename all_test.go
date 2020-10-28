package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
	_ "github.com/proullon/ramsql/driver"
)

var loginValues = []struct {
	input          User
	expectedResult string
}{
	{User{Username: "Anil", Password: "admin"}, `{"authenticated":"true"}`},
	{User{Username: "asd", Password: "asd"}, `{"authenticated":"false"}`},
	{User{Username: "", Password: "admin"}, `{"authenticated":"false"}`},
	{User{Username: "Anil", Password: ""}, `{"authenticated":"false"}`},
	{User{Username: "", Password: ""}, `{"authenticated":"false"}`},
}
var addValues = []struct {
	input          Source
	expectedResult string
}{
	{Source{Host: "http://www.google.com", Desired: 200, Interval: 30, Method: "GET", Proxy: "", LastCode: 0}, `{"status":"OK"}`},
}
var statusValues = []struct {
	expectedResult string
}{
	{`[{"host":"https://www.google.com","desired":200,"interval":30,"method":"GET","proxy":"","lastCode":200}]`},
}

func TestAll(t *testing.T) {
	//	t.Parallel() // multithreading tests
	for _, login := range loginValues {
		t.Run("Login", testLogin(login.input, login.expectedResult))
	}
	for _, add := range addValues {
		t.Run("Add", testAdd(add.input, add.expectedResult))
		t.Run("Delete", testDelete(add.input, add.expectedResult))
	}
	for _, status := range statusValues {
		t.Run("Login", testStatus(status.expectedResult))
	}
}

func testLogin(user User, expected string) func(*testing.T) {
	return func(t *testing.T) {
		data, _ := json.Marshal(user)
		req, err := http.NewRequest("POST", "/login/", bytes.NewBuffer(data))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(Login)

		handler.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}
		//expected := `{"authenticated":"false"}`
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
		}
	}
}

func testAdd(source Source, expected string) func(*testing.T) {
	return func(t *testing.T) {

		batch := []string{
			`create table status (host TEXT,desired INTEGER,interval INTEGER,method TEXT,proxy TEXT,lastCode INTEGER);`,
			`insert into status (host,desired,interval,method,proxy,lastCode) values ('https://www.google.com','200','30','GET','','200');`,
		}

		db, err := sql.Open("ramsql", "testAdd")
		if err != nil {
			t.Fatalf("sql.Open : Error : %s\n", err)
		}
		defer db.Close()

		for _, b := range batch {
			_, err = db.Exec(b)
			if err != nil {
				t.Fatalf("sql.Exec: Error: %s\n", err)
			}
		}

		dtbs = db
		data, _ := json.Marshal(source)
		req, err := http.NewRequest("POST", "/add/", bytes.NewBuffer(data))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(Add)

		handler.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
		}
	}
}
func testDelete(source Source, expected string) func(*testing.T) {
	return func(t *testing.T) {

		batch := []string{
			`create table status (host TEXT,desired INTEGER,interval INTEGER,method TEXT,proxy TEXT,lastCode INTEGER);`,
			`insert into status (host,desired,interval,method,proxy,lastCode) values ('https://www.google.com','200','30','GET','','200');`,
		}

		db, err := sql.Open("ramsql", "testDelete")
		if err != nil {
			t.Fatalf("sql.Open : Error : %s\n", err)
		}
		defer db.Close()

		for _, b := range batch {
			_, err = db.Exec(b)
			if err != nil {
				t.Fatalf("sql.Exec: Error: %s\n", err)
			}
		}

		dtbs = db
		data, _ := json.Marshal(source)
		req, err := http.NewRequest("POST", "/delete/", bytes.NewBuffer(data))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(Add)

		handler.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
		}
	}
}
func testStatus(expected string) func(*testing.T) {
	return func(t *testing.T) {
		batch := []string{
			`create table status (host TEXT,desired INTEGER,interval INTEGER,method TEXT,proxy TEXT,lastCode INTEGER);`,
			`insert into status (host,desired,interval,method,proxy,lastCode) values ('https://www.google.com','200','30','GET','','200');`,
		}

		db, err := sql.Open("ramsql", "testStatus")
		if err != nil {
			t.Fatalf("sql.Open : Error : %s\n", err)
		}
		defer db.Close()
		for _, b := range batch {
			_, err = db.Exec(b)
			if err != nil {
				t.Fatalf("sql.Exec: Error: %s\n", err)
			}
		}

		dtbs = db

		s := httptest.NewServer(http.HandlerFunc(WsStatus))
		defer s.Close()

		// Convert http://127.0.0.1 to ws://127.0.0.
		u := "ws" + strings.TrimPrefix(s.URL, "http")
		// Connect to the server
		ws, _, err := websocket.DefaultDialer.Dial(u, nil)
		if err != nil {
			t.Fatalf("%v", err)
		}
		defer ws.Close()

		if err := ws.WriteMessage(websocket.TextMessage, []byte("hello")); err != nil {
			t.Fatalf("%v", err)
		}
		_, p, err := ws.ReadMessage()
		if err != nil {
			t.Fatalf("%v", err)
		}
		if string(p) != expected {
			t.Errorf("handler returned unexpected body: got %v want %v", string(p), expected)
		}
	}
}
