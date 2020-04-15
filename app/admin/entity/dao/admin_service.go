package dao

type AdminService struct {
	ID          int64  `json:"id" sql:"id"`
	Name        string `json:"name" sql:"name"`
	Alias       string `json:"alias" sql:"alias"`
	Description string `json:"description" sql:"description"`
	Status      int64  `json:"status" sql:"status"`
	ExpireAt    int64  `json:"expire_at" sql:"expire_at"`
	CreateAt    int64  `json:"create_at" sql:"create_at"`
	UpdateAt    int64  `json:"update_at" sql:"update_at"`
}

type AdminServiceSearch struct {
	Name         string `json:"name" sql:"name"`
	Alias        string `json:"alias" sql:"alias"`
	Status       int64  `json:"status" sql:"status"`
	ExpireStatus int64
}
