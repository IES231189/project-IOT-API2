package domain

import (
	
	"go.mongodb.org/mongo-driver/bson/primitive"

)

type StrongBox struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre            string             `bson:"nombre" json:"nombre"`
	Estado            string             `bson:"estado" json:"estado"`
	CodigoProducto    string             `bson:"codigo_producto" json:"codigo_producto"`
	UsuariosConAcceso []UserStrongBox    `bson:"usuarios_con_acceso" json:"usuarios_con_acceso"`
}
