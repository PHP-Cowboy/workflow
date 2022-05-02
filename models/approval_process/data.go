package approval_process

import "workflow/models"

//审批备份数据
type ApprovalBackupsData struct {
	models.Base
	Type           string `gorm:"type:varchar(50);not null;comment:类型"`
	Identification string `gorm:"type:varchar(50);not null;default:'';comment:标识"`
	Title          string `gorm:"type:varchar(32);not null;comment:任务名称"`
	Data           string `gorm:"type:MEDIUMTEXT;not null;comment:备份数据"`
	Status         int    `gorm:"type:tinyint;default:2;comment:审批状态:0:拒绝,1:通过,2:待处理,3:撤销"`
	models.Creator
}

// 审批历史记录表
type ApprovalProcessHistory struct {
	models.Base
	SaasId                string `gorm:"type:char(8);not null;comment:组织SaasID"`
	ApprovalProcessDataId int    `gorm:"type:int(11) unsigned;not null;comment:工作流任务数据表id"`
	ApproverId            int    `gorm:"type:int(11) unsigned;not null;comment:审批人id"`
	ApproverName          string `gorm:"type:varchar(255);not null;comment:审批人名称"`
	Type                  int    `gorm:"type:tinyint(2) unsigned;not null;default:0;comment:审批操作类型,0:拒绝,1:通过,4:重置"`
	ResetNum              int    `gorm:"type:tinyint;not null;default:-1;comment:重置到的number编号"`
	ApprovalProcessInfoId int    `gorm:"type:int(11) unsigned;not null;comment:工作流审批基础信息表id"`
}
