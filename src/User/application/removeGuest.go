package application

import "api/src/User/domain"


type RemoverInvitado struct {
	repo domain.UserRepository
}


func NewRemoverInvitado(repo domain.UserRepository) *RemoverInvitado {
	return &RemoverInvitado{repo: repo}
}


func (uc *RemoverInvitado) Ejecutar(userID string, guestID string) error {
	return uc.repo.RemoveGuest(userID, guestID)
}
