package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/me.li/mvc/controllers"
	"github.com/me.li/mvc/models"
)

type pageDataStruct struct {
	Title    string
	ShortUrl string
	LongUrl  string
}

var pageData pageDataStruct
var t *template.Template

func init() {
	t = template.Must(template.ParseGlob("mvc/views/*.html"))
	models.RedisDbInit()
}

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	pageData = pageDataStruct{Title: "404"}
	t.ExecuteTemplate(w, "error.html", pageData)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/urls", controllers.CreateUrl).Methods("POST")
	router.HandleFunc("/urls/bulk", controllers.CreateUrls).Methods("POST")
	router.HandleFunc("/urls/{id}", controllers.GetUrl).Methods("GET")
	router.HandleFunc("/urls/{id}", controllers.DeleteUrl).Methods("DELETE")
	router.HandleFunc("/{id}", controllers.HomeHandler)
	router.HandleFunc("/error404/{id}", ErrorHandler)
	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":8011", router))
}
