package logic

import "github.com/muskong/GoWechat/app/user/entity"

func UserInfo(userId int64) (user interface{}, err error) {

	user, err = entity.UserModel.GetUser(userId)

	return
}
