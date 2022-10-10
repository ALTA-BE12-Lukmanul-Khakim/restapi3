package delivery

import (
	"clean_architecture/config"
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
	e.POST("/register", ctl.AddUser())
	e.POST("/login", ctl.Login())
}

func (uc *UserControll) AddUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input RegisterFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}

		cnv := ToDomain(input)
		res, err := uc.srv.Register(cnv)

		if err != nil {
			log.Error("user	 handler error: ", err.Error())
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}

		return c.JSON(http.StatusCreated, SuccessResponse("berhasil register", ToResponse(res, "reg")))
	}
}

func (uc *UserControll) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var login domain.User
		if err := c.Bind(&login); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "invalid input",
			})
		}

		res, err := uc.srv.Login(login.Nama, login.HP)

		if err != nil {
			log.Error("user handler error: ", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}

		resToken := config.GenerateToken(login.ID)

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get specific data",
			"data":    res,
			"token":   resToken,
		})
	}
}
