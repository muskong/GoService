package entity

import (
	"database/sql"

	gormDb "github.com/muskong/GoPkg/gorm"
	"github.com/muskong/GoPkg/zaplog"
	"gorm.io/gorm"
)

type (
	order      struct{}
	StoryOrder struct {
		gorm.Model
		ID              int
		StoryAttitudeId int
		StoryId         int
		UserId          int
		State           string
		CreatedAt       sql.NullTime
		UpdatedAt       sql.NullTime
		DeletedAt       sql.NullTime
		Story           Stories `gorm:"foreignkey:ID;references:StoryId"`
	}
)

var Order = &order{}

func (m *order) StateUnpaid() string {
	return "unpaid"
}
func (m *order) StatePaid() string {
	return "paid"
}
func (m *order) StateFinish() string {
	return "finish"
}
func (m *order) StateRefund() string {
	return "refund"
}
func (m *order) StateRefunded() string {
	return "refunded"
}

func (m *order) OrderList(userId, page, limit int) (list []*StoryOrder, count int64, err error) {
	db := gormDb.ClientNew().Model(StoryOrder{})

	err = db.Where("user_id=?", userId).Preload("Story").Count(&count).Order("id desc").Limit(limit).Offset(page).Find(&list).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}

func (m *order) OrderDetail(orderId int) (*StoryOrder, error) {
	db := gormDb.ClientNew().Model(StoryOrder{})

	var order StoryOrder
	err := db.Where("id = ?", orderId).First(&order).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return &order, err
}

func (m *order) OrderCreate(storyAttitudeId, storyId, userId int, state string) (*StoryOrder, error) {
	db := gormDb.ClientNew().Model(StoryOrder{})

	order := StoryOrder{
		StoryAttitudeId: storyAttitudeId,
		StoryId:         storyId,
		UserId:          userId,
		State:           state,
	}
	err := db.Create(&order).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return &order, err
}
