package cmd

import (
	"go-rest/config"
	"go-rest/repo"
	"go-rest/rest"
	"go-rest/svc"
)

func serveRest() {
	appConfig := config.GetApp()
	stdRepo := repo.NewStudentRepo()
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
