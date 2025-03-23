package application

import (
	"api/src/User/domain"
)

type EliminarUsuario struct {
	repo domain.UserRepository
}

// NewEliminarUsuario crea una nueva instancia del caso de uso EliminarUsuario
func NewEliminarUsuario(repo domain.UserRepository) *EliminarUsuario {
	return &EliminarUsuario{repo: repo}
}


func (uc *EliminarUsuario) Ejecutar(ID string) error {
	// Llamar al repositorio para eliminar el usuario
	return uc.repo.DeleteUser(ID)
}
