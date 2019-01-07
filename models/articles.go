package models

import (
	"blog/util"
	"github.com/ilibs/gosql"
	"log"
	"time"
)

type Articles struct {
	Id          int       `from:"id" db:"id" json:"id"`
	Title       string    `from:"title" db:"title" json:"title"`
	CategoryId  int       `from:"category_id" db:"category_id" json:"category_id"`
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

// 文章的赞扬反对数
type ArticleEvaluate struct {
	PraiseNum  int `json:"praise_num"`
	AgainstNum int `json:"against_num"`
}

type ArticleInfo struct {
	Articles
	CategoryInfo    *Category `json:"category_info" db:"-" relation:"category_id,id"`
	ArticleEvaluate `json:"article_evaluate" db:"-"`
	Prev            Articles `json:"prev"`
	Next            Articles `json:"next"`
}

type ArticleList struct {
	Articles
	CategoryInfo *Category `json:"category_info" db:"-" relation:"category_id,id"`
}

// 文章列表
type ArticleResp struct {
	List []*ArticleList `json:"list"`
	Page *util.Page     `json:"page"`
}

func GetArticleList(article *Articles, page int, num int, keyword string, startTime string, endTime string) (*ArticleResp, error) {
	start := (page - 1) * num
	args := make([]interface{}, 0)
	where := " 1 = 1 "

	if article.Id > 0 {
		where += " and articles.id = ? "
		args = append(args, article.Id)
	}

	if keyword != "" {
		where += " and articles.title like ? "
		args = append(args, "%"+keyword+"%")
	}

	if article.CategoryId > 0 {
		where += " and articles.category_id = ? "
		args = append(args, article.CategoryId)
	}

	if startTime != "" && endTime != "" {
		where += " and articles.created_at between ? and ? "
		args = append(args, startTime, endTime)
	}
	total, err := gosql.Model(&Articles{}).Where(where, args...).Count()
	if err != nil {
		return nil, err
	}
	args = append(args, start, num)
	log.Print("sql begin")
	var articles = make([]*ArticleList, 0)
	if err = gosql.Model(&articles).Where(where).Limit(num).All(); err != nil {
		return nil, err
	}
	pageStruct := &util.Page{
		Current:  page,
		PageSize: num,
		Total:    int(total),
	}
	return &ArticleResp{articles, pageStruct}, nil
}
