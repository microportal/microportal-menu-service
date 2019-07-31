package repository

import (
	"microportal-menu-service/models"
	"microportal-menu-service/utils"
)

type MenuRepository struct{}

func (m MenuRepository) FindMenuByModuleID(menu models.Menu, menus [] models.Menu, moduleID int) (interface{}, error) {
	rows, err := GetDB().Query("SELECT * FROM menu WHERE moduleId=$1", moduleID)
	if err != nil {
		return []models.Menu{}, err
	}

	for rows.Next() {
		err = rows.Scan(&menu.ID, &menu.ModuleId, &menu.Name, &menu.Order, &menu.ParentId)
		menus = append(menus, menu)
	}

	return utils.ResultData(menus, []models.Menu{}, err)
}

func (m MenuRepository) FindMenuByID(menu models.Menu, id int) (interface{}, error) {
	row := GetDB().QueryRow("SELECT * FROM menu WHERE id=&1", id)
	err := row.Scan(&menu.ID, &menu.ModuleId, &menu.Name, &menu.Order, &menu.ParentId)

	return utils.ResultData(menu, models.Menu{}, err)
}
