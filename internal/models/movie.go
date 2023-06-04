package models

type Movie struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Price         float32 `json:"price"`
	Creation_date string  `json:"creation_date"`
}
