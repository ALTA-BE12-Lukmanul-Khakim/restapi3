package delivery

import (
	"clean_architecture/config"
	"clean_architecture/feature/logistic/domain"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

type LogisticControll struct {
	srv domain.LogisticService
}

func NewV(e *echo.Echo, srv *domain.LogisticService, ap *config.AppConfig) {
	ctl := LogisticControll{srv: *srv}
	e.POST("/logistics", ctl.AddLogistic(), middleware.JWT(ap.JWTSecret))
	e.GET("/logistics", ctl.ShowAllLogistic())
	e.GET("/logistics/:expedisi", ctl.ShowExpLogistic())
}

func (vc *LogisticControll) AddLogistic() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input domain.Logistic
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "invalid input",
			})
		}
		res, err := vc.srv.AddLogistic(input)

		if err != nil {
			log.Error("Logistic handler error: ", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "sucsess input new Logistic",
			"data":    res,
		})
	}
}

func (vc *LogisticControll) ShowAllLogistic() echo.HandlerFunc {
	return func(c echo.Context) error {
		var res []domain.Logistic
		res, err := vc.srv.ShowAllLogistic()

		if err != nil {
			log.Error("Logistic handler error: ", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "sucsess input new Logistic",
			"data":    res,
		})
	}
}

func (vc *LogisticControll) ShowExpLogistic() echo.HandlerFunc {
	return func(c echo.Context) error {
		var res []domain.Logistic
		res, err := vc.srv.ShowExpLogistic("Expedisi")

		if err != nil {
			log.Error("Logistic handler error: ", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "sucsess input new Logistic",
			"data":    res,
		})

	}
}
