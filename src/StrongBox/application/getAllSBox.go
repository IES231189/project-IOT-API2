package application

import "api/src/StrongBox/domain"

type GetAllStrongBox struct {
	repo domain.StrongBoxRepository
}

func NewGetAllStrongBox(repo domain.StrongBoxRepository) *GetAllStrongBox {
	return &GetAllStrongBox{repo: repo}
}

func (uc *GetAllStrongBox) Execute() ([]domain.StrongBox, error) {
	return uc.repo.GetAllStrongBoxes()
}
