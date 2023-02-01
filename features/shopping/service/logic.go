package service

import (
	"haiooId/shoppingChart/features/shopping"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/gommon/log"
)

type shoppingService struct {
	shoppingRepository shopping.RepositoryInterface
	validate           *validator.Validate
}

func New(repo shopping.RepositoryInterface) shopping.ServiceInterface {
	return &shoppingService{
		shoppingRepository: repo,
		validate:           validator.New(),
	}
}

func (service *shoppingService) Create(input shopping.ShoppingCore) (err error) {
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}

	dataProduct, errFindProduct := service.shoppingRepository.FindProduct(input.ProductName)
	if errFindProduct != nil {
		return errFindProduct
	}

	if dataProduct.ProductName == input.ProductName {
		input.ProductName = dataProduct.ProductName
		errCreate := service.shoppingRepository.Update(input, dataProduct.ProductCode)
		if errCreate != nil {
			log.Error(errCreate.Error())
			return err
		}
	} else {
		errCreate := service.shoppingRepository.Create(input)
		if errCreate != nil {
			log.Error(errCreate.Error())
			return err
		}
	}

	return nil
}

func (service *shoppingService) Get(product_name, qty string) (data []shopping.ShoppingCore, err error) {
	data, err = service.shoppingRepository.Get(product_name, qty)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return data, err

}

func (service *shoppingService) Delete(productCode uint) error {
	err := service.shoppingRepository.Delete(productCode)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}
