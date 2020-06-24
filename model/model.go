package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

func (TcUser) TableName() string {
	return "TC_USERS"
}

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string
}
type TcUser struct {
	KoNm  string    `json:"koNm" gorm:"column:KO_NM"`
	EnNm  string    `json:"enNm" gorm:"column:EN_NM"`
	TcLv  int       `json:"tcLv" gorm:"column:TC_LV"`
	TcSt  int       `json:"tcSt" gorm:"column:TC_ST"`
	UpYmd time.Time `json:"upYmd" gorm:"column:UP_YMD"`
	GrYmd time.Time `json:"grYmd" gorm:"column:GR_YMD"`
}
