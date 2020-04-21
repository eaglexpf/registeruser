package request

type RequestRegisterJob struct {
	URI         string `form:"uri" json:"uri" xml:"uri" binding:"required"`
	Method      string `form:"method" json:"method" xml:"method" binding:"required"`
	Data        string `form:"data" json:"data" xml:"data" binding:"required"`
	Delay       int64  `form:"delay" json:"delay" xml:"delay" binding:"required"`
	Bomb        bool   `form:"bomb" json:"bomb" xml:"bomb"`
	SuccessData string `form:"success_data" json:"success_data" xml:"success_data"`
}
