package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/cenk-kalpakoglu/go-vulnerable-api/internal/config"
	"github.com/cenk-kalpakoglu/go-vulnerable-api/internal/router"
	"github.com/cenk-kalpakoglu/go-vulnerable-api/internal/store"
)

func New() *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	cfg := config.Load()
	ms, _ := store.NewMemoryStore()
	router.Register(e, cfg, ms)

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		c.Logger().Error(err)
		_ = c.JSON(http.StatusInternalServerError, map[string]string{"error": http.StatusText(http.StatusInternalServerError)})
	}

	return e
}

func Start(e *echo.Echo, addr string) error {
	return e.Start(addr)
}

