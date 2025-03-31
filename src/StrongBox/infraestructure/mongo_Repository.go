package infraestructure

import (
	//User "api/src/User/domain"           
	"api/src/StrongBox/domain" 
	"api/src/core"
	"context"
	"log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"fmt"

)

type MongoStrongBoxRepository struct {
	collection *mongo.Collection
	userCollection *mongo.Collection  // Agregamos la referencia a la colección de usuarios
}

// NewMongoStrongBoxRepository inicializa un nuevo repositorio de MongoDB para StrongBox
func NewMongoStrongBoxRepository() *MongoStrongBoxRepository {
	client := core.GetMongoClient()
	if client == nil {
		log.Fatal("No se pudo obtener el cliente de MongoDB")
	}
	collection := client.Database("base_iot_db").Collection("StrongBox")
	userCollection := client.Database("base_iot_db").Collection("usuarios") // Inicializamos la colección de usuarios
	return &MongoStrongBoxRepository{collection: collection, userCollection: userCollection}
}

func (r *MongoStrongBoxRepository) CreateStrongBox(strongBox *domain.StrongBox) (string, error) {
	if strongBox == nil {
		return "", fmt.Errorf("la caja fuerte no puede ser nil")
	}

	// Asegúrate de que el UsuarioID sea válido
	if strongBox.UsuarioID == primitive.NilObjectID {
		return "", fmt.Errorf("se requiere un UsuarioID válido")
	}

	// Inserta la caja fuerte en la base de datos
	result, err := r.collection.InsertOne(context.TODO(), strongBox)
	if err != nil {
		log.Printf("Error al crear la caja fuerte: %v", err)
		return "", err
	}

	objectID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("no se pudo convertir el ID a ObjectID")
	}
	return objectID.Hex(), nil
}


// DeleteStrongBox elimina una caja fuerte de la base de datos por ID
func (r *MongoStrongBoxRepository) DeleteStrongBox(ID string) error {
	objectID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		log.Printf("Error al convertir ID a ObjectID: %v", err)
		return fmt.Errorf("ID inválido")
	}

	_, err = r.collection.DeleteOne(context.TODO(), bson.M{"_id": objectID})
	if err != nil {
		log.Printf("Error al eliminar la caja fuerte: %v", err)
		return err
	}

	return nil
}

// GetAllStrongBoxes obtiene todas las cajas fuertes de la base de datos
func (r *MongoStrongBoxRepository) GetAllStrongBoxes() ([]domain.StrongBox, error) {
	var strongBoxes []domain.StrongBox

	cursor, err := r.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Printf("Error al obtener las cajas fuertes: %v", err)
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var strongBox domain.StrongBox
		if err := cursor.Decode(&strongBox); err != nil {
			log.Printf("Error al decodificar la caja fuerte: %v", err)
			return nil, err
		}
		strongBoxes = append(strongBoxes, strongBox)
	}

	if err := cursor.Err(); err != nil {
		log.Printf("Error en el cursor: %v", err)
		return nil, err
	}

	return strongBoxes, nil
}

// GetStrongBoxByID busca una caja fuerte por su ID
func (r *MongoStrongBoxRepository) GetStrongBoxByID(ID string) (*domain.StrongBox, error) {
	objectID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		log.Printf("Error al convertir ID a ObjectID: %v", err)
		return nil, fmt.Errorf("ID inválido")
	}

	var strongBox domain.StrongBox
	err = r.collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&strongBox)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("Caja fuerte con ID %s no encontrada", ID)
			return nil, nil
		}
		log.Printf("Error al buscar la caja fuerte por ID: %v", err)
		return nil, err
	}

	log.Printf("Caja fuerte encontrada: %+v", strongBox)
	return &strongBox, nil
}



func (r *MongoStrongBoxRepository) AddUserToStrongBox(boxID string, userID string) error {
    objectID, err := primitive.ObjectIDFromHex(boxID)
    if err != nil {
        return fmt.Errorf("ID de caja inválido")
    }

    userObjectID, err := primitive.ObjectIDFromHex(userID)
    if err != nil {
        return fmt.Errorf("ID de usuario inválido")
    }

    // Aquí asociamos el usuario a la caja fuerte
    update := bson.M{
        "$push": bson.M{
            "usuarios_con_acceso": domain.UserStrongBox{
                UsuarioID: userObjectID, // Cambiado de UserID a UsuarioID
            },
        },
    }

    _, err = r.collection.UpdateOne(context.TODO(), bson.M{"_id": objectID}, update)
    if err != nil {
        return fmt.Errorf("error al agregar usuario a la caja fuerte: %v", err)
    }
    return nil
}




// RemoveUserFromStrongBox elimina un usuario de la caja fuerte
func (r *MongoStrongBoxRepository) RemoveUserFromStrongBox(boxID string, userID string) error {
	objectBoxID, err := primitive.ObjectIDFromHex(boxID)
	if err != nil {
		log.Printf("Error al convertir boxID a ObjectID: %v", err)
		return fmt.Errorf("ID de la caja inválido")
	}

	objectUserID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Printf("Error al convertir userID a ObjectID: %v", err)
		return fmt.Errorf("ID de usuario inválido")
	}

	update := bson.M{
		"$pull": bson.M{
			"usuarios_con_acceso": bson.M{"usuario_id": objectUserID}, // Elimina el usuario del array
		},
	}

	_, err = r.collection.UpdateOne(context.TODO(), bson.M{"_id": objectBoxID}, update)
	if err != nil {
		log.Printf("Error al eliminar usuario de la caja: %v", err)
		return err
	}

	return nil
}
