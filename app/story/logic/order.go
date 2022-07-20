package logic

import "github.com/muskong/GoWechat/app/story/entity"

type (
	LikeOrder struct {
		StoryAttitudeId int
		StoryId         int
		UserId          int
		State           string
	}
)

func StoryLike(lo LikeOrder) (result interface{}, err error) {

	lo.State = entity.Order.StateUnpaid()
	result, err = entity.Order.OrderCreate(lo.StoryAttitudeId, lo.StoryId, lo.UserId, lo.State)

	return
}
