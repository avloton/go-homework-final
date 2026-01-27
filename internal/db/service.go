package db

import (
	"database/sql"
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
			"completed"	TEXT NOT NULL DEFAULT 'No',
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
		INSERT INTO popular_products (name, description, type, price, units, image_path) VALUES ('Булочка с корицей', '', '', 75, 'шт', 'images/cinnamon_bun.avif');`
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

func SelectAllPopularProducts() []models.Product {
	db := DbConnect()
	defer db.Close()
	rows, err := db.Query(`select id, name, description, type, price, units, image_path from popular_products`)
	if err != nil {
		log.Println("Query SelectAllPopularProducts error: ", err)
		return make([]models.Product, 0)
	}
	defer rows.Close()
	
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