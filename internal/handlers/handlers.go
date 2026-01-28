package handlers

import (
	"fmt"
	"log"
	"mywebsite/internal/db"
	"mywebsite/internal/models"
	"net/http"
	"path"
	"text/template"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if r.URL.Path == "/" {
			Products := db.SelectAllPopularProducts()
			tmpl, err := template.ParseFiles("./web/templates/index.html")
			if err != nil {
				log.Println(err)
				return
			}
			tmpl.Execute(w, Products)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("./web/templates/about.html")
		if err != nil {
			log.Println(err)
			return
		}
		tmpl.Execute(w, nil)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func ContactsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("./web/templates/contacts.html")
		if err != nil {
			log.Println(err)
			return
		}
		tmpl.Execute(w, nil)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func MenuHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("./web/templates/menu.html")
		if err != nil {
			log.Println(err)
			return
		}
		tmpl.Execute(w, nil)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func OrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		Products := db.SelectAllPopularProducts()
		tmpl, err := template.ParseFiles("./web/templates/order.html")
		if err != nil {
			log.Println(err)
			return
		}
		tmpl.Execute(w, Products)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func ShowOrdersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		OrdersInfo := db.CountAllOrders()
		Orders := db.SelectAllOrders()
		data := map[string]interface{}{"OrdersInfo": OrdersInfo, "Orders": Orders}
		tmpl, err := template.ParseFiles("./web/templates/show_orders.html")
		if err != nil {
			log.Println(err)
			return
		}
		tmpl.Execute(w, data)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func GetImages(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		filePath := fmt.Sprintf("./web/img/%s", path.Base(r.URL.String()))
		http.ServeFile(w, r, filePath)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var newOrder models.Order
	if r.Method == http.MethodPost {
		newOrder.CustomerName = r.FormValue("name")
		newOrder.Telephone = r.FormValue("phone")
		newOrder.Email = r.FormValue("email")
		newOrder.Address = r.FormValue("address")
		newOrder.DeliveryDate = r.FormValue("delivery-date")
		newOrder.DeliveryTime = r.FormValue("delivery-time")
		newOrder.OrderList = r.FormValue("order-items")
		newOrder.Comments = r.FormValue("comments")
		newOrder.PaymentMethod = r.FormValue("payment")
		err := db.InsertNewOrder(&newOrder)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		html := `
				<!DOCTYPE html>
				<html>
				<head>
					<title>Спасибо за заказ!</title>
				</head>
				<body>
					<div class="thank-you-message">
					
						<h2>Спасибо за ваш заказ!</h2>
						<p>Мы получили вашу заявку и свяжемся с вами в ближайшее время для подтверждения.</p>

						<form action="/show_orders" method="get" style="display: inline-block;">
            				<button type="submit" style="margin-right: 10px;">
                				Отследить заказ
            				</button>
        				</form>
        
        				<form action="/" method="get" style="display: inline-block;">
            				<button type="submit">
                				Вернуться на главную
            				</button>
        				</form>

					</div>
				</body>
				</html>
			`
		fmt.Fprintf(w, html)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func CreateFeedback(w http.ResponseWriter, r *http.Request) {
	var newFeedback models.Feedback
	if r.Method == http.MethodPost {
		newFeedback.CustomerName = r.FormValue("name")
		newFeedback.Email = r.FormValue("email")
		newFeedback.Message = r.FormValue("message")
		newFeedback.Subject = r.FormValue("subject")
		err := db.InsertNewFeedback(&newFeedback)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		html := `
				<!DOCTYPE html>
				<html>
				<head>
					<title>Спасибо!</title>
				</head>
				<body>
					<div class="thank-you-message">
						<h2>Спасибо за ваше сообщение, мы обязательно ответим!</h2>
						<form action="/" method="get">
    						<button type="submit">
								Вернуться на главную
							</button>
						</form>
					</div>
				</body>
				</html>
			`
		fmt.Fprintf(w, html)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
