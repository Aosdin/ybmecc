package route

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"ybmecc/api"
	"ybmecc/model"
)

func Init() *echo.Echo {
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	r := e.Group("/api")
	config := middleware.JWTConfig{
		Claims:     &model.JwtCustomClaims{},
		SigningKey: []byte("secret"),
		ErrorHandlerWithContext: func(err error, c echo.Context) error {
			return c.JSON(http.StatusOK, model.ExcuseError{MsgTxt: "로그인생성에 실패했습니다.", Success: false})
		},
	}

	r.Use(middleware.JWTWithConfig(config))

	//e.GET("/", api.Home)
	//v1 := app.Group("/api/v1")
	r.GET("/tcUser", api.GetTcUsers)
	r.POST("/tcUser", api.AddTcUsers)
	r.PUT("/tcUser", api.PutTcUsers)
	r.DELETE("/tcUser", api.DelTcUsers)
	e.POST("/signin", api.Signin)
	return e
}
