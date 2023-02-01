package repository

import (
	"errors"
	"haiooId/shoppingChart/features/shopping"

	"gorm.io/gorm"
)

type shoppingRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) shopping.RepositoryInterface {
	return &shoppingRepository{
		db: db,
	}
}

func (repo *shoppingRepository) Create(input shopping.ShoppingCore) error {
	shoppingGorm := fromCore(input)
	tx := repo.db.Create(&shoppingGorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed")
	}
	return nil
}

func (repo *shoppingRepository) Update(inputQty shopping.ShoppingCore, product_code uint) error {
	shoppingGorm := fromCoreUpdate(inputQty)
	var shopping Shopping
	tx := repo.db.Model(&shopping).Where("product_code = ?", product_code).Updates(&shoppingGorm) // proses update
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("update failed")
	}
	return nil
}

func (repo *shoppingRepository) Get(product_name, qty string) (data []shopping.ShoppingCore, err error) {
	var shopping []Shopping

	tx := repo.db.Where("product_name LIKE ? AND qty LIKE ?", "%"+product_name+"%", "%"+qty+"%").Find(&shopping)

	if tx.Error != nil {
		return data, tx.Error
	}

	if tx.RowsAffected == 0 {
		return data, tx.Error
	}

	var dataCore = toCoreList(shopping)
	return dataCore, nil
}

func (repo *shoppingRepository) Delete(product_code uint) error {
	var shopping Shopping
	tx := repo.db.Delete(&shopping, product_code)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("delete failed")
	}
	return nil
}

func (repo *shoppingRepository) FindProduct(product string) (result shopping.ShoppingCore, err error) {
	var shoppingData Shopping
	tx := repo.db.Where("product_name = ?", product).First(&shoppingData)
	if tx.Error != nil {
		return shopping.ShoppingCore{}, tx.Error
	}

	result = shoppingData.toCore()

	return result, nil
}
