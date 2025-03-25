
package repository


import (
	"api/src/User/domain/entities"
	
)


type UserRepository interface {
	CreateUser(user *entities.User) (string, error)       // Crear un nuevo usuario
	DeleteUser(ID string) error                  // Eliminar un usuario por ID
	GetAllUsers() ([]entities.User, error)                // Obtener todos los usuarios
	GetUserByPin(Pin string) (*entities.User, error)      
	AddGuest(userID string, guest entities.Invitado) error
	//RemoveGuest(userID string, guestID string) error
}
