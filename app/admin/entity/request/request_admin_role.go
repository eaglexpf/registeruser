package request

type RequestAdminRoleRegister struct {
	Name        string `form:"name" json:"name" binding:"required"`
	Description string `form:"description" json:"description" binding:"required"`
}

type RequestAdminRoleUpdate struct {
	ID          int64  `form:"id" json:"id"`
	Name        string `form:"name" json:"name" binding:"required"`
	Description string `form:"description" json:"description" binding:"required"`
}
