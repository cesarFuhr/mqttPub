package main

import (
	"testing"
)

func TestRegister(t *testing.T) {
	var chanSlice []chan Message
	streamer := OBDStreamer{
		channels: chanSlice,
	}
	t.Run("Should return nil if the channel was registred", func(t *testing.T) {
		c := make(chan Message, 100)
		got := streamer.Register(c)

		if got != nil {
			t.Errorf("want %v, got %v", nil, got)
		}
	})
	t.Run("Should have the channel registred in the channels slice", func(t *testing.T) {
		c := make(chan Message, 100)
		streamer.Register(c)

		assertInsideChannels(t, streamer.channels, c)
	})
	t.Run("Registred channel", func(t *testing.T) {
		t.Run("should receive Messages on every change", func(t *testing.T) {
			want := Message{
				Topic: "param",
				Msg:   "value",
			}

			c := make(chan Message)
			streamer.Register(c)

			go streamer.event(want.Topic, want.Msg)

			got := <-c

			if got != want {
				t.Errorf("want %v, got %v", want, got)
			}
		})
	})
}

func TestStart(t *testing.T) {
	streamer := OBDStreamer{}
	t.Run("Should reutrn nil if succesfuly started", func(t *testing.T) {
		got := streamer.Start()

		if got != nil {
			t.Errorf("want %v, got %v", nil, got)
		}
	})
}

func assertInsideChannels(t *testing.T, a []chan Message, want chan Message) {
	t.Helper()
	has := false
	for _, v := range a {
		if v == want {
			has = true
		}
	}
	if !has {
		t.Errorf("Did not found: %v, of type %T in %v", want, want, a)
	}
}
