package entity

import (
	"database/sql"

	gormDb "github.com/muskong/GoPkg/gorm"
	"github.com/muskong/GoPkg/zaplog"
	"gorm.io/gorm"
)

type (
	progress      struct{}
	StoryProgress struct {
		gorm.Model
		ID        int
		StoryId   int
		Title     string
		Content   string
		Date      string
		CreatedAt string
		UpdatedAt sql.NullTime
		DeletedAt sql.NullTime
	}
)

var Progress = &progress{}

func (m *progress) ProgressList(page, limit int) (list []*StoryProgress, count int64, err error) {
	db := gormDb.ClientNew().Model(StoryProgress{})

	err = db.Count(&count).Order("id desc").Limit(limit).Offset(page).Find(&list).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}

func (m *progress) ProgressDetail(progressId int) (*StoryProgress, error) {
	db := gormDb.ClientNew().Model(Stories{})

	var progress StoryProgress
	err := db.Where("id = ?", progressId).First(&progress).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return &progress, err
}
