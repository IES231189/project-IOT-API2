package application

import "api/src/User/domain/repository"


type RemoverInvitado struct {
	repo repository.UserRepository
}


func NewRemoverInvitado(repo repository.UserRepository) *RemoverInvitado {
	return &RemoverInvitado{repo: repo}
}


func (uc *RemoverInvitado) Ejecutar(userID string, guestID string) error {
	return uc.repo.RemoveGuest(userID, guestID)
}
