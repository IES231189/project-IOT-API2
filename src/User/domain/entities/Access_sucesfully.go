
package entities

import (
    "time"
    //"go.mongodb.org/mongo-driver/bson/primitive"

)

type AccesoExitoso struct {
    Pin     string    `bson:"pin" json:"pin"` 
    Usuario *User     `bson:"usuario" json:"usuario"`
    Estado  string    `bson:"estado" json:"estado"`
    Fecha   time.Time `bson:"fecha" json:"fecha"`
}