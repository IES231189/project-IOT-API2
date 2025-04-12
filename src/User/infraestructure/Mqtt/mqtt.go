package Mqtt

import (
	"api/src/User/application"
	"api/src/User/infraestructure/model"
	"api/src/User/infraestructure/repository"
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"time"
)

type MqttService struct {
	client     mqtt.Client
	useCase    *application.ObtenerUsuarioPorPin
	accesoRepo *repository.MongoAccesoRepository
}

// Estructura para el mensaje entrante
type MensajeAcceso struct {
	DeviceId  string `json:"deviceId"`
	Pin       string `json:"pin"`
	Timestamp int64  `json:"timestamp,omitempty"`
}

func NewMqttService(useCase *application.ObtenerUsuarioPorPin, accesoRepo *repository.MongoAccesoRepository) *MqttService {
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://54.236.196.93:1883")
	opts.SetClientID("api_server")
	opts.SetUsername("dev1")
	opts.SetPassword("devpublisher")

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Error al conectar con MQTT: %v", token.Error())
	}

	service := &MqttService{client: client, useCase: useCase, accesoRepo: accesoRepo}
	service.subscribeToTopic()
	return service
}

func (s *MqttService) subscribeToTopic() {
	topic := "caja/accesos"
	token := s.client.Subscribe(topic, 1, func(client mqtt.Client, msg mqtt.Message) {
		s.handleMessage(string(msg.Payload()))
	})

	token.Wait()
	if token.Error() != nil {
		log.Fatalf("Error al suscribirse al tópico: %v", token.Error())
	}
	log.Printf("Suscrito al tópico: %s", topic)
}

func (s *MqttService) handleMessage(payload string) {
	log.Printf("Mensaje recibido en MQTT: %s", payload)

	var mensaje MensajeAcceso
	var pin string
	var deviceId string

	err := json.Unmarshal([]byte(payload), &mensaje)
	if err != nil {

		pin = payload
		deviceId = "unknown"
	} else {
		// Extraer datos del JSON
		pin = mensaje.Pin
		deviceId = mensaje.DeviceId
	}

	// Validar el PIN
	user, invitado, err := s.useCase.Ejecutar(pin)
	if err != nil {
		log.Printf("Error al buscar usuario: %v", err)

		if err := s.accesoRepo.GuardarAccesoIncorrecto(pin, deviceId, "error_sistema", time.Now()); err != nil {
            log.Printf("Error al registrar fallo del sistema: %v", err)
        }

		s.publishResult(model.Resultado{
			Estado:   "error",
			Pin:      pin,
			DeviceId: deviceId,
		}, "caja/resultados")
		return
	}

	// Determinar resultado
	estado := "incorrecto"
	nombreUsuario := ""
	if user != nil {
		estado = "correcto"
		if invitado != nil {
			nombreUsuario = invitado.Nombre
			log.Printf("Invitado encontrado: %+v", invitado)
		} else {
			nombreUsuario = user.Nombre
			log.Printf("Usuario principal encontrado: %+v", user)
		}

		if err := s.accesoRepo.GuardarAccesoExitoso(pin, nombreUsuario, estado, time.Now()); err != nil {
			log.Printf("Error al guardar acceso exitoso: %v", err)
		}
	} else {
		log.Println("Usuario no encontrado para PIN:", pin)
		if err := s.accesoRepo.GuardarAccesoIncorrecto(pin, deviceId, "intrusión", time.Now()); err != nil {
			log.Printf("Error al guardar acceso incorrecto: %v", err)
		}
	}

	// Publicar respuesta
	s.publishResult(model.Resultado{
		Estado:   estado,
		Pin:      pin,
		DeviceId: deviceId,
	}, "caja/resultados")
}

func (s *MqttService) publishResult(result model.Resultado, topic string) {
	payload, err := json.Marshal(result)
	if err != nil {
		log.Printf("Error al convertir el resultado a JSON: %v", err)
		return
	}

	token := s.client.Publish(topic, 1, false, payload)
	token.Wait()

	if token.Error() != nil {
		log.Printf("Error al publicar en el tópico %s: %v", topic, token.Error())
	} else {
		log.Printf("Resultado publicado en %s: %s", topic, string(payload))
	}
}
