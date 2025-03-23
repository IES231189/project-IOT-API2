package application

import "api/src/StrongBox/domain"

type GetStrongBoxByID struct {
	repo domain.StrongBoxRepository
}

func NewGetStrongBoxByID(repo domain.StrongBoxRepository) *GetStrongBoxByID {
	return &GetStrongBoxByID{repo: repo}
}

func (uc *GetStrongBoxByID) Execute(ID string) (*domain.StrongBox, error) {
	return uc.repo.GetStrongBoxByID(ID)
}
