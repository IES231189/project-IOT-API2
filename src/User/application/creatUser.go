package application

import (
	"api/src/User/domain"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// CrearUsuario es el caso de uso para crear un nuevo usuario
type CrearUsuario struct {
	repo domain.UserRepository
}

// NewCrearUsuario crea una nueva instancia del caso de uso CrearUsuario
func NewCrearUsuario(repo domain.UserRepository) *CrearUsuario {
	return &CrearUsuario{repo: repo}
}

// Ejecutar crea un nuevo usuario en la base de datos y devuelve su ID
func (uc *CrearUsuario) Ejecutar(user *domain.User) (string, error) {
	// Verificar que la contraseña no esté vacía
	if user.Contraseña == "" {
		return "", errors.New("la contraseña no puede estar vacía")
	}

	// Hashear la contraseña con bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Contraseña), bcrypt.DefaultCost)
	if err != nil {
		return "", err // Si ocurre un error al hashear la contraseña, se devuelve el error
	}
	user.Contraseña = string(hashedPassword) // Asignar el hash al campo de la contraseña

	// Guardar el usuario con la contraseña encriptada en la base de datos
	return uc.repo.CreateUser(user) // Devuelve el ID y error
}
