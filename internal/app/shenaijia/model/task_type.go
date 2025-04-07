package model

type TaskType struct {
	Id       uint   `json:"id"       orm:"id"       description:""`
	Name     string `json:"name"     orm:"name"     description:""`
	Standard bool   `json:"standard" orm:"standard" description:"是否为标准类型"`
}
