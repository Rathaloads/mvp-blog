package model

const TableNameSysUserRole = "sys_user_role"

type SysUserRole struct {
	Id     int64 `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserId int64 `gorm:"column:user_id;not null" json:"user_id"`
	RoleId int64 `gorm:"column:role_id;not null" json:"role_id"`
}

func (l *SysUserRole) TableName() string {
	return TableNameSysUserRole
}
