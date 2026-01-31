package main

import (
	"log"
	"mywebsite/internal/db"
	"mywebsite/internal/handlers"
	"mywebsite/internal/middleware"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	//Режим инициализации БД
	if len(os.Args) == 2 && os.Args[1] == "populate" {
		db.PopulateDb()
		os.Exit(0)
	}
	//Режим запуска веб-сервиса
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

	//Отдаем метрики через эндпоинт /metrics
	m.Handle("/metrics", promhttp.Handler())

	//Метрики будут собираться через middleware
	handlerWithMetrics := middleware.MetricsMiddleware(m)

	log.Println("Server started on port :8080 ...")
	log.Println("Metrics available at :8080/metrics")
	if err := http.ListenAndServe(":8080", handlerWithMetrics); err != nil {
		log.Fatal(err)
	}
}
