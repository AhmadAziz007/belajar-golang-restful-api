package implement

import (
	"azizdev/helper"
	"azizdev/model/domain"
	"azizdev/repository"
	"context"
	"database/sql"
	"errors"
)

type CategoryRepositoryImpl struct{}

func NewCategoryRepository() repository.CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "insert into category(category_name) values (?)"
	result, err := tx.ExecContext(ctx, SQL, category.CategoryName)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.CategoryId = int(id)
	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "update category set category_name = ? where category_id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.CategoryName, category.CategoryId)
	helper.PanicIfError(err)

	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	SQL := "delete from category where category_id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.CategoryId)
	helper.PanicIfError(err)
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	SQL := "select category_id, category_name from category where category_id = ?"
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.PanicIfError(err)
	defer rows.Close()

	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.CategoryId, &category.CategoryName)
		helper.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("category is not found")
	}
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := "select category_id, category_name from category"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.CategoryId, &category.CategoryName)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}
	return categories
}
