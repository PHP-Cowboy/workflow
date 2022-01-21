package models

//工作流审批基础信息表
type ApprovalProcessInfo struct {
	Base
	Creator
	TaskInfoId     int    `gorm:"type:int(11) unsigned;not null;comment:工作流任务id"`
	SaasId         string `gorm:"type:char(8);not null;comment:组织SaasID"`
	Type           string `gorm:"type:varchar(64);not null;comment:类型"`
	Identification string `gorm:"type:varchar(64);not null;default:'';comment:标识"`
	Status         int    `gorm:"type:tinyint;default:2;comment:审批状态:0:拒绝,1:通过,2:待处理,3:撤销"`
}

//工作流审批数据状态表
type ApprovalProcessStatus struct {
	Base
	TaskInfoId            int    `gorm:"type:int(11) unsigned;not null;comment:工作流任务id"`
	SaasId                string `gorm:"type:char(8);not null;comment:组织SaasID"`
	TaskApprovalProcessId int    `gorm:"type:int(11) unsigned;not null;comment:工作流审批流程任务表id"`
	Number                int    `gorm:"type:tinyint unsigned;not null;comment:编号"`
	Status                int    `gorm:"type:tinyint;default:2;comment:审批状态:0:拒绝,1:通过,2:待处理,3:撤销"`
	ApprovalProcessInfoId int    `gorm:"type:int(11) unsigned;not null;comment:工作流审批基础信息表id"`
}

// 工作流审批数据表
type ApprovalProcessData struct {
	//工作流审批流程任务表 中的一条可能会对应这里的多条，number中体现
	Base
	TaskInfoId            int    `gorm:"type:int(11) unsigned;not null;comment:工作流任务id"`
	ApproverId            int    `gorm:"type:int(11) unsigned;not null;comment:审批人id"`
	ApproverName          string `gorm:"type:varchar(255);not null;comment:审批人名称"`
	Number                int    `gorm:"type:tinyint unsigned;not null;comment:编号"`
	ApprovalMethod        int    `gorm:"type:tinyint;default:10;comment:多人审批时审批方式:10:依次审批,20:会签（须所有审批人同意）,30:或签（一名审批人同意或拒绝即可）"`
	Status                int    `gorm:"type:tinyint;default:2;comment:审批状态:0:拒绝,1:通过,2:待处理,3:撤销"`
	PrevStatus            bool   `gorm:"not null;default:0;comment:前一审人（或组织）审批批状态:0:未通过,1:通过"`
	SaasId                string `gorm:"type:char(8);not null;comment:组织SaasID"`
	Type                  string `gorm:"type:varchar(50);not null;comment:类型"`
	Identification        string `gorm:"type:varchar(50);not null;default:'';comment:标识"`
	ApprovalProcessInfoId int    `gorm:"type:int(11) unsigned;not null;comment:工作流审批基础信息表id"`
}

// 审批历史记录表
type ApprovalProcessHistory struct {
	Base
	SaasId                string `gorm:"type:char(8);not null;comment:组织SaasID"`
	ApprovalProcessDataId int    `gorm:"type:int(11) unsigned;not null;comment:工作流任务数据表id"`
	ApproverId            int    `gorm:"type:int(11) unsigned;not null;comment:审批人id"`
	ApproverName          string `gorm:"type:varchar(255);not null;comment:审批人名称"`
	Type                  int    `gorm:"type:tinyint(2) unsigned;not null;default:0;comment:审批操作类型,0:拒绝,1:通过,2:重置"`
	ResetNum              int    `gorm:"type:tinyint;not null;default:-1;comment:重置到的number编号"`
	ApprovalProcessInfoId int    `gorm:"type:int(11) unsigned;not null;comment:工作流审批基础信息表id"`
}

//评论
type Comment struct {
	Base
	Saas
	Creator
	Content               string `gorm:"type:varchar(128);not null;comment:内容"`
	ApprovalProcessDataId int    `gorm:"type:int(11) unsigned;not null;comment:工作流审批数据表id"`
	ApprovalProcessInfoId int    `gorm:"type:int(11) unsigned;not null;comment:工作流审批基础信息表id"`
}
