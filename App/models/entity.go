package models

type ProductList struct {
	id              int
	ProductName     string
	ProductDesc     string
	ProductPrice    int
	ProductQuantity int
	CreatedDate     string
}

type InsertProductRequest struct {
	ProductName     string `json:"ProductName"`
	ProductDesc     string `json:"ProductDesc"`
	ProductPrice    int    `json:"ProductPrice"`
	ProductQuantity int    `json:"ProductQuantity"`
}

type SortProductResponse struct {
	Id              int    `json:"Id"`
	ProductName     string `json:"ProductName"`
	ProductDesc     string `json:"ProductDesc"`
	ProductPrice    int    `json:"ProductPrice"`
	ProductQuantity int    `json:"ProductQuantity"`
}
