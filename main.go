package main

import (
	"azizdev/app"
	"azizdev/controller/implement"
	"azizdev/helper"
	"azizdev/middleware"
	implement2 "azizdev/repository/implement"
	implement3 "azizdev/service/implement"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := implement2.NewCategoryRepository()
	categoryService := implement3.NewCategoryService(categoryRepository, db, validate)
	categoryController := implement.NewCategoryController(categoryService)

	menuRepository := implement2.NewMenuRepository()
	menuService := implement3.NewMenuService(menuRepository, db, validate)
	menuController := implement.NewMenuController(menuService)

	router := app.NewRouter(categoryController, menuController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
