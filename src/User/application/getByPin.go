package application

import "api/src/User/domain/repository"
import "api/src/User/domain/entities"
// ObtenerUsuarioPorPin es el caso de uso para obtener un usuario por su PIN
type ObtenerUsuarioPorPin struct {
	repo repository.UserRepository
}

// NewObtenerUsuarioPorPin crea una nueva instancia del caso de uso ObtenerUsuarioPorPin
func NewObtenerUsuarioPorPin(repo repository.UserRepository) *ObtenerUsuarioPorPin {
	return &ObtenerUsuarioPorPin{repo: repo}
}

// Ejecutar obtiene un usuario de la base de datos por su PIN
func (uc *ObtenerUsuarioPorPin) Ejecutar(pin string) (*entities.User, error) {
	// Llamar al repositorio para buscar el usuario por PIN
	return uc.repo.GetUserByPin(pin)
}
