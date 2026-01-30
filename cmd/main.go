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
	m.HandleFunc("/show_orders", handlers.ShowOrdersHandler)
	m.HandleFunc("/show_feedbacks", handlers.ShowFeedbacksHandler)
	m.HandleFunc("/images/", handlers.GetImages)
	m.HandleFunc("/create_order", handlers.CreateOrder)
	m.HandleFunc("/create_feedback", handlers.CreateFeedback)
	m.HandleFunc("/finish_order/", handlers.FinishOrder)
	m.HandleFunc("/return_order/", handlers.ReturnOrder)
	m.HandleFunc("/delete_feedback/", handlers.DeleteFeedback)

	log.Println("Server started on port :8080 ...")
	if err := http.ListenAndServe(":8080", m); err != nil {
		log.Fatal(err)
	}
}
