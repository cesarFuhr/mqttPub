package main

import (
	"fmt"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
)

type Publisher struct {
	c MQTT.Client
}

type Message struct {
	Topic string
	Msg   string
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

func (p *Publisher) Publish(m Message) {
	token := p.c.Publish(m.Topic, 0, false, m.Msg)
	token.Wait()
}
