package main

import (
	"github.com/julienschmidt/httprouter"
	//"gopkg.in/mgo.v2"
	"net/http"
	"gopkg.in/mgo.v2/bson"
	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
	"context"
	"fmt"
	"log"
	"io"
	"encoding/json"	
	//"google.golang.org/api/iterator"

	//"./controllers"
)

var ctx = context.Background()
//var client *firestore.Client
var opt = option.WithCredentialsFile("../../../../service_accounts/042/trim-diode-344014-af63a3b83268.json")
//var opt = option.WithCredentialsFile("firestore-sa@trim-diode-344014.iam.gserviceaccount.com")
//var opt = option.WithCredentialsFile("trim-diode-344014@appspot.gserviceaccount.com")
//var client, err = firestore.NewClient(ctx, "trim-diode-344014", opt)
var client, err = firestore.NewClient(ctx, "trim-diode-344014@appspot.gserviceaccount.com", opt)


type User struct {
	Id     bson.ObjectId `json:"id" bson:"_id"`
	Name   string        `json:"name" bson:"name"`
	Gender string        `json:"gender" bson:"gender"`
	Age    int           `json:"age" bson:"age"`
}

type server struct {
    router    *httprouter.Router
    firestore *firestore.Client
}

func newServer(client *firestore.Client) *server {
	s := &server{router: httprouter.New(), firestore: client}
	s.routes()
	return s
}

func (s *server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	s.router.ServeHTTP(writer, request)
}

func (s *server) routes() {

   s.router.GET("/", Index) 
   s.router.GET("/user/:id", GetUser) 
   s.router.POST("/user", CreateUser) 
   s.router.DELETE("/user/:id", DeleteUser) 
}

func main() {

	if err != nil {
		 log.Fatalf("firestore new error:%s\n", err)
	}
	defer client.Close()

	s := newServer(client)
	server := http.Server{
        Addr:         ":" + "8080",
        Handler:      s,
	}    
	server.ListenAndServe()
}

func Index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	io.WriteString(w, "huhuu")
}

func GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	//fmt.Println(p.ByName("id"))
	//fmt.Println(r)
	//fmt.Println(ctx)
	//fmt.Println(client)
	//iterator := userCollection.createIterator()
	/*
	iter := client.Collection("users").Documents(ctx)
	for {
        doc, err := iter.Next()
        if err == iterator.Done {
                break
        }
        if err != nil {
                log.Fatalf("Failed to iterate: %v", err)
        }
        fmt.Println(doc.Data())
}
*/
	/*
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound) // 404
		return
	}
*/
	//oid := bson.ObjectIdHex(id)

	//u := models.User{}

	u, err := client.Collection("users").Doc(p.ByName("id")).Get(ctx)
	if u != nil {
	 //fmt.Println(u)
	 /*
	 for {
		doc, err := u.Next()
		if err == iterator.Done {
				break
		}
		if err != nil {
				//return fmt.Errorf("documents iterator: %v", err)
				io.WriteString(w, err.Error())
		}
		//fmt.Fprintf(w, "%s: %s", doc.Ref.ID, doc.Data()["name"])
		io.WriteString(w, doc.Data()["name"])
	}
	*/

	 //io.WriteString(w, u.Data().id)
	 //io.WriteString(w, u.string())
	}
	
	//fmt.Println(err)
	//w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(http.StatusOK) // 200

	io.WriteString(w, u.Data()["name"].(string))
	io.WriteString(w, err.Error())
	//io.WriteString(w, json.Marshal(client))
	
	/*
	if err := client.session.DB("go-web-dev-db").C("users").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		return
	}
*/
/*
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
	*/
}

func CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {	
	
	fmt.Println("Add user")

	u := User{}
	json.NewDecoder(r.Body).Decode(&u)

	u.Id = bson.NewObjectId()

//	client.session.DB("go-web-dev-db").C("users").Insert(u)

	_, _, err := client.Collection("users").Add(ctx,u)

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}


func DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	fmt.Println("Delete user")

	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(id)

	// Delete user
	/*
	if err := uc.session.DB("go-web-dev-db").C("users").RemoveId(oid); err != nil {
		w.WriteHeader(404)
		return
	}
*/
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Deleted user", oid, "\n")
}
