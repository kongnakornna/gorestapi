package main

import (
	"fmt"
	"sync"
	"time"
)

type Queue struct {
    items []int
    mu    sync.Mutex
    cond  *sync.Cond
    // channel สำหรับ external trigger (อาจเป็น nil)
    notifyCh chan struct{}
}

func NewQueue() *Queue {
    q := &Queue{
        notifyCh: make(chan struct{}, 1),
    }
    q.cond = sync.NewCond(&q.mu)
    return q
}

// EnableNotify เปิดการแจ้งเตือนผ่าน channel (ตั้ง notifyCh ให้ไม่เป็น nil)
func (q *Queue) EnableNotify() {
    q.mu.Lock()
    defer q.mu.Unlock()
    if q.notifyCh == nil {
        q.notifyCh = make(chan struct{}, 1)
    }
}

// DisableNotify ปิดการแจ้งเตือนผ่าน channel (ตั้ง notifyCh = nil)
func (q *Queue) DisableNotify() {
    q.mu.Lock()
    defer q.mu.Unlock()
    q.notifyCh = nil
}

// Push เพิ่มข้อมูลและแจ้งเตือนทั้ง Cond และ channel (ถ้า channel ไม่เป็น nil)
func (q *Queue) Push(v int) {
    q.mu.Lock()
    q.items = append(q.items, v)
    q.mu.Unlock()

    q.cond.Signal() // ปลุก Wait

    // ส่ง signal ผ่าน channel ถ้ามี
    q.mu.Lock()
    ch := q.notifyCh
    q.mu.Unlock()
    if ch != nil {
        select {
        case ch <- struct{}{}:
        default:
        }
    }
}

// PopWithChannel รอรับข้อมูลผ่าน channel หรือ Cond
func (q *Queue) PopWithChannel(ctx <-chan struct{}) (int, error) {
    // ตัวแปร channel ที่อาจเป็น nil
    var notifyCh <-chan struct{} = q.notifyCh

    for {
        q.mu.Lock()
        if len(q.items) > 0 {
            v := q.items[0]
            q.items = q.items[1:]
            q.mu.Unlock()
            return v, nil
        }
        q.mu.Unlock()

        // รอจากหลายแหล่ง
        select {
        case <-ctx:
            return 0, fmt.Errorf("cancelled")
        case <-notifyCh:
            // ได้รับ signal จาก channel (ถ้าไม่เป็น nil)
            continue
        default:
            // ใช้ Cond Wait แทน
            q.mu.Lock()
            q.cond.Wait()
            q.mu.Unlock()
        }
    }
}

func main() {
    q := NewQueue()
    q.EnableNotify() // เปิด channel notification

    // ตัวรับ
    go func() {
        for i := 0; i < 5; i++ {
            val, _ := q.PopWithChannel(nil)
            fmt.Println("received:", val)
        }
    }()

    // ตัวส่ง
    for i := 1; i <= 5; i++ {
        q.Push(i)
        time.Sleep(500 * time.Millisecond)
    }

    // ปิด channel notification แล้วลองอีกครั้ง
    q.DisableNotify()
    fmt.Println("disable notify, push 6")
    q.Push(6)
    time.Sleep(1 * time.Second) // จะยังไม่มีใครรับ เพราะ channel ปิดและ Cond ก็ไม่ถูก Signal? (ต้อง Signal ด้วย)
    // แต่ cond ถูก Signal ใน Push อยู่แล้ว ดังนั้นยังรับได้
}