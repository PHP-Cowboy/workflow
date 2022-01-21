package main

import (
	"workflow/global"
	"workflow/initialize"
	"workflow/models"
)

func main() {

	initialize.InitConfig()
	initialize.InitDB()

	db := global.DB

	_ = db.Set(models.TableOptions, models.GetOptions("工作流基础信息模板表")).AutoMigrate(&models.BaseInfoTemplate{})

	_ = db.Set(models.TableOptions, models.GetOptions("工作流审批流程模板表")).AutoMigrate(&models.ApprovalProcessTemplate{})

	_ = db.Set(models.TableOptions, models.GetOptions("工作流抄送模板表")).AutoMigrate(&models.CcPeopleTemplate{})

	_ = db.Set(models.TableOptions, models.GetOptions("工作流任务信息表")).AutoMigrate(&models.TaskInfo{})

	_ = db.Set(models.TableOptions, models.GetOptions("工作流任务发起人表")).AutoMigrate(&models.TaskInitiator{})

	_ = db.Set(models.TableOptions, models.GetOptions("工作流审批流程任务表")).AutoMigrate(&models.TaskApprovalProcess{})

	_ = db.Set(models.TableOptions, models.GetOptions("工作流抄送任务表")).AutoMigrate(&models.TaskCc{})

	_ = db.Set(models.TableOptions, models.GetOptions("工作流审批基础信息表")).AutoMigrate(&models.ApprovalProcessInfo{})

	_ = db.Set(models.TableOptions, models.GetOptions("工作流审批数据状态表")).AutoMigrate(&models.ApprovalProcessStatus{})

	_ = db.Set(models.TableOptions, models.GetOptions("工作流审批数据表")).AutoMigrate(&models.ApprovalProcessData{})

	_ = db.Set(models.TableOptions, models.GetOptions("审批备份数据")).AutoMigrate(&models.ApprovalBackupsData{})

	_ = db.Set(models.TableOptions, models.GetOptions("审批模块表头")).AutoMigrate(&models.ApprovalModuleHeader{})

	_ = db.Set(models.TableOptions, models.GetOptions("审批历史记录表")).AutoMigrate(&models.ApprovalProcessHistory{})

	_ = db.Set(models.TableOptions, models.GetOptions("评论")).AutoMigrate(&models.Comment{})

}
