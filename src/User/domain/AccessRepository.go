package domain

import(
	
	"time"
)


type AccesoRepository interface {
    GuardarAccesoExitoso(pin string, usuario string, estado string, fecha time.Time) error
    ObtenerAccesosExitosos() ([]AccesoExitoso, error)
	GuardarAccesoIncorrecto(pin string, deviceId string, estado string, fecha time.Time) error
	ObtenerAccesosIncorrectos() ([]AccesoIncorrecto, error)
}