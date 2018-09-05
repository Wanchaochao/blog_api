package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "blog/config"
	"blog/server"
	"github.com/ilibs/gosql"
	"blog/config"
)

func main() {
	// 连接数据库
	gosql.Connect(config.App.Db)
	//command server
	cliServ := server.NewCliServer()
	cliServ.Run()
}