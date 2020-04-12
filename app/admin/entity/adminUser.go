// 项目数据定义
package entity

// adminUser后台用户类型
type AdminUser struct {
	ID                 int64  `json:"-" sql:"id"`
	UUID               string `json:"uuid" sql:"uuid"`
	UserName           string `json:"username" sql:"username"`
	PasswordHash       string `json:"-" sql:"password_hash"`
	PasswordResetToken string `json:"-" sql:"password_reset_token"`
	Email              string `json:"email" sql:"email"`
	Nickname           string `json:"nickname" sql:"nickname"`
	AvatarUrl          string `json:"avatar_url" sql:"avatar_url"`
	Status             int    `json:"status" sql:"status"`
	CreateAt           int64  `json:"create_at" sql:"create_at"`
	UpdateAt           int64  `json:"update_at" sql:"update_at"`
	DeleteAt           int64  `json:"delete_at" sql:"delete_at"`
}
