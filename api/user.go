package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"ybmecc/db"
	"ybmecc/model"

	"github.com/labstack/echo/v4"
)

func GetTcUsers(c echo.Context) error {
	db := db.DbManager()
	users := []model.TcUser{}
	db.Find(&users)
	return c.JSON(http.StatusOK, users)
}
func AddTcUsers(c echo.Context) error {
	s, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}
	b := model.TcUser{}
	if err := json.Unmarshal(s, &b); err != nil {
		return err
	}
	b.UpYmd = time.Now().UTC()
	b.GrYmd = time.Now().UTC()
	db := db.DbManager()
	err = db.Debug().Create(&b).Error
	if err != nil {
		return c.JSON(http.StatusOK, model.ExcuseError{MsgTxt: "선생님등록에 실패했습니다.", Success: false})
	}
	return c.JSON(http.StatusOK, model.ExcuseSuccess{Success: true})
}
func PutTcUsers(c echo.Context) error {
	s, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}
	b := model.TcUser{}
	fmt.Println(b)
	if err := json.Unmarshal(s, &b); err != nil {
		return err
	}
	b.UpYmd = time.Now().UTC()
	db := db.DbManager()
	db.Model(&b).Where("Idx = ?", b.Idx).Update(b)
	return c.JSON(http.StatusOK, model.ExcuseSuccess{Success: true})
}
func DelTcUsers(c echo.Context) error {
	s, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}
	b := model.TcUser{}
	fmt.Println(b)
	if err := json.Unmarshal(s, &b); err != nil {
		return err
	}
	db := db.DbManager()
	db.Model(&b).Where("Idx = ?", b.Idx).Update(model.TcUser{TcSt: 0, UpYmd: time.Now().UTC()})
	return c.JSON(http.StatusOK, model.ExcuseSuccess{Success: true})
}
