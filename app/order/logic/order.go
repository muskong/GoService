package logic

import "github.com/muskong/GoWechat/entity"

type (
	Page struct {
		entity.Page
	}
	LikeOrder struct {
		StoryAttitudeId int
		StoryId         int
		UserId          int
		State           string
	}
)

func StoryLike(lo LikeOrder) (result interface{}, err error) {

	lo.State = entity.NewOrder().StateUnpaid()
	result, err = entity.NewOrder().OrderCreate(lo.StoryAttitudeId, lo.StoryId, lo.UserId, lo.State)

	return
}

func UserStories(userId int, page *Page) (result entity.Result, err error) {

	list, count, err := entity.NewOrder().OrderList(userId, page.Offset(), page.Limit)

	result.Data = list
	result.Pagination.Limit = page.Limit
	result.Pagination.Page = page.Page
	result.Pagination.Total = count

	return
}
