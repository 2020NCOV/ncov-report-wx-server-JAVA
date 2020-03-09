package service

import (
	"github.com/gin-gonic/gin"
	"github.com/medivhzhan/weapp/v2"
	"Miniprogram-server-Golang/model"
	"Miniprogram-server-Golang/serializer"
	"os"
)

// UserOpenIdService 获取用户token服务
type UserOpenIdService struct {
	Code string `form:"code" json:"code"`
}

// gecode 用户登录函数，获取openid和sessionkey，作为之后操作的验证
func (service *UserOpenIdService) GetCode(c *gin.Context) serializer.Response {
	res, err := weapp.Login(os.Getenv("APP_ID"), os.Getenv("APP_SECRET"), service.Code)

	if err != nil {
		//处理错误
		return serializer.ParamErr("获取openid失败", err)
	}

	if err := res.GetResponseError(); err != nil {
		//处理小程序传送的错误信息
		return serializer.ParamErr("小程序报错", err)
	}

	info := model.Code{
		ErrCode: res.ErrCode,
		Uid:   res.OpenID,
		Token: res.SessionKey,
		Code:  service.Code,
	}

	//查看数据库中是否已有token信息
	count := 0
	if model.DB.Model(&model.Code{}).Where("uid = ? and token = ?", res.OpenID, res.SessionKey).Count(&count); count == 0 {
		//将之前的数据删除
		model.DB.Where("uid = ?", res.OpenID).Delete(model.Code{})

		// 记录用户本次token信息
		if err := model.DB.Create(&info).Error; err != nil {
			return serializer.ParamErr("token记录失败", err)
		}
	}

	return serializer.BuildStatusResponse(info)
}
