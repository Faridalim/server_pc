package routes

import (
	"server_pc/controllers"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {

	e := echo.New()
	e.GET("/pc", controllers.FetchAllPc)
	e.POST("/pc", controllers.InsertPc)
	e.PUT("/pc", controllers.UpdatePc)
	e.POST("/pc/last_active", controllers.UpdateLastActive)
	e.DELETE("/pc", controllers.DeletePc)
	e.POST("/pc/restart_all", controllers.RestartAllPc)
	e.POST("/pc/shutdown_all", controllers.ShutdownAllPc)

	// config := middleware.JWTConfig{
	// 	Claims:     &helpers.JwtCustomClaims{},
	// 	SigningKey: []byte("secret"),
	// }

	//public
	e.GET("/", hello)

	// r := e.Group("/restricted")

	// r.Use(middleware.JWTWithConfig(config))
	// r.GET("/pegawai", controllers.FetchAllPegawai)
	// r.GET("/updateakses", controllers.UpdateAccess)

	return e

}

func hello(c echo.Context) error {
	return c.String(200, "OK")
}
