package golang_concurrency

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// kita membuat func dengan nama RunAsynchronous yang akan di running dengan goroutine, kemudian kita hanya perlu menambahkan parameter *sync.WaitGroup,
// untuk menambahkan proses gunakan method Add(int) dan dibawahnnya barulah masukkan kode program kita,
// jangan lupa juga untuk menandai proses telah selesai dengan method Done() menggunakan defer agar bisa dipastikan method ini akan dijalankan setelah function selesai.

func RunAsynchronous(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)

	fmt.Println("Hello")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsynchronous(group)
	}

	group.Wait()
	fmt.Println("Selesai")
}

// Once adalah fitur digolang yang bisa kita gunakan untuk memastikan bahwa sebuah function hanya dieksekusi hanya sekali.

var counter = 0

func OnlyOnce() {
	counter++
}

func TestOnce(t *testing.T) {
	once := sync.Once{}
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go func() {
			group.Add(1)
			once.Do(OnlyOnce)
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter :", counter)
}

// Pool ini digunakan untuk menyimpan data, selanjutnya untuk menggunakan datanya kita bisa mengambil data dari Pool, dan setelah selesai menggunakan datanya, kita bisa menyimpan kembali ke Pool nya
func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			return "Default"
		},
	}
	group := sync.WaitGroup{}

	pool.Put("Muhammad")
	pool.Put("Deril")

	for i := 0; i < 10; i++ {
		group.Add(1)
		go func() {
			defer group.Done()

			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	group.Wait()
	fmt.Println("Selesai")
}
