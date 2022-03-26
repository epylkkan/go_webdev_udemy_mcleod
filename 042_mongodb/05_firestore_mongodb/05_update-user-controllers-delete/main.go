// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// [START gae_go111_app]

// Firestore Example instead of MongoDB

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"cloud.google.com/go/firestore"	
	"context"
	"github.com/julienschmidt/httprouter"
	"encoding/json"	
)

var ctx = context.Background()
var client, err = firestore.NewClient(ctx, "trim-diode-344014")

type User struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
}


func main() {

	mux := httprouter.New()
	mux.GET("/", index) 
	mux.GET("/user/:id", GetUser) 
	mux.POST("/user", CreateUser) 
	mux.DELETE("/user/:id", DeleteUser) 
	http.ListenAndServe(":8080", mux)
		
	// [START setting_port]
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}


func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Welcome to Firestore Training App !")
}



func GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	u, err := client.Collection("users").Doc(p.ByName("id")).Get(ctx)

	if err != nil {
		fmt.Fprint(w, "Error GetUser")
		return
	}
	fmt.Fprintln(w, u.Data()["name"])
	fmt.Fprintln(w, u.Data()["gender"])
	fmt.Fprintln(w, u.Data()["age"])
}


func CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	u := User{}
	json.NewDecoder(r.Body).Decode(&u)

	_, _, err := client.Collection("users").Add(ctx,u)

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Fprint(w, "Error CreateUser")
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)

}


func DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

 	id := p.ByName("id")	
	_, err := client.Collection("users").Doc(id).Delete(ctx)

	if err != nil {
		w.WriteHeader(404)
		fmt.Fprint(w, "Error DeleteUser")
		return
	}

	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Deleted user", id, "\n")

}

