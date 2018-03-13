package account

import (
	"time"
)

type UserGroup struct {
	Id        uint64    `json:"id"`
	GroupName string    `json:"group_name"`
	Remarks   string    `json:"remarks"`
	Created   time.Time `json:"created", orm:"auto_now_add;type(datetime)"`
	Updated   time.Time `json:"updated", orm:"auto_now;type(datetime)"`
	IsDel     uint8     `json:"is_del",orm:"default(0)"`
}

type UserGroupCheck struct {
	Id        uint64    `json:"id"`
	GroupName string    `json:"group_name"`
	Remarks   string    `json:"remarks"`
	Created   time.Time `json:"created", orm:"auto_now_add;type(datetime)"`
	Updated   time.Time `json:"updated", orm:"auto_now;type(datetime)"`
	IsDel     uint8     `json:"is_del", orm:"default(0)"`
	Check     bool      `json:"check"`
}
