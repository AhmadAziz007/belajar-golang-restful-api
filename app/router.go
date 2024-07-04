package app

import (
	"azizdev/controller"
	"azizdev/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controller.CategoryController, menuController controller.MenuController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.GET("/api/menu", menuController.FindAll)
	router.POST("/api/menu", menuController.Create)

	router.PanicHandler = exception.ErrorHandler

	return router
}
