package main

import "fmt"

func main() {
	// for i := 1; i <= 3; i++ {
	// 	fmt.Println(i)
	// }

	// var sum int
	// for i := 1; i <= 10; i++ {
	// 	sum += i
	// }

	// // For go v1.22+
	for i := range 3 { // [0, 2]
		fmt.Println(i)
	}

	 // Traditional for loop
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }
    
    // While-style
    count := 0
    for count < 3 {
        fmt.Println("Count:", count)
        count++
    }
    
    // Range กับ slice
    numbers := []int{10, 20, 30}
    for index, value := range numbers {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }
    
    // Range กับ map
    scores := map[string]int{"สมชาย": 85, "สมหญิง": 90}
    for name, score := range scores {
        fmt.Printf("%s ได้ %d คะแนน\n", name, score)
    }
	 // break - ออกจาก loop
    for i := 0; i < 10; i++ {
        if i == 5 {
            break
        }
        fmt.Println(i)
    }
    
    // continue - ข้ามรอบนั้น
    for i := 0; i < 5; i++ {
        if i%2 == 0 {
            continue
        }
        fmt.Println("Odd:", i)
    }
    
    // defer - ทำงานตอนจบ function (LIFO)
    defer fmt.Println("World")
    fmt.Println("Hello")
    // ผลลัพธ์: Hello World


	
}
