package delivery

import (
	"haiooId/shoppingChart/features/shopping"
	"haiooId/shoppingChart/utils/helper"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type shoppingDelivery struct {
	shoppingService shopping.ServiceInterface
}

func New(service shopping.ServiceInterface, e *echo.Echo) {
	handler := &shoppingDelivery{
		shoppingService: service,
	}
	e.POST("/shopping-charts", handler.Create)
	e.GET("/shopping-charts", handler.Get)
	e.DELETE("/shopping-charts/:product_code", handler.Delete)
}

func (delivery *shoppingDelivery) Create(c echo.Context) error {
	shoppingInput := shoppingRequest{}
	errBind := c.Bind(&shoppingInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data. "+errBind.Error()))
	}

	dataCore := toCore(shoppingInput)
	err := delivery.shoppingService.Create(dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed insert data. "+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("Success create data"))
}

func (delivery *shoppingDelivery) Get(c echo.Context) error {
	product_name := c.QueryParam("product_name")
	qty := c.QueryParam("qty")

	log.Println("product_name:", product_name)
	log.Println("qty:", qty)

	results, err := delivery.shoppingService.Get(product_name, qty)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	dataResponse := fromCoreList(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success read user.", dataResponse))
}

func (delivery *shoppingDelivery) Delete(c echo.Context) error {
	productCodeParam, _ := strconv.Atoi(c.Param("product_code"))
	err := delivery.shoppingService.Delete(uint(productCodeParam))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Success delete data."))
}
