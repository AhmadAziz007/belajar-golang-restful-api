package response

type MenuResponse struct {
	MenuId     int     `json:"menu_id"`
	CategoryId int     `json:"category_id"`
	MenuName   string  `json:"menu_name"`
	Price      float64 `json:"price"`
}
