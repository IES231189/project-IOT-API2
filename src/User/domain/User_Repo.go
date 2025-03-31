package domain


type UserRepository interface {
	CreateUser(user *User) (string, error)    
	DeleteUser(ID string) error                 
	GetAllUsers() ([]User, error)                
	GetUserByPin(Pin string) (*User, error)
	GetGuestsByUserID(userID string) ([]Invitado, error) 
	LoginUser(correo, contraseña string) (*User, error)      
	AddGuest(userID string, guest Invitado) error
	RemoveGuest(userID string, guestID string) error
}
