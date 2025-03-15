package core

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
