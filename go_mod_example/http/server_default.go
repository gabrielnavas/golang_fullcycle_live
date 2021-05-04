package http

import (
	"net/http"

	"github.com/gabrielnavas/go_mod_example/model"
	"github.com/labstack/echo/v4"
)

type WebServer struct {
	Products []model.Product
}

func NewWebServer() *WebServer {
	return &WebServer{}
}

func (w WebServer) Serve() {
	e := echo.New()
	e.GET("/product", w.getAll)
	e.Logger.Fatal(e.Start(":8585"))
}

func (w WebServer) getAll(c echo.Context) error {
	return c.JSON(http.StatusOK, w.Products)
}
