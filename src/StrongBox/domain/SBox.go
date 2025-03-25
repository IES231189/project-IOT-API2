package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StrongBox struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre            string             `bson:"nombre" json:"nombre"`
	Estado            string             `bson:"estado" json:"estado"`
	UltimaActividad   primitive.DateTime `bson:"ultima_actividad" json:"ultima_actividad"`
	UsuariosConAcceso []UserStrongBox    `bson:"usuarios_con_acceso" json:"usuarios_con_acceso"`
}

