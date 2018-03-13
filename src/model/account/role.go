package account

type Role struct {
	Id      uint64 `json:"id"`
	Pid     uint64 `json:"pid"`
	Name    string `json:"name"`
	Module  string `json:"module", orm:"size(50)"`
	Action  string `json:"action", orm:"size(50)"`
	RoleUrl string `json:"role_url"`
	IsMenu  uint8  `json:"is_menu"`
	Remarks string `json:"remarks"`
}

type RoleTree struct {
	Id      uint64 `json:"id"`
	Pid     uint64 `json:"pId"` // ztree: 简单模式的 JSON 数据需要使用 id / pId 表示节点的父子包含关系
	Name    string `json:"name"`
	Open    bool   `json:"open"`
	Checked bool   `json:"checked"`
	RoleUrl string `json:"role_url"`
	Click   string `json:"click"`
}
