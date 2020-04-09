package entity

type AdminUser struct {
	ID             int64  `json:"-"`
	UUID           string `json:"uuid"`
	UserName       string `json:"username"`
	AuthKey        string `json:"auth_key"`
	PassHash       string `json:"-"`
	ResetPassToken string `json:"-"`
	Email          string `json:"email"`
	Mobile         string `json:"mobile"`
	Nickname       string `json:"nickname"`
	AvatarUrl      string `json:"avatar_url"`
	Status         int    `json:"status"`
	CreateAt       int64  `json:"create_at"`
	UpdateAt       int64  `json:"update_at"`
	DeleteAt       int64  `json:"delete_at"`
}
