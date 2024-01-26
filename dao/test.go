package dao

import (
	"fmt"
	"gin_api/internal"
	"gin_api/model"
)

type TestDao struct {
} //声明一个简单的结构体用于作为数据查询的入口

func NewTestDao() TestDao {
	return TestDao{} //对外查询的函数
}
func (dao *TestDao) TestInfo(id int64) (User *model.User, err error) {
	//fist函数用于gorm返回只返回一条数据但是如果查不到数据会报错，如果是获取列表的话建议用find函数
	err = internal.DB.Model(&model.User{}).Where("id=?", id).
		First(&User).Error
	if err != nil {
		fmt.Println(err)
	}
	return
}
