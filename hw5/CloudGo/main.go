package main

import (
	"CloudGo/service"
	"github.com/spf13/pflag"
	"os"
)

const (
	PORT string = "8080"
)

func main(){
	port := os.Getenv("PORT")

	if len(port) == 0{
		port = PORT
	}

	pPort :=  pflag.StringP("port", "p", PORT, "PORT for httpd listening")
	pflag.Parse()
	if len(*pPort)!=0{
		port = *pPort
	}

	service.NewServer(port)
}
