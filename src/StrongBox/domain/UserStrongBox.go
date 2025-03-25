package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserStrongBox struct {
	UsuarioID primitive.ObjectID `bson:"usuario_id,omitempty" json:"usuario_id"`
	Rol       string             `bson:"rol" json:"rol"`
	Pin       string             `bson:"pin" json:"pin"`
}