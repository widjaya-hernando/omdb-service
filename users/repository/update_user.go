package repository

import (
	"log"

	"github.com/widjaya-hernando/prod-trans-user-service/models"

	"gorm.io/gorm"
)

func (r *Repository) UpdateUser(tx *gorm.DB, user *models.User) error {
	db := r.db
	if tx != nil {
		db = tx
	}
	err := db.
		Debug().
		Updates(&user).Error

	if err == gorm.ErrRecordNotFound {
		log.Println("err-update-user: ", err)
		return nil
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println("err-update-user: ", err)
		return err
	}

	return nil
}
