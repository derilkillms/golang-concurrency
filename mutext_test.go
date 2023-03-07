package golang_concurrency

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// goroutine nya akan balapan merubah goroutine nya. (race condition)
func TestRaceCondition(t *testing.T) {
	x := 0
	var mutex sync.Mutex

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {

				//sebelum dilakukan increment pada variable x kita lock terlebih dahulu mutex nya dan
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
				//setalah dilakukan increment baru kita unlock kembali.
				//Maka 1000 goroutine ini akan mencoba melakukan locking dan cuma satu goroutine saja yang diperbolehkan melakukan lock nya, setelah satu goroutine itu berhasil melakukan lock maka 999 goroutine lainnya akan menunggu sampai mock nya di unlock.

			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter = ", x)
}

//sebenar nya bisa menggunakan mutex saja namun nanti akan terjadi rebutan antara proses membaca dan mengubah.
//Di golang telah disediakan struct RWMutex(Read Write Mutex) untuk menangani hal ini, dimana mutex jenis ini memiliki dua lock, lock untuk Read dan lock untuk Write.

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amaunt int) {
	account.RWMutex.Lock()
	account.Balance += amaunt
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()
	return balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}
	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()

	}
	time.Sleep(5 * time.Second)
	fmt.Println("Total Balance", account.GetBalance())
}
