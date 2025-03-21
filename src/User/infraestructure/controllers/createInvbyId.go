package controllers

import (
	"api/src/User/domain"
	"api/src/User/infraestructure"
	"net/http"
    "github.com/gin-gonic/gin"
	
)

func AddGuestHandler(c *gin.Context) {
    userID := c.Param("userID")
    var guest domain.Invitado

    if err := c.ShouldBindJSON(&guest); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Error en el formato del JSON"})
        return
    }

    repo := infraestructure.NewMongoUserRepository()
    err := repo.AddGuest(userID, guest)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Invitado agregado correctamente"})
}
