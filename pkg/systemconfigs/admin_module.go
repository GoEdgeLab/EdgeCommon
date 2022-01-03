package systemconfigs

// AdminModule 管理用户模块权限
type AdminModule struct {
	Code     string   `json:"code"`     // 模块代号
	AllowAll bool     `json:"allowAll"` // 允许所有的动作
	Actions  []string `json:"actions"`  // 只允许的动作
}
