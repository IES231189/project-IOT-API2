package Mqtt

import (
	"api/src/User/application"
	"api/src/User/infraestructure/model"
	mqtt "github.com/eclipse/paho.mqtt.golang"

	"log"
	"encoding/json"
)

// MqttService maneja la conexión y suscripción a MQTT
type MqttService struct {
	client  mqtt.Client
	useCase *application.ObtenerUsuarioPorPin
}

// creamos la instancia de mqtt service
func NewMqttService(useCase *application.ObtenerUsuarioPorPin) *MqttService {
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://54.236.196.93:1883") 
	opts.SetClientID("api_server")
	opts.SetUsername("dev1")  
	opts.SetPassword("devpublisher") 


	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Error al conectar con MQTT: %v", token.Error())
	}

	service := &MqttService{client: client, useCase: useCase}
	service.subscribeToTopic()
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
		s.publishResult(model.Resultado{Estado:"error" , Pin:payload}, "caja/resultados")
		return
	}

	var estado string

	if user == nil {
		 estado = "incorrecto"
		log.Println("Usuario no encontrado")
		s.publishResult(model.Resultado{Estado:estado , Pin:payload} , "caja/resultados")
		
	} else {
		estado = "correcto"
		log.Printf("Usuario encontrado: %+v", user)
		s.publishResult(model.Resultado{Estado:estado , Pin:payload} , "caja/resultados")
	}
}

func (s *MqttService) publishResult(result interface{} , topic string ){
payload , err := json.Marshal(result)
if err != nil{
		log.Printf("Error al conevertir el resultado a JSON:%v" , err)
		return
	}


	token := s.client.Publish(topic ,1 , false , payload)

	token.Wait()
	if token.Error() != nil{
		log.Printf("Error al publicar el topico %s:%v", topic , token)
	}else{
		log.Printf("Resiltado publicado en el topico %s:%s", topic , string(payload))
	}

}

