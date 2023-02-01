package main

import (
	"fmt"
	"haiooId/shoppingChart/config"
	"haiooId/shoppingChart/factory"
	"haiooId/shoppingChart/utils/database/mysql"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.GetConfig()
	db := mysql.InitDB(cfg)

	e := echo.New()
	factory.InitFactory(e, db)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.SERVER_PORT)))
}
