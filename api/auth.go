package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"ybmecc/db"
	"ybmecc/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func Signin(c echo.Context) error {
	s, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}
	b := model.Signin{}
	if err := json.Unmarshal(s, &b); err != nil {
		return err
	}
	fmt.Println(b)
	u := model.TcUser{}
	fmt.Println(u)
	db := db.DbManager()
	err = db.Where(&model.TcUser{EnNm: b.Id, TcPwd: b.Pwd}).First(&u).Error
	if err != nil {
		return c.JSON(http.StatusOK, model.ExcuseError{MsgTxt: "로그인에 실패 했습니다.", Success: false})
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = b.Id
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.JSON(http.StatusOK, model.ExcuseError{MsgTxt: "토큰생성에 실패했습니다. 다시 시도해주세요.", Success: false})
	}
	return c.JSON(http.StatusOK, model.SigninExcuseSuccess{Token: t, Success: true})
}
