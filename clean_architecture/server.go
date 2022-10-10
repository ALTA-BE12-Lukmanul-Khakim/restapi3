package main

import (
	"clean_architecture/config"

	dataUser "clean_architecture/feature/user/data"
	userDelivery "clean_architecture/feature/user/delivery"
	userService "clean_architecture/feature/user/services"

	dataLogistic "clean_architecture/feature/logistic/data"
	logisticDelivery "clean_architecture/feature/logistic/delivery"
	logisticService "clean_architecture/feature/logistic/services"
	"clean_architecture/utils/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	//pemanggialn echo
	e := echo.New()

	//pemanggilan config
	cfg := config.NewConfig()
	db := database.InitDB(cfg)

	database.MigrateDB(db)

	mdl := dataUser.NewU(db)
	mdv := dataLogistic.NewV(db)

	uServices := userService.New(mdl)
	vServices := logisticService.New(mdv)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))

	userDelivery.New(e, &uServices)
	logisticDelivery.NewV(e, &vServices, cfg)

	log.Fatal(e.Start(":8000"))

}
