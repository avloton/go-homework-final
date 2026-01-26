package main

import (
	"log"
	"mywebsite/internal/handlers"
	"net/http"
)

func main() {
	m := http.NewServeMux()
	m.HandleFunc("/", handlers.IndexHandler)
	m.HandleFunc("/about", handlers.AboutHandler)
	m.HandleFunc("/contacts", handlers.ContactsHandler)
	m.HandleFunc("/menu", handlers.MenuHandler)
	m.HandleFunc("/order", handlers.OrderHandler)
	m.HandleFunc("/images/", handlers.GetImages)
	log.Println("Server started ...")
	http.ListenAndServe(":8080", m)
}