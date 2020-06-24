package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"ybmecc/db"
	"ybmecc/model"

	"github.com/labstack/echo"
)

func GetUsers(c echo.Context) error {
	db := db.DbManager()
	users := []model.User{}
	db.Find(&users)
	// spew.Dump(json.Marshal(users))
	// return c.JSON(http.StatusOK, users)
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
	b.UpYmd = time.Now()
	b.GrYmd = time.Now()
	fmt.Println(b)
	db := db.DbManager()
	db.NewRecord(b)
	db.Create(&b)
	return c.JSON(http.StatusOK, b)
}
