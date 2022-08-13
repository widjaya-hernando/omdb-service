package repository

import (
	"github.com/widjaya-hernando/prod-trans-user-service/models"

	"gorm.io/gorm"
)

type RepoMethod interface {
	AddProduct(tx *gorm.DB, product *models.Product) error
	GetProduct(tx *gorm.DB, productId uint64) (*models.Product, error)
	GetProductByTransactionId(tx *gorm.DB, transactionId uint64) (*models.Product, error)
	ListProducts(tx *gorm.DB) ([]models.Product, error)
	UpdateProduct(tx *gorm.DB, product *models.Product) error
}

type Repository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) RepoMethod {
	return &Repository{
		db: db,
	}
}
