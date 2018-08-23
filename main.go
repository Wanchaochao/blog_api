package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "app/config"
	"app/server"
)

func main() {
	//command server
	cliServ := server.NewCliServer()
	cliServ.Run()
}