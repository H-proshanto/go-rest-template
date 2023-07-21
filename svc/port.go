package svc

type StudentRepo interface {
	GetStudent(id string) *Student
	CreateStudent(std *Student)
	UpdateStudent(id string, std *Student) *Student
	DeleteStudent(id string) *Student
}

type Service interface {
	GetStudent(id string) *Student
	CreateStudent(std *Student)
	UpdateStudent(id string, std *Student) *Student
	DeleteStudent(id string) *Student
}
