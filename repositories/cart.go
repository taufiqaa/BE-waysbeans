package repositories

import (
	"waysbeans/models"

	"gorm.io/gorm"
)

type CartRepository interface {
	GetCart(ID int) (models.Cart, error)
	CreateCart(cart models.Cart) (models.Cart, error)
	UpdateCart(cart models.Cart) (models.Cart, error)
	FindCartId(CartId int) ([]models.Cart, error)
	UpdateCartTransaction(models.Cart) (models.Cart, error)
	DeleteCart(cart models.Cart) (models.Cart, error)
}

func RepositoryCart(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetCart(ID int) (models.Cart, error) {
	var cart models.Cart
	err := r.db.Preload("Product").First(&cart, ID).Error

	return cart, err
}

func (r *repository) CreateCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Create(&cart).Error

	return cart, err
}

func (r *repository) UpdateCartTransaction(cart models.Cart) (models.Cart, error) {
	err := r.db.Save(&cart).Error

	return cart, err
}

func (r *repository) UpdateCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Save(&cart).Error

	return cart, err
}

func (r *repository) DeleteCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Delete(&cart).Error

	return cart, err
}

func (r *repository) FindCartId(UserID int) ([]models.Cart, error) {
	var cart []models.Cart
	err := r.db.Preload("Product").Find(&cart, "user_id = ?", UserID).Find(&cart, "transaction_id IS NULL").Error

	return cart, err
}
