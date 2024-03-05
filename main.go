package main

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

func (fifo_channel FiFoChannel) PullMessage() string {
	return <-fifo_channel.channel
}

func (fifo_channel FiFoChannel) GetSize() int {
	return len(fifo_channel.channel)
}

func main() {
	fifo_channel := InitChannel(7)

	fifo_channel.PushMessage("hello")
	fifo_channel.PushMessage("world")

	println(fifo_channel.GetSize())
	println(fifo_channel.PullMessage())
	println(fifo_channel.PullMessage())
}
