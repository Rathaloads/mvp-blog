package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameSysRole = "sys_role"

type SysRole struct {
	Id       int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	RoleName string    `gorm:"column:role_name;size:255;not null" json:"role_name"`     // 角色名称
	RoleKey  string    `gorm:"column:role_key;size:255;not null" json:"role_key"`       // 角色权限字符
	RoleDesc string    `gorm:"column:role_desc;size:255;default:null" json:"role_desc"` // 描述
	Status   int8      `gorm:"column:status;default:null" json:"status"`                // 状态: 0-禁用,1-正常
	CreateAt time.Time `gorm:"column:create_at;default:null" json:"create_at"`
	UpdateAt time.Time `gorm:"column:update_at;default:null" json:"update_at"`
}

func (l *SysRole) TableName() string {
	return TableNameSysRole
}

func GetSysRoleList(db *gorm.DB, pageIndex, pageSize int) (list []SysRole, count int64, err error) {
	result := db.Model(SysRole{}).Count(&count).Scopes(Paginate(pageIndex, pageSize)).Order("create_at DESC").Find(&list)
	return list, count, result.Error
}

func SaveOrUpdateSysRole(db *gorm.DB, roleId int64, rolename, roleKey, roleDesc string, status int8) (*SysRole, error) {
	var result *gorm.DB
	sysRole := &SysRole{
		Id:       roleId,
		RoleName: rolename,
		RoleKey:  roleKey,
		RoleDesc: roleDesc,
	}
	if roleId > 0 {
		result = db.Updates(sysRole)
	} else {
		result = db.Create(sysRole)
	}
	return sysRole, result.Error
}

func DeleteSysRole(db *gorm.DB, ids []int64) (int64, error) {
	result := db.Where("id IN ?", ids).Delete(SysRole{})
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
