//gorm的扩展
package db

import (
	"gorm.io/gorm"
)

//扩展的结构体
type ExtendDB struct {
	*gorm.DB
}

type PageInfo struct {
	CurrPage  int
	PageSize  int
	DataCount int64
}

/*
	扩展分页方法 必须指定 Model
	后期尝试修改为插件
	pageInfo :分页信息
	dest :为查询出来的数据
*/
func (extendDB *ExtendDB) Page(pageInfo *PageInfo, dest interface{}) *ExtendDB {

	//默认值处理
	if pageInfo.CurrPage == 0 {
		pageInfo.CurrPage = 1
	}
	if pageInfo.PageSize == 0 {
		pageInfo.PageSize = 10
	}
	if pageInfo.PageSize == 0 {
		pageInfo.PageSize = 10
	}

	//针对当前页的偏移量
	offset := (pageInfo.CurrPage - 1) * pageInfo.PageSize

	//总数据量
	if pageInfo.DataCount == 0 {
		//查询数量
		db := extendDB.Count(&(pageInfo.DataCount))

		if db.Error != nil {
			extendDB.DB = db
			return extendDB
		}

		if pageInfo.DataCount == 0 {
			extendDB.Error = gorm.ErrRecordNotFound
			return extendDB
		}
	} else {
		//如果偏移量大于总数据量 那说明没有更多数据
		if int64(offset) > pageInfo.DataCount {
			extendDB.Error = gorm.ErrRecordNotFound
			return extendDB
		}
	}

	extendDB.DB = extendDB.Offset(offset).Limit(pageInfo.PageSize).Find(dest)
	return extendDB
}
