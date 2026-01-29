package db

import (
	"database/sql"
	"fmt"
	"log"
	"mywebsite/internal/models"
)


func PopulateDb() {
	log.Println("Begin populate database ...")
	db := DbConnect()
	defer db.Close()
	query := `
		CREATE TABLE "feedbacks" (
			"id"	INTEGER NOT NULL,
			"customer_name"	TEXT NOT NULL,
			"email"	TEXT NOT NULL,
			"subject"	TEXT,
			"message"	TEXT NOT NULL,
			PRIMARY KEY("id" AUTOINCREMENT)
		);

		CREATE TABLE "orders" (
			"id"	INTEGER NOT NULL,
			"customer_name"	TEXT NOT NULL,
			"telephone"	TEXT NOT NULL,
			"email"	TEXT,
			"address"	TEXT NOT NULL,
			"delivery_date"	TEXT NOT NULL,
			"delivery_time"	TEXT NOT NULL,
			"order_list"	TEXT NOT NULL,
			"comments"	TEXT,
			"payment_method"	TEXT NOT NULL,
			"status"	TEXT NOT NULL DEFAULT 'new',
			PRIMARY KEY("id" AUTOINCREMENT)
		);

		CREATE TABLE "popular_products" (
			"id"	INTEGER NOT NULL,
			"name"	TEXT NOT NULL,
			"description"	TEXT,
			"type"	TEXT,
			"price"	INTEGER,
			"units"	TEXT DEFAULT 'шт',
			"image_path"	TEXT,
			PRIMARY KEY("id" AUTOINCREMENT)
		);

		INSERT INTO popular_products (name, description, type, price, units, image_path) VALUES ('Ржаной хлеб', '', '', 120, '500г', 'images/rue_bread.avif');
		INSERT INTO popular_products (name, description, type, price, units, image_path) VALUES ('Круассан с шок.', '', '', 90, 'шт', 'images/croissant_choco.avif');
		INSERT INTO popular_products (name, description, type, price, units, image_path) VALUES ('Яблочный пирог', '', '', 350, '800г', 'images/apple_pie.avif');
		INSERT INTO popular_products (name, description, type, price, units, image_path) VALUES ('Булочка с корицей', '', '', 75, 'шт', 'images/cinnamon_bun.avif');
		
		INSERT INTO orders (customer_name,telephone,email,address,delivery_date,delivery_time,order_list,comments,payment_method,status) VALUES
	 	('Иван Петров','+7912345678','ivan@mail.ru','ул. Ленина 1','2026-01-29','12:00-15:00','Круассан с шоколадом 1 шт','','card','new'),
	 	('Сергей Рябинин','+7374912123','my@mail.ru','ул. Пушкина 5','2026-01-30','9:00-12:00','Яблочный пирог 1 шт.','Код от домофона 1234','online','delivered');`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Finish")
}

func InsertNewOrder(order *models.Order) error {
	db := DbConnect()
	defer db.Close()
	_, err := db.Exec("INSERT INTO orders (customer_name, telephone, email, address, delivery_date, delivery_time, order_list, comments, payment_method) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)", 
						order.CustomerName, order.Telephone, order.Email, order.Address, order.DeliveryDate, order.DeliveryTime, order.OrderList, order.Comments, order.PaymentMethod)
	if err != nil {
		return err
	}
	return nil
}

func InsertNewFeedback(feedback *models.Feedback) error {
	db := DbConnect()
	defer db.Close()
	_, err := db.Exec("INSERT INTO feedbacks (customer_name, email, subject, message) VALUES (?, ?, ?, ?)", 
						feedback.CustomerName, feedback.Email, feedback.Subject, feedback.Message)
	if err != nil {
		return err
	}
	return nil
}

func CountAllOrders() (models.OrdersInfo) {
	db := DbConnect()
	defer db.Close()
	var ordersInfo models.OrdersInfo
	
	rowAll := db.QueryRow(`SELECT count(*) FROM orders`)	
	err := rowAll.Scan(&ordersInfo.CountAll)
	if err != nil {
		log.Println(err)
	}

	rowNew := db.QueryRow(`SELECT count(*) FROM orders WHERE status = 'new'`)
	err = rowNew.Scan(&ordersInfo.CountNew)
	if err != nil {
		log.Println(err)
	}

	return ordersInfo
}

func SelectAllOrders() []models.Order {
	db := DbConnect()
	defer db.Close()
	rows, err := db.Query(`SELECT 
								id, 
								customer_name, 
								telephone, 
								email, 
								address, 
								delivery_date, 
								delivery_time, 
								order_list, 
								comments, 
								payment_method, 
								status 
							FROM orders`,
						)
	defer rows.Close()
	if err != nil {
		log.Println("Query SelectAllOrders error: ", err)
		return make([]models.Order, 0)
	}

	Orders := []models.Order{}

	for rows.Next() {
		order := models.Order{}
		var OrderEmail sql.NullString
		var OrderComments sql.NullString

		err := rows.Scan(
			&order.Id, 
			&order.CustomerName, 
			&order.Telephone, 
			&OrderEmail, 
			&order.Address, 
			&order.DeliveryDate, 
			&order.DeliveryTime,
			&order.OrderList,
			&OrderComments,
			&order.PaymentMethod,
			&order.Status,
		)
		if err != nil {
			log.Println("Error with SelectAllOrders scan: ", err)
			continue
		}
		if OrderEmail.Valid {
			order.Email = OrderEmail.String
		}
		if OrderComments.Valid {
			order.Comments = OrderComments.String
		}

		order.DeliveryDateTime = fmt.Sprintf("%s %s", order.DeliveryDate, order.DeliveryTime)

		if order.Status == "new" {
			order.StatusText = "Новый"
		}
		if order.Status == "delivered" {
			order.StatusText = "Доставлен"
		}

		Orders = append(Orders, order)
	}
	return Orders
}

func SelectAllPopularProducts() []models.Product {
	db := DbConnect()
	defer db.Close()
	rows, err := db.Query(`SELECT id, name, description, type, price, units, image_path FROM popular_products`)
	defer rows.Close()
	if err != nil {
		log.Println("Query SelectAllPopularProducts error: ", err)
		return make([]models.Product, 0)
	}
	
	Products := []models.Product{}
	
	for rows.Next() {
		product := models.Product{}

		var ProductDesc sql.NullString
		var ProductType sql.NullString
		var ProductPrice sql.NullFloat64
		var ProductUnits sql.NullString
		var ProductsImagePath sql.NullString

		err := rows.Scan(&product.Id, &product.Name, &ProductDesc, &ProductType, &ProductPrice, &ProductUnits, &ProductsImagePath)
		if err != nil {
			log.Println("Error with popular_products scan: ", err)
			continue
		}

		if ProductDesc.Valid {
			product.Description = ProductDesc.String
		}
		if ProductType.Valid {
			product.Type = ProductType.String
		}
		if ProductPrice.Valid {
			product.Price = ProductPrice.Float64
		}
		if ProductUnits.Valid {
			product.Units = ProductUnits.String
		}
		if ProductsImagePath.Valid {
			product.ImagePath = ProductsImagePath.String
		}

		Products = append(Products, product)
	}
	return Products
}