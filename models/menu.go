package models

type Menu struct {
	ID       int    `json:"id"`
	ModuleId int    `json:"moduleId"`
	Name     string `json:"name"`
	Order    string `json:"order"`
	ParentId int    `json:"parentId"`
}
