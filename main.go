package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/joho/godotenv"
)

func main() {
	if err := bootstrap(); err != nil {
		fmt.Println(err)
	}

	p := Publisher{}
	run(p)
}

func run(pub Publisher) {
	pub.Connect(os.Getenv("MQTT_BROKER_URL"))

	exit := pleaseLeave()
	finish := make(chan struct{})

	go startPublishing(pub, exit, finish)
	<-finish
}

func startPublishing(pub Publisher, exit chan struct{}, finish chan struct{}) {
	bufSize, _ := strconv.Atoi(os.Getenv("OBD_STREAM_BUFFER_SIZE"))
	messages := make(chan Message, bufSize)
	streamer := OBDStreamer{}
	_ = streamer.Register(messages)
	_ = streamer.Start()
	for {
		select {
		case <-exit:
			fmt.Println("Goodbye...")
			pub.c.Disconnect(250)
			close(finish)
			return
		case msg := <-messages:
			fmt.Println("Published!")
			pub.Publish(msg)
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
