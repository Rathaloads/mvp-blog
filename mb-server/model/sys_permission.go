package model

import (
	"time"
)

const TableNameSysPermission = "sys_permission"

type SysPermission struct {
	Id             int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	ParentId       int64     `gorm:"column:parent_id;default:0" json:"parent_id"`                     // 父权限id,没有则为0
	PermissionName string    `gorm:"column:permission_name;size:255;not null" json:"permission_name"` // 权限名称
	PermissionKey  string    `gorm:"column:permission_key;size:255;not null" json:"permission_key"`   // 权限符号
	PermissionType int8      `gorm:"column:permission_type;default:0" json:"permission_type"`         // 权限类型(1:目录,2:菜单,3:按钮,4:接口)
	Path           string    `gorm:"column:path;size:255;default:null" json:"path"`                   // 控制路径(菜单或api)
	Icon           string    `gorm:"column:icon;size:255;default:null" json:"icon"`                   // 图标
	Sort           int32     `gorm:"column:sort;default:null" json:"sort"`                            // 排序
	CreateAt       time.Time `gorm:"column:create_at;default:null" json:"create_at"`
	UpdateAt       time.Time `gorm:"column:update_at;default:null" json:"update_at"`
}

func (l *SysPermission) TableName() string {
	return TableNameSysPermission
}
