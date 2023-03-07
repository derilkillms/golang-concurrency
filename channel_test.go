package golang_concurrency

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		//mengirim data
		channel <- "Muhammad Deril"
		fmt.Println("selesai mengirim data ke channel")
	}()
	data := <-channel

	fmt.Println(data)

	time.Sleep(5 * time.Second)

	defer close(channel)
}
