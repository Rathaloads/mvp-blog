package services

import (
	"fmt"
	"mb-server/common/db"
	"mb-server/common/utils/token"
	"mb-server/model"
	"time"
)

const TB_AdminUsers = "admin_users"

func AdminUserLogin(email string, password string) (string, error) {
	var model model.AdminUser
	c, err := db.MysqlGet(TB_AdminUsers, "email", email, &model)
	if err != nil {
		return "", err
	}
	if c == 0 {
		return "", fmt.Errorf("user no exist")
	}
	tokenStr, cliam, err := token.CreateJwt(email)
	model.LastLoginToken = model.LoginToken
	model.LoginToken = tokenStr
	model.LoginTime = time.Now().Unix()
	model.ExpirationTime = cliam.ExpiresAt.Unix()

	if err != nil {
		return "", err
	}

	err = db.MysqlDB.Where("email=?", email).Select("login_token", "last_login_token", "expiration_time", "login_time").Updates(model).Error
	if err != nil {
		return "", err
	}
	expTime := cliam.ExpiresAt.Time
	delayTime := expTime.Sub(time.Now())

	db.RedisSetEx("token", email, tokenStr, delayTime)
	return tokenStr, nil
}
func AdminUserRegister(email string, password string) error {
	var model model.AdminUser
	c, err := db.MysqlGet(TB_AdminUsers, "email", email, &model)
	if err != nil {
		return err
	}
	if c != 0 {
		return fmt.Errorf("email has exisit!")
	}
	model.Email = email
	model.Password = password

	err = db.MysqlCreate(TB_AdminUsers, &model)
	if err != nil {
		return err
	}
	return nil
}
