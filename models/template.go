package models

// 工作流基础信息模板表
type BaseInfoTemplate struct {
	Base
	Title         string `gorm:"type:varchar(50);not null;comment:工作流标题"`
	GroupId       int    `gorm:"type:tinyint;not null;comment:所属分组id"`
	GroupName     string `gorm:"type:varchar(50);not null;comment:所属分组名称"`
	Content       string `gorm:"type:varchar(200);default:'';comment:表单说明"`
	InitiatorType int    `gorm:"type:tinyint(1) unsigned;not null;comment:发起人类型:0:个人,1:组织"`
	InitiatorId   int    `gorm:"type:int(11) unsigned;not null;comment:发起人id/组织id"`
	InitiatorName string `gorm:"type:varchar(20);not null;comment:发起人名/组织名称"`
	OwnerType     int    `gorm:"type:tinyint(1) unsigned;not null;comment:所属人类型:0:个人,1:组织"`
	OwnerId       int    `gorm:"type:int(11) unsigned;not null;comment:所属人id/所属组织id"`
	OwnerName     string `gorm:"type:varchar(20);not null;comment:所属人名/所属组织名称"`
	SaasId        string `gorm:"type:char(8);not null;comment:组织SaasID"`
	IsUsed        bool   `gorm:"default:0;comment:是否已使用:0:否,1:是"`
	Operator
}

// 工作流审批流程模板表
type ApprovalProcessTemplate struct {
	Base
	BaseInfoId        int    `gorm:"type:int(11) unsigned;not null;comment:模板id"`
	Title             string `gorm:"type:varchar(50);not null;default:审核人;comment:流程标题"`
	ApproverType      int    `gorm:"type:tinyint;not null;comment:审批人员类型"`
	ApproverTypeParam string `gorm:"type:varchar(255);default:'';comment:审批人员类型参数"`
	ApprovalMethod    int    `gorm:"type:tinyint;default:10;comment:多人审批时审批方式:10:依次审批,20:会签（须所有审批人同意）,30:或签（一名审批人同意或拒绝即可）"`
}

// 工作流抄送模板表
type CcPeopleTemplate struct {
	Base
	Title        string `gorm:"type:varchar(50);not null;default:审核人;comment:抄送标题"`
	BaseInfoId   int    `gorm:"type:int(11) unsigned;not null;comment:模板id"`
	CcPeopleId   string `gorm:"type:varchar(255);comment:抄送人id"`
	CcPeopleName string `gorm:"type:varchar(255);comment:抄送人昵称"`
}
