package model

type TaskTemplet struct {
	Id       uint   `json:"id"   orm:"id"   description:""`
	Name     string `json:"name" orm:"name" description:""`
	Type     uint   `json:"type" orm:"type" description:""`
	Standard bool   `json:"standard" orm:"standard" description:""`
}
