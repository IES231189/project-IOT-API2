package domain

type StrongBoxRepository interface {
	CreateStrongBox(box *StrongBox) (string, error) 
	DeleteStrongBox(ID string) error                 
	GetAllStrongBoxes() ([]StrongBox, error)         
	GetStrongBoxByID(ID string) (*StrongBox, error)  
}
