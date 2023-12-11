package test

import (
	"fmt"
	"log"
	"testing"

	"github.com/nsqio/go-nsq"
)

type MyMessageHandler struct{}

func (h *MyMessageHandler) HandleMessage(msg *nsq.Message) error {
	fmt.Println("Received message:", string(msg.Body))
	return nil
}

func TestA(t *testing.T) {
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("my_topic", "my_channel", config)
	if err != nil {
		log.Fatal(err)
	}

	consumer.AddHandler(&MyMessageHandler{})

	err = consumer.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		log.Fatal(err)
	}

	producer, err := nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 10; i++ {

		message := fmt.Sprintf("index:%d", i)

		err := producer.Publish("my_topic", []byte(message))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func MyMethod() error {
	return nil
}

// 装饰器函数，接收一个函数作为参数，并返回一个新的函数
func Decorator(f func() error) func() {
	return func() {
		err := f() // 调用原始方法
		if err != nil {
			fmt.Println("err")
		}
		fmt.Println("Executing decorator logic")
	}
}

func TestName(t *testing.T) {
	decoratedMethod := Decorator(MyMethod)
	decoratedMethod()
}
