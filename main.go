package main

import (
	"errors"
	"fmt"
)

type FiFoChannel struct {
	channel chan string
}

func InitChannel(chan_size int) FiFoChannel {
	return FiFoChannel{
		make(chan string, chan_size),
	}
}

func (fifo_channel FiFoChannel) PushMessage(message string) {
	select {
	case fifo_channel.channel <- message:
	default:
		fmt.Println("Channel buffer is full, skipping push:", message)
	}
}

func (fifo_channel *FiFoChannel) PullMessage() (string, error) {
	select {
	case msg := <-fifo_channel.channel:
		return msg, nil
	default:
		return "", errors.New("channel is empty")
	}
}

func (fifo_channel FiFoChannel) GetSize() int {
	return len(fifo_channel.channel)
}

func main() {
	fifo_channel := InitChannel(3)

	fifo_channel.PushMessage("hello")
	fifo_channel.PushMessage("world")

	println(fifo_channel.GetSize())
	msg1, err1 := fifo_channel.PullMessage()
	if err1 != nil {
		fmt.Println("Error pulling message:", err1)
	} else {
		fmt.Println("Pulled Message:", msg1)
	}

	msg2, err2 := fifo_channel.PullMessage()
	if err2 != nil {
		fmt.Println("Error pulling message:", err2)
	} else {
		fmt.Println("Pulled Message:", msg2)
	}

	println(fifo_channel.GetSize())
}
