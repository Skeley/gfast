// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// TaskTemplet is the golang structure for table task_templet.
type TaskTemplet struct {
	Id   uint   `json:"id"   orm:"id"   description:""`
	Name string `json:"name" orm:"name" description:""`
	Type uint   `json:"type" orm:"type" description:""`
}
