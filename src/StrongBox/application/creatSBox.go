
package application

import (
	"api/src/StrongBox/application/services"
	"api/src/StrongBox/domain"
	"log"
	"go.mongodb.org/mongo-driver/bson/primitive" 
)

type CrearStrongBoxService struct {
	repository     domain.StrongBoxRepository
	eventService   *services.EventService
}

func NewCreateStrongBoxService(repository domain.StrongBoxRepository, eventService *services.EventService) *CrearStrongBoxService {
	return &CrearStrongBoxService{
		repository:   repository,
		eventService: eventService,
	}
}

func (s *CrearStrongBoxService) Execute(box *domain.StrongBox) (string, error) {
	strongBoxID, err := s.repository.CreateStrongBox(box)
	if err != nil {
		return "", err
	}


	objectID, err := primitive.ObjectIDFromHex(strongBoxID)
	if err != nil {
		log.Printf("Error converting ID to ObjectID: %v", err)
		return strongBoxID, nil 
	}
	box.ID = objectID

	
	if err := s.eventService.PublishStrongBoxCreatedEvent(box); err != nil {
		log.Printf("Error publishing event: %v", err)
	}

	return strongBoxID, nil
}