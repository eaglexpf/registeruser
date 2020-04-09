package entity

type RequestAdminInsert struct {
	Username  string `form:"username" json:"username" xml:"username" binding:"username"`
	Password  string `form:"username" json:"password" xml:"password" binding:"username"`
	RepeatPwd string `form:"repeat_pwd" json:"repeat_pwd" xml:"repeat_pwd" binding:"required"`
	Email     string `form:"email" json:"email" xml:"email" binding:"email" binding:"required"`
}
