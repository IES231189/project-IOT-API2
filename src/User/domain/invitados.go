package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Invitado struct {
	InvitadoID primitive.ObjectID `bson:"invitado_id,omitempty" json:"invitado_id"`
	Nombre     string             `bson:"nombre" json:"nombre"`
	Pin        string             `bson:"pin" json:"pin"`
}
