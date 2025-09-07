package model

const TableNameSysRolePermission = "sys_role_permission"

type SysRolePermission struct {
	Id           int64 `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	RoleId       int64 `gorm:"column:role_id;not null" json:"role_id"`
	PermissionId int64 `gorm:"column:permission_id;not null" json:"permission_id"`
}

func (l *SysRolePermission) TableName() string {
	return TableNameSysRolePermission
}
