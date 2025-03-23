// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// TaskTemplate is the golang structure for table task_template.
type TaskTemplate struct {
	Type uint   `json:"type" orm:"type" description:"类型"`
	Name string `json:"name" orm:"name" description:"默认项目名"`
}
