package apiserver

type APIServer struct {
	Config *Config
}

func New(conf *Config) *APIServer {
	return &APIServer{
		Config: conf,
	}
}

func (s *APIServer) Start() error {
	return nil
}
