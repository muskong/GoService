package logic

import (
	"errors"

	"github.com/muskong/GoPkg/jwt"
	"github.com/muskong/GoPkg/we"
	"github.com/muskong/GoPkg/zaplog"
	"github.com/muskong/GoWechat/app/user/entity"
	"golang.org/x/crypto/bcrypt"
)

type (
	LoginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	JwtData struct {
		Token         string
		Name          string
		Avatar        string
		NickName      string
		AccountAmount float64
	}
)

func LoginVerify(data LoginData) (jwtData JwtData, err error) {

	user, err := entity.UserModel.GetUserName(data.Username)
	if user.Id <= 0 || err != nil {
		err = errors.New("用户或密码出错1")
		return
	}

	is := checkHashPassword(data.Password, user.Password)
	if !is {
		err = errors.New("用户或密码出错2")
		return
	}

	jwtData.Name = user.Name
	jwtData.Avatar = user.Avatar
	jwtData.NickName = user.NickName
	jwtData.AccountAmount = user.AccountAmount
	jwtData.Token = jwt.Jwt.GenerateToken(user.Id)

	return
}

func LoginWeChat(code string) (jwtData JwtData, err error) {

	openid, err := we.We.CodeToOpenid(code)
	if err != nil {
		zaplog.Sugar.Error(err)
		return
	}
	user, err := entity.UserModel.GetOpenid(openid)
	if err != nil {
		zaplog.Sugar.Error(err)
		return
	}

	if user.Id <= 0 {
		err = errors.New("微信用户登录失败")
		return
	}

	jwtData.Name = user.Name
	jwtData.Avatar = user.Avatar
	jwtData.NickName = user.NickName
	jwtData.AccountAmount = user.AccountAmount
	jwtData.Token = jwt.Jwt.GenerateToken(user.Id)

	return
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}

func checkHashPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
