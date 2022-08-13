package repository

import (
	"log"

	"github.com/widjaya-hernando/prod-trans-user-service/models"

	"gorm.io/gorm"
)

func (r *Repository) UpdateProduct(tx *gorm.DB, product *models.Product) error {
	db := r.db
	if tx != nil {
		db = tx
	}
	err := db.
		Updates(&product).Error

	if err == gorm.ErrRecordNotFound {
		log.Println("err-update-product: ", err)
		return nil
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println("err-update-product: ", err)
		return err
	}

	return nil
}
