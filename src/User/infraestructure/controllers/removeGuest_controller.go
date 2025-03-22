package controllers

import (
	"api/src/User/application"
	"api/src/User/infraestructure"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RemoveGuestHandler maneja la eliminación de un invitado de un usuario
func RemoveGuestHandler(c *gin.Context) {
	log.Println("Método recibido: DELETE")

	// Obtener el ID del usuario y el ID del invitado desde la URL
	userID := c.Param("userID")
	guestID := c.Param("guestID")

	if userID == "" || guestID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userID y guestID son obligatorios"})
		return
	}

	// Convertir guestID a primitive.ObjectID
	guestObjectID, err := primitive.ObjectIDFromHex(guestID)
	if err != nil {
		log.Printf("Error al convertir guestID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "guestID no válido"})
		return
	}

	// Convertir guestObjectID a string (Hex)
	guestIDString := guestObjectID.Hex()

	// Crear una instancia del repositorio para el caso de uso
	repo := infraestructure.NewMongoUserRepository()

	// Crear el caso de uso para eliminar el invitado
	useCase := application.NewRemoverInvitado(repo)

	// Llamar al caso de uso para eliminar el invitado
	err = useCase.Ejecutar(userID, guestIDString)
	if err != nil {
		log.Printf("Error al eliminar el invitado: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar el invitado"})
		return
	}

	log.Println("Invitado eliminado correctamente")
	c.JSON(http.StatusOK, gin.H{"message": "Invitado eliminado correctamente"})
}
