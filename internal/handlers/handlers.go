package handlers

import (
	"fmt"
	"log"
	"mywebsite/internal/db"
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

func GetImages(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		filePath := fmt.Sprintf("./web/img/%s", path.Base(r.URL.String()))
		http.ServeFile(w, r, filePath)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
