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

func UserStories(userId int, page *Page) (result Result, err error) {

	list, count, err := entity.Order.OrderList(userId, page.Offset(), page.Limit)

	result.Data = list
	result.Pagination.Limit = page.Limit
	result.Pagination.Page = page.Page
	result.Pagination.Total = count

	return
}
