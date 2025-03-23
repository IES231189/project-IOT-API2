package application

import (
	"api/src/User/domain"
	"errors"
)

type UserService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{repo: repo}
}


func (s *UserService) AddGuest(userID string, invitado domain.Invitado) error {
	if userID == "" {
		return errors.New("el ID del usuario no puede estar vacío")
	}
	if invitado.Nombre == "" || invitado.Pin == "" {
		return errors.New("el invitado debe tener nombre y pin")
	}

	
	return s.repo.AddGuest(userID, invitado)
}
