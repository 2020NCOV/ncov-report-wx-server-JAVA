package serializer

import "Miniprogram-server-Golang/model"

// Usertoken 用户token
type Status struct {
	ErrCode int `json:"errcode"`
	Uid     string `json:"uid"`
	Token   string `json:"token"`
}

// student 用户序列化器
type Student struct {
	Uid          string `json:"uid"`
	Name         string `json:"name"`
	PhoneNum     string `json:"phone_num"`
	UserId       string `json:"userid"`
	Corpname     string `json:"corpname"`
	TypeCorpname string `json:"type_corpname"`
	TypeUsername string `json:"type_username"`
}

// Corp 表单模板序列化器
type Corp struct {
	Corpname     string `json:"corpname"`
	TypeCorpname string `json:"type_corpname"`
	TypeUsername string `json:"type_username"`
	TemplateCode string `json:"template_code"`
}

// isRegistered 用户序列化器
type IsRegistered struct {
	IsRegistered int `json:"is_registered"`
}

type CheckUser struct {
	IsExist int    `json:"is_exist"`
	UserId  string `json:"userid"`
	Corpid  string `json:"corpid"`
}

// BuildUserInfo 序列化
func BuildUserCheck(x int, corpid string, userid string) CheckUser {
	return CheckUser{
		IsExist: x,
		Corpid:  corpid,
		UserId:  userid,
	}
}

// BuildUserInfo 序列化
func BuildUserInfo(user model.Student) Student {
	return Student{
		Uid:          user.Uid,
		Name:         user.Name,
		PhoneNum:     user.PhoneNum,
		UserId:       user.UserId,
		Corpname:     user.Corpid,
		TypeCorpname: "组织编号",
		TypeUsername: "学号",
	}
}

// BuildCorp 序列化status
func BuildCorp(corp model.Corp) Corp {
	return Corp{
		Corpname:     corp.Corpid,
		TypeCorpname: "组织编号",
		TypeUsername: "学号",
		TemplateCode: "default",
	}
}

// BuildStatus 序列化status
func BuildStatus(info model.Code) Status {
	return Status{
		ErrCode: info.ErrCode,
		Uid:     info.Uid,
		Token:   info.Token,
	}
}

// BuildstatusResponse 序列化status响应，返回token、uid等信息
func BuildStatusResponse(info model.Code) Response {
	return Response{
		Data: BuildStatus(info),
	}
}

// BuildCorpResponse 序列化status响应
func BuildCorpResponse(corp model.Corp) Response {
	return Response{
		Data: BuildCorp(corp),
	}
}

// BuildIsRegisteredResponse 序列化用户注册响应
func BuildIsRegisteredResponse(x int) Response {
	return Response{
		Data: IsRegistered{IsRegistered: x},
	}
}

// BuildUserCheckResponse 序列化验证用户是否响应
func BuildUserCheckResponse(x int, corpid string, userid string) Response {
	return Response{
		Data: BuildUserCheck(x, corpid, userid),
	}
}

// BuildUserInfoResponse 序列化用户信息响应
func BuildUserInfoResponse(user model.Student) Response {
	return Response{
		Data: BuildUserInfo(user),
	}
}
