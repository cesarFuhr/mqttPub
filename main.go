package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := bootstrap()
	if err != nil {
		fmt.Println(err)
	}

	publisher := Publisher{}
	publisher.Connect(os.Getenv("MQTT_BROKER_URL"))

	ticker := time.NewTicker(5 * time.Second)
	exit := pleaseLeave()

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

func bootstrap() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
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
