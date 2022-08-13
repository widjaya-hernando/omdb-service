package repository

import (
	"log"

	"github.com/widjaya-hernando/prod-trans-user-service/models"

	"gorm.io/gorm"
)

func (r *Repository) GetProduct(tx *gorm.DB, productId uint64) (*models.Product, error) {
	var result *models.Product

	db := r.db
	if tx != nil {
		db = tx
	}
	err := db.
		Where("id = ?", productId).
		First(&result).Error

	if err == gorm.ErrRecordNotFound {
		log.Println("err-get-product: ", err)
		return nil, nil
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println("err-get-product: ", err)
		return nil, err
	}

	return result, nil
}
