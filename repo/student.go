package repo

import "go-rest/svc"

type StudentRepo interface {
	svc.StudentRepo
}

type studentRepo struct{}

func NewStudentRepo() StudentRepo {
	return &studentRepo{}
}

func (r *studentRepo) GetStudent(id string) *svc.Student {
	std := &svc.Student{
		ID:    "1",
		Name:  "Proshanto",
		Email: "proshantolal@gmail.com",
	}

	return std
}
