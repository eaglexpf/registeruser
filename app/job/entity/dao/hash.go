package dao

type Hash struct {
	UUID        string `json:"uuid"`
	URI         string `json:"uri"`
	Method      string `json:"method"`
	Data        string `json:"data"`
	Time        int64  `json:"time"`      // 任务第一次执行时间
	LastTime    int64  `json:"last_time"` // 任务下一次执行时间
	Delay       int64  `json:"delay"`     // 延时时间
	Bomb        bool   `json:"bomb"`
	Num         int64  `json:"num"`          // 程序执行次数
	SuccessData string `json:"success_data"` // 成功的内容
}
