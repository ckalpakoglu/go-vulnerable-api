package router

import (
	"net/http"

	jwtmw "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"

	"github.com/golang-jwt/jwt/v5"

	"github.com/labstack/echo-contrib/pprof"

	"github.com/cenk-kalpakoglu/go-vulnerable-api/internal/config"
	"github.com/cenk-kalpakoglu/go-vulnerable-api/internal/handlers"
	"github.com/cenk-kalpakoglu/go-vulnerable-api/internal/store"
	"github.com/cenk-kalpakoglu/go-vulnerable-api/internal/util"
)

func Register(e *echo.Echo, cfg config.AppConfig, store *store.MemoryStore) {
	// Health
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	// pprof
	pprof.Register(e)

	otpHandler := handlers.NewOTPHandler(store)
	e.POST("/forgot-password", otpHandler.ForgotPassword)

	authHandler := handlers.NewAuthHandler(store, cfg)
	e.POST("/login", authHandler.Login)

	usersHandler := handlers.NewUsersHandler(store)
	reportsHandler := handlers.NewReportsHandler()

	// JWT middleware
	jwtConfig := jwtmw.Config{
		SigningKey:    []byte(cfg.JWTSecret),
		NewClaimsFunc: func(c echo.Context) jwt.Claims { return &util.Claims{} },
		TokenLookup:   "header:Authorization:Bearer ",
	}

	g := e.Group("/users")
	g.Use(jwtmw.WithConfig(jwtConfig))
	g.GET("", usersHandler.List)
	g.GET("/mobile", usersHandler.GetMobile)
	g.GET("/hello", usersHandler.Hello)

	// Reports endpoint (still path traversal vulnerable by design) with JWT auth
	rg := e.Group("/reports")
	rg.Use(jwtmw.WithConfig(jwtConfig))
	rg.GET("/download", reportsHandler.Download)
}
