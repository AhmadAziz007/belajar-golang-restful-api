package implement

import (
	"azizdev/helper"
	"azizdev/model/domain"
	"azizdev/repository"
	"context"
	"database/sql"
	"errors"
)

type MenuRepositoryImpl struct{}

func NewMenuRepository() repository.MenuRepository {
	return &MenuRepositoryImpl{}
}

func (repository MenuRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, menuId int) (domain.Menu, error) {
	SQL := "select menu_id, category_id, menu_name, price from menu where id=?"
	rows, err := tx.QueryContext(ctx, SQL, menuId)
	helper.PanicIfError(err)
	defer rows.Close()

	menu := domain.Menu{}
	if rows.Next() {
		err := rows.Scan(&menu.MenuId, &menu.CategoryId, &menu.MenuName, &menu.Price)
		helper.PanicIfError(err)
		return menu, nil
	} else {
		return menu, errors.New("Menu not found")
	}
}

func (repository MenuRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, menu domain.Menu) domain.Menu {
	SQL := "INSERT INTO menu (category_id, menu_name, price) VALUES (?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, menu.CategoryId, menu.MenuName, menu.Price)
	helper.PanicIfError(err)

	menuId, err := result.LastInsertId()
	helper.PanicIfError(err)

	menu.MenuId = int(menuId)
	return menu
}

func (repository MenuRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, menu domain.Menu) domain.Menu {
	SQL := "update menu set menu_name=?, price=? where menu_id = ?"
	_, err := tx.QueryContext(ctx, SQL, menu.MenuName, menu.Price, menu.MenuId)
	helper.PanicIfError(err)

	return menu
}

func (repository MenuRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, menu domain.Menu) {
	SQL := "delete from menu where menu_id = ?"
	_, err := tx.ExecContext(ctx, SQL, menu.MenuId)
	helper.PanicIfError(err)
}

func (repository MenuRepositoryImpl) FindByName(ctx context.Context, tx *sql.Tx, menuName string, categoryName string) (domain.Menu, error) {
	SQL := `SELECT a.menu_id, a.category_id, b.category_name, a.menu_name, a.price 
			FROM menu AS a 
			INNER JOIN category AS b ON a.category_id = b.category_id 
			WHERE a.menu_name LIKE ? 
			AND b.category_name LIKE ? 
			ORDER BY a.menu_id, a.category_id, b.category_name, a.menu_name, a.price ASC`

	menuNameLike := "%" + menuName + "%"
	categoryNameLike := "%" + categoryName + "%"

	row := tx.QueryRowContext(ctx, SQL, menuNameLike, categoryNameLike)

	var menu domain.Menu
	if err := row.Scan(&menu.MenuId, &menu.CategoryId, &menu.CategoryName, &menu.MenuName, &menu.Price); err != nil {
		if err == sql.ErrNoRows {
			return menu, errors.New("menu not found")
		}
		return menu, err
	}
	return menu, nil
}

func (repository MenuRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Menu, error) {
	SQL := `select a.menu_id, a.category_id, b.category_name, a.menu_name, a.price 
			from menu as a 
			inner join category as b on a.category_id = b.category_id 
			order by a.menu_id, a.category_id, b.category_name, a.menu_name, a.price asc `
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var menuLists []domain.Menu
	for rows.Next() {
		menu := domain.Menu{}
		err := rows.Scan(&menu.MenuId, &menu.CategoryId, &menu.CategoryName, &menu.MenuName, &menu.Price)
		helper.PanicIfError(err)
		menuLists = append(menuLists, menu)
	}
	// Check for errors from iterating over rows
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return menuLists, nil
}
