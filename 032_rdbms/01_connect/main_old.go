package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"net/http"
)

var db *sql.DB
var err error

func main() {
	// user:password@tcp(localhost:5555)/dbname?charset=utf8
/*
AWS
	db, err = sql.Open("mysql", "awsuser:mypassword@tcp(mydbinstance.cakwl95bxza0.us-west-1.rds.amazonaws.com:3306)/test02?charset=utf8")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)
*/


	var dbUser = "root"          // e.g. 'my-db-user'
	var dbPwd = "Purpletree462" // e.g. 'my-db-password'
	var instanceConnectionName = "heroic-oarlock-340615:europe-north1:gotraining" // e.g. 'project:region:instance'
	var dbName = "test" // e.g. 'my-database'
	var socketDir = "/cloudsql"
	var dbURI = fmt.Sprintf("%s:%s@unix(/%s/%s)/%s?parseTime=true", dbUser, dbPwd, socketDir, instanceConnectionName, dbName)

	dbPool, err := sql.Open("mysql", dbURI)

	if err != nil {
//			return fmt.Errorf("sql.Open: %v", err)
	}

	
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	//err := http.ListenAndServe(":8080", nil)
	http.ListenAndServe(":8080", nil)
	// check(err)
	
}

func index(w http.ResponseWriter, req *http.Request) {
	_, err = io.WriteString(w, "Successfully completed.")
	check(err)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
