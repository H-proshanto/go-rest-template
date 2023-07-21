package svc

type StudentRepo interface {
	GetStudent(id string) *Student
}

type Service interface {
	GetStudent(id string) *Student
}
