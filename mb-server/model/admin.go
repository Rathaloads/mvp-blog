package model

type AdminUser struct {
	Email          string `gorm:"column:email"`
	Password       string `gorm:"column:password"`
	LoginToken     string `gorm:"column:login_token"`
	LastLoginToken string `gorm:"columnlast_login_token"`
	ExpirationTime string `gorm:"expiration_time"`
	LoginTime      string `gorm:"login_time"`
}
