package main

import (
	"fmt"
	"time"
)

func main() {
    // สร้าง channel สำหรับรับงาน
    workCh := make(chan string)
    
    // channel สำหรับควบคุมการเปิด/ปิด (toggle)
    toggleCh := make(chan bool)
    
    // nil channel เริ่มต้น (ปิดอยู่)
    var activeWorkCh chan string = nil
    
    go func() {
        for {
            select {
            case work, ok := <-activeWorkCh:
                if !ok {
                    fmt.Println("work channel closed, exiting")
                    return
                }
                fmt.Printf("✅ รับงาน: %s\n", work)
                
            case toggle := <-toggleCh:
                if toggle {
                    activeWorkCh = workCh // เปิด: ใช้ channel จริง
                    fmt.Println("🟢 เปิดรับงาน")
                } else {
                    activeWorkCh = nil    // ปิด: ใช้ nil channel (ไม่รับข้อมูล)
                    fmt.Println("🔴 ปิดรับงาน")
                }
            }
        }
    }()
    
    // ส่งงานไปเรื่อยๆ
    go func() {
        i := 1
        for {
            workCh <- fmt.Sprintf("งานที่ %d", i)
            i++
            time.Sleep(500 * time.Millisecond)
        }
    }()
    
    // ตัวอย่างการ toggle
    toggleCh <- true   // เปิดรับ
    time.Sleep(2 * time.Second)
    toggleCh <- false  // ปิดรับ (งานที่ส่งมาจะไม่ถูกประมวลผล)
    time.Sleep(2 * time.Second)
    toggleCh <- true   // เปิดอีกครั้ง
    time.Sleep(2 * time.Second)
    
    close(workCh)
    close(toggleCh)
}