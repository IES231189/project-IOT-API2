
package repository

import (
    "api/src/User/domain"
    "context"
    "fmt"
    "log"
    "time"
    
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type MongoAccesoRepository struct {
    exitososCollection   *mongo.Collection
    incorrectosCollection *mongo.Collection
}

func NewMongoAccesoRepository(client *mongo.Client) *MongoAccesoRepository {
    db := client.Database("proyecto")
    return &MongoAccesoRepository{
        exitososCollection:   db.Collection("accesos_exitosos"),
        incorrectosCollection: db.Collection("accesos_incorrectos"),
    }
}

func (r *MongoAccesoRepository) GuardarAccesoExitoso(pin string, usuario string, estado string, fecha time.Time) error {
    acceso := domain.AccesoExitoso{
        Pin:     pin,
        Usuario: usuario,
        Estado:  estado,
        Fecha:   fecha,
    }

    _, err := r.exitososCollection.InsertOne(context.TODO(), acceso)
    if err != nil {
        log.Printf("Error al guardar acceso exitoso: %v", err)
        return fmt.Errorf("error al guardar acceso exitoso")
    }

    log.Printf("Acceso exitoso guardado: %+v", acceso)
    return nil
}

func (r *MongoAccesoRepository) ObtenerAccesosExitosos() ([]domain.AccesoExitoso, error) {
    var accesos []domain.AccesoExitoso

    opts := options.Find().
        SetSort(bson.D{{"fecha", -1}}).
        SetLimit(5)

    cursor, err := r.exitososCollection.Find(context.TODO(), bson.M{}, opts)
    if err != nil {
        log.Printf("Error al obtener accesos exitosos: %v", err)
        return nil, fmt.Errorf("error al obtener accesos exitosos")
    }
    defer cursor.Close(context.TODO())

    for cursor.Next(context.TODO()) {
        var acceso domain.AccesoExitoso
        if err := cursor.Decode(&acceso); err != nil {
            log.Printf("Error al decodificar acceso exitoso: %v", err)
            continue
        }
        accesos = append(accesos, acceso)
    }

    if err := cursor.Err(); err != nil {
        log.Printf("Error en el cursor: %v", err)
        return nil, fmt.Errorf("error en el cursor")
    }

    return accesos, nil
}

func (r *MongoAccesoRepository) GuardarAccesoIncorrecto(pin string, deviceId string, estado string, fecha time.Time) error {
    // Verificar intentos previos en los últimos 5 minutos
    filtro := bson.M{
        "deviceId": deviceId,
        "fecha": bson.M{"$gt": time.Now().Add(-5 * time.Minute)},
    }
    
    count, err := r.incorrectosCollection.CountDocuments(context.TODO(), filtro)
    if err != nil {
        log.Printf("Error al contar intentos previos: %v", err)
        count = 0
    }

    acceso := domain.AccesoIncorrecto{
        Pin:         pin,
        DeviceId:    deviceId,
        Estado:      estado,
        Fecha:       fecha,
        EsIntrusion: count >= 2, // Marcar como intrusión después de 3 intentos
        Intento:     int(count) + 1,
    }

    _, err = r.incorrectosCollection.InsertOne(context.TODO(), acceso)
    if err != nil {
        log.Printf("Error al guardar acceso incorrecto: %v", err)
        return fmt.Errorf("error al guardar acceso incorrecto")
    }

    log.Printf("Acceso incorrecto registrado - PIN: %s, Device: %s, Intento: %d, Intrusión: %v", 
        pin, deviceId, acceso.Intento, acceso.EsIntrusion)
    
    return nil
}

func (r *MongoAccesoRepository) ObtenerAccesosIncorrectos() ([]domain.AccesoIncorrecto, error) {
    var accesos []domain.AccesoIncorrecto

    opts := options.Find().
        SetSort(bson.D{{"fecha", -1}}).
        SetLimit(10)

    cursor, err := r.incorrectosCollection.Find(context.TODO(), bson.M{}, opts)
    if err != nil {
        log.Printf("Error al obtener accesos incorrectos: %v", err)
        return nil, fmt.Errorf("error al obtener accesos incorrectos")
    }
    defer cursor.Close(context.TODO())

    for cursor.Next(context.TODO()) {
        var acceso domain.AccesoIncorrecto
        if err := cursor.Decode(&acceso); err != nil {
            log.Printf("Error al decodificar acceso incorrecto: %v", err)
            continue
        }
        accesos = append(accesos, acceso)
    }

    if err := cursor.Err(); err != nil {
        log.Printf("Error en el cursor: %v", err)
        return nil, fmt.Errorf("error en el cursor")
    }

    return accesos, nil
}