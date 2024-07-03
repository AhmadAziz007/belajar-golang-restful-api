package implement

import (
	"azizdev/exception"
	"azizdev/helper"
	"azizdev/helper/model"
	"azizdev/model/domain"
	"azizdev/model/web/create"
	"azizdev/model/web/response"
	"azizdev/model/web/update"
	"azizdev/repository"
	"azizdev/service"
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
)

type MenuServiceImpl struct {
	MenuRepository repository.MenuRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewMenuService(menuRepository repository.MenuRepository, DB *sql.DB, validate *validator.Validate) service.MenuService {
	return &MenuServiceImpl{
		MenuRepository: menuRepository,
		DB:             DB,
		Validate:       validate}
}

func (service *MenuServiceImpl) Create(ctx context.Context, request create.MenuCreateRequest) response.MenuResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	menu := domain.Menu{
		CategoryId: request.CategoryId,
		MenuName:   request.MenuName,
		Price:      request.Price,
	}
	savedMenu := service.MenuRepository.Save(ctx, tx, menu)

	return model.ToMenuResponse(savedMenu)
}

func (service MenuServiceImpl) Update(ctx context.Context, request update.MenuUpdateRequest) response.MenuResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	menu, err := service.MenuRepository.FindById(ctx, tx, request.MenuId)
	helper.PanicIfError(err)

	menu.CategoryId = request.CategoryId
	menu.MenuName = request.MenuName
	menu.Price = request.Price

	menu = service.MenuRepository.Update(ctx, tx, menu)

	return model.ToMenuResponse(menu)
}

func (service MenuServiceImpl) Delete(ctx context.Context, menuId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	menu, err := service.MenuRepository.FindById(ctx, tx, menuId)
	helper.PanicIfError(err)
	service.MenuRepository.Delete(ctx, tx, menu)
}

func (service MenuServiceImpl) FindByName(ctx context.Context, menuName string, categorynName string) response.MenuResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	menuList, err := service.MenuRepository.FindByName(ctx, tx, menuName, categorynName)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return model.ToMenuResponse(menuList)
}

func (service MenuServiceImpl) FindAll(ctx context.Context) []response.MenuResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	menuList, err := service.MenuRepository.FindAll(ctx, tx)
	return model.ToMenuResponses(menuList)
}
