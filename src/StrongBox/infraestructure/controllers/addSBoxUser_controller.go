package controllers

import (
    "api/src/StrongBox/application"
    "api/src/StrongBox/infraestructure"
    "github.com/gin-gonic/gin"
    "net/http"
    "log"
)

func AddUserToStrongBox(c *gin.Context) {
    // Crear el repositorio para la caja fuerte
    repo := infraestructure.NewMongoStrongBoxRepository() // Crea el repositorio

    // Crear el servicio para agregar un invitado a la caja fuerte
    addUserUC := application.NewAddUserToStrongBoxService(repo) // Crea el servicio con el repo

    // Obtener los parámetros de la URL: boxID y guestID
    boxID := c.Param("boxID")
    guestID := c.Param("guestID")

    // Validar que los IDs sean válidos
    if len(boxID) == 0 || len(guestID) == 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID de caja fuerte o invitado inválido"})
        return
    }

    // Llamar al servicio para agregar el invitado a la caja fuerte
    err := addUserUC.Execute(boxID, guestID)
    if err != nil {
        // Si ocurre un error, devolver un mensaje con el error
        log.Println("Error al agregar invitado:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Si todo fue bien, devolver una respuesta exitosa
    c.JSON(http.StatusOK, gin.H{"message": "Invitado agregado exitosamente a la caja fuerte"})
}
