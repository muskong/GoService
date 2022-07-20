package entity

import (
	"database/sql"

	"github.com/muskong/GoPkg/gorm"
	"github.com/muskong/GoPkg/zaplog"
)

type (
	userModel struct{}
	User      struct {
		Id            int64
		WechatOpenid  string
		Name          string
		Avatar        string
		NickName      string
		Password      string
		AccountAmount float64
		CreatedAt     sql.NullTime
		UpdatedAt     sql.NullTime
	}
)

var UserModel = &userModel{}

func (m *userModel) TableName() string {
	return "users"
}

func (m *userModel) GetUserName(username string) (*User, error) {
	db := gorm.ClientNew()

	var user User
	err := db.Model(User{}).Where("name=?", username).First(&user).Error
	if err != nil {
		zaplog.Sugar.Error("[userModel]get data error:", err)
	}
	return &user, err
}

func (m *userModel) GetUser(userId int64) (*User, error) {
	db := gorm.ClientNew().Model(User{})

	var user User
	err := db.Where("id=?", userId).First(&user).Error
	if err != nil {
		zaplog.Sugar.Error("[userModel]get data error:", err)
	}
	return &user, err
}

func (m *userModel) GetOpenid(openid string) (*User, error) {
	db := gorm.ClientNew().Model(User{})

	var user User
	err := db.Debug().Where(User{
		WechatOpenid: openid,
	}).FirstOrCreate(&user).Error

	if err != nil {
		zaplog.Sugar.Error("[userModel]get data error:", err)
	}
	return &user, err
}
