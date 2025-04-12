// En domain/acceso.go
package domain

import "time"

type AccesoIncorrecto struct {
    Pin         string    `bson:"pin" json:"pin"`
    DeviceId    string    `bson:"deviceId" json:"deviceId"`
    Estado      string    `bson:"estado" json:"estado"` // "incorrecto", "intrusión", etc.
    Fecha       time.Time `bson:"fecha" json:"fecha"`
    EsIntrusion bool      `bson:"esIntrusion" json:"esIntrusion"`
    Intento     int       `bson:"intento" json:"intento"` // Número de intento consecutivo
}