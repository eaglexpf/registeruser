package dao

type AdminPermission struct {
	ParentID     int64  `json:"parent_id" sql:"parent_id"`
	ParentType   string `json:"parent_type" sql:"parent_type"`
	ChildrenID   int64  `json:"children_id" sql:"children_id"`
	ChildrenType string `json:"children_type" sql:"children_type"`
}
