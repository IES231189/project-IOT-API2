package infraestructure

import (
	User "api/src/User/domain"           
	"api/src/StrongBox/domain" 
	"api/src/core"
	"context"
	"log"
	"strings"
	 "strconv"

	"go.mongodb.org/mongo-driver/mongo/options"
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
	collection := client.Database("proyecto").Collection("StrongBox")
	userCollection := client.Database("proyecto").Collection("User") // Inicializamos la colección de usuarios
	return &MongoStrongBoxRepository{collection: collection, userCollection: userCollection}
}

// CreateStrongBox crea una nueva caja fuerte en la base de datos
func (r *MongoStrongBoxRepository) CreateStrongBox(strongBox *domain.StrongBox) (string, error) {
	if strongBox == nil {
		return "", fmt.Errorf("la caja fuerte no puede ser nil")
	}

	// Asignar un código autoincremental antes de insertar la caja
	codigoProducto, err := r.ObtenerSiguienteCodigoProducto()
	if err != nil {
		return "", fmt.Errorf("error generando código de producto: %v", err)
	}
	strongBox.CodigoProducto = codigoProducto

	// Insertar la caja en la base de datos
	result, err := r.collection.InsertOne(context.TODO(), strongBox)
	if err != nil {
		log.Printf("Error al crear la caja fuerte: %v", err)
		return "", err
	}

	// Convertir el ID insertado a ObjectID
	objectID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("no se pudo convertir el ID a ObjectID")
	}
	return objectID.Hex(), nil
}

// ObtenerSiguienteCodigoProducto obtiene el último código de producto y genera el siguiente
func (r *MongoStrongBoxRepository) ObtenerSiguienteCodigoProducto() (string, error) {
	// Variable para almacenar la última caja fuerte encontrada
	var ultimaCaja domain.StrongBox

	// Buscar la última caja fuerte ordenando por `codigo_producto` de mayor a menor
	err := r.collection.FindOne(
		context.TODO(),
		bson.M{}, // Sin filtros, buscamos cualquier caja fuerte existente
		options.FindOne().SetSort(bson.D{{"codigo_producto", -1}}),
	).Decode(&ultimaCaja)

	// Si no hay documentos, comenzamos desde "CP-001"
	if err == mongo.ErrNoDocuments {
		return "CP-001", nil
	} else if err != nil {
		log.Printf("Error al obtener la última caja fuerte: %v", err)
		return "", fmt.Errorf("error al obtener la última caja fuerte: %v", err)
	}

	// Extraer la parte numérica del código (ejemplo: "CP-045" → "045")
	parteNumerica := strings.TrimPrefix(ultimaCaja.CodigoProducto, "CP-")

	// Convertir a entero para incrementar
	ultimoNumero, err := strconv.Atoi(parteNumerica)
	if err != nil {
		return "CP-001", nil // Si falla la conversión, empezamos desde CP-001
	}

	// Generar el siguiente código con formato CP-XXX (ejemplo: "CP-046")
	nuevoCodigo := fmt.Sprintf("CP-%03d", ultimoNumero+1)
	return nuevoCodigo, nil
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
	// Convertir el boxID a ObjectID
	objectBoxID, err := primitive.ObjectIDFromHex(boxID)
	if err != nil {
		log.Printf("Error al convertir boxID a ObjectID: %v", err)
		return fmt.Errorf("ID de la caja inválido")
	}

	// Convertir el userID (invitado_id) a ObjectID
	objectUserID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Printf("Error al convertir userID a ObjectID: %v", err)
		return fmt.Errorf("ID del usuario inválido")
	}

	// Buscar en la colección "User" un usuario que tenga un invitado con ese ID
	var user User.User
	filter := bson.M{"invitados.invitado_id": objectUserID}
	err = r.userCollection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("Error: Usuario invitado no encontrado en ningún usuario.")
			return fmt.Errorf("usuario invitado no encontrado")
		}
		log.Printf("Error al buscar usuario invitado en la base de datos: %v", err)
		return err
	}

	// Buscar el usuario dentro del array de invitados
	var invitedUser User.Invitado
	for _, invitado := range user.MisInvitados {
		if invitado.InvitadoID == objectUserID {
			invitedUser = invitado
			break
		}
	}

	if invitedUser.InvitadoID.IsZero() {
		log.Println("Error: No se encontró el usuario invitado en la lista de invitados.")
		return fmt.Errorf("usuario invitado no encontrado en la lista")
	}

	// Crear el objeto de usuario para agregarlo a la caja fuerte
	userToAdd := &domain.UserStrongBox{
		UsuarioID: invitedUser.InvitadoID, // ID del invitado
		Nombre:    invitedUser.Nombre,     // Nombre del invitado
		Pin:       invitedUser.Pin,        // PIN del invitado
	}

	// Agregar el usuario al array `usuarios_con_acceso` en la caja fuerte
	update := bson.M{
		"$push": bson.M{"usuarios_con_acceso": userToAdd},
	}

	// Realizar la actualización en la base de datos
	_, err = r.collection.UpdateOne(context.TODO(), bson.M{"_id": objectBoxID}, update)
	if err != nil {
		log.Printf("Error al agregar usuario a la caja: %v", err)
		return err
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
