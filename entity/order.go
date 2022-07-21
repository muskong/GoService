package entity

import (
	"database/sql"

	gormDb "github.com/muskong/GoPkg/gorm"
	"github.com/muskong/GoPkg/zaplog"
	"gorm.io/gorm"
)

type (
	_order struct{}

	Order struct {
		gorm.Model
		ID              int
		StoryAttitudeId int
		StoryId         int
		UserId          int
		State           string
		CreatedAt       sql.NullTime
		UpdatedAt       sql.NullTime
		DeletedAt       sql.NullTime
		StoryAttitude   StoryAttitude `gorm:"foreignkey:StoryAttitudeId;references:ID"`
		Story           Stories       `gorm:"foreignkey:StoryId;references:ID"`
	}
)

func NewOrder() *_order {
	return &_order{}
}

func (m *_order) StateUnpaid() string {
	return "unpaid"
}
func (m *_order) StatePaid() string {
	return "paid"
}
func (m *_order) StateFinish() string {
	return "finish"
}
func (m *_order) StateRefund() string {
	return "refund"
}
func (m *_order) StateRefunded() string {
	return "refunded"
}

func (m *_order) OrderList(userId, page, limit int) (list []*Order, count int64, err error) {
	db := gormDb.ClientNew().Model(Order{})

	err = db.Where("user_id=?", userId).Preload("StoryAttitude").Preload("Story").Count(&count).Order("id desc").Limit(limit).Offset(page).Find(&list).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}

func (m *_order) OrderDetail(orderId int) (*Order, error) {
	db := gormDb.ClientNew().Model(Order{})

	var order Order
	err := db.Where("id = ?", orderId).Preload("StoryAttitude").Preload("Story").First(&order).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return &order, err
}

func (m *_order) OrderCreate(storyAttitudeId, storyId, userId int, state string) (*Order, error) {
	db := gormDb.ClientNew().Model(Order{})

	order := Order{
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
