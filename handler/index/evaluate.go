package index

import (
	"blog/core"
	"blog/models"
	"database/sql"
	"github.com/ilibs/gosql"
	"strconv"
)

var Evaluate core.HandlerFunc = func(c *core.Context) core.Response {
	evaluate := &models.Evaluate{}
	args := make([]interface{}, 0)

	fkey, _ := strconv.Atoi(c.DefaultQuery("id", ""))
	evaluate.ForeignKey = fkey
	args = append(args, fkey)

	ftype, _ := strconv.Atoi(c.DefaultQuery("type", ""))
	evaluate.Type = ftype
	args = append(args, ftype)

	evaluate.Ip = c.ClientIP()
	args = append(args, evaluate.Ip)

	praise, _ := strconv.Atoi(c.DefaultQuery("praise", ""))
	evaluate.Praise = praise
	args = append(args, praise)

	msg := ""

	// 先判断是否为取消点踩点赞(重复点)
	if err := gosql.Model(&models.Evaluate{}).Where("foreign_key = ? and type = ? and ip = ? and praise = ? ", args...).Get(); err != nil {
		if err == sql.ErrNoRows {
			// 不是重复点踩点赞,判断是否为踩变赞，赞变踩
			err := gosql.Model(&models.Evaluate{}).Where("foreign_key = ? and type = ? and ip = ? ", args[0:3]...).Get()
			if err != nil {
				if err != sql.ErrNoRows {
					return c.Fail(201, err)
				} else {
					// 没点过踩和赞，新建
					_, err := gosql.Model(evaluate).Create()
					if err != nil {
						return c.Fail(202, err)
					}
				}

			} else {
				// 踩赞互换
				if _, err := gosql.Model(&models.Evaluate{Praise: praise}).Where("foreign_key = ? and type = ? and ip = ? ", args[0:3]...).Update("praise"); err != nil {
					return c.Fail(203, err)
				}
			}
			if praise == 1 {
				msg = "点赞成功!"
			} else {
				msg = "踩了一下!"
			}
		} else {
			// 其他错误
			return c.Fail(204, err)
		}
	} else {
		// 查到了,取消点赞点踩
		if _, err := gosql.Model(evaluate).Delete(); err != nil {
			return c.Fail(205, "删除失败")
		}
		if praise == 1 {
			msg = "取消赞!"
		} else {
			msg = "取消踩!"
		}
	}

	resp := struct {
		PraiseNum  int `json:"praise_num"`
		AgainstNum int `json:"against_num"`
	}{}
	p, err := gosql.Model(&models.Evaluate{}).Where("foreign_key = ? and type = ? and praise = 1", fkey, ftype).Count()
	if err != nil {
		return c.Fail(206, err)
	}
	a, err := gosql.Model(&models.Evaluate{}).Where("foreign_key = ? and type = ? and praise = 0", fkey, ftype).Count()
	if err != nil {
		return c.Fail(206, err)
	}
	resp.PraiseNum = int(p)
	resp.AgainstNum = int(a)
	return c.Success(resp, msg)
}
