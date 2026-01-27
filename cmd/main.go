package main

import (
	"fmt"
	"log"
	"mywebsite/internal/handlers"
	"mywebsite/internal/db"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) == 2 && os.Args[1] == "populate" {
		fmt.Println("Populate!")
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
	log.Println("Server started ...")
	http.ListenAndServe(":8080", m)
}