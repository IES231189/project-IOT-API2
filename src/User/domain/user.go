package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id"` // Asegúrate de usar _id para MongoDB
	Nombre string             `bson:"Nombre" json:"nombre"`
	Correo string             `bson:"Correo" json:"correo"`
	Pin    string             `bson:"Pin" json:"pin"`
}