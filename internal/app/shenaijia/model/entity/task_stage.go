// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// TaskStage is the golang structure for table task_stage.
type TaskStage struct {
	Id        uint64 `json:"id"        orm:"id"         description:"stage id"`
	TaskId    int    `json:"taskId"    orm:"task_id"    description:""`
	TempletId uint64 `json:"templetId" orm:"templet_id" description:"所属任务模板"`
	Name      string `json:"name"      orm:"name"       description:"阶段名"`
	Icon      string `json:"icon"      orm:"icon"       description:"图标"`
	Comment   string `json:"comment"   orm:"comment"    description:"描述"`
	Position  uint   `json:"position"  orm:"position"   description:"任务流位置"`
}
