package handlers

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type ReportsHandler struct{}

func NewReportsHandler() *ReportsHandler {
	return &ReportsHandler{}
}

// Download is intentionally vulnerable to path traversal by directly using the provided file path.
// Example: /reports/download?file=../../../../etc/passwd
func (h *ReportsHandler) Download(c echo.Context) error {
	file := c.QueryParam("file")
	if file == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "missing file"})
	}
	// Intentionally unsafe: no sanitization, direct read
	data, err := os.ReadFile("./reports/" + file)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "file not found"})
	}
	return c.Blob(http.StatusOK, "application/octet-stream", data)
}
