package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameSysUsers = "sys_users"

type SysUsers struct {
	Id       int64      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Username string     `gorm:"column:username;size:255;not null" json:"username"`
	Password string     `gorm:"column:password;size:255;not null" json:"password"`
	Email    string     `gorm:"column:email;size:255;not null" json:"email"`
	Status   int8       `gorm:"column:status;default:1" json:"status"`       // 状态: 0-禁用,1-正常
	UserType int8       `gorm:"column:user_type;default:1" json:"user_type"` // 用户类型: (0:超级管理员,1:普通用户)
	CreateAt *time.Time `gorm:"column:create_at;default:null" json:"create_at"`
	UpdateAt *time.Time `gorm:"column:update_at;default:null" json:"update_at"`
}

func (l *SysUsers) TableName() string {
	return TableNameSysUsers
}

func GetSysUserList(db *gorm.DB, pageIndex int, pageSize int, username string, email string, status int) (list []SysUsers, count int64, err error) {
	db = db.Model(SysUsers{})
	if username != "" {
		db = db.Where("username=?", username)
	}
	if email != "" {
		db = db.Where("email=?", username)
	}
	db.Where("status=?", status)
	result := db.Count(&count).Order("create_at DESC, id DESC").Scopes(Paginate(pageIndex, pageSize)).Find(&list)
	return list, count, result.Error
}

func CreateSysUser(db *gorm.DB, username, password, email string, status int8, usertype int8) (*SysUsers, error) {
	sysUsers := &SysUsers{
		Username: username,
		Password: password,
		Email:    email,
		Status:   status,
		UserType: usertype,
	}
	result := db.Create(sysUsers)
	return sysUsers, result.Error
}

func SaveOrUpdateSysUser(db *gorm.DB, id int64, username string, status int8, usertype int8) (*SysUsers, error) {
	sysuser := &SysUsers{
		Id:       id,
		Username: username,
		Status:   status,
		UserType: usertype,
	}
	var result *gorm.DB
	if id > 0 {
		result = db.Updates(sysuser)
	} else {
		result = db.Create(sysuser)
	}
	return sysuser, result.Error
}

func DeleteSysUser(db *gorm.DB, ids []int64) (int64, error) {
	result := db.Where("id IN ?", ids).Delete(SysUsers{})
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
