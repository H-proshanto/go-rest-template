package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) getStudent(ctx *gin.Context) {
	studentId := ctx.Param("id")

	id := s.svc.GetStudent(studentId)

	ctx.JSON(http.StatusOK, id)

}
