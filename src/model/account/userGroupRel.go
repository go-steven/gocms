package account

type UserGroupRel struct {
	Id      uint64
	UserId  uint64
	GroupId uint64
	IsDel   uint8 `orm:"default(0)"`
}
