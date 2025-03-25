
package Mqtt

import (
    "api/src/User/application"
    "api/src/User/infraestructure/repository"
    "api/src/User/infraestructure/model"
    mqtt "github.com/eclipse/paho.mqtt.golang"
    "log"
    "time"
    "encoding/json"
)

type MqttService struct {
    client     mqtt.Client
    useCase    *application.ObtenerUsuarioPorPin
    accesoRepo *repository.MongoAccesoRepository
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
    service.subscribeToTopic() // Suscribirse al tópico MQTT
    return service
}

// subscribeToTopic se suscribe al tópico "caja/accesos"
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

// handleMessage procesa los mensajes recibidos
func (s *MqttService) handleMessage(payload string) {
    log.Printf("Mensaje recibido en MQTT: %s", payload)

    user, err := s.useCase.Ejecutar(payload)
    if err != nil {
        log.Printf("Error al buscar usuario: %v", err)
        s.publishResult(model.Resultado{Estado: "error", Pin: payload}, "caja/resultados")
        return
    }

    var estado string
    if user == nil {
        estado = "incorrecto"
        log.Println("Usuario no encontrado")
    } else {
        estado = "correcto"
        log.Printf("Usuario encontrado: %+v", user)
        // Almacenar acceso exitoso en MongoDB
        if err := s.accesoRepo.GuardarAccesoExitoso(payload, user, estado, time.Now()); err != nil {
            log.Printf("Error al guardar acceso exitoso: %v", err)
        }
    }

    resultado := model.Resultado{
        Estado: estado,
        Pin:    payload,
    }

    s.publishResult(resultado, "caja/resultados")
}

// publishResult publica el resultado en un tópico MQTT
func (s *MqttService) publishResult(result interface{}, topic string) {
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
        log.Printf("Resultado publicado en el tópico %s: %s", topic, string(payload))
    }
}