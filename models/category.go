package models

type Category struct {
	Id          int    `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	Cover       string `db:"cover" json:"cover"`
	Avatar      string `db:"avatar" json:"avatar"`
	Title       string `db:"title" json:"title"`
	Description string `db:"description" json:"description"`
}

func (c *Category) DbName() string {
	return "default"
}

func (c *Category) TableName() string {
	return "category"
}

func (c *Category) PK() string {
	return "id"
}
