package main

import (
	"fmt"
	"musiclib/cmd/server"
	"musiclib/database"

	"musiclib/internal/controllers"
	"musiclib/internal/errors"
	"musiclib/internal/service"
)


func main(){
	cfg := InitConfig()
	db, err := database.IntitDb(cfg.DataBaseCfg)
	if err != nil{
		panic(err)
	}
	service := service.NewService(db)
	errs := errors.NewError()
	handler := controllers.NewHandle(service, errs)
	router := controllers.NewRouter(handler)

	server := server.NewServer()

	server.AddEndpoint("/info", router.Song)
	
	server.Init(cfg.ServerCfg)
	fmt.Println(server.Rgstr)
	server.Run()
}