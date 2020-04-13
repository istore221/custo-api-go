package main

import (
	"flag"
	"fmt"
	"github.com/istore221/custo-api/config"
	"log"
	"net/http"
)

// Version indicates the current version of the application.
var Version = "1.0.0"

var flagConfig = flag.String("config", "./config/development.yml", "path to the config file")

//
func main() {

	flag.Parse()

	// load application configurations
	// flag set based on APP_ENV ENV Variable
	cfg, err := config.LoadConfiguration(*flagConfig)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cfg)

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello world "+cfg.Environment))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

//https://gist.github.com/enricofoltran/10b4a980cd07cb02836f70a4ab3e72d7

//https://github.com/gilcrest/go-api-basic/blob/master/main.go

//https://github.com/cobbinma/example-go-api

//https://medium.com/@eminetto/clean-architecture-using-golang-b63587aa5e3f
