package service

import (
	"context"
	"gin_api/dao"
	"gin_api/model"
)

type TestService struct { //入参函数
	Id int64 `json:"id" form:"id"` //用于form-data接收或者json接收
}

func (service *TestService) TestInfo(ctx context.Context) *model.UserResp {
	var res model.UserResp
	res.Code = 400
	res.Mess = "数据返回错误"

	TestDao := dao.NewTestDao()
	data, err := TestDao.TestInfo(service.Id)
	if err != nil {
		res.Code = 400
		res.Mess = "查询失败"
		return &res
	}
	res.Code = 200
	res.Mess = "成功"
	res.Data = data
	return &res
}
