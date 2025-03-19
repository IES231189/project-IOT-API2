package application

import "api/src/User/domain"

// ObtenerTodosLosUsuarios es el caso de uso para obtener todos los usuarios
type ObtenerTodosLosUsuarios struct {
	repo domain.UserRepository
}

// NewObtenerTodosLosUsuarios crea una nueva instancia del caso de uso ObtenerTodosLosUsuarios
func NewObtenerTodosLosUsuarios(repo domain.UserRepository) *ObtenerTodosLosUsuarios {
	return &ObtenerTodosLosUsuarios{repo: repo}
}

// Ejecutar obtiene todos los usuarios de la base de datos
func (uc *ObtenerTodosLosUsuarios) Ejecutar() ([]domain.User, error) {
	// Llamar al repositorio para obtener todos los usuarios
	return uc.repo.GetAllUsers()
}
