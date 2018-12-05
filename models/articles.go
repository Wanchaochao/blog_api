package models

import (
	"github.com/ilibs/gosql"
	"github.com/verystar/golib/pagination"
	"log"
	"time"
)

type Articles struct {
	Id          int       `from:"id" db:"id" json:"id"`
	Title       string    `from:"title" db:"title" json:"title"`
	CategoryId  string    `from:"category_id" db:"category_id" json:"category_id"`
	Description string    `from:"description" db:"description" json:"description"`
	Author      string    `from:"author" db:"author" json:"author"`
	Content     string    `from:"content" db:"content" json:"content"`
	CreatedAt   time.Time `db:"created_at" json:"created_at" time_format:"2006-01-02 15:04:05"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at" time_format:"2006-01-02 15:04:05"`
}

func (a *Articles) DbName() string {
	return "default"
}

func (a *Articles) TableName() string {
	return "articles"
}

func (a *Articles) PK() string {
	return "id"
}

type ArticleResp struct {
	List    []*ArticleList `json:"list"`
	Total   int            `json:"total"`
	Current int            `json:"current"`
}

type ArticleList struct {
	Articles
	Category string `db:"name"`
}

func GetArticleList(article *Articles, page int, num int, keyword string, startTime string, endTime string) (*ArticleResp, error) {
	var articles = make([]*ArticleList, 0)
	start := (page - 1) * num
	args := make([]interface{}, 0)
	where := " 1 = 1 "

	if article.Id > 0 {
		where += " and a.id = ? "
		args = append(args, article.Id)
	}

	if keyword != "" {
		where += " and a.title like ? "
		args = append(args, "%"+keyword+"%")
	}

	if article.CategoryId != "" {
		where += " and a.category_id = ? "
		args = append(args, article.CategoryId)
	}

	if startTime != "" && endTime != "" {
		where += " and a.created_at between ? and ? "
		args = append(args, startTime, endTime)
	}
	total, err := gosql.Model(&Articles{}).Where(where, args...).Count()
	if err != nil {
		return nil, err
	}
	args = append(args, start, num)
	log.Print("sql begin")
	rows, err := gosql.Queryx("select a.*,c.name from articles a left join category c on c.id = a.category_id  where "+where+" order by a.id desc limit ?,?", args...)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		v := &ArticleList{}
		err := rows.StructScan(v)
		if err != nil {
			return nil, err
		}
		articles = append(articles, v)
	}
	pages := pagination.New(int(total), num, page, 5)

	return &ArticleResp{articles, pages.Total(), 0}, nil
}
