package repository

import (
	"log"

	"github.com/widjaya-hernando/prod-trans-user-service/models"

	"gorm.io/gorm"
)

func (r *Repository) GetUser(tx *gorm.DB, userId uint64) (*models.User, error) {
	var result *models.User

	db := r.db
	if tx != nil {
		db = tx
	}
	err := db.
		Where("id = ?", userId).
		First(&result).Error

	if err == gorm.ErrRecordNotFound {
		log.Println("err-get-user: ", err)
		return nil, nil
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println("err-get-user: ", err)
		return nil, err
	}

	return result, nil
}
