package main

import (
	"log"
	"mywebsite/internal/db"
	"mywebsite/internal/handlers"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) == 2 && os.Args[1] == "populate" {
		db.PopulateDb()
		os.Exit(0)
	}
	m := http.NewServeMux()
	m.HandleFunc("/", handlers.IndexHandler)
	m.HandleFunc("/about", handlers.AboutHandler)
	m.HandleFunc("/contacts", handlers.ContactsHandler)
	m.HandleFunc("/menu", handlers.MenuHandler)
	m.HandleFunc("/order", handlers.OrderHandler)
	m.HandleFunc("/images/", handlers.GetImages)
	m.HandleFunc("/create_order", handlers.CreateOrder)
	
	log.Println("Server started on port :8080 ...")
	http.ListenAndServe(":8080", m)
}
