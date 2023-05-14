package apiserver

import (
	"github.com/IlyaVishnikin/http-rest-api/internal/handler"
	"github.com/IlyaVishnikin/http-rest-api/internal/logger"
	"net/http"
)

type APIServer struct {
	Config *Config
}

func New(conf *Config) *APIServer {
	return &APIServer{
		Config: conf,
	}
}

func (s *APIServer) Start() error {
	logger.Info.Println("Starting server")
	http.HandleFunc("/", handler.Index)
	return http.ListenAndServe(s.Config.BindAddr, nil)
}
