package model

type TaskStage struct {
	Id        uint   `json:"id"      orm:"id"      description:"stage id"`
	TempletId uint   `json:"templetId" orm:"templet_id" description:"所属任务模板"`
	Name      string `json:"name"    orm:"name"    description:"阶段名"`
	Icon      string `json:"icon"    orm:"icon"    description:"图标"`
}
