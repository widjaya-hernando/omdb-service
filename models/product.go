package models

import (
	pb "github.com/widjaya-hernando/prod-trans-user-service/gen"
)

type Product struct {
	ID           uint64        `json:"id"`
	ProductName  string        `json:"product_name"`
	Quantity     string        `json:"quantity"`
	Price        string        `json:"price"`
	CreatedAt    int64         `json:"created_at"`
	UpdatedAt    int64         `json:"updated_at,omitempty"`
	Transactions []Transaction `json:"transactions" gorm:"foreignKey:product_id"`
}

func ProductToGrpcMessage(product *Product) *pb.Product {
	return &pb.Product{
		Id:          product.ID,
		ProductName: product.ProductName,
		Quantity:    product.Quantity,
		Price:       product.Price,
	}
}
