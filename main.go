package main

import (
	//"api/src/User/application"
	//"api/src/User/infraestructure"
	//"api/src/User/infraestructure/Mqtt"
	userRoutes "api/src/User/infraestructure/routers" 
	strongBoxRoutes "api/src/StrongBox/infraestructure/routers" 
	"api/src/core"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// Función para verificar la conexión a MongoDB
func handler(c *gin.Context) {
	client := core.GetMongoClient()
	databases, err := client.ListDatabaseNames(c, nil)
	if err != nil {
		log.Println("Error al obtener bases de datos:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al conectar con la base de datos"})
		return
	}
	log.Printf("Bases de datos disponibles: %v", databases)
	c.JSON(http.StatusOK, gin.H{"message": "Conexión exitosa a MongoDB"})
}

func main() {
	// Inicialización de Gin
	r := gin.Default()

	// Configuración CORS para permitir peticiones desde el frontend
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Configurar rutas de usuarios
	userRoutes.SetupRoutes(r) // Usamos el nombre 'userRoutes' para las rutas de usuario

	// Configurar rutas de StrongBox
	strongBoxRoutes.SetupStrongBoxRoutes(r) // Usamos el nombre 'strongBoxRoutes' para las rutas de StrongBox

	// Ruta de prueba para verificar conexión con MongoDB
	r.GET("/testMongo", handler)

	// Inicializar la conexión MQTT y la suscripción
	//repo := infraestructure.NewMongoUserRepository()
	//useCase := application.NewObtenerUsuarioPorPin(repo)
	//Mqtt.NewMqttService(useCase) // Inicia la suscripción a MQTT

	// Iniciar servidor en el puerto 8080
	log.Println("Servidor escuchando en el puerto 8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}

