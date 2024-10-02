package database

type Config struct{
	host string
	port string 
	username string
	password string 
	database string	
}

func NewConfig() *Config{
	return &Config{
		host: "localhost",
		port: "5432",
		username: "postgres",
		password: "Kra888452",
		database: "postgres",
	}
}

