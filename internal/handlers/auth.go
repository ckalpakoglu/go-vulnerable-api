package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/cenk-kalpakoglu/go-vulnerable-api/internal/config"
	"github.com/cenk-kalpakoglu/go-vulnerable-api/internal/store"
	"github.com/cenk-kalpakoglu/go-vulnerable-api/internal/util"
)

type AuthHandler struct {
	Store  *store.MemoryStore
	Config config.AppConfig
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewAuthHandler(store *store.MemoryStore, cfg config.AppConfig) *AuthHandler {
	return &AuthHandler{Store: store, Config: cfg}
}

func (h *AuthHandler) Login(c echo.Context) error {
	var req loginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	user, ok := h.Store.Authenticate(req.Email, req.Password)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid credentials"})
	}
	token, err := util.GenerateJWT(h.Config.JWTSecret, user.ID, user.Email, 24*time.Hour)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not generate token"})
	}
	return c.JSON(http.StatusOK, map[string]string{"token": token})
}
