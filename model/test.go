package model

import "gorm.io/gorm"

/*
此model层主要用来存放数据组字段结构体，表和数据库中表名一一对应，resp结尾的未返回的数据，后期可以统一封装一个统一结构体返回，不然的话比较麻烦。
*/
type User struct { //测试表
	gorm.Model        //此直接调用gorm数据结构体
	Name       string `json:"name" gorm:"name"`
}

//type UserResp struct {
//	Code int    `json:"code,omitempty"`
//	Mess string `json:"mess,omitempty"`
//	Data *User  `json:"data,omitempty"`
//}
