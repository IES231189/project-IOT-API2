package application

import "api/src/User/domain"

// ObtenerInvitadosPorID es el caso de uso para obtener los invitados de un usuario por su ID
type ObtenerInvitadosPorID struct {
	repo domain.UserRepository
}

// NewObtenerInvitadosPorID crea una nueva instancia del caso de uso ObtenerInvitadosPorID
func NewObtenerInvitadosPorID(repo domain.UserRepository) *ObtenerInvitadosPorID {
	return &ObtenerInvitadosPorID{repo: repo}
}

// Ejecutar obtiene los invitados de un usuario de la base de datos por su ID
func (uc *ObtenerInvitadosPorID) Ejecutar(userID string) ([]domain.Invitado, error) {
	// Llamar al repositorio para obtener los invitados por el ID del usuario
	return uc.repo.GetGuestsByUserID(userID)
}
