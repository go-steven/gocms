package account

import (
	"time"
)

type Shop struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`

	Created time.Time `json:"created", orm:"auto_now_add;type(datetime)"`
	Updated time.Time `json:"updated", orm:"auto_now;type(datetime)"` // 更新时间
	IsDel   int8      `json:"is_del", orm:"default(0)"`
}
