package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"ABC-Live-Web-API/fetcher"
);

type Data struct {
	Data	any	`json:"data"`
}

type MixData struct {
	News	[]fetcher.News	`json:"news"`
	Stats	[2][7]string	`json:"stats"`
}

var e *echo.Echo

func init() {
	e = echo.New()
	e.Use(middleware.Gzip())
	e.Use(middleware.ContextTimeoutWithConfig(middleware.ContextTimeoutConfig{
		Timeout: 15 * time.Second,
	}))

	e.GET("/", func(c echo.Context) error {
		et, bl, rbi := fetcher.Fetch()
		return c.JSON(http.StatusOK, MixData{
			append(et, bl...),
			rbi,
		})
	})
}

func Handler(w http.ResponseWriter, r *http.Request) {
	e.ServeHTTP(w, r)
}