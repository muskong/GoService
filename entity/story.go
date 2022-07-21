package entity

import (
	"database/sql"

	gormDb "github.com/muskong/GoPkg/gorm"
	"github.com/muskong/GoPkg/zaplog"
	"gorm.io/gorm"
)

type (
	story   struct{}
	Stories struct {
		gorm.Model
		ID            int
		Title         string
		Code          string
		Content       string
		State         string
		Rate          int
		Price         float64
		CreatedAt     string
		UpdatedAt     sql.NullTime
		DeletedAt     sql.NullTime
		StoryAttitude []StoryAttitude `gorm:"foreignkey:StoryId;references:ID"`
	}
)

var Story = &story{}

func (m *story) StoryList(page, limit int) (list []*Stories, count int64, err error) {
	db := gormDb.ClientNew().Model(Stories{})

	err = db.Count(&count).Order("id desc").Limit(limit).Offset(page).Find(&list).Error
	if err != nil {
		zaplog.Sugar.Error("[storiesModel]get data error:", err)
	}
	return
}

func (m *story) StoryDetail(storyId int) (*Stories, error) {
	db := gormDb.ClientNew().Model(&Stories{})

	var story Stories
	err := db.Where("id = ?", storyId).Preload("StoryAttitude").First(&story).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return &story, err
}
