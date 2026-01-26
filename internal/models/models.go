package models

type Product struct {
	Id 				int
	Name 			string
	Description 	string
	Type 			string
	Price 			float64
	Units 			string
	ImagePath 		string
}

type Order struct {
	Id 				int
	CustomerName 	string
	Telephone 		string
	Email 			string
	Address 		string
	DeliveryDate 	string
	DeliveryTime 	string
	OrderList 		string
	Comments 		string
	PaymentMethod 	string
	Completed 		bool
}

type Feedback struct {
	Id 				int
	CustomerName 	string
	Subject 		string
	Message 		string
}