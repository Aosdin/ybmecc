package route

import (
	"github.com/labstack/echo"
	"ybmecc/api"
)

func Init() *echo.Echo {
	e := echo.New()

	//e.GET("/", api.Home)
	//v1 := app.Group("/api/v1")
	e.GET("/users", api.GetUsers)
	e.POST("/tcUser", api.AddTcUsers)
	return e
}
