package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"registeruser/app/job/entity/request"
	"registeruser/app/job/entity/response"
	"registeruser/app/job/server/task"
	"registeruser/app/job/server/timer"
	"registeruser/app/job/service"
)

var jobService service.JobService

func init() {
	jobService = service.NewJobService()
}

func handler(handler func(c *gin.Context) error) func(c *gin.Context) {
	return func(c *gin.Context) {
		if err := handler(c); err != nil {
			result := response.Response.ErrorResponse(err)
			c.AbortWithStatusJSON(http.StatusOK, result)
		}
	}
}

func RegisterRouter(router *gin.Engine) {
	job := router.Group("/job")
	job.POST("/", handler(createJob))
	job.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "success")
	})
	go timer.Run()
	go task.Run()
	//go sub.Run()
}

/**
 * @apiDefine api_job_101 1.延时服务
 */

/**
 * @api {post} /job/ 1.添加任务
 * @apiDescription 按id查询服务信息
 * @apiGroup api_job_101
 * @apiVersion 1.0.0
 *
 * @apiParam {string} uri 任务的请求地址
 * @apiParam {string} method 任务的请求方式
 * @apiParam {string} data 任务的请求数据
 * @apiParam {string} [success_data] 请求成功后的返回数据
 * @apiParam {int} delay 延时时间
 * @apiParam {bool} bomb 是否是定时任务
 *
 * @apiSuccess {int} code 状态值
 * @apiSuccess {string} msg 状态描述
 * @apiSuccess {object} data 返回数据
 *
 **/
func createJob(c *gin.Context) (err error) {
	var data request.RequestRegisterJob
	if err = c.ShouldBind(&data); err != nil {
		return
	}
	// 执行服务
	ok, err := jobService.CreateJob(c, &data)
	if err != nil {
		return
	}
	if !ok {
		err = errors.New("任务添加失败")
		return
	}
	c.JSON(http.StatusOK, response.Response.Success())
	return
}
