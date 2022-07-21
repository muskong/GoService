package logic

import "github.com/muskong/GoWechat/entity"

type Page struct {
	entity.Page
}

func Stories(page *Page) (result entity.Result, err error) {

	list, count, err := entity.Story.StoryList(page.Offset(), page.Limit)

	result.Data = list
	result.Pagination.Limit = page.Limit
	result.Pagination.Page = page.Page
	result.Pagination.Total = count

	return
}

func Story(storyId int) (result interface{}, err error) {

	result, err = entity.Story.StoryDetail(storyId)

	return
}
