// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// TaskType is the golang structure for table task_type.
type TaskType struct {
	Type uint   `json:"type" orm:"type" description:"类型"`
	Name string `json:"name" orm:"name" description:"类型名"`
}
