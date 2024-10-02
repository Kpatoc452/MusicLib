package server

type Config struct{ 
	Addr string
}

func NewConfig() *Config{
	return &Config{
		Addr: "localhost:8080",
	}
}