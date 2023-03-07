package golang_concurrency

import (
	"fmt"
	"strconv"
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

// Buffered channel
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 2)
	defer close(channel)

	go func() {
		channel <- "Muhammad"
		channel <- "Deril"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Selesai")
}

// Range channel
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	//mengirim data ke channel
	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		//close biar ga deadlock
		close(channel)
	}()

	//mengambil data dari channel
	for data := range channel {
		fmt.Println("Menerima Data", data)
	}

	fmt.Println("Selesai")
}

// Select channel
func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Muhammad Deril"
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari Channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari Channel 2", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}
}

// default select untuk mencegah deadlock karena menunggu data dari channel
func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari Channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari Channel 2", data)
			counter++
		default:
			fmt.Println("Menunggu data")
		}

		if counter == 2 {
			break
		}
	}
}
