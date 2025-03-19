// application/crear_strongbox_service.go
package application

import "api/src/StrongBox/domain"

type CrearStrongBoxService struct {
	repository domain.StrongBoxRepository 
}

func NewCreateStrongBoxService(repository domain.StrongBoxRepository) *CrearStrongBoxService {
	return &CrearStrongBoxService{
		repository: repository,
	}
}

func (s *CrearStrongBoxService) Execute(box *domain.StrongBox) (string, error) {
	return s.repository.CreateStrongBox(box)
}
