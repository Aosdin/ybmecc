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
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://127.0.0.1:8081", "http://localhost:8081"},
		//AllowHeaders:     []string{"authorization", "Content-Type"},
		AllowCredentials: true,
		AllowMethods:     []string{echo.OPTIONS, echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		AllowHeaders:     []string{echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))
	r := e.Group("/api")
	config := middleware.JWTConfig{
		Claims:     &model.JwtCustomClaims{},
		SigningKey: []byte("secret"),
		ErrorHandlerWithContext: func(err error, c echo.Context) error {
			return c.JSON(http.StatusOK, model.ExcuseError{MsgTxt: "토큰생성에 실패했습니다.", Success: false})
		},
	}

	r.Use(middleware.JWTWithConfig(config))
	// Users
	r.GET("/tcSeniorUser", api.GetSeniorTcList)
	r.GET("/tcUser", api.GetTcUsers)
	r.POST("/tcUser", api.AddTcUsers)
	r.PUT("/tcUser", api.PutTcUsers)
	r.DELETE("/tcUser", api.DelTcUsers)
	r.GET("/sdUser", api.GetSdUsers)
	r.POST("/sdUser", api.AddSdUsers)
	r.PUT("/sdUser", api.PutSdUsers)
	r.DELETE("/sdUser", api.DelSdUsers)

	// Auth
	e.POST("/api/signin", api.Signin)
	return e
}
