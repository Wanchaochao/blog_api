package models

import "time"

type Comments struct {
	Id        int       `db:"id" json:"id"`
	ArticleId int       `db:"article_id" json:"article_id"`
	Nickname  string    `db:"nickname" json:"nickname"`
	Content   string    `db:"content" json:"content"`
	Pid       int       `db:"p_id" json:"p_id"`
	CreatedAt time.Time `db:"created_at" json:"created_at" time_format:"2006-01-02 15:04:05"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at" time_format:"2006-01-02 15:04:05"`
}

func (c *Comments) DbName() string {
	return "default"
}

func (c *Comments) TableName() string {
	return "comments"
}

func (c *Comments) PK() string {
	return "id"
}

type CommentsInfo struct {
	Comments
	PraiseNum  int `json:"praise_num"`
	AgainstNum int `json:"against_num"`
}
