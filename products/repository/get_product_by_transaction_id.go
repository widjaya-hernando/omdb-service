package repository

import (
	"log"

	"github.com/widjaya-hernando/prod-trans-user-service/models"
	"github.com/widjaya-hernando/prod-trans-user-service/utils/errors"

	"gorm.io/gorm"
)

func (r *Repository) GetProductByTransactionId(tx *gorm.DB, transactionId uint64) (*models.Product, error) {
	var result models.Transaction

	db := r.db
	if tx != nil {
		db = tx
	}
	err := db.
		Where("id = ?", transactionId).
		Preload("Product").
		First(&result).Error

	if err == gorm.ErrRecordNotFound {
		log.Println("err-get-product-by-transaction-id: ", err)
		return nil, errors.ErrNotFound
	}
	if err != nil {
		log.Println("err-get-product-by-transaction-id: ", err)
		return nil, err
	}

	return result.Product, nil
}
