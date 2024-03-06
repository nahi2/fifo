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
	fifo_channel.channel <- message
}

func (fifo_channel *FiFoChannel) PullMessage() (string, error) {
	msg, ok := <-fifo_channel.channel
	if !ok {
		return "", errors.New("channel is empty")
	}
	return msg, nil
}

func (fifo_channel FiFoChannel) GetSize() int {
	return len(fifo_channel.channel)
}

func main() {
	fifo_channel := InitChannel(2)

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
