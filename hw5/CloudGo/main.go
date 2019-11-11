package main

import (
	"CloudGo/service"
	"github.com/spf13/pflag"
	"os"
)

const (
	//set a default port
	PORT string = "8080"
)

func main(){
	port := os.Getenv("PORT")
	//use default port
	if len(port) == 0{
		port = PORT
	}
	//allow user set another port
	pPort :=  pflag.StringP("port", "p", PORT, "PORT for httpd listening")
	pflag.Parse()
	//if user set a port, use it
	if len(*pPort)!=0{
		port = *pPort
	}
	//run server
	service.NewServer(port)
}
