package routes

import (
	"api/src/StrongBox/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

// SetupStrongBoxRoutes configura las rutas para la gestión de cajas fuertes (StrongBox)
func SetupStrongBoxRoutes(r *gin.Engine) {
	// Cambié AddUserToStrongBox por AddInvitadoToStrongBox
	r.POST("/strongbox/:boxID/invitados/:guestID", controllers.AddUserToStrongBox) // Nueva ruta para agregar invitados
	r.POST("/strongbox", controllers.CrearStrongBoxHandler)                
	r.GET("/strongboxes", controllers.ObtenerTodasLasCajasFuertesHandler)              
	r.GET("/strongboxes/:id", controllers.ObtenerStrongBoxByIDHandler)       
	r.DELETE("/strongboxes/:id", controllers.EliminarStrongBoxHandler)
	r.DELETE("/strongbox/:boxID/users/:userID", controllers.RemoveUserFromStrongBoxHandler)
}
