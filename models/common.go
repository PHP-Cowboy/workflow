package models

//审批备份数据
type ApprovalBackupsData struct {
	Base
	Type           string `gorm:"type:varchar(50);not null;comment:类型"`
	Identification string `gorm:"type:varchar(50);not null;default:'';comment:标识"`
	Data           string `gorm:"type:MEDIUMTEXT;not null;comment:备份数据"`
	Status         int    `gorm:"type:tinyint;default:2;comment:审批状态:0:拒绝,1:通过,2:待处理,3:撤销"`
	Saas
}

//审批模块表头
type ApprovalModuleHeader struct {
	Base
	Type           string `gorm:"type:varchar(50);index:idx;not null;comment:类型"`
	Identification string `gorm:"type:varchar(50);index:idx;not null;default:'';comment:标识"`
	Data           string `gorm:"type:TEXT;not null;comment:表头数据"`
	Version        int    `gorm:"type:tinyint unsigned;index:idx;not null;default:1;comment:版本"`
	Saas
}

// 审批消息表
//type ApprovalProcessMessage struct {
//	Base
//	SaasId                string `gorm:"type:char(8);not null;comment:组织SaasID"`
//	ApprovalProcessInfoId int    `gorm:"type:int(11) unsigned;not null;comment:工作流审批基础信息表id"`
//	Type                  int    `gorm:"type:tinyint(2) unsigned;not null;default:0;comment:消息类型,0:审批,1:评论"`
//	ApproverId            string `gorm:"type:varchar(255);not null;comment:审核人id"`
//	ApproverName          string `gorm:"type:varchar(255);not null;comment:审核人名称"`
//	Content               string `gorm:"type:varchar(255);comment:消息内容"`
//	IsRead                bool   `gorm:"not null;default:0;comment:是否已读:0:否,1:是"`
//}
