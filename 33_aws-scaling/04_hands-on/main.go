package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {

	var err error
	db, err = sql.Open("mysql", "dbuser:password@tcp()/tes02?charset=utf8")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/instance", instance)
	http.HandleFunc("/amigos", amigos)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":80", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello from AWS")
}

func ping(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "ok")
}

func instance(w http.ResponseWriter, req *http.Request) {
	s := getInstance()
	io.WriteString(w, s)
}

func amigos(w http.ResponseWriter, req *http.Request) {

	rows, err := db.Query(`SELECT amigo_name FROM amigo`)
	check(err)

	// data to be used in query
	s := getInstance()
	var name string
	s += "\nRETRIEVED RECORDS:\n"

	// query
	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		s += name + "\n"
	}
	fmt.Fprintln(w, s)
}

func getInstance() string {

	resp, err := http.Get("http://169.254.169.254/latest/meta-data/instance-id")
	check(err)

	bs := make([]byte, resp.ContentLength)
	resp.Body.Read(bs)
	resp.Body.Close()

	return string(bs)
}
func check(err error) {

	if err != nil {
		fmt.Println(err)
	}
}

/*
Run Results:

*/
