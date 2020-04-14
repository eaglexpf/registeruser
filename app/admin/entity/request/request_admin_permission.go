package request

type RequestPermissionRegister struct {
	ParentID     int64  `form:"parent_id" json:"parent_id" binding:"required"`
	ParentType   string `form:"parent_type" json:"parent_type" binding:"required"`
	ChildrenID   int64  `form:"children_id" json:"children_id" binding:"required"`
	ChildrenType string `form:"children_type" json:"children_type" binding:"required"`
}

type RequestPermissionSearch struct {
	ParentID     int64  `form:"parent_id" json:"parent_id"`
	ParentType   string `form:"parent_type" json:"parent_type"`
	ChildrenID   int64  `form:"children_id" json:"children_id"`
	ChildrenType string `form:"children_type" json:"children_type"`
	Page         int64  `form:"page" json:"page"`
	PageSize     int64  `form:"page_size" json:"page_size"`
}

type RequestPermissionDelete struct {
	ParentID     int64  `form:"parent_id" json:"parent_id"`
	ParentType   string `form:"parent_type" json:"parent_type"`
	ChildrenID   int64  `form:"children_id" json:"children_id"`
	ChildrenType string `form:"children_type" json:"children_type"`
}
