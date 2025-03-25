package application

import (
	"errors"
	"api/src/User/domain/repository"
	"api/src/User/domain/entities"
)


type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}


func (s *UserService) AddGuest(userID string, invitado entities.Invitado) error {
	if userID == "" {
		return errors.New("el ID del usuario no puede estar vacío")
	}
	if invitado.Nombre == "" || invitado.Pin == "" {
		return errors.New("el invitado debe tener nombre y pin")
	}

	
	return s.repo.AddGuest(userID, invitado)
}
