package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "blog/config"
	"blog/server"
)

func main() {
	//command server
	cliServ := server.NewCliServer()
	cliServ.Run()
}