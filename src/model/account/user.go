package account

import (
	"time"
)

type User struct {
	Id         uint64    `json:"id"`
	Name       string    `json:"name", orm:"unique"`
	Mail       string    `json:"mail"`
	Realname   string    `json:"realname"`
	Phone      string    `json:"phone"`
	Department string    `json:"department"`
	Passwd     string    `json:"passwd"`
	Created    time.Time `json:"created", orm:"auto_now_add;type(datetime)"`
	Updated    time.Time `json:"updated", orm:"auto_now;type(datetime)"` // 更新时间
	IsDel      uint8     `json:"is_del", orm:"default(0)"`
}
