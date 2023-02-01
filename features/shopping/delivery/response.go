package delivery

import "haiooId/shoppingChart/features/shopping"

type ShoppingResponse struct {
	ProductCode uint   `json:"product_code"`
	ProductName string `json:"product_name"`
	Qty         uint   `json:"qty"`
}

func fromCore(dataCore shopping.ShoppingCore) ShoppingResponse {
	return ShoppingResponse{
		ProductCode: dataCore.ProductCode,
		ProductName: dataCore.ProductName,
		Qty:         dataCore.Qty,
	}
}

func fromCoreList(dataCore []shopping.ShoppingCore) []ShoppingResponse {
	var dataResponse []ShoppingResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
