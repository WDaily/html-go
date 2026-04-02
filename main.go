package main

import (
	"html/template"
	"fmt"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)

type Book struct{
	ID 		int
	Title 	string 
	Author 	string 
}

var books = []Book{
			{ID:1,Title:"title",Author: "author"},
			{ID:2,Title:"title",Author: "author"},
			{ID:3,Title:"title",Author: "author"},
			{ID:4,Title:"title",Author: "author"},
			{ID:5,Title:"title",Author: "author"},
		}
	

func main(){
	
	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("./front/static/"))

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/",fs))

	r.Use(methodMiddleware)

	r.HandleFunc("/",getBooks).Methods("GET")
	r.HandleFunc("/book",createBook).Methods("POST")
	r.HandleFunc("/update/book/{id:[0-9]+}",updateBook).Methods("POST","PUT")
	r.HandleFunc("/books/{id:[0-9]+}",deleteBook).Methods("POST","DELETE")
	
	fmt.Println("server starting")

	log.Fatal(http.ListenAndServe(":8080",r))
}


func methodMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter,r*http.Request){

		if r.Method == http.MethodPost{
			method := r.PostFormValue("correct_method")
			if method == http.MethodPut || method == http.MethodDelete{
				r.Method = method
			}
		}
		next.ServeHTTP(w,r)
	})
}

func getBooks(w http.ResponseWriter, r*http.Request){
	t :=template.Must(template.ParseFiles("front/base.html","front/index.html"))

	if err := t.Execute(w,books); err != nil{ 
		log.Fatalf("failed to execute:%v",err)
	}
}

func createBook (w http.ResponseWriter, r*http.Request){

	id := rand.Intn(100)
	title := r.PostFormValue("Title")
	author := r.PostFormValue("Author")

	books = append(books,Book{ID:id,Title:title,Author:author})


	http.Redirect(w,r,"/",http.StatusFound)
}

func updateBook(w http.ResponseWriter, r*http.Request){
	param := mux.Vars(r)

	id,_ := strconv.Atoi(param["id"])

	for index,item := range books{
		if item.ID == id {
			
			books = append(books[:index],books[index+1:]...)

			var book Book

			book.ID = id
			book.Title = r.PostFormValue("Title")
			book.Author = r.PostFormValue("Author")

			books = append(books,book)
		}
	}
	http.Redirect(w,r,"/",http.StatusFound)
}

func deleteBook(w http.ResponseWriter, r*http.Request){
	param := mux.Vars(r)

	id,_ := strconv.Atoi(param["id"])

	for index,item := range books{
		if item.ID == id{
			books = append(books[:index],books[index+1:]...)
			break
		}
	}
	http.Redirect(w,r,"/", http.StatusFound)
}