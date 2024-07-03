package update

type MenuUpdateRequest struct {
	MenuId     int     `validate:"required"`
	CategoryId int     `validate:"required"`
	MenuName   string  `validate:"required,min=1,max=100" json:"menuName"`
	Price      float64 `validate:"required" json:"price"`
}
