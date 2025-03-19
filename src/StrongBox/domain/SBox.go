package domain

import "time"

type StrongBox struct {
	ID             string    `bson:"_id,omitempty" json:"id"`
	Nombre         string    `bson:"nombre" json:"nombre"`
	Estado         string    `bson:"estado" json:"estado"`
	UltimaActividad time.Time `bson:"ultima_actividad" json:"ultima_actividad"`
}
