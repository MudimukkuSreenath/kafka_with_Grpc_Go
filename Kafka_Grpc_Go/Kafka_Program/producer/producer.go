package producer

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	topic         = "TestTopic"
	brokerAddress = "localhost:9092"
)

func main() {
	ctx := context.Background()
	produce(ctx)
}

func produce(ctx context.Context) {
	i := "Priyanshi Panchal"

	l := log.New(os.Stdout, "kafka writer: ", 0)
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		Logger: l,
	})


	// each kafka message has a key and value. The key is used
	// to decide which partition (and consequently, which broker)
	// the message gets published on
	err := w.WriteMessages(ctx, kafka.Message{
		Key: []byte(i),
		// create an arbitrary message payload for the value
		Value: []byte("this is message" + i),
	})
	if err != nil {
		panic("could not write message " + err.Error())
	}

	// log a confirmation once the message is written
	fmt.Println("writes:", i)
	// sleep for a second
	time.Sleep(time.Second)

}

