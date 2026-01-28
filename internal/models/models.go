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
	Id 					int
	CustomerName 		string
	Telephone 			string
	Email 				string
	Address 			string
	DeliveryDate 		string
	DeliveryTime 		string
	OrderList 			string
	Comments 			string
	PaymentMethod 		string
	Status				string
	StatusText			string
	DeliveryDateTime	string
}

type OrdersInfo struct {
	CountAll		int
	CountNew		int
}

type Feedback struct {
	Id 				int
	CustomerName 	string
	Email			string
	Subject 		string
	Message 		string
}