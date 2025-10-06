package handlers

import (
	"net/http"

	"github.com/cenk-kalpakoglu/go-vulnerable-api/internal/store"
	"github.com/cenk-kalpakoglu/go-vulnerable-api/internal/util"

	"github.com/labstack/echo/v4"
)

type OTPHandler struct {
	Store *store.MemoryStore
}

func NewOTPHandler(store *store.MemoryStore) *OTPHandler {
	return &OTPHandler{Store: store}
}

type forgotPasswordRequest struct {
	Email  string `json:"email"`
	Mobile string `json:"mobile"`
}

func (h *OTPHandler) ForgotPassword(c echo.Context) error {
	var req forgotPasswordRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	if h.Store.CheckUserAndMobile(req.Email, req.Mobile) {
		token := util.GenerateRandomOTP(req.Email)
		otp, ok := h.Store.AddNewOTP(token)
		if !ok {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid otp"})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "OTP has been sent",
			"OTP":     otp.Value,
		})
	}

	return c.JSON(http.StatusNotFound, map[string]string{"error": "user and mobile not found"})
}
