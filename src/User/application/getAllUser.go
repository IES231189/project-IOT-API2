package application

import "api/src/User/domain/repository"
import "api/src/User/domain/entities"

// ObtenerTodosLosUsuarios es el caso de uso para obtener todos los usuarios
type ObtenerTodosLosUsuarios struct {
	repo repository.UserRepository
}

// NewObtenerTodosLosUsuarios crea una nueva instancia del caso de uso ObtenerTodosLosUsuarios
func NewObtenerTodosLosUsuarios(repo repository.UserRepository) *ObtenerTodosLosUsuarios {
	return &ObtenerTodosLosUsuarios{repo: repo}
}

// Ejecutar obtiene todos los usuarios de la base de datos
func (uc *ObtenerTodosLosUsuarios) Ejecutar() ([]entities.User, error) {
	// Llamar al repositorio para obtener todos los usuarios
	return uc.repo.GetAllUsers()
}
