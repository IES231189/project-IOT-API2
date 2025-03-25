package application

import "api/src/User/domain/repository"
import "api/src/User/domain/entities"

// CrearUsuario es el caso de uso para crear un nuevo usuario
type CrearUsuario struct {
	repo repository.UserRepository
}

// NewCrearUsuario crea una nueva instancia del caso de uso CrearUsuario
func NewCrearUsuario(repo repository.UserRepository) *CrearUsuario {
	return &CrearUsuario{repo: repo}
}

// Ejecutar crea un nuevo usuario en la base de datos y devuelve su ID
func (uc *CrearUsuario) Ejecutar(user *entities.User) (string, error) {
	return uc.repo.CreateUser(user) // Devuelve ID y error
}
