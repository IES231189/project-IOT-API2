package rabbit

import (
	"log"
	"time"
	"errors"
	"github.com/streadway/amqp"
)

type RabbitMQClient struct {
	conn         *amqp.Connection
	channel      *amqp.Channel
	confirmChan  chan amqp.Confirmation
	done         chan bool
	amqpURI      string
	reconnectDelay time.Duration
}

func NewRabbitMQClient(amqpURI string) (*RabbitMQClient, error) {
	client := &RabbitMQClient{
		amqpURI:       amqpURI,
		reconnectDelay: 5 * time.Second,
		done:          make(chan bool),
	}

	if err := client.connect(); err != nil {
		return nil, err
	}

	go client.handleReconnect()

	return client, nil
}

func (c *RabbitMQClient) connect() error {
	var err error

	c.conn, err = amqp.Dial(c.amqpURI)
	if err != nil {
		return err
	}

	c.channel, err = c.conn.Channel()
	if err != nil {
		c.conn.Close()
		return err
	}

	// Habilitar confirmaciones de publicación
	if err := c.channel.Confirm(false); err != nil {
		c.channel.Close()
		c.conn.Close()
		return err
	}

	c.confirmChan = c.channel.NotifyPublish(make(chan amqp.Confirmation, 1))

	log.Println("Conexión a RabbitMQ establecida exitosamente")
	return nil
}

func (c *RabbitMQClient) handleReconnect() {
	for {
		select {
		case <-c.done:
			return
		case <-c.conn.NotifyClose(make(chan *amqp.Error)):
			log.Println("Conexión perdida, intentando reconectar...")
			
			// Esperar antes de reconectar
			time.Sleep(c.reconnectDelay)

			if err := c.connect(); err != nil {
				log.Printf("Error al reconectar: %v", err)
				continue
			}
		}
	}
}

func (c *RabbitMQClient) Publish(queueName string, body []byte) error {
	_, err := c.channel.QueueDeclare(
		queueName,
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return err
	}

	err = c.channel.Publish(
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent, // Mensaje persistente
			ContentType: "application/json",
			Body:        body,
			Timestamp:   time.Now(),
		})

	if err != nil {
		return err
	}

	// Esperar confirmación
	select {
	case confirmed := <-c.confirmChan:
		if !confirmed.Ack {
			return errors.New("mensaje no confirmado por RabbitMQ")
		}
		return nil
	case <-time.After(5 * time.Second):
		return errors.New("tiempo de espera para confirmación agotado")
	}
}

func (c *RabbitMQClient) Close() error {
	close(c.done)
	
	if c.channel != nil {
		if err := c.channel.Close(); err != nil {
			log.Printf("Error al cerrar canal: %v", err)
		}
	}
	if c.conn != nil {
		if err := c.conn.Close(); err != nil {
			log.Printf("Error al cerrar conexión: %v", err)
		}
	}
	
	log.Println("Conexión RabbitMQ cerrada correctamente")
	return nil
}