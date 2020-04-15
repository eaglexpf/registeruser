package request

type RequestServiceSearch struct {
	Name     string `form:"name" json:"name"`
	Alias    string `form:"alias" json:"alias"`
	Status   int64  `form:"status" json:"status"`
	Page     int64  `form:"page" json:"page"`
	PageSize int64  `form:"page_size" json:"page_size"`
}

type RequestServiceRegister struct {
	Name        string `form:"name" json:"name" binding:"required"`
	Alias       string `form:"alias" json:"alias" binding:"required"`
	Description string `form:"description" json:"description"`
	Status      int64  `form:"status" json:"status" binding:"required"`
	ExpireAt    int64  `form:"expire_at" json:"expire_at"`
}

type RequestServiceUpdate struct {
	Name        string `form:"name" json:"name" binding:"required"`
	Alias       string `form:"alias" json:"alias" binding:"required"`
	Description string `form:"description" json:"description"`
	Status      int64  `form:"status" json:"status" binding:"required"`
	ExpireAt    int64  `form:"expire_at" json:"expire_at"`
}
