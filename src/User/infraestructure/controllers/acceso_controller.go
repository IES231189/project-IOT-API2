package controllers

import (
    "api/src/User/infraestructure/repository"
    "github.com/gin-gonic/gin"
    "net/http"
)

var accesoRepo *repository.MongoAccesoRepository

// InitAccesoController inicializa el repositorio de accesos exitosos
func InitAccesoController(repo *repository.MongoAccesoRepository) {
    accesoRepo = repo
}

// GetAccesosExitososHandler devuelve la lista de accesos exitosos
func GetAccesosExitososHandler(c *gin.Context) {
    if accesoRepo == nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Repositorio no inicializado"})
        return
    }

    accesos, err := accesoRepo.ObtenerAccesosExitosos()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, accesos)
}