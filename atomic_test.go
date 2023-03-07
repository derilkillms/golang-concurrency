package golang_concurrency

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

// Paket atom menyediakan primitif memori atom tingkat rendah yang berguna untuk mengimplementasikan algoritma sinkronisasi.
// Fungsi-fungsi ini membutuhkan kehati-hatian agar dapat digunakan dengan benar. Kecuali untuk aplikasi khusus tingkat rendah,
// sinkronisasi lebih baik dilakukan dengan saluran atau fasilitas paket sinkronisasi. Berbagi memori dengan berkomunikasi;
// jangan berkomunikasi dengan berbagi memori.

func TestAtomic(t *testing.T) {
	var group sync.WaitGroup
	var counter int64 = 0

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			defer group.Done()
			for j := 0; j < 100; j++ {
				atomic.AddInt64(&counter, 1)
			}
		}()
	}

	group.Wait()
	fmt.Println("Counter :", counter)
}

// Kode program diatas sama seperti yang sebelumnya kita buat menggunakan Mutex, bedanya kali ini kita tidak melakukan lock tetapi menggunakan package atomic,
// baris 24 kita melakukan increment pada variable counter dengan lebih mudah dan simple.
// jika kita akan melakukan manipulasi data primitive seperti number yang mana data nya diakses banyak goroutine maka kita bisa menggunakan atomic.
