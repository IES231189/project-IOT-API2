package domain


type UserRepository interface {
	CreateUser(user *User) (string, error)       // Crear un nuevo usuario
	DeleteUser(ID string) error                  // Eliminar un usuario por ID
	GetAllUsers() ([]User, error)                // Obtener todos los usuarios
	GetUserByPin(Pin string) (*User, error)      // Buscar un usuario por PIN
}
