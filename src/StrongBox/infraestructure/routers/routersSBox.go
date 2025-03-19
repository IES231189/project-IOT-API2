package routes

import (
	"api/src/StrongBox/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

// SetupStrongBoxRoutes configura las rutas para la gestión de cajas fuertes (StrongBox)
func SetupStrongBoxRoutes(r *gin.Engine) {
	// Rutas para la gestión de cajas fuertes
	r.POST("/strongbox", controllers.CrearStrongBoxHandler)                
	r.GET("/strongboxes", controllers.ObtenerStrongBoxByIDHandler)              
	r.GET("/strongboxes/:id", controllers.ObtenerStrongBoxByIDHandler)       
	r.DELETE("/strongboxes/:id", controllers.EliminarStrongBoxHandler)         
}
