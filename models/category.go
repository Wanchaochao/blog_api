package models
type Category struct {
	Id   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
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

