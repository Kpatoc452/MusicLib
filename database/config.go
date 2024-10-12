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
		port: "5436",
		username: "musicpguser",
		password: "musicpgpswd",
		database: "musicdb",
	}
}

