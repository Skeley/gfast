// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// TaskStep is the golang structure for table task_step.
type TaskStep struct {
	Id      uint   `json:"id"      orm:"id"       description:"步骤id"`
	StageId uint   `json:"stageId" orm:"stage_id" description:"stage id"`
	Name    string `json:"name"    orm:"name"     description:"步骤名"`
	Comment string `json:"comment" orm:"comment"  description:"描述"`
	Order   uint   `json:"order"   orm:"order"    description:"排序位置"`
}
