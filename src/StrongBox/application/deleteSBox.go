package application

import "api/src/StrongBox/domain"

type DeleteStrongBox struct {
	repo domain.StrongBoxRepository
}

func NewDeleteStrongBox(repo domain.StrongBoxRepository) *DeleteStrongBox {
	return &DeleteStrongBox{repo: repo}
}

func (uc *DeleteStrongBox) Execute(ID string) error {
	return uc.repo.DeleteStrongBox(ID)
}
