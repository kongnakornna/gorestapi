package main

import (
    "fmt"
    "sync"
    "time"
)

// merge รวมข้อมูลจาก 2 channels เข้าด้วยกัน
// ใช้ nil channel เพื่อปิด branch ที่ถูกปิดแล้ว
func merge(ch1, ch2 <-chan int) <-chan int {
    out := make(chan int)
    
    var wg sync.WaitGroup
    wg.Add(2)
    
    // Goroutine 1: อ่านจาก ch1
    go func() {
        defer wg.Done()
        for v := range ch1 {
            out <- v
        }
    }()
    
    // Goroutine 2: อ่านจาก ch2
    go func() {
        defer wg.Done()
        for v := range ch2 {
            out <- v
        }
    }()
    
    // Goroutine ปิด out channel เมื่อทั้ง ch1 และ ch2 ปิด
    go func() {
        wg.Wait()
        close(out)
    }()
    
    return out
}

// mergeWithNil แสดงการใช้ nil channel แบบคลาสสิก
func mergeWithNil(a, b <-chan int) <-chan int {
    out := make(chan int)
    
    go func() {
        // ปิด out เมื่อ goroutine จบ
        defer close(out)
        
        // ทำการ merge จนกว่าทั้ง a และ b จะเป็น nil
        for a != nil || b != nil {
            select {
            case v, ok := <-a:
                // ถ้า a ปิดแล้ว ให้ตั้ง a = nil เพื่อปิด branch นี้
                if !ok {
                    a = nil // สำคัญ: ปิด branch a
                    continue
                }
                out <- v
                
            case v, ok := <-b:
                // ถ้า b ปิดแล้ว ให้ตั้ง b = nil เพื่อปิด branch นี้
                if !ok {
                    b = nil // สำคัญ: ปิด branch b
                    continue
                }
                out <- v
            }
        }
    }()
    
    return out
}

// ฟังก์ชัน helper สำหรับสร้าง test channel
func asChan(values ...int) <-chan int {
    ch := make(chan int)
    go func() {
        for _, v := range values {
            ch <- v
            time.Sleep(100 * time.Millisecond) // simulate work
        }
        close(ch)
    }()
    return ch
}

func main() {
    fmt.Println("=== Merge Channel Example ===")
    
    // สร้าง test data
    ch1 := asChan(1, 3, 5, 7)
    ch2 := asChan(2, 4, 6, 8)
    
    // รวม channel
    merged := mergeWithNil(ch1, ch2)
    
    // รับข้อมูลจาก merged channel
    fmt.Print("Merged values: ")
    for v := range merged {
        fmt.Printf("%d ", v)
    }
    fmt.Println("\n\nMerge completed!")
}