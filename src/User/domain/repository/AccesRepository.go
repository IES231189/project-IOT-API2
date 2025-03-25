package repository

import(
	"api/src/User/domain/entities"
	"time"
)


type AccesoRepository interface {
    GuardarAccesoExitoso(pin string, usuario *entities.User, estado string, fecha time.Time) error
    ObtenerAccesosExitosos() ([]entities.AccesoExitoso, error)
}
