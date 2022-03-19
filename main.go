package main

import (
	"fmt"
	"user/domain/repository"
	"user/handler"

	"github.com/jinzhu/gorm"
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("user"),
		service.Version("latest"),
	)

	// Init service
	srv.Init()

	// Init database
	db, err := gorm.Open("test.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	db.SingularTable(true)
	rp := repository.NewUserRepository(db)
	rp.InitTable()

	// Service instance
	// userDatasService := service.NewUserDatasService(repository.NewUserRepository(db))
	// err := user.RegisterUserHandler(srv.Server(), &handler.User{UserdatasService: userDatasService})
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// Register handler
	// pb.RegisterUserHandler(srv.Server(), new(handler.User))
	srv.Handle(new(handler.User))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
