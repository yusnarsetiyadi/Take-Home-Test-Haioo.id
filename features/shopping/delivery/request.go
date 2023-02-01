package delivery

import "haiooId/shoppingChart/features/shopping"

type shoppingRequest struct {
	ProductName string `json:"product_name" form:"product_name"`
	Qty         uint   `json:"qty" form:"qty"`
}

func toCore(shoppingInput shoppingRequest) shopping.ShoppingCore {
	shoppingCoreData := shopping.ShoppingCore{
		ProductName: shoppingInput.ProductName,
		Qty:         shoppingInput.Qty,
	}
	return shoppingCoreData
}
