// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Community is the golang structure for table community.
type Community struct {
	Id            uint   `json:"id"            orm:"id"             description:""`
	Pid           uint   `json:"pid"           orm:"pid"            description:""`
	CommunityName string `json:"communityName" orm:"community_name" description:"小区名"`
}
