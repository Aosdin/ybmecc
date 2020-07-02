package model

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"time"
)

func (TcUser) TableName() string {
	return "TC_USERS"
}

type ExcuseError struct {
	//Msg error `json:"msg"`
	MsgTxt  string `json:"msg"`
	Success bool   `json:"success"`
}
type ExcuseSuccess struct {
	MsgTxt  string `json:"msg"`
	Success bool   `json:"success"`
}
type SigninExcuseSuccess struct {
	Token   string `json:"token"`
	Success bool   `json:"success"`
}
type User struct {
	gorm.Model
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string
}
type JwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}
type TcUser struct {
	Idx   int       `json:"idx" gorm:"column:IDX"`
	KoNm  string    `json:"koNm" gorm:"column:KO_NM"`
	EnNm  string    `json:"enNm" gorm:"column:EN_NM"`
	TcPwd string    `json:"tcPwd" gorm:"column:TC_PWD"`
	TcLv  int       `json:"tcLv" gorm:"column:TC_LV"`
	TcSt  int       `json:"tcSt" gorm:"column:TC_ST"`
	BdYmd time.Time `json:"dbYmd" gorm:"column:BD_YMD"`
	UpYmd time.Time `json:"upYmd" gorm:"column:UP_YMD"`
	GrYmd time.Time `json:"grYmd" gorm:"column:GR_YMD"`
}
type SdUser struct {
	Idx        int       `json:"idx" gorm:"column:IDX"`
	KoNm       string    `json:"koNm" gorm:"column:KO_NM"`
	EnNm       string    `json:"enNm" gorm:"column:EN_NM"`
	PrPhoneNum string    `json:"prPhoneNum" gorm:"column:PR_PHONE_NUM"`
	SdPhoneNum string    `json:"sdPhoneNum" gorm:"column:SD_PHONE_NUM"`
	BdYmd      time.Time `json:"dbYmd" gorm:"column:BD_YMD"`
	UpYmd      time.Time `json:"upYmd" gorm:"column:UP_YMD"`
	GrYmd      time.Time `json:"grYmd" gorm:"column:GR_YMD"`
}

type Signin struct {
	Id  string `json:"id"`
	Pwd string `json:"pwd"`
}
