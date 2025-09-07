package model

import "gorm.io/gorm"

func Paginate(pageIndex, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pageIndex <= 0 {
			pageIndex = 1
		}
		if pageSize <= 0 {
			pageSize = 10
		}
		if pageSize >= 100 {
			pageSize = 100
		}
		offset := (pageIndex - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
