package repository

import (
	"log"

	"github.com/widjaya-hernando/prod-trans-user-service/models"

	"gorm.io/gorm"
)

func (r *Repository) ListTransactions(tx *gorm.DB) ([]models.Transaction, error) {
	var results []models.Transaction

	db := r.db
	if tx != nil {
		db = tx
	}
	err := db.
		Find(&results).Error

	if err == gorm.ErrRecordNotFound {
		log.Println("err-list-transactions: ", err)
		return nil, nil
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println("err-list-transactions: ", err)
		return nil, err
	}

	return results, nil
}
