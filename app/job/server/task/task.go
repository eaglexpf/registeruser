package task

import "registeruser/app/job/service"

var jobService service.JobService

func init() {
	jobService = service.NewJobService()
}
func Run() {
	go jobService.RunTask()
}
