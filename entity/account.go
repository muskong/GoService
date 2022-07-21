package entity

import (
	"github.com/muskong/GoPkg/gorm"
	"github.com/muskong/GoPkg/zaplog"
)

type (
	account     struct{}
	UserAccount struct {
		Id        int
		UserId    int
		Before    float64
		Change    float64
		After     float64
		Remark    string
		Table     string
		TableId   int
		CreatedAt string
	}
)

var Account = &account{}

func (m *account) AccountList(page, limit int) (list []*UserAccount, count int64, err error) {
	db := gorm.ClientNew().Model(UserAccount{})

	err = db.Count(&count).Order("id desc").Limit(limit).Offset(page).Find(&list).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}

func (m *account) AccountDetail(accountId int) (*UserAccount, error) {
	db := gorm.ClientNew().Model(UserAccount{})

	var account UserAccount
	err := db.Where("id = ?", accountId).First(&account).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return &account, err
}
