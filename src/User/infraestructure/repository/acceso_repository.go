
package repository

import (
    "api/src/User/domain/entities"
    "context"
    "fmt"
    "log"
    "time"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

type MongoAccesoRepository struct {
    collection *mongo.Collection
}

func NewMongoAccesoRepository(client *mongo.Client) *MongoAccesoRepository {
    collection := client.Database("base_iot_db").Collection("accesos_exitosos")
    return &MongoAccesoRepository{collection: collection}
}

func (r *MongoAccesoRepository) GuardarAccesoExitoso(pin string, usuario *entities.User, estado string, fecha time.Time) error {
    acceso := entities.AccesoExitoso{
        Pin:    pin, // Usar el pin directamente como string
        Usuario: usuario,
        Estado: estado,
        Fecha:  fecha,
    }

    _, err := r.collection.InsertOne(context.TODO(), acceso)
    if err != nil {
        log.Printf("Error al guardar acceso exitoso: %v", err)
        return fmt.Errorf("error al guardar acceso exitoso")
    }

    log.Printf("Acceso exitoso guardado: %+v", acceso) // Agregar log para depuración
    return nil
}

func (r *MongoAccesoRepository) ObtenerAccesosExitosos() ([]entities.AccesoExitoso, error) {
    var accesos []entities.AccesoExitoso

    cursor, err := r.collection.Find(context.TODO(), bson.M{})
    if err != nil {
        log.Printf("Error al obtener accesos exitosos: %v", err)
        return nil, fmt.Errorf("error al obtener accesos exitosos")
    }
    defer cursor.Close(context.TODO())

    for cursor.Next(context.TODO()) {
        var acceso entities.AccesoExitoso
        if err := cursor.Decode(&acceso); err != nil {
            log.Printf("Error al decodificar acceso exitoso: %v", err)
            return nil, fmt.Errorf("error al decodificar acceso exitoso")
        }
        accesos = append(accesos, acceso)
    }

    if err := cursor.Err(); err != nil {
        log.Printf("Error en el cursor: %v", err)
        return nil, fmt.Errorf("error en el cursor")
    }

    log.Printf("Accesos exitosos recuperados: %+v", accesos) // Agregar log para depuración
    return accesos, nil
}