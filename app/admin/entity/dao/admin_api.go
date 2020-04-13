package dao

type DaoAdminApi struct {
	ID       int64  `json:"id" sql:"id"`
	Method   string `json:"method" sql:"method"`
	Path     string `json:"path" sql:"path"`
	CreateAt int64  `json:"create_at" sql:"create_at"`
	UpdateAt int64  `json:"update_at" sql:"update_at"`
}
