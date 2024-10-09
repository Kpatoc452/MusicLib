package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)


type DataBase struct{
	db *sql.DB
}

func IntitDb(cfg *Config) (*DataBase, error ){
	connstr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.host, cfg.port, cfg.username, cfg.password, cfg.database)

	db, err := sql.Open("postgres", connstr)
	if err != nil {
		log.Println("Error init database. InitDb")
		return nil, err
	} 

	err = db.Ping()
	if err != nil{
		log.Println("Error ping db. InitDb")
		return nil, err
	}

	fmt.Println("Succesfully connected")
	return &DataBase{
		db: db,
	}, nil
}
