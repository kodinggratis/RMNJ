package service

import "github.com/kodinggratis/RMNJ.git/domain"

type userService struct{}

func NewUserService() domain.UserService {
	return &userService{}
}

func (s *userService) GetByID(id int) (domain.User, error) {
	// Di sini tempat logic: validasi, perhitungan, dll.
	return domain.User{ID: id, Name: "Koding Gratis User"}, nil
}
