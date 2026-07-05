package handler;

import (
	"net/http"
	"ABC-Live-Web-API/fetchers"
	"github.com/labstack/echo/v4"
);

type Data struct {
	Data	any	`json:"data"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	e := echo.New();

	e.GET("/rbi", func (c echo.Context) error {
		data := Data{fetchers.RBIFetcher()};
		return c.JSON(http.StatusOK, data);
	});

	e.GET("/et_bfsi", func (c echo.Context) error {
		data := Data{fetchers.ETFetcher()};
		return c.JSON(http.StatusOK, data);
	});

	e.GET("/business_line", func (c echo.Context) error {
		data := Data{fetchers.BLFetcher()};
		return c.JSON(http.StatusOK, data);
	});

	e.ServeHTTP(w, r)
}