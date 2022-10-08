package delivery

import (
	"clean_architecture/feature/vendor/domain"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type VendorControll struct {
	srv domain.VendorService
}

func New(e *echo.Echo, srv *domain.VendorService) {
	ctl := VendorControll{srv: *srv}
	e.POST("/vendors", ctl.AddVendor())
	e.GET("/vendors", ctl.ShowAllVendor())
	//e.GET("/vendors/:expedisi", ctl.ShowExpVendor())
}

func (vc *VendorControll) AddVendor() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input domain.Vendor
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "invalid input",
			})
		}
		res, err := vc.srv.AddVendor(input)

		if err != nil {
			log.Error("vendor handler error: ", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "sucsess input new vendor",
			"data":    res,
		})
	}
}

func (vc *VendorControll) ShowAllVendor() echo.HandlerFunc {
	return func(c echo.Context) error {
		var res []domain.Vendor
		res, err := vc.srv.ShowAllVendor()

		if err != nil {
			log.Error("vendor handler error: ", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "sucsess input new vendor",
			"data":    res,
		})
	}
}

// func (vc *VendorControll) ShowExpVendor() echo.HandlerFunc {

// }
