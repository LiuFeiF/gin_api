package service

import (
	"context"
	"fmt"
	"gin_api/dao"
	"gin_api/pkg/res"
)

type TestService struct { //入参函数
	Id int64 `json:"id" form:"id"` //用于form-data接收或者json接收
}

func (service *TestService) TestInfo(ctx context.Context) res.Response {
	TestDao := dao.NewTestDao()
	data, err := TestDao.TestInfo(service.Id)
	if err != nil {
		return res.Response{
			Status: 400,
			Msg:    "查询数据失败",
			Data:   nil,
		}
	}
	fmt.Println(data)
	return res.Response{
		Status: 200,
		Msg:    "成功",
		Data: res.Responses{
			ItemList: data,
		},
	}
}
