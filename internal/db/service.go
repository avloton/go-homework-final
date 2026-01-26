package db

import (
	"database/sql"
	"log"
	"mywebsite/internal/models"
)


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