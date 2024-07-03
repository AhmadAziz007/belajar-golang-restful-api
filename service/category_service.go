package service

import (
	"azizdev/model/web/create"
	"azizdev/model/web/response"
	"azizdev/model/web/update"
	"context"
)

type CategoryService interface {
	Create(ctx context.Context, request create.CategoryCreateRequest) response.CategoryResponse
	Update(ctx context.Context, request update.CategoryUpdateRequest) response.CategoryResponse
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) response.CategoryResponse
	FindAll(ctx context.Context) []response.CategoryResponse
}
