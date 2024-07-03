package repository

import (
	"azizdev/model/domain"
	"context"
	"database/sql"
)

type MenuRepository interface {
	Save(ctx context.Context, tx *sql.Tx, menu domain.Menu) domain.Menu
	Update(ctx context.Context, tx *sql.Tx, menu domain.Menu) domain.Menu
	Delete(ctx context.Context, tx *sql.Tx, menu domain.Menu)
	FindById(ctx context.Context, tx *sql.Tx, menuId int) (domain.Menu, error)
	FindByName(ctx context.Context, tx *sql.Tx, menuName string, categoryName string) (domain.Menu, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Menu, error)
}
