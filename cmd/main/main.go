package main

import (
	"flag"
	"github.com/IlyaVishnikin/http-rest-api/internal/apiserver"
	"github.com/IlyaVishnikin/http-rest-api/internal/logger"
)

const (
	infoLogPath = "log/info.log"
	warnLogPath = "log/warn.log"
	errLogPath  = "log/err.log"
)

var (
	configPath = flag.String("conf", "config/apiserver.json", "Server configuration file location")
)

func main() {
	flag.Parse()
	logger.Init(infoLogPath, warnLogPath, errLogPath)
	defer logger.Close()
	conf := apiserver.NewConfig(*configPath)
	serv := apiserver.New(conf)
	serv.Start()
}
