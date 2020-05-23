package main

import (
	"flag"
	"fmt"
	"github.com/istore221/custo-api/config"
	"github.com/istore221/custo-api/internal/log"
	"net/http"
	"os"
	"time"
)

// Version indicates the current version of the application.
var Version = "1.0.0"

var flagConfig = flag.String("config", "./config/development.yml", "path to the config file")




func main() {

	flag.Parse()

	

	//logz.SetFormatter(&logz.JSONFormatter{})
	//logz.WithFields(logz.Fields{
	//	"animal": "walrus",
	//	"number": 1,
	//	"size":   10,
	//}).Info("A walrus appears")


	//create root log tagged with server version

	//    var log = &logrus.Logger{
	//      Out: os.Stderr, // console log writter
	//      Formatter: new(logrus.JSONFormatter),
	//      Hooks: make(logrus.LevelHooks),
	//      Level: logrus.DebugLevel,
	//    }


	log.GetLogger().WithFields(log.Fields{
			"animal": "walrus",
			"number": 1,
			"size":   10,
	}).Warn("A walrus appears")

	fmt.Println(log.GetLogger().GetLevel())


	//fmt.Println("CHANGED LEVEL ",logger.GetLevel())
	//
	//switch v := logger.(type) {// type switch
	//case *log.MockLogger:
	//	fmt.Println("x is MockLogger", v)
	//case *log.LogrusLogger:
	//
	//	fmt.Println("x LogrusLogger",)
	//default:
	//	fmt.Println("type unknown", v)
	//}

	
	// load application configurations
	// flag set based on APP_ENV ENV Variable
	cfg, err := config.LoadConfiguration(*flagConfig)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}


	hs := &http.Server{
		Addr:    ":"+cfg.Server.Port,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		log.GetLogger().WithFields(log.Fields{
			"name" : "cat",
		}).Debug("/ served")
		writer.Write([]byte("/ Hello world "+cfg.Environment))
	})
	mux.HandleFunc("/go", func(writer http.ResponseWriter, request *http.Request) {
		log.GetLogger().Info("/go served")
		writer.Write([]byte("/go Hello world "+cfg.Environment))
	})

	hs.Handler = mux


	
	if err := hs.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		//error
		log.GetLogger().Info("error occured")
		os.Exit(-1)
	}


}

//https://gist.github.com/enricofoltran/10b4a980cd07cb02836f70a4ab3e72d7

//https://github.com/gilcrest/go-api-basic/blob/master/main.go

//https://github.com/cobbinma/example-go-api

//https://medium.com/@eminetto/clean-architecture-using-golang-b63587aa5e3f


//https://golangbot.com/custom-errors/

//https://gist.github.com/peterhellberg/e36274f213f7a2e2b89a3d837fbafbe1

//https://github.com/sirupsen/logrus/blob/master/logger.go
//https://github.com/sirupsen/logrus/blob/master/entry.go