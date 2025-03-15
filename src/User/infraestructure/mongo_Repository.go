package infraestructure

import (
	"api/src/User/domain"
	"api/src/core"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	
)

type MongoUserRepository struct {
	collection *mongo.Collection
}

func NewMongoUserRepository() *MongoUserRepository {
	client := core.GetMongoClient()
	collection := client.Database("proyecto").Collection("Usuarios")
	return &MongoUserRepository{collection: collection}
}

// CreateUser crea un nuevo usuario en la base de datos
func (r *MongoUserRepository) CreateUser(user *domain.User) (string, error) {
	result, err := r.collection.InsertOne(context.TODO(), user)
	if err != nil {
		return "", err
	}
	// Retorna el ID del nuevo usuario
	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

// DeleteUser elimina un usuario de la base de datos por ID
func (r *MongoUserRepository) DeleteUser(ID string) error {
	// Convierte el ID de string a ObjectID
	id, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = r.collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	return err
}

// GetAllUsers obtiene todos los usuarios de la base de datos
func (r *MongoUserRepository) GetAllUsers() ([]domain.User, error) {
	var users []domain.User
	cursor, err := r.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var user domain.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// GetUserByPin busca un usuario por su PIN
func (r *MongoUserRepository) GetUserByPin(Pin string) (*domain.User, error) {
	var user domain.User
	err := r.collection.FindOne(context.TODO(), bson.M{"pin": Pin}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // No se encontró el usuario
		}
		return nil, err // Otro tipo de error
	}
	return &user, nil
}
