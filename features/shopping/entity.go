package shopping

type ShoppingCore struct {
	ProductCode uint
	ProductName string
	Qty         uint
}

type ServiceInterface interface {
	Create(input ShoppingCore) error
	Get(product_name, qty string) (data []ShoppingCore, err error)
	Delete(id uint) error
}

type RepositoryInterface interface {
	Create(input ShoppingCore) error
	Get(product_name, qty string) (data []ShoppingCore, err error)
	Delete(product_code uint) error
	Update(input ShoppingCore, product_code uint) error
	FindProduct(product string) (data ShoppingCore, err error)
}
