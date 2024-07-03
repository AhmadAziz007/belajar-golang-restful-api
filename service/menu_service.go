package service

import (
	"azizdev/model/web/create"
	"azizdev/model/web/response"
	"azizdev/model/web/update"
	"context"
)

type MenuService interface {
	Create(ctx context.Context, request create.MenuCreateRequest) response.MenuResponse
	Update(ctx context.Context, request update.MenuUpdateRequest) response.MenuResponse
	Delete(ctx context.Context, menuId int)
	FindByName(ctx context.Context, menuName string, categoryName string) response.MenuResponse
	FindAll(ctx context.Context) []response.MenuResponse
}
