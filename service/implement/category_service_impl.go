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

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) service.CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 DB,
		Validate:           validate}
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request create.CategoryCreateRequest) response.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		CategoryName: request.CategoryName,
	}

	category = service.CategoryRepository.Save(ctx, tx, category)

	return model.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request update.CategoryUpdateRequest) response.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, request.CategoryId)
	helper.PanicIfError(err)

	category.CategoryName = request.CategoryName

	category = service.CategoryRepository.Update(ctx, tx, category)

	return model.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)

	service.CategoryRepository.Delete(ctx, tx, category)
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) response.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return model.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []response.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)
	return model.ToCategoryResponses(categories)
}
