package implement

import (
	"azizdev/helper"
	"azizdev/model/web"
	"azizdev/model/web/create"
	"azizdev/model/web/update"
	"azizdev/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type MenuControllerImpl struct {
	MenuService service.MenuService
}

func NewMenuController(menuService service.MenuService) *MenuControllerImpl {
	return &MenuControllerImpl{
		MenuService: menuService}
}

func (controller *MenuControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	menuCreateRequest := create.MenuCreateRequest{}
	helper.ReadFromRequestBody(request, &menuCreateRequest)

	menuResponse := controller.MenuService.Create(request.Context(), menuCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   menuResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *MenuControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	menuUpdateRequest := update.MenuUpdateRequest{}
	helper.ReadFromRequestBody(request, &menuUpdateRequest)

	menuId := params.ByName("menuId")
	id, err := strconv.Atoi(menuId)
	helper.PanicIfError(err)

	menuUpdateRequest.MenuId = id

	menuResponse := controller.MenuService.Update(request.Context(), menuUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   menuResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *MenuControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	menuId := params.ByName("menuId")
	id, err := strconv.Atoi(menuId)
	helper.PanicIfError(err)

	controller.MenuService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *MenuControllerImpl) FindbyName(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	menuName := params.ByName("name")
	categoryName := params.ByName("categoryName")

	menuResponse := controller.MenuService.FindByName(request.Context(), menuName, categoryName)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   menuResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *MenuControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	menuResponse := controller.MenuService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   menuResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
