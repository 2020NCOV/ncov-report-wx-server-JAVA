package service

import (
	"Miniprogram-server-Golang/model"
	"Miniprogram-server-Golang/serializer"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

// SaveDailyInfoService 管理每日上传信息服务
type SaveDailyInfoService struct {
	IsReturnSchool            int    `form:"is_return_school" json:"is_return_school"`
	ReturnTime                string `form:"return_time" json:"return_time"`
	ReturnDormNum             string `form:"return_dorm_num" json:"return_dorm_num"`
	ReturnTrafficInfo         string `form:"return_traffic_info" json:"return_traffic_info"`
	CurrentHealthValue        int    `form:"current_health_value" json:"current_health_value"`
	CurrentContagionRiskValue int    `form:"current_contagion_risk_value" json:"current_contagion_risk_value"`
	ReturnDistrictValue       int    `form:"return_district_value" json:"return_district_value"`
	CurrentDistrictValue      int    `form:"current_district_value" json:"current_district_value"`
	CurrentTemperature        int    `form:"current_temperature" json:"current_temperature"`
	PsyStatus                 int    `form:"psy_status" json:"psy_status"`
	PsyDemand                 int    `form:"psy_demand" json:"psy_demand"`
	PsyKnowledge              int    `form:"psy_knowledge" json:"psy_knowledge"`
	Remarks                   string `form:"remarks" json:"remarks"`
	PlanCompanyDate           string `form:"plan_company_date" json:"plan_company_date"`
	Uid                       int    `form:"uid" json:"uid"`
	Token                     string `form:"token" json:"token"`
	TemplateCode              string `form:"template_code" json:"template_code"`
}

// isRegistered 判断用户是否存在
func (service *SaveDailyInfoService) SaveDailyInfo(c *gin.Context) serializer.Response {
	if !model.CheckToken(strconv.Itoa(service.Uid), service.Token) {
		return serializer.ParamErr("token验证错误", nil)
	}

	//判断是否重复提交

	//err:= model.DB2.QueryRow("select userID from report_record_company where userID =?， time = ?",service.Uid,time.Now().Format("2006-01-02"))
	//if err==nil {
	//   return serializer.ParamErr("今日您已提交，请勿重复提交", nil)
	//}

	var count int
	if model.DB2.QueryRow("select count(userID) from report_record_company where userID =?， time = ?", service.Uid, time.Now().Format("2006-01-02")).Scan(&count); count > 0 {
		return serializer.ParamErr("今日您已提交，请勿重复提交", nil)
	}
	//保存信息
	var time = time.Now().Format("2006-01-02")
	queryStr := "insert into report_record_company(is_return_school,current_health_value,current_contagion_risk_value,return_district_value," +
		"current_district_value,current_temperature,remarks,psy_status,psy_demand,psy_knowledge,plan_company_date,return_dorm_num,return_time," +
		"return_traffic_info,userID,time)" + "values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	if err := model.DB2.QueryRow(queryStr, service.IsReturnSchool, service.CurrentHealthValue, service.CurrentContagionRiskValue,
		service.ReturnDistrictValue, service.CurrentDistrictValue, service.CurrentTemperature, service.Remarks,
		service.PsyStatus, service.PsyDemand, service.PsyKnowledge, service.PlanCompanyDate, service.ReturnDormNum, service.ReturnTime,
		service.ReturnTrafficInfo, service.Uid, time); err != nil {
		return serializer.ParamErr("上传失败", nil)
	}
	return serializer.BuildSuccessSave()
}

// Gorm版

////isRegistered 判断用户是否存在
//func (service *SaveDailyInfoService) SaveDailyInfo(c *gin.Context) serializer.Response {
//	if !model.CheckToken(strconv.Itoa(service.Uid), service.Token) {
//		return serializer.ParamErr("token验证错误", nil)
//	}
//	dailyInfo := model.Record{
//		IsReturnSchool:            service.IsReturnSchool,
//		CurrentHealthValue:        service.CurrentHealthValue,
//		CurrentContagionRiskValue: service.CurrentContagionRiskValue,
//		ReturnDistrictValue:       service.ReturnDistrictValue,
//		CurrentDistrictValue:      service.CurrentDistrictValue,
//		CurrentTemperature:        service.CurrentTemperature,
//		Remarks:                   service.Remarks,
//		PsyStatus:                 service.PsyStatus,
//		PsyDemand:                 service.PsyDemand,
//		PsyKnowledge:              service.PsyKnowledge,
//		PlanCompanyDate:           service.PlanCompanyDate,
//		ReturnDormNum:             service.ReturnDormNum,
//		ReturnTime:                service.ReturnTime,
//		ReturnTrafficInfo:         service.ReturnTrafficInfo,
//		Uid:                       service.Uid,
//		SaveDate:                  time.Now().Format("2006-01-02"),
//	}
//	判断该用户这天是否已经提交过
//	count := 0
//	if model.DB.Model(&model.DailyInfo{}).Where("uid = ? and save_date = ?", service.Uid, time.Now().Format("2006-01-02")).Count(&count); count > 0 {
//		return serializer.ParamErr("今日您已提交，请勿重复提交", nil)
//	}
//	记录用户当日信息
//	if err := model.DB.Create(&dailyInfo).Error; err != nil {
//		return serializer.ParamErr("上传失败", err)
//	}
//	return serializer.BuildSuccessSave()
//}
