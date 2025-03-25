package application

import "api/src/StrongBox/domain"

type AddUserToStrongBoxService struct {
	repository domain.StrongBoxRepository
}

// ✅ Recibe el repositorio como dependencia
func NewAddUserToStrongBoxService(repository domain.StrongBoxRepository) *AddUserToStrongBoxService {
	return &AddUserToStrongBoxService{
		repository: repository,
	}
}

// Lógica del caso de uso (sin acceso a la base de datos)
func (s *AddUserToStrongBoxService) Execute(boxID string, user *domain.UserStrongBox) error {
	return s.repository.AddUserToStrongBox(boxID, user)
}
