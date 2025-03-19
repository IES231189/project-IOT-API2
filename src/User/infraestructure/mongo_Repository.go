package infraestructure

import (
	"api/src/User/domain"
	"api/src/core"
	"context"
	"log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"fmt"
)

type MongoUserRepository struct {
	collection *mongo.Collection
}

// NewMongoUserRepository inicializa un nuevo repositorio de MongoDB
func NewMongoUserRepository() *MongoUserRepository {
	client := core.GetMongoClient()
	if client == nil {
		log.Fatal("No se pudo obtener el cliente de MongoDB")
	}
	collection := client.Database("base_iot_db").Collection("usuarios")
	return &MongoUserRepository{collection: collection}
}

// CreateUser crea un nuevo usuario en la base de datos
func (r *MongoUserRepository) CreateUser(user *domain.User) (string, error) {
	if user == nil {
		return "", fmt.Errorf("el usuario no puede ser nil")
	}

	result, err := r.collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Printf("Error al crear usuario: %v", err)
		return "", err
	}

	// Retorna el ID del nuevo usuario
	objectID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("no se pudo convertir el ID a ObjectID")
	}
	return objectID.Hex(), nil
}

// DeleteUser elimina un usuario de la base de datos por ID
func (r *MongoUserRepository) DeleteUser(ID string) error {
	objectID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		log.Printf("Error al convertir ID a ObjectID: %v", err)
		return fmt.Errorf("ID inválido")
	}

	_, err = r.collection.DeleteOne(context.TODO(), bson.M{"_id": objectID})
	if err != nil {
		log.Printf("Error al eliminar usuario: %v", err)
		return err
	}

	return nil
}

// GetAllUsers obtiene todos los usuarios de la base de datos
func (r *MongoUserRepository) GetAllUsers() ([]domain.User, error) {
	var users []domain.User

	cursor, err := r.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Printf("Error al obtener usuarios: %v", err)
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var user domain.User
		if err := cursor.Decode(&user); err != nil {
			log.Printf("Error al decodificar usuario: %v", err)
			return nil, err
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		log.Printf("Error en el cursor: %v", err)
		return nil, err
	}

	return users, nil
}

// GetUserByPin busca un usuario por su PIN
func (r *MongoUserRepository) GetUserByPin(Pin string) (*domain.User, error) {
	if Pin == "" {
		return nil, fmt.Errorf("el PIN no puede estar vacío")
	}

	var user domain.User
	err := r.collection.FindOne(context.TODO(), bson.M{"Pin": Pin}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("Usuario con PIN %s no encontrado", Pin)
			return nil, nil
		}
		log.Printf("Error al buscar usuario por PIN: %v", err)
		return nil, err
	}

	log.Printf("Usuario encontrado: %+v", user)
	return &user, nil
}