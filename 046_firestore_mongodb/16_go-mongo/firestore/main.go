package main

import (
	//"fmt"
	//"log"
	"net/http"
	//"os"
	"errors"
	"cloud.google.com/go/firestore"	
	"google.golang.org/api/iterator"
	"context"
	//"github.com/julienschmidt/httprouter"
	//"encoding/json"	
	//"strconv"
	"html/template"
)

var ctx = context.Background()
var client, err = firestore.NewClient(ctx, "trim-diode-344014")
var TPL *template.Template

type Book struct {
	// add ID and tags if you need them
	// ID     bson.ObjectId // `json:"id" bson:"_id"`
	Isbn   string  // `json:"isbn" bson:"isbn"`
	Title  string  // `json:"title" bson:"title"`
	Author string  // `json:"author" bson:"author"`
	Price  float64 // `json:"price" bson:"price"`
}

func init() {
	TPL = template.Must(template.ParseGlob("templates/*.gohtml"))
}


func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/books", Index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/books/show", Show)
//	http.HandleFunc("/books/create", Create)
//	http.HandleFunc("/books/create/process", CreateProcess)
//	http.HandleFunc("/books/update", Update)
//	http.HandleFunc("/books/update/process", UpdateProcess)
//  http.HandleFunc("/books/delete/process", DeleteProcess)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/books", http.StatusSeeOther)
}

func Index(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	bks, err := AllBooks()

	if err != nil {
		http.Error(w, http.StatusText(500)+err.Error(), http.StatusInternalServerError)
		return
	}

	TPL.ExecuteTemplate(w, "books.gohtml", bks)
}


func Show(w http.ResponseWriter, r *http.Request) {
	
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	bk, err := OneBook(r)		
	
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	TPL.ExecuteTemplate(w, "show.gohtml", bk)
}

/*
func Create(w http.ResponseWriter, r *http.Request) {
	config.TPL.ExecuteTemplate(w, "create.gohtml", nil)
}

func CreateProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	bk, err := PutBook(r)
	if err != nil {
		http.Error(w, http.StatusText(406), http.StatusNotAcceptable)
		return
	}

	config.TPL.ExecuteTemplate(w, "created.gohtml", bk)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	bk, err := OneBook(r)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	config.TPL.ExecuteTemplate(w, "update.gohtml", bk)
}

func UpdateProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	bk, err := UpdateBook(r)
	if err != nil {
		http.Error(w, http.StatusText(406), http.StatusBadRequest)
		return
	}

	config.TPL.ExecuteTemplate(w, "updated.gohtml", bk)
}

func DeleteProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	err := DeleteBook(r)
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/books", http.StatusSeeOther)
}
*/

func AllBooks() ([]Book, error) {
	
	var price_f float64 
	bks_from_db := client.Collection("books").Documents(ctx)
	books := []Book{}

	defer bks_from_db.Stop()
	for {
		book, err := bks_from_db.Next()
		if err == iterator.Done {
			break
		}

		price_f, _ = book.Data()["Price"].(float64)
		books = append(books, Book{book.Data()["Isbn"].(string), book.Data()["Title"].(string), book.Data()["Author"].(string), price_f })		
	}

	return books, err
}


func OneBook(r *http.Request) (Book, error) {	

	var price_f float64 
	var bk Book

	isbn := r.FormValue("isbn")
	
	if isbn == "" {
		return Book{}, errors.New("400. Bad Request.")				
	}	
	
	// Mongo: err := config.Books.Find(bson.M{"isbn": isbn}).One(&bk)

	bk_map := client.Collection("books").Where("Isbn", "==", isbn).Documents(ctx)

	for {
        doc, err := bk_map.Next()
        if err == iterator.Done {
                break
        }
        if err != nil {
                return Book{}, errors.New("400. Bad Request.")	
        }
		price_f = doc.Data()["Price"].(float64)
		bk = Book{ doc.Data()["Isbn"].(string), doc.Data()["Title"].(string), doc.Data()["Author"].(string), price_f }		
	
		break
	}
	
	return bk, err
}

/*
func PutBook(r *http.Request) (Book, error) {
	// get form values
	bk := Book{}
	bk.Isbn = r.FormValue("isbn")
	bk.Title = r.FormValue("title")
	bk.Author = r.FormValue("author")
	p := r.FormValue("price")

	// validate form values
	if bk.Isbn == "" || bk.Title == "" || bk.Author == "" || p == "" {
		return bk, errors.New("400. Bad request. All fields must be complete.")
	}

	// convert form values
	f64, err := strconv.ParseFloat(p, 32)
	if err != nil {
		return bk, errors.New("406. Not Acceptable. Price must be a number.")
	}
	bk.Price = float32(f64)

	// insert values
	err = config.Books.Insert(bk)
	if err != nil {
		return bk, errors.New("500. Internal Server Error." + err.Error())
	}
	return bk, nil
}

func UpdateBook(r *http.Request) (Book, error) {
	// get form values
	bk := Book{}
	bk.Isbn = r.FormValue("isbn")
	bk.Title = r.FormValue("title")
	bk.Author = r.FormValue("author")
	p := r.FormValue("price")

	if bk.Isbn == "" || bk.Title == "" || bk.Author == "" || p == "" {
		return bk, errors.New("400. Bad Request. Fields can't be empty.")
	}

	// convert form values
	f64, err := strconv.ParseFloat(p, 32)
	if err != nil {
		return bk, errors.New("406. Not Acceptable. Enter number for price.")
	}
	bk.Price = float32(f64)

	// update values
	err = config.Books.Update(bson.M{"isbn": bk.Isbn}, &bk)
	if err != nil {
		return bk, err
	}
	return bk, nil
}

func DeleteBook(r *http.Request) error {
	isbn := r.FormValue("isbn")
	if isbn == "" {
		return errors.New("400. Bad Request.")
	}

	err := config.Books.Remove(bson.M{"isbn": isbn})
	if err != nil {
		return errors.New("500. Internal Server Error")
	}
	return nil
}

*/