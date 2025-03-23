// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// UserInfo is the golang structure for table user_info.
type UserInfo struct {
	Id               uint   `json:"id"               orm:"id"                 description:"用户id"`
	Type             string `json:"type"             orm:"type"               description:"用户类型"`
	CommunityMajorId uint   `json:"communityMajorId" orm:"community_major_id" description:"小区主id"`
	CommunityMinorId uint   `json:"communityMinorId" orm:"community_minor_id" description:"小区次id"`
	Inviter          uint   `json:"inviter"          orm:"inviter"            description:"邀请人id"`
}
