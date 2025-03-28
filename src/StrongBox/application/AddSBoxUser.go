package application

import "api/src/StrongBox/domain"


type AddUserToStrongBoxService struct {
    repo domain.StrongBoxRepository
}


func NewAddUserToStrongBoxService(repo domain.StrongBoxRepository) *AddUserToStrongBoxService {
    return &AddUserToStrongBoxService{repo: repo}
}


func (s *AddUserToStrongBoxService) Execute(boxID string, userID string) error {
    return s.repo.AddUserToStrongBox(boxID, userID)
}
