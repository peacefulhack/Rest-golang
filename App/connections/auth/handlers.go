package auth

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"restArchitecture/mikail/App/controllers"
	"restArchitecture/mikail/App/models"
	"restArchitecture/mikail/App/utils/enum"
)

type HandlersInterface interface {
	InsertProductHandlers() echo.HandlerFunc
	SortProductHandlers() echo.HandlerFunc
}

type Handlers struct {
	controllers controllers.ProductControllers
}

func NewAuthHandlers(ct controllers.ProductControllers) Handlers {
	return Handlers{controllers: ct}
}

func (m *Handlers) InsertProductHandlers() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := &models.InsertProductRequest{}
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		}

		if req.ProductName == "" || req.ProductDesc == "" || req.ProductPrice == 0 {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		}

		err := m.controllers.InsertProductData(req)
		if err != nil {
			return c.JSON(enum.ErrorResponse(err))
		}
		return c.JSON(http.StatusCreated, "Product telah ditambahkan")
	}

}

func (m *Handlers) SortProductHandlers() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := m.controllers.SortProductData(c.Param("key"), c.Param("order"))
		if err != nil {
			fmt.Println(err)
			return c.JSON(enum.ErrorResponse(err))
		}
		return c.JSON(http.StatusCreated, res)
	}
}
