// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// TaskStage is the golang structure for table task_stage.
type TaskStage struct {
	Id      uint   `json:"id"      orm:"id"      description:"stage id"`
	Type    uint   `json:"type"    orm:"type"    description:"所属任务类型"`
	Name    string `json:"name"    orm:"name"    description:"阶段名"`
	Icon    string `json:"icon"    orm:"icon"    description:"图标"`
	Comment string `json:"comment" orm:"comment" description:"描述"`
	Order   uint   `json:"order"   orm:"order"   description:"任务流位置"`
}
