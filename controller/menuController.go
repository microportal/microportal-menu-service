package controller

import (
	"github.com/gorilla/mux"
	"microportal-menu-service/models"
	"microportal-menu-service/repository"
	"microportal-menu-service/utils"
	"net/http"
	"strconv"
)

type MenuController struct{}

var (
	menus []models.Menu
	mRepo repository.MenuRepository
)

func init() {
	mRepo = repository.MenuRepository{}
}

func (c MenuController) FindMenuByModuleID(writer http.ResponseWriter, request *http.Request) {
	var menu models.Menu
	menus = []models.Menu{}
	params := mux.Vars(request)
	moduleId, _ := strconv.Atoi(params["moduleId"])
	data, err := mRepo.FindMenuByModuleID(menu, menus, moduleId)
	utils.SendResult(writer, data, err)
}
