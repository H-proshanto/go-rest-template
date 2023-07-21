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

func (s *service) CreateStudent(std *Student) {
	s.studentRepo.CreateStudent(std)
}

func (s *service) UpdateStudent(id string, std *Student) *Student {
	return s.studentRepo.UpdateStudent(id, std)
}

func (s *service) DeleteStudent(id string) *Student {
	return s.studentRepo.DeleteStudent(id)
}
