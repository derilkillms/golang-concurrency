package golang_concurrency

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Timer adalah representasi suatu kejadian, ketika waktu timer sudah expired, maka event akan dikirim ke dalam channel, untuk membuat timer kita bisa menggunakan time.NewTimer(duration).
func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())

	time := <-timer.C
	fmt.Println(time)
}

// Kadang kita hanya butuh channel nya saja, tidak membutuhkan data Timer nya, untuk melakukan hal ini kita bisa menggunakan function time.After(duration)
func TestAfter(t *testing.T) {
	channel := time.After(1 * time.Second)
	time := <-channel
	fmt.Println(time)
}

// Kadang ada kebutuhan kita ingin menjalankan sebuah function dengan delay waktu tertentu, kita bisa memanfaatkan Timer dengan menggunakan function timer.AfterFunc(),
func TestAfterFunc(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(1)

	time.AfterFunc(1*time.Second, func() {
		fmt.Println("Execute after 1 second")
		group.Done()
	})
	group.Wait()
}

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)

	for tick := range ticker.C {
		fmt.Println(tick)
	}

}

// Kadang kita tidak butuh data Ticker nya, kita hanya butuh channel nya saja, jika demikian kita bisa menggunakan function timer.Tick(duration), function ini tidak akan mengembalikan Ticker, hanya mengembalikan channel timer nya saja.

func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)

	for tick := range channel {
		fmt.Println(tick)

	}
}
