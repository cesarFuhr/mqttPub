package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	exit := pleaseLeave()

	publisher := Publisher{}
	publisher.Connect("tcp://0.0.0.0:1883")

	ticker := time.NewTicker(5 * time.Second)

	for {
		select {
		case <-exit:
			fmt.Println("Goodbye...")
			publisher.c.Disconnect(250)
			return
		case <-ticker.C:
			fmt.Println("Published!")
			publisher.Publish("testTopic", "testing")
		}
	}
}

func pleaseLeave() chan struct{} {
	sigs := make(chan os.Signal, 1)
	done := make(chan struct{}, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		close(done)
	}()
	return done
}
