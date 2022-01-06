package controllers

import (
	"net/http"
	"server_pc/security"

	"github.com/labstack/echo/v4"
)

func UpdateAccess(c echo.Context) error {
	security.UpdateAccess()
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Akses Diupdate",
	})
}
