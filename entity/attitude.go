package entity

import (
	"database/sql"

	gormDb "github.com/muskong/GoPkg/gorm"
	"github.com/muskong/GoPkg/zaplog"
	"gorm.io/gorm"
)

type (
	attitude      struct{}
	StoryAttitude struct {
		gorm.Model
		ID        int
		Pid       int
		StoryId   int
		Title     string
		State     string
		CreatedAt string
		UpdatedAt sql.NullTime
		DeletedAt sql.NullTime
	}
)

var Attitude = &attitude{}

func (m *attitude) StateAllow() string {
	return "allow" //允许 allow
}
func (m *attitude) StateDeny() string {
	return "deny" //禁止 deny
}

func (m *attitude) AttitudeList(page, limit int) (list []*StoryAttitude, count int64, err error) {
	db := gormDb.ClientNew().Model(StoryAttitude{})

	err = db.Where("state=?", m.StateAllow()).Count(&count).Order("id desc").Limit(limit).Offset(page).Find(&list).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}

func (m *attitude) AttitudeListByStoryId(storyId int) (list []*StoryAttitude, count int64, err error) {
	db := gormDb.ClientNew().Model(StoryAttitude{})

	err = db.Where("state=? AND story_id=?", m.StateAllow(), storyId).Count(&count).Order("id desc").Find(&list).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}

func (m *attitude) AttitudeDetail(attitudeId int) (*StoryAttitude, error) {
	db := gormDb.ClientNew().Model(Stories{})

	var attitude StoryAttitude
	err := db.Where("state=? AND id = ?", m.StateAllow(), attitudeId).First(&attitude).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return &attitude, err
}
