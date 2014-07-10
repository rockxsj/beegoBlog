package models

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Users struct {
	Id       int
	Username string
	Password string
	Email    string
	Nickname string
	Avatar   string
	Addtime  time.Time
	ModTime  time.Time
}

/**
 * 检查登陆
 */
func (this *Users) CheckLogin(username string, password string) bool {
	err := o.Raw("SELECT * FROM users WHERE username = ? AND password = ?", username, password).QueryRow(&this)

	if err == nil {
		return true
	} else {
		return false
	}
}
