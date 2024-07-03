package model

import (
	"azizdev/model/domain"
	"azizdev/model/web/response"
)

func ToMenuResponse(menu domain.Menu) response.MenuResponse {
	return response.MenuResponse{
		MenuId:     menu.MenuId,
		CategoryId: menu.CategoryId,
		MenuName:   menu.MenuName,
		Price:      menu.Price,
	}
}

func ToMenuResponses(menus []domain.Menu) []response.MenuResponse {
	var menuResponses []response.MenuResponse
	for _, menuList := range menus {
		menuResponses = append(menuResponses, ToMenuResponse(menuList))
	}
	return menuResponses
}
