package main

import (
	"errors"
	"fmt"
	"sync"
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
	var wg sync.WaitGroup
	fifo_channel := InitChannel(2)

	wg.Add(2)
	go func() {
		defer wg.Done()
		fifo_channel.PushMessage("hello")
	}()
	go func() {
		defer wg.Done()
		fifo_channel.PushMessage("world")
	}()

	wg.Wait()

	fmt.Println(fifo_channel.GetSize())

	wg.Add(2)
	go func() {
		defer wg.Done()
		msg, err := fifo_channel.PullMessage()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(msg)
		}
	}()
	go func() {
		defer wg.Done()
		msg, err := fifo_channel.PullMessage()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(msg)
		}
	}()

	wg.Wait()
	fmt.Println(fifo_channel.GetSize())
}
