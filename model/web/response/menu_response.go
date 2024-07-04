package response

type MenuResponse struct {
	MenuId       int     `json:"menu_id"`
	CategoryId   int     `json:"category_id"`
	CategoryName string  `json:"category_name"`
	MenuName     string  `json:"menu_name"`
	Price        float64 `json:"price"`
}
