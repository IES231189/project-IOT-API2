package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre      string             `bson:"nombre" json:"nombre"`
	Correo      string             `bson:"correo" json:"correo"`
	Contraseña  string             `bson:"contraseña" json:"contraseña"`  
	Pin         string             `bson:"pin" json:"pin"`
	MisInvitados []Invitado        `bson:"misinvitados" json:"mis_invitados"`
}
