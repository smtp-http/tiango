package main

import(
	"fmt"
	"github.com/smtp-http/tiango/config"
	"github.com/smtp-http/tiango/api"
)


func main(){
	loader := config.GetLoader()
	loader.Load("./config.json",config.GetConfig())

	port := config.GetConfig().HttpPort

	fmt.Printf("tiango: %v\n",port)

	server := &api.GinServer{}

	server.StartHttpServer()
}