package application

import "api/src/User/domain"

// ObtenerUsuarioPorPin es el caso de uso para obtener un usuario por su PIN
type ObtenerUsuarioPorPin struct {
	repo domain.UserRepository
}

// NewObtenerUsuarioPorPin crea una nueva instancia del caso de uso ObtenerUsuarioPorPin
func NewObtenerUsuarioPorPin(repo domain.UserRepository) *ObtenerUsuarioPorPin {
	return &ObtenerUsuarioPorPin{repo: repo}
}

// Ejecutar obtiene un usuario de la base de datos por su PIN
func (uc *ObtenerUsuarioPorPin) Ejecutar(pin string) (*domain.User, error) {
	// Llamar al repositorio para buscar el usuario por PIN
	return uc.repo.GetUserByPin(pin)
}
