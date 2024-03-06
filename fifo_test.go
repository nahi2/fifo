package main

import "testing"

func TestInitChannel(t *testing.T) {
	fifo := InitChannel(1)

	if cap(fifo.channel) != 1 {
		t.Errorf("Expected channel size: 1, got: %d", cap(fifo.channel))
	}

	fifo = InitChannel(20)

	if cap(fifo.channel) != 20 {
		t.Errorf("Expected channel size: 20, got: %d", cap(fifo.channel))
	}
}

func TestFiFoChannel_PushMessage(t *testing.T) {
	fifo := InitChannel(5)

	fifo.PushMessage("hello")

	if fifo.GetSize() != 1 {
		t.Errorf("Want size: 1, got: %d", fifo.GetSize())
	}
}

func TestFiFoChannel_PullMessage(t *testing.T) {
	fifo := InitChannel(5)

	fifo.PushMessage("hello")
	message, err := fifo.PullMessage()

	if err != nil {
		t.Errorf("PullMessage failed: %v", err)
	}

	if message != "hello" {
		t.Errorf("Want message: hello, got: %s", message)
	}
}

func TestFiFoChannel_GetSize(t *testing.T) {
	fifo := InitChannel(5)

	fifo.PushMessage("hello")

	if fifo.GetSize() != 1 {
		t.Errorf("Want size: 1, got: %d", fifo.GetSize())
	}
}
