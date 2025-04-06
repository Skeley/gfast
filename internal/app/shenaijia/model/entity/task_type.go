// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// TaskType is the golang structure for table task_type.
type TaskType struct {
	Id       uint   `json:"id"       orm:"id"       description:""`
	Name     string `json:"name"     orm:"name"     description:""`
	Standard int    `json:"standard" orm:"standard" description:"是否为标准类型"`
}
