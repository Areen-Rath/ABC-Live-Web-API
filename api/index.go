package handler;

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"ABC-Live-Web-API/fetchers"
);

type Data struct {
	Data	any	`json:"data"`
}

var e *echo.Echo

func init() {
	e = echo.New()
	e.Use(middleware.Gzip())
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: 15 * time.Second,
	}))

	e.GET("/rbi", func(c echo.Context) error {
		data := Data{fetchers.RBIFetcher()}
		return c.JSON(http.StatusOK, data)
	})

	e.GET("/et_bfsi", func(c echo.Context) error {
		data := Data{fetchers.ETFetcher()}
		return c.JSON(http.StatusOK, data)
	})

	e.GET("/business_line", func(c echo.Context) error {
		data := Data{fetchers.BLFetcher()}
		return c.JSON(http.StatusOK, data)
	})
}

func Handler(w http.ResponseWriter, r *http.Request) {
	e.ServeHTTP(w, r)
}