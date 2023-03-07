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
		//memberikan nilai default supaya tidak nil ketika sedang menunggu data
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

//Golang memiliki sebuah struct bernama sync.Map,
//Map ini mirip dengan golang Map yang sudah kita coba sebelumnya,
// namun yang membedakan Map ini aman untuk menggunakan concurrent menggunakan goroutine.
//Store(key, value)	Untuk menyimpan data ke map
// Load(key)	Untuk mengambil data dari map menggunakan key
// Delete(key)	Untuk menghapus data di map menggunakan key
// Range(function(key, value))	Digunakan untuk melakukan iterasi seluruh data di map

func AddToMap(data *sync.Map, value int, group *sync.WaitGroup) {
	defer group.Done()
	data.Store(value, value)
}

func TestMap(t *testing.T) {
	data := &sync.Map{}
	group := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go AddToMap(data, i, group)
	}
	group.Wait()
	data.Range(func(key, value any) bool {
		fmt.Println(key, ":", value)
		return true
	})
}

// Cond adalah implementasi locking berbasis kondisi, cond membutuhkan locker(Mutex atau RWMutex) untuk implementasi locking nya, namun berbeda dengan locker biasanya, di cond terdapat function Wait() untuk menunggu apakah perlu menunggu atau tidak
// Jadi nanti saat setelah kita melakukan locking di dalam condition ini kita akan memanggil Wait(), jadi nanti jika menurut condition tersebut kita harus menunggu maka kita akan menunggu.
// Setelah menggunakan Wait(), kita bisa menggunakan function Signal() untuk memberitahu sebuah goroutine agar tidak perlu menunggu lagi, sedangkan function Broadcast() digunakan untuk memberitahu semua goroutine untuk tidak perlu menunggu lagi. Untuk membuat cond kita bisa menggunakan sync.NewCond(Locker).

var locker = sync.Mutex{}
var cond = sync.NewCond(&locker)
var group = sync.WaitGroup{}

func WaitCondition(value int) {
	defer group.Done()

	cond.L.Lock()
	cond.Wait()

	fmt.Println("Done", value)
	cond.L.Unlock()
}

func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ {
		group.Add(1)
		go WaitCondition(i)
	}
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)

			//signal untuk menunggu dan terhenti sementara
			cond.Signal()

			//broadcast tidak perlu menunggu lagi
			// cond.Broadcast()
		}
	}()
	group.Wait()
}
