// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Community is the golang structure for table community.
type Community struct {
	MajorId       uint   `json:"majorId"       orm:"major_id"       description:"主id"`
	MinorId       uint   `json:"minorId"       orm:"minor_id"       description:"次id"`
	CommunityName string `json:"communityName" orm:"community_name" description:"小区名"`
}
