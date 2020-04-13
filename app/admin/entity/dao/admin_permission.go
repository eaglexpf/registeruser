package dao

type DaoAdminPermission struct {
	ParentID     int64  `json:"parent_type" sql:"parent_type"`
	ParentType   string `json:"parent_id" sql:"parent_id"`
	ChildrenID   int64  `json:"children_id" sql:"children_id"`
	ChildrenType string `json:"children_type" sql:"children_type"`
}
