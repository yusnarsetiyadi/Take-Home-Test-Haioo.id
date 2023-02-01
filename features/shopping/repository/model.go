package repository

import (
	"haiooId/shoppingChart/features/shopping"

	"gorm.io/gorm"
)

type Shopping struct {
	gorm.Model
	ProductCode uint `gorm:"primaryKey;autoIncrement;not_null"`
	ProductName string
	Qty         uint
}

func fromCore(dataCore shopping.ShoppingCore) Shopping {
	shoppingGorm := Shopping{
		ProductName: dataCore.ProductName,
		Qty:         dataCore.Qty,
	}
	return shoppingGorm
}

func fromCoreUpdate(dataCore shopping.ShoppingCore) Shopping {
	qty := dataCore.Qty
	shoppingGorm := Shopping{
		Qty: qty,
	}
	return shoppingGorm
}

func (dataModel *Shopping) toCore() shopping.ShoppingCore {
	return shopping.ShoppingCore{
		ProductCode: dataModel.ID,
		ProductName: dataModel.ProductName,
		Qty:         dataModel.Qty,
	}
}

func toCoreList(dataModel []Shopping) []shopping.ShoppingCore {
	var dataCore []shopping.ShoppingCore
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}
