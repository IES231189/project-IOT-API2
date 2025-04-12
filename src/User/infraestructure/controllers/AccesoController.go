// controllers/acceso_controller.go
package controllers

import (
    "api/src/User/infraestructure/repository"
    "github.com/gin-gonic/gin"
    "net/http"
)

var accesoRepo *repository.MongoAccesoRepository

func InitAccesoController(repo *repository.MongoAccesoRepository) {
    accesoRepo = repo
}

func GetAccesosExitososHandler(c *gin.Context) {
    if accesoRepo == nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Repositorio no inicializado"})
        return
    }

    accesos, err := accesoRepo.ObtenerAccesosExitosos()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
            "message": "Error al obtener accesos exitosos",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "total": len(accesos),
        "data": accesos,
    })
}

func GetAccesosIncorrectosHandler(c *gin.Context) {
    if accesoRepo == nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Repositorio no inicializado"})
        return
    }

    accesos, err := accesoRepo.ObtenerAccesosIncorrectos()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
            "message": "Error al obtener accesos incorrectos",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "total": len(accesos),
        "data": accesos,
    })
}