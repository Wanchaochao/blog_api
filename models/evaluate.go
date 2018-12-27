package models

import "time"

type Evaluate struct {
	Id          int       `db:"id" json:"id"`
	ForeignKey  int       `db:"foreign_key" json:"foreign_key"`
	Type        int       `db:"type" json:"type"`
	Praise      int       `db:"praise" json:"praise"`
	Ip          string    `db:"ip" json:"ip"`
	Description string    `db:"description" json:"description"`
	CreatedAt   time.Time `db:"created_at" json:"created_at" time_format:"2006-01-02 15:04:05"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at" time_format:"2006-01-02 15:04:05"`
}

func (c *Evaluate) DbName() string {
	return "default"
}

func (c *Evaluate) TableName() string {
	return "evaluate"
}

func (c *Evaluate) PK() string {
	return "id"
}
