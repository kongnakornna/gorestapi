package main

import (
    "fmt"
    "os"
)

func readFile() {
    f, err := os.Open("file.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer f.Close() // รับประกันว่าจะปิดไฟล์เมื่อฟังก์ชันจบ

    // อ่านข้อมูลจากไฟล์...
    data := make([]byte, 100)
    n, _ := f.Read(data)
    fmt.Printf("Read %d bytes: %s\n", n, string(data[:n]))
}

func main() {
    readFile()
}