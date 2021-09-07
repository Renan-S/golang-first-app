package main

//To run app ./mini-cli-app flag (ip or server) for default or ./mini-cli-app flag (ip or server) --host anyhost.com

import (
	"go-basics/mini-cli-app/app"
	"log"
	"os"
)

func main() {
	application := app.Create()                          //Calling Create from app/app.go
	if error := application.Run(os.Args); error != nil { //os.Args is a default value to get OS command line
		log.Fatal(error)
	}
}
