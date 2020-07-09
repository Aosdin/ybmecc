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

func GetSeniorTcList(c echo.Context) error {
	db := db.DbManager()
	users := []model.TcUser{}
	db.Where("TC_ST = ? AND TC_LV = ?", 1, 1).Find(&users)
	return c.JSON(http.StatusOK, users)
}

func GetTcUsers(c echo.Context) error {
	db := db.DbManager()
	users := []model.TcUser{}
	db.Where("TC_ST = ?", 1).Find(&users)
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
	t := model.TcUser{}
	if db.Where("TC_ID = ?", b.TcId).First(&t).RecordNotFound() {
		err = db.Debug().Create(&b).Error
		if err != nil {
			return c.JSON(http.StatusOK, model.ExcuseError{MsgTxt: "선생님등록에 실패했습니다.", Success: false})
		} else {
			return c.JSON(http.StatusOK, model.ExcuseSuccess{Success: true})
		}
	} else {
		return c.JSON(http.StatusOK, model.ExcuseError{MsgTxt: "동일한 ID가 사용중입니다.", Success: false})
	}
}
func PutTcUsers(c echo.Context) error {
	s, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}
	b := model.TcUser{}
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

func GetSdUsers(c echo.Context) error {
	db := db.DbManager()
	users := []model.SdUser{}
	db.Where("SD_ST = ?", 1).Find(&users)
	return c.JSON(http.StatusOK, model.ExcuseDataSuccess{Data: users, Success: true})
}

func AddSdUsers(c echo.Context) error {
	s, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}
	b := model.SdUser{}
	if err := json.Unmarshal(s, &b); err != nil {
		return err
	}
	b.UpYmd = time.Now().UTC()
	b.GrYmd = time.Now().UTC()
	db := db.DbManager()
	sd := model.SdUser{}
	if db.Where("PR_PHONE_NUM = ? AND KO_NM = ?", b.PrPhoneNum, b.KoNm).First(&sd).RecordNotFound() {
		err = db.Debug().Create(&b).Error
		if err != nil {
			return c.JSON(http.StatusOK, model.ExcuseError{MsgTxt: "학생등록에 실패했습니다.", Success: false})
		} else {
			return c.JSON(http.StatusOK, model.ExcuseSuccess{Success: true})
		}
	} else {
		return c.JSON(http.StatusOK, model.ExcuseError{MsgTxt: "동일한 데이터가 있습니다.", Success: false})
	}
}

func PutSdUsers(c echo.Context) error {
	s, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}
	b := model.SdUser{}
	if err := json.Unmarshal(s, &b); err != nil {
		return err
	}
	b.UpYmd = time.Now().UTC()
	db := db.DbManager()
	db.Model(&b).Where("Idx = ?", b.Idx).Update(b)
	return c.JSON(http.StatusOK, model.ExcuseSuccess{Success: true})
}

func DelSdUsers(c echo.Context) error {
	s, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}
	b := model.SdUser{}
	if err := json.Unmarshal(s, &b); err != nil {
		return err
	}
	db := db.DbManager()
	db.Model(&b).Where("Idx = ?", b.Idx).Update(map[string]interface{}{"SdSt": 0, "UpYmd": time.Now().UTC()})
	return c.JSON(http.StatusOK, model.ExcuseSuccess{Success: true})
}
