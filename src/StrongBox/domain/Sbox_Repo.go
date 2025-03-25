package domain

type StrongBoxRepository interface {
	CreateStrongBox(box *StrongBox) (string, error) 
	DeleteStrongBox(ID string) error                 
	GetAllStrongBoxes() ([]StrongBox, error)         
	GetStrongBoxByID(ID string) (*StrongBox, error)
	
	AddUserToStrongBox(boxID string, user *UserStrongBox) error
	RemoveUserFromStrongBox(boxID string, userID string) error  
}
