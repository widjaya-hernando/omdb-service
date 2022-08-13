package repository

import (
	"log"

	"github.com/widjaya-hernando/prod-trans-user-service/models"

	"gorm.io/gorm"
)

func (r *Repository) ListUsers(tx *gorm.DB) ([]models.User, error) {
	var results []models.User

	db := r.db
	if tx != nil {
		db = tx
	}
	err := db.
		Find(&results).Error

	if err == gorm.ErrRecordNotFound {
		log.Println("err-list-users: ", err)
		return nil, nil
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println("err-list-users: ", err)
		return nil, err
	}

	return results, nil
}
