package infraestructure

import (
	"api/src/User/domain"
	"api/src/core"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
	collection := client.Database("proyecto").Collection("Usuarios")
	return &MongoUserRepository{collection: collection}
}

// CreateUser crea un nuevo usuario en la base de datos
func (r *MongoUserRepository) CreateUser(user *domain.User) (string, error) {
	if user == nil {
		return "", fmt.Errorf("el usuario no puede ser nil")
	}

	// Asigna un nuevo ObjectID al usuario si aún no tiene uno
	user.ID = primitive.NewObjectID()

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

// UpdateUser actualiza los datos de un usuario por su ID
func (r *MongoUserRepository) UpdateUser(ID string, updatedUser *domain.User) error {
	if updatedUser == nil {
		return fmt.Errorf("los datos del usuario no pueden ser nil")
	}

	objectID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		log.Printf("Error al convertir ID a ObjectID: %v", err)
		return fmt.Errorf("ID inválido")
	}

	update := bson.M{
		"$set": bson.M{
			"nombre": updatedUser.Nombre,
			"correo": updatedUser.Correo,
			"pin":    updatedUser.Pin,
		},
	}

	_, err = r.collection.UpdateOne(context.TODO(), bson.M{"_id": objectID}, update)
	if err != nil {
		log.Printf("Error al actualizar usuario: %v", err)
		return err
	}

	return nil
}

// AddGuest agrega un nuevo invitado a la lista de mis_invitados de un usuario
func (r *MongoUserRepository) AddGuest(userID string, guest domain.Invitado) error {
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Printf("Error al convertir userID a ObjectID: %v", err)
		return fmt.Errorf("ID inválido")
	}

	guest.InvitadoID = primitive.NewObjectID() // Generamos un nuevo ID para el invitado

	update := bson.M{
		"$push": bson.M{"MisInvitados": guest}, // Agrega el nuevo invitado a la lista
	}

	_, err = r.collection.UpdateOne(context.TODO(), bson.M{"_id": objectID}, update)
	if err != nil {
		log.Printf("Error al agregar invitado: %v", err)
		return err
	}

	return nil
}
