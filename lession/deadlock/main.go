package main

import (
	"fmt"
	"sync"
	"time"
)

type Wallet struct {
	ID      int
	mu      sync.Mutex
	Balance int
}

// BUG! Transfer นี้มีโอกาส Deadlock สูง เพราะไม่ได้กำหนดลำดับการ Lock
func TransferBad(from, to *Wallet, amount int) {
    // Goroutine 1: Transfer(A, B) จะ Lock A -> B
    // Goroutine 2: Transfer(B, A) จะ Lock B -> A
    from.mu.Lock()
    fmt.Printf("🔒 Lock %d\n", from.ID)
    // จำลองการทำงานระหว่าง Lock ทั้งสองตัว เพิ่มโอกาส Deadlock
    time.Sleep(1 * time.Millisecond)
    to.mu.Lock()
    fmt.Printf("🔒 Lock %d\n", to.ID)

    // Critical Section
    from.Balance -= amount
    to.Balance += amount

    to.mu.Unlock()
    from.mu.Unlock()
    fmt.Printf("✅ Transfer %d -> %d success\n", from.ID, to.ID)
}

func main() {
    // สร้าง Wallet สองใบ
    walletA := &Wallet{ID: 1, Balance: 100}
    walletB := &Wallet{ID: 2, Balance: 100}

    // พยายามโอนเงินพร้อมกันคนละทิศทาง
    go TransferBad(walletA, walletB, 10) // A -> B
    go TransferBad(walletB, walletA, 20) // B -> A

    // รอให้ Goroutine ทั้งคู่ทำงานเสร็จ
    time.Sleep(2 * time.Second)
    fmt.Printf("Final Balance: A=%d, B=%d\n", walletA.Balance, walletB.Balance)
}
