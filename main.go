package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "blog/config"
	"blog/server"
	"github.com/ilibs/gosql"
	"blog/config"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	// 连接数据库
	gosql.Connect(config.App.Db)
	//command server
	cliServ := server.NewCliServer()
	setPid(os.Getpid())
	cliServ.Run()
}

func setPid(pid int) {
	d := []byte(strconv.Itoa(pid))
	err := ioutil.WriteFile("./blog.pid", d, 0644)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}
}