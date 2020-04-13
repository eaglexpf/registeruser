package dao

type AdminRole struct {
	ID          int64  `json:"id" sql:"id"`
	Name        string `json:"name" sql:"name"`
	Description string `json:"description" sql:"description"`
	CreateAt    int64  `json:"create_at" sql:"create_at"`
	UpdateAt    int64  `json:"update_at" sql:"update_at"`
}
