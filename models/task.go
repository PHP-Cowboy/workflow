package models

//工作流任务表
type TaskInfo struct {
	Base
	//BaseInfoId     int    `gorm:"type:int(11) unsigned;not null;comment:模板id"`
	Title          string `gorm:"type:varchar(32);not null;comment:任务名称"`
	Type           string `gorm:"type:varchar(64);not null;comment:类型"`
	Identification string `gorm:"type:varchar(64);not null;default:'';comment:标识"`
	DepartmentId   int    `gorm:"type:int(11) unsigned;not null;comment:部门id"`
	Department     string `gorm:"type:varchar(32);not null;default:'';comment:部门名称"`
	//Status         int    `gorm:"type:tinyint;default:2;comment:审批状态:0:拒绝,1:通过,2:待处理,3:撤销"`
	Creator
	Saas
}

//工作流任务发起人表
type TaskInitiator struct {
	Base
	Type       string `gorm:"type:varchar(255);not null;comment:类型(1:部门,2:用户)"`
	ParamId    string `gorm:"type:varchar(255);not null;comment:参数id(类型为部门时表示部门id，类型为用户时表示用户id)"`
	ParamName  string `gorm:"type:varchar(255);not null;comment:名称(类型为部门时表示部门名称，类型为用户时表示用户昵称)"`
	TaskInfoId int    `gorm:"type:int(11) unsigned;not null;comment:工作流任务id"`
	Saas
}

//工作流审批流程任务表
type TaskApprovalProcess struct {
	Base
	TaskInfoId int `gorm:"type:int(11) unsigned;not null;comment:工作流任务id"`
	//BaseInfoId     int    `gorm:"type:int(11) unsigned;not null;comment:模板id"`
	Type           string `gorm:"type:varchar(255);not null;comment:类型(1:部门,2:用户)"`
	ApproverId     string `gorm:"type:varchar(255);not null;comment:审核人id(类型为部门时为部门id,为用户时是用户id,多个用,分割)"`
	ApproverName   string `gorm:"type:varchar(255);not null;comment:审核人名称(类型为部门时为部门名称,为用户时是用户昵称,多个用,分割)"`
	Number         int    `gorm:"type:tinyint unsigned;not null;comment:编号"`
	ApprovalMethod int    `gorm:"type:tinyint;default:10;comment:多人审批时审批方式:10:依次审批,20:会签（须所有审批人同意）,30:或签（一名审批人同意或拒绝即可）"`
	Status         int    `gorm:"type:tinyint;default:2;comment:审批状态:0:拒绝,1:通过,2:待处理,3:撤销"` //后期用不上可以删掉
	Saas
}

// 工作流抄送任务表
type TaskCc struct {
	Base
	TaskInfoId   int    `gorm:"type:int(11) unsigned;not null;comment:工作流任务id"`
	BaseInfoId   int    `gorm:"type:int(11) unsigned;not null;comment:模板id"`
	Type         string `gorm:"type:varchar(255);not null;comment:类型(1:部门,2:用户)"`
	CcPeopleId   string `gorm:"type:varchar(255);comment:抄送人(类型为部门时为部门id,为用户时是用户id,多个用,分割)"`
	CcPeopleName string `gorm:"type:varchar(255);comment:抄送人昵称(类型为部门时为部门名称,为用户时是用户昵称,多个用,分割)"`
	Saas
}
