package account

type GroupRoleRel struct {
	Id      uint64 `json:"id"`
	GroupId uint64 `json:"group_id"`
	RoleId  uint64 `json:"role_id"`
	IsDel   uint8  `json:"is_del", orm:"default(0)"`
}
