package factory

import (
	shoppingDelivery "haiooId/shoppingChart/features/shopping/delivery"
	shoppingRepo "haiooId/shoppingChart/features/shopping/repository"
	shoppingService "haiooId/shoppingChart/features/shopping/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	shoppingRepoFactory := shoppingRepo.New(db)
	shoppingServiceFactory := shoppingService.New(shoppingRepoFactory)
	shoppingDelivery.New(shoppingServiceFactory, e)
}
