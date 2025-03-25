package application

import "api/src/StrongBox/domain"

type RemoveUserFromStrongBoxService struct {
	repository domain.StrongBoxRepository
}

func NewRemoveUserFromStrongBoxService(repository domain.StrongBoxRepository) *RemoveUserFromStrongBoxService {
	return &RemoveUserFromStrongBoxService{repository: repository}
}

func (s *RemoveUserFromStrongBoxService) Execute(boxID string, userID string) error {
	return s.repository.RemoveUserFromStrongBox(boxID, userID)
}
