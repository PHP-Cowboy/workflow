package initialize

import (
	"github.com/gofiber/fiber/v2"
	"workflow/router"
)

func InitRoute(app *fiber.App) {
	g := app.Group("/api")

	//审批
	approval := g.Group("/approval")
	router.InitApproval(approval)

	//任务
	task := g.Group("/task")
	router.InitTask(task)

	//公共
	common := g.Group("/common")
	router.InitCommon(common)
}
