package routes

import (
	"echo_rest/controllers"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {

	e := echo.New()

	e.GET("/", hello)
	e.GET("/pegawai", controllers.FetchAllPegawai)
	e.POST("/pegawai", controllers.StorePegawai)
	e.PUT("/pegawai", controllers.UpdatePegawai)
	e.DELETE("/pegawai", controllers.DeletePegawai)

	return e
	// e.Logger.Fatal(e.Start("127.0.0.1:3000"))
}

func hello(c echo.Context) error {
	return c.String(200, "OK")
}
