package handlers

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"

	"github.com/cenk-kalpakoglu/go-vulnerable-api/internal/store"
	"github.com/cenk-kalpakoglu/go-vulnerable-api/internal/util"
)

type UsersHandler struct {
	Store *store.MemoryStore
}

func NewUsersHandler(store *store.MemoryStore) *UsersHandler {
	return &UsersHandler{Store: store}
}

func (h *UsersHandler) List(c echo.Context) error {
	// Only admin can list all users
	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}
	claims, ok := token.Claims.(*util.Claims)
	if !ok || claims.Email != "admin" {
		return c.JSON(http.StatusForbidden, map[string]string{"error": "forbidden"})
	}
	return c.JSON(http.StatusOK, h.Store.ListUsers())
}

// GetMobile now accepts a query parameter 'mobile' and returns the full user
func (h *UsersHandler) GetMobile(c echo.Context) error {
	mobile := c.QueryParam("mobile")
	if mobile == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "missing mobile"})
	}
	u, ok := h.Store.GetUserByMobile(mobile)
	if !ok {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "user not found"})
	}
	return c.JSON(http.StatusOK, u)
}

// Hello returns information only for the authenticated user
func (h *UsersHandler) Hello(c echo.Context) error {
	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}
	claims, ok := token.Claims.(*util.Claims)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}
	// Return minimal info for current user
	u, ok := h.Store.GetUserByID(claims.UserID)
	if !ok {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "user not found"})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "hello",
		"user":    u,
	})
}

//func (h *UsersHandler) DownloadReports(c echo.Context) error {
//	reportYear := c.QueryParam("year")
//	if date == "" {
//		return c.JSON(http.StatusBadRequest, map[string]string{"error": "missing report date"})
//	}
//
//	token, ok := c.Get("user").(*jwt.Token)
//	if !ok {
//		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
//	}
//	claims, ok := token.Claims.(*util.Claims)
//	if !ok {
//		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
//	}
//  filename := fmt.Sprintf("reports/%s/%s.pdf", claims.UserID, reportYear)
//}
