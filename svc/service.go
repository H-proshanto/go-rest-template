package svc

type service struct {
	studentRepo StudentRepo
}

func NewService(studentRepo StudentRepo) Service {
	return &service{
		studentRepo: studentRepo,
	}
}

func (s *service) GetStudent(id string) *Student {
	return s.studentRepo.GetStudent(id)
}
