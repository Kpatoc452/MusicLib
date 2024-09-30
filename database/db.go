package database

import (
	"database/sql"
	"fmt"
)


type DataBase struct{
	db *sql.DB
}

func IntitDb() *DataBase {
	cfg := newConfig()
	connstr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.host, cfg.port, cfg.username, cfg.password, cfg.database)

	db, err := sql.Open(cfg.database, connstr)
	if err != nil {
		panic(err)
	} 

	err = db.Ping()
	if err != nil{
		panic(err)
	}

	fmt.Println("Succesfully connected")
	return &DataBase{
		db: db,
	}
}
