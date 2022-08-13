package repository

import (
	"log"

	"github.com/widjaya-hernando/prod-trans-user-service/models"

	"gorm.io/gorm"
)

func (r *Repository) AddUser(tx *gorm.DB, user *models.User) error {
	db := r.db
	if tx != nil {
		db = tx
	}
	err := db.Create(user).Error
	if err != nil {
		log.Println("err-add-user: ", err)
		return err
	}
	return nil
}
