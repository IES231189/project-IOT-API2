package application

import "api/src/User/domain"

// CrearUsuario es el caso de uso para crear un nuevo usuario
type CrearUsuario struct {
	repo domain.UserRepository
}

// NewCrearUsuario crea una nueva instancia del caso de uso CrearUsuario
func NewCrearUsuario(repo domain.UserRepository) *CrearUsuario {
	return &CrearUsuario{repo: repo}
}

// Ejecutar crea un nuevo usuario en la base de datos y devuelve su ID
func (uc *CrearUsuario) Ejecutar(user *domain.User) (string, error) {
	return uc.repo.CreateUser(user) // Devuelve ID y error
}
