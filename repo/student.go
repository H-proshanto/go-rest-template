package repo

import (
	"fmt"
	"go-rest/svc"
	"time"

	"gorm.io/gorm"
)

type StudentRepo interface {
	svc.StudentRepo
}

type studentRepo struct {
	db *gorm.DB
}

func NewStudentRepo(db *gorm.DB) StudentRepo {
	return &studentRepo{
		db: db,
	}
}

func (r *studentRepo) GetStudent(id string) *svc.Student {
	var std *svc.Student
	result := r.db.Where(fmt.Sprintf("id = %s", id)).First(&std)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			fmt.Println("User not found")
		} else {
			fmt.Println("Error while fetching user:", result.Error)
		}
	}

	return std
}

func (r *studentRepo) CreateStudent(std *svc.Student) {
	r.db.Create(std)
}

func (r *studentRepo) UpdateStudent(id string, newStd *svc.Student) *svc.Student {
	var std *svc.Student
	result := r.db.Where(fmt.Sprintf("id = %s", id)).First(&std)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			fmt.Println("User not found")
		} else {
			fmt.Println("Error while fetching user:", result.Error)
		}
	}

	std.Name = newStd.Name
	std.Email = newStd.Email
	std.UpdatedAt = time.Now()

	r.db.Save(&std)

	return std
}

func (r *studentRepo) DeleteStudent(id string) *svc.Student {
	var std *svc.Student
	result := r.db.Where(fmt.Sprintf("id = %s", id)).First(&std)

	std.DeletedAt = gorm.DeletedAt{
		Time: time.Now(),
	}

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			fmt.Println("User not found")
		} else {
			fmt.Println("Error while fetching user:", result.Error)
		}
	}

	r.db.Unscoped().Delete(&std)

	return std
}
