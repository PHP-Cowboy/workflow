package main

import (
	"workflow/global"
	"workflow/initialize"
	"workflow/models"
	"workflow/models/approval_process"
)

func main() {

	initialize.InitConfig()
	initialize.InitDB()

	db := global.DB

	_ = db.Set(models.TableOptions, models.GetOptions("工作流任务表")).AutoMigrate(&approval_process.TaskInfo{})

	_ = db.Set(models.TableOptions, models.GetOptions("工作流任务发起人表")).AutoMigrate(&approval_process.TaskInitiator{})

	_ = db.Set(models.TableOptions, models.GetOptions("工作流审批流程任务表")).AutoMigrate(&approval_process.TaskApprovalProcess{})

	_ = db.Set(models.TableOptions, models.GetOptions("工作流抄送任务表")).AutoMigrate(&approval_process.TaskCc{})

	_ = db.Set(models.TableOptions, models.GetOptions("审批列表权限")).AutoMigrate(&approval_process.ApprovalListAuth{})

	_ = db.Set(models.TableOptions, models.GetOptions("工作流审批基础信息表")).AutoMigrate(&approval_process.ApprovalProcessInfo{})

	_ = db.Set(models.TableOptions, models.GetOptions("工作流审批数据表")).AutoMigrate(&approval_process.ApprovalProcessData{})

	_ = db.Set(models.TableOptions, models.GetOptions("审批备份数据")).AutoMigrate(&approval_process.ApprovalBackupsData{})

	_ = db.Set(models.TableOptions, models.GetOptions("审批历史记录表")).AutoMigrate(&approval_process.ApprovalProcessHistory{})
}
