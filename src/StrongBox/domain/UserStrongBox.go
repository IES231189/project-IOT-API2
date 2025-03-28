package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserStrongBox struct {
	UsuarioID primitive.ObjectID `bson:"usuario_id,omitempty" json:"usuario_id"`
	Nombre    string             `bson:"nombre" json:"nombre"`
	Pin       string             `bson:"pin" json:"pin"`
}
