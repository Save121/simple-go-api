package entity

type Movie struct {
	ID            string  `db:"id"`
	Name          string  `db:"name"`
	Description   string  `db:"description"`
	Price         float32 `db:"price"`
	Created_by    string  `db:"created_by"`
	Creation_date string  `db:"creation_date"`
}
