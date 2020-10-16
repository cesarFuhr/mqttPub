package main

import (
	"time"
)

type OBDStreamer struct {
	channels []chan Message
}

func (s *OBDStreamer) Register(c chan Message) error {
	s.channels = append(s.channels, c)
	return nil
}

func (s *OBDStreamer) Start() error {
	ticker := time.NewTicker(time.Second)
	go func(ticker *time.Ticker) {
		for {
			<-ticker.C
			s.event("testTopic", "testing")
		}
	}(ticker)
	return nil
}

func (s *OBDStreamer) event(t, v string) {
	message := Message{
		Topic: t,
		Msg:   v,
	}
	for _, v := range s.channels {
		v <- message
	}
}
