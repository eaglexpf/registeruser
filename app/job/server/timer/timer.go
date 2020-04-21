package timer

import (
	"registeruser/app/job/service"
	"time"
)

var jobService service.JobService

func init() {
	jobService = service.NewJobService()
}
func Run() {
	for {
		time.Sleep(1 * time.Second)
		go jobService.ExpireTask()
		//go jobService.SetList()
	}
}
