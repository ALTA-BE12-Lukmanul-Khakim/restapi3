package delivery

import (
	"clean_architecture/feature/user/domain"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type UserControll struct {
	srv domain.UserService
}

func New(e *echo.Echo, srv *domain.UserService) {
	ctl := UserControll{srv: *srv}
	e.POST("/users", ctl.AddUser())
}

func (uc *UserControll) AddUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input domain.User
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "invalid input",
			})
		}

		res, err := uc.srv.AddUser(input)

		if err != nil {
			log.Error("user handler error: ", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "sucsess register data",
			"data":    res,
		})
	}
}
