package main

type OBDStreamer struct {
	channels []chan Message
}

func (s *OBDStreamer) Register(c chan Message) error {
	s.channels = append(s.channels, c)
	return nil
}

func (s *OBDStreamer) Start() error {
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
