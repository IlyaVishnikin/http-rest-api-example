package apiserver

import (
	"encoding/json"
	"github.com/IlyaVishnikin/http-rest-api/internal/logger"
	"io"
	"os"
)

type Config struct {
	BindAddr string `json:"bind_addr"`
}

func NewConfig(path string) *Config {
	conf := Config{}
	confFile, err := os.Open(path)
	if err != nil {
		logger.Err.Fatalf("Cannot open configuration file %s: %v\n", path, err)
	}
	defer confFile.Close()
	data, err := io.ReadAll(confFile)
	if err != nil {
		logger.Err.Fatalf("Cannot read configuration file: %s: %v\n", path, err)
	}
	err = json.Unmarshal([]byte(data), &conf)
	if err != nil {
		logger.Err.Fatalf("Parsing error %s: %v\n", path, err)
	}
	return &conf
}
