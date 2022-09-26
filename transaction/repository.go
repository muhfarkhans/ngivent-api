package transaction

import (
	"ngevent-api/helper"

	"gorm.io/gorm"
)

type Repository interface {
	Save(transaction Transaction) (Transaction, error)
	Update(transaction Transaction) (Transaction, error)
	FindAll(params helper.Pagination, userType string, userId int) ([]Transaction, *helper.Pagination, error)
	FindById(id int, userType string, userId int) (Transaction, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(transaction Transaction) (Transaction, error) {
	err := r.db.Create(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *repository) Update(transaction Transaction) (Transaction, error) {
	err := r.db.Save(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *repository) FindAll(params helper.Pagination, userType string, userId int) ([]Transaction, *helper.Pagination, error) {
	var transactions []Transaction
	db := r.db

	if userType == "member" {
		db = db.Where("user_id = ?", userId)
	}

	err := db.Preload("Event").Preload("User").Scopes(helper.Paginate(transactions, &params, r.db)).Find(&transactions)
	if err != nil {
		return transactions, &params, err.Error
	}

	return transactions, &params, nil
}

func (r *repository) FindById(id int, userType string, userId int) (Transaction, error) {
	var user Transaction
	db := r.db

	if userType == "member" {
		db.Where("user_id = ?", userId)
	}

	err := db.Preload("Event").Preload("User").Where("id = ?", id).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
