// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// CommunityUnion is the golang structure for table community_union.
type CommunityUnion struct {
	Id               uint `json:"id"               orm:"id"                 description:""`
	UserId           uint `json:"userId"           orm:"user_id"            description:"用户id"`
	CommunityMajorId uint `json:"communityMajorId" orm:"community_major_id" description:"小区主id"`
	CommunityMinorId uint `json:"communityMinorId" orm:"community_minor_id" description:"小区次id"`
}
