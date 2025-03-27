package core

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
	"errors"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/dgrijalva/jwt-go"
)

var (
	clientInstance *mongo.Client
	clientOnce     sync.Once
)

// GetMongoClient devuelve una instancia única de la conexión a MongoDB
func GetMongoClient() *mongo.Client {
	clientOnce.Do(func() {
		// Cargar variables de entorno desde el archivo .env
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error al cargar el archivo .env: %v", err)
		}

		// Obtener variables de entorno
		mongoURI := os.Getenv("MONGO_URI")
		if mongoURI == "" {
			log.Fatal("La variable de entorno MONGO_URI no está definida")
		}

		// Configurar opciones del cliente
		clientOptions := options.Client().ApplyURI(mongoURI)

		// Crear cliente
		client, err := mongo.NewClient(clientOptions)
		if err != nil {
			log.Fatalf("Error al crear el cliente de MongoDB: %v", err)
		}

		// Establecer contexto con timeout
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Conectar a MongoDB
		err = client.Connect(ctx)
		if err != nil {
			log.Fatalf("Error al conectar con MongoDB: %v", err)
		}

		// Verificar conexión
		err = client.Ping(ctx, nil)
		if err != nil {
			log.Fatalf("No se pudo conectar a MongoDB: %v", err)
		}

		fmt.Println("✅ Conexión exitosa a MongoDB")
		clientInstance = client
	})

	return clientInstance
}

// GetMongoDatabase devuelve la base de datos configurada en el .env
func GetMongoDatabase() *mongo.Database {
	dbName := os.Getenv("MONGO_DB")
	if dbName == "" {
		log.Fatal("La variable de entorno MONGO_DB no está definida")
	}

	client := GetMongoClient()
	return client.Database(dbName)
}

func GenerateJWT(userID string) (string, error) {
	// Obtén el secreto desde el entorno (asegúrate de que JWT_SECRET esté configurado en tu entorno)
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", errors.New("JWT_SECRET no está configurado")
	}

	// Creamos un token con claims (reclamos) que incluyen el ID del usuario
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 24).Unix(), // El token expira en 24 horas
	})

	// Firmamos el token
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}