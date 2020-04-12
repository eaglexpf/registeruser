package request

type RequestRegisterAdminUser struct {
	Username  string `form:"username" json:"username" xml:"username" binding:"required"`
	Password  string `form:"password" json:"password" xml:"password" binding:"required"`
	RepeatPwd string `form:"repeat_pwd" json:"repeat_pwd" xml:"repeat_pwd" binding:"required"`
	Email     string `form:"email" json:"email" xml:"email" binding:"email" binding:"required"`
	Nickname  string `form:"nickname" json:"nickname" xml:"nickname" binding:"required"`
	AvatarUrl string `form:"avatar_url" json:"avatar_url" xml:"avatar_url"`
}

type RequestAdminUserLogin struct {
	Username string `form:"username" json:"username" xml:"username" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

type RequestAdminUserUpdateInfo struct {
	UUID      string `form:"password" json:"uuid" xml:"uuid"`
	Nickname  string `form:"nickname" json:"nickname" xml:"nickname" binding:"required"`
	AvatarUrl string `form:"avatar_url" json:"avatar_url" xml:"avatar_url" binding:"required"`
}

type RequestAdminUserResetPwd struct {
	UUID      string `form:"password" json:"uuid" xml:"uuid"`
	Password  string `form:"password" json:"password" xml:"password" binding:"required"`
	PwdRepeat string `form:"pwd_repeat" json:"pwd_repeat" xml:"pwd_repeat" binding:"required"`
	OldPwd    string `form:"old_pwd" json:"old_pwd" xml:"old_pwd" binding:"required"`
}