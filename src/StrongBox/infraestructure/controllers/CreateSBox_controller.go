package controllers

import (
	"api/src/StrongBox/application"
	"api/src/StrongBox/application/services"
	"api/src/StrongBox/domain"
	"api/src/StrongBox/infraestructure"
	"api/src/StrongBox/infraestructure/rabbit" 
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CrearStrongBoxHandler(c *gin.Context) {
	repo := infraestructure.NewMongoStrongBoxRepository()
	
	rabbitMQURI := os.Getenv("RABBITMQ_URI")
	if rabbitMQURI == "" {
		rabbitMQURI = "amqp://admin:admin@100.28.135.252:5672/"
	}
	
	rabbitClient, err := rabbit.NewRabbitMQClient(rabbitMQURI)
	if err != nil {
		http.Error(c.Writer, "Error al conectar con RabbitMQ", http.StatusInternalServerError)
		return
	}
	defer rabbitClient.Close()
	
	
	eventService := services.NewEventService(rabbitClient)
	crearStrongBoxUC := application.NewCreateStrongBoxService(repo, eventService)

	var strongBox domain.StrongBox
	if err := json.NewDecoder(c.Request.Body).Decode(&strongBox); err != nil {
		http.Error(c.Writer, "Error al decodificar JSON", http.StatusBadRequest)
		return
	}

	
	if strongBox.UsuarioID == primitive.NilObjectID {
		http.Error(c.Writer, "El ID del usuario es obligatorio", http.StatusBadRequest)
		return
	}

	strongBoxID, err := crearStrongBoxUC.Execute(&strongBox)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":    "Caja fuerte creada exitosamente",
		"strongBoxID": strongBoxID,
	})
}