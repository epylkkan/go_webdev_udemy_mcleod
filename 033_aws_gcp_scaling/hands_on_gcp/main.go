package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"net/http"
	/*
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
	*/
)

var db *sql.DB

func main() {
	var err error
	//db, err = sql.Open("mysql", "root:mypassword@tcp(127.0.0.1:3306)/test")
	db, err = sql.Open("mysql", "root:mypassword@tcp(35.228.143.206:3306)/test")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/instance", instance)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/amigos", amigos)
	http.ListenAndServe(":80", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello from AWS.")
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
//	ctx := context.Background()

	/*
	c, err := google.DefaultClient(ctx, compute.CloudPlatformScope)
	if err != nil {
			log.Fatal(err)
	}

	computeService, err := compute.New(c)
	if err != nil {
			log.Fatal(err)
	}

	// Project ID for this request.
	project := "my-project" // TODO: Update placeholder value.

	// The name of the zone for this request.
	zone := "my-zone" // TODO: Update placeholder value.

	// Name of the instance resource to return.
	instance := "my-instance" // TODO: Update placeholder value.
*/
/*
	resp, err := computeService.Instances.Get(project, zone, instance).Context(ctx).Do()
	if err != nil {
			log.Fatal(err)
	}

	// TODO: Change code below to process the `resp` object:
	fmt.Printf("%#v\n", resp)

	*/
	/*
	resp, err := http.Get("http://169.254.169.254/latest/meta-data/instance-id")
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}

	bs := make([]byte, resp.ContentLength)
	resp.Body.Read(bs)
	resp.Body.Close()
	return string(bs)
	*/
	return ""
}
