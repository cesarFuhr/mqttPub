package main

import (
	"fmt"

	"github.com/google/uuid"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)


type Publisher struct {
	c MQTT.Client
}

func (p *Publisher) Connect(brokerUrl string) {
	opts := MQTT.NewClientOptions().AddBroker(brokerUrl)
	id, _ := uuid.NewUUID()
	clientName := fmt.Sprintf("Pub-%s", id)
	opts.SetClientID(clientName)

	p.c = MQTT.NewClient(opts)
	if token := p.c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}


func (p *Publisher) Publish(topic, msg string) {
	token := p.c.Publish(topic, 0, false, msg)
	token.Wait()
}
