package services

import (
	"encoding/json"
	"errors"
	"log"
	"time"
	"fmt"
	"api/src/StrongBox/domain"
	"api/src/StrongBox/infraestructure/rabbit"
)

type EventService struct {
	rabbitMQClient *rabbit.RabbitMQClient
}

func NewEventService(rabbitMQClient *rabbit.RabbitMQClient) *EventService {
	return &EventService{
		rabbitMQClient: rabbitMQClient,
	}
}

func (s *EventService) PublishStrongBoxCreatedEvent(strongBox *domain.StrongBox) error {
	if strongBox == nil {
		return errors.New("strongBox cannot be nil")
	}

	event := map[string]interface{}{
		"event_type":      "StrongBoxCreated",
		"id":             strongBox.ID.Hex(),
		"nombre":         strongBox.Nombre,
		"estado":         strongBox.Estado,
		"usuario_id":     strongBox.UsuarioID.Hex(),
		"codigo_producto": strongBox.CodigoProducto,
		"timestamp":      time.Now().UTC().Format(time.RFC3339),
	}

	eventBytes, err := json.Marshal(event)
	if err != nil {
		return err
	}

	// Intentar publicación con reintentos
	maxRetries := 3
	for i := 0; i < maxRetries; i++ {
		err = s.rabbitMQClient.Publish("ValidarProducto", eventBytes)
		if err == nil {
			log.Printf("Evento publicado exitosamente: %+v", event)
			return nil
		}

		log.Printf("Intento %d: Error al publicar evento: %v", i+1, err)
		time.Sleep(2 * time.Second * time.Duration(i+1))
	}

	return fmt.Errorf("no se pudo publicar el evento después de %d intentos: %v", maxRetries, err)
}