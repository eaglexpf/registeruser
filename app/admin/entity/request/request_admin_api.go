package request

type RequestApiRegister struct {
	Method      string `form:"method" json:"method" binding:"required"`
	Path        string `form:"path" json:"path" binding:"required"`
	Description string `form:"description" json:"description"`
}

type RequestApiUpdate struct {
	Method      string `form:"method" json:"method" binding:"required"`
	Path        string `form:"path" json:"path" binding:"required"`
	Description string `form:"description" json:"description"`
}
