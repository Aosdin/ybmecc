package api

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
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

	db := db.DbManager()
	err = db.Where(&model.TcUser{TcId: b.Id, TcPwd: sha(b.Pwd)}).First(&u).Error
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
	if u.TcSt == 0 {
		return c.JSON(http.StatusOK, model.ExcuseError{MsgTxt: "계정이 비활성화 상태입니다. 관리자에게 문의하세요.", Success: false})
	}
	return c.JSON(http.StatusOK, model.SigninExcuseSuccess{Token: t, Data: model.TcUserRespons{Idx: u.Idx, KoNm: u.KoNm, EnNm: u.EnNm, TcLv: u.TcLv, BdYmd: u.BdYmd}, Success: true})
}

func sha(d string) string {
	secret := "mysecret"
	fmt.Printf("Secret: %s Data: %s\n", secret, d)
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(d))
	sha := hex.EncodeToString(h.Sum(nil))
	fmt.Println("Result: " + sha)
	return sha
}
