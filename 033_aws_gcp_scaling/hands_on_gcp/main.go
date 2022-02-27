package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"net/http"
	"cloud.google.com/go/compute/metadata"

)

var db *sql.DB

func main() {
	var err error
	//db, err = sql.Open("mysql", "root:mypassword@tcp(127.0.0.1:3306)/test")

	// use GCP CloudSQL private IP (public IP was not removed)
	db, err = sql.Open("mysql", "root:mypassword@tcp(10.31.32.3:3306)/test") 
	check(err)
	// defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/instance", instance)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/amigos", amigos)
	http.ListenAndServe(":80", nil)
	defer db.Close()
}

func index(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello from GCP.")
}

func ping(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "OK")
}

func instance(w http.ResponseWriter, req *http.Request) {
	s := getInstance()
	io.WriteString(w, s)
}

func amigos(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Query(`SELECT aName FROM amigos;`)
	check(err)

	// data to be used in query
	s := getInstance()
	s += "\nRETRIEVED RECORDS:\n"
	var name string

	// query
	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		s += name + "\n"
	}
	fmt.Fprintln(w, s)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func getInstance() string {
	
	// GCP version
	
	c, err := metadata.InstanceName()
	if err != nil {
		fmt.Println(err)
	}

	return c
}
