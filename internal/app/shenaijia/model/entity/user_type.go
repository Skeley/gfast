// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// UserType is the golang structure for table user_type.
type UserType struct {
	Type uint   `json:"type" orm:"type" description:"用户类型"`
	Name string `json:"name" orm:"name" description:"name"`
}
