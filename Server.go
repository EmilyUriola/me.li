package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/me.li/mvc/controllers"
	"github.com/me.li/mvc/models"
)

func init() {
	models.RedisDbInit()
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/urls", controllers.CreateUrl).Methods("POST")
	router.HandleFunc("/urls/bulk", controllers.CreateUrls).Methods("POST")
	router.HandleFunc("/urls/{id}", controllers.GetUrl).Methods("GET")
	router.HandleFunc("/urls/{id}", controllers.DeleteUrl).Methods("DELETE")
	router.HandleFunc("/{id}", controllers.HomeHandler)
	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":8011", router))
}
