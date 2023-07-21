package rest

import (
	"fmt"
	"go-rest/config"
	"go-rest/svc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router    *gin.Engine
	svc       svc.Service
	appConfig *config.Application
}

func NewServer(svc svc.Service, appConfig *config.Application) (*Server, error) {
	server := &Server{
		svc:       svc,
		appConfig: appConfig,
	}

	server.setupRouter()

	return server, nil
}

func (s *Server) setupRouter() {
	router := gin.Default()

	s.router = router

	router.GET("/api/test", test)
	router.POST("/api/students", s.createStudent)
	router.GET("/api/students/:id", s.getStudent)
	router.PATCH("/api/students/:id", s.updateStudent)
	router.DELETE("/api/students/:id", s.deleteStudent)
}

func (s *Server) Start() error {
	return s.router.Run(fmt.Sprintf("%s:%s", s.appConfig.Host, s.appConfig.Port))
}

func test(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Hello world")
}
