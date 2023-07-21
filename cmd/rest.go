package cmd

import (
	"go-rest/config"
	"go-rest/database"
	"go-rest/repo"
	"go-rest/rest"
	"go-rest/svc"
)

func serveRest() {
	appConfig := config.GetApp()
	db := database.NewDatabase()
	stdRepo := repo.NewStudentRepo(db)
	svc := svc.NewService(stdRepo)

	server, err := rest.NewServer(svc, appConfig)

	if err != nil {
		panic(err)
	}

	err = server.Start()

	if err != nil {
		panic(err)
	}
}
