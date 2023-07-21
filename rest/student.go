package rest

import (
	"encoding/json"
	"go-rest/svc"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) getStudent(ctx *gin.Context) {
	studentId := ctx.Param("id")

	id := s.svc.GetStudent(studentId)

	ctx.JSON(http.StatusOK, id)

}

func (s *Server) createStudent(ctx *gin.Context) {
	var std svc.Student
	body, err := ctx.GetRawData()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
		return
	}

	err = json.Unmarshal(body, &std)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	s.svc.CreateStudent(&std)

	ctx.JSON(http.StatusOK, std)
}

func (s *Server) updateStudent(ctx *gin.Context) {
	var std svc.Student

	studentId := ctx.Param("id")
	body, err := ctx.GetRawData()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
		return
	}

	err = json.Unmarshal(body, &std)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	updatedStd := s.svc.UpdateStudent(studentId, &std)

	ctx.JSON(http.StatusOK, updatedStd)
}

func (s *Server) deleteStudent(ctx *gin.Context) {
	studentId := ctx.Param("id")

	std := s.svc.DeleteStudent(studentId)

	ctx.JSON(http.StatusOK, std)
}
