package controllers

import (
	"fmt"
	"net/http"
	"server_pc/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

func checkError(err error, c echo.Context) {
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
}

func FetchAllPc(c echo.Context) error {
	res, err := models.FetchAllPc()
	checkError(err, c)

	return c.JSON(http.StatusOK, res)

}

func InsertPc(c echo.Context) error {
	nama := c.FormValue("nama")
	lokasi := c.FormValue("lokasi")

	res, err := models.InsertPc(nama, lokasi)
	checkError(err, c)

	return c.JSON(http.StatusOK, res)

}

func UpdatePc(c echo.Context) error {
	nama := c.FormValue("nama")
	lokasi := c.FormValue("lokasi")
	id := c.FormValue("id")

	id_int, err := strconv.Atoi(id)
	checkError(err, c)

	res, err := models.UpdatePc(id_int, nama, lokasi)
	checkError(err, c)

	return c.JSON(200, res)
}

func UpdateLastActive(c echo.Context) error {
	id := c.FormValue("id")
	fmt.Println(id)
	id_int, err := strconv.Atoi(id)
	checkError(err, c)

	result, err := models.UpdateLastActive(id_int)
	checkError(err, c)
	return c.JSON(200, result)
}

func DeletePc(c echo.Context) error {
	id := c.FormValue("id")
	id_int, err := strconv.Atoi(id)
	checkError(err, c)

	result, err := models.DeletePc(id_int)
	checkError(err, c)
	return c.JSON(200, result)
}

func RestartAllPc(c echo.Context) error {

	result, err := models.RestartAllPc()
	checkError(err, c)
	return c.JSON(200, result)
}

func ShutdownAllPc(c echo.Context) error {
	res, err := models.ShutdownAllPc()
	checkError(err, c)

	return c.JSON(200, res)
}
