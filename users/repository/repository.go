package repository

import (
	"github.com/widjaya-hernando/prod-trans-user-service/models"

	"gorm.io/gorm"
)

type RepoMethod interface {
	AddUser(tx *gorm.DB, user *models.User) error
	GetUser(tx *gorm.DB, userId uint64) (*models.User, error)
	GetUserByTransactionId(tx *gorm.DB, transactionId uint64) (*models.User, error)
	ListUsers(tx *gorm.DB) ([]models.User, error)
	UpdateUser(tx *gorm.DB, user *models.User) error
}

type Repository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) RepoMethod {
	return &Repository{
		db: db,
	}
}
