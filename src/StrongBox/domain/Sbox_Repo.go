package domain


type StrongBoxRepository interface {
    CreateStrongBox(box *StrongBox) (string, error)           // Crea una nueva caja fuerte
    DeleteStrongBox(ID string) error                          // Elimina una caja fuerte existente
    GetAllStrongBoxes() ([]StrongBox, error)                  // Obtiene todas las cajas fuertes
    GetStrongBoxByID(ID string) (*StrongBox, error)           // Obtiene una caja fuerte por ID
    
	AddUserToStrongBox(boxID string, userID string) error 
    RemoveUserFromStrongBox(boxID string, userID string) error  // Elimina un usuario de la caja fuerte
}
