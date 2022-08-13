package models

import pb "github.com/widjaya-hernando/prod-trans-user-service/gen"

type Transaction struct {
	ID        uint64   `json:"id"`
	UserID    uint64   `json:"user_id"`
	ProductID uint64   `json:"product_id"`
	Quantity  string   `json:"quantity"`
	Total     string   `json:"total"`
	CreatedAt int64    `json:"created_at"`
	Product   *Product `json:"product" gorm:"foreignKey:product_id"`
	User      *User    `json:"user" gorm:"foreignKey:user_id"`
}

func TransactionToGrpcMessage(transaction *Transaction) *pb.Transaction {
	return &pb.Transaction{
		Id:        transaction.ID,
		UserId:    transaction.UserID,
		ProductId: transaction.ProductID,
		Quantity:  transaction.Quantity,
		Total:     transaction.Total,
	}
}
