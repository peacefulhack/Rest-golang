package auth

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

func Run(h Handlers) error {
	e := echo.New()

	api := e.Group("/product")

	api.POST("/insert", h.InsertProductHandlers())
	api.GET("/sort/:key/:order", h.SortProductHandlers())

	err := e.Start(":8080")
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
