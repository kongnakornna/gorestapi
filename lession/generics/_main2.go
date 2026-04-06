package main

import "fmt"
// ชือโครงสร้างหลัก Person
type Person struct {
    Name string  //ตัวแปร  Name
    Age  int  //ตัวแปร  Age
}

// Pointer receiver – สามารถแก้ไขค่าต้นฉบับได้
// รับค่า (p *Person) จาก โครงสร้างหลัก Person มาใช้ใน ชือฟังก์ชั่น  HaveBirthday
func (p *Person) HaveBirthday() {
    p.Age++ // แก้ไข field ของ instance ดั้งเดิม  เพิ่มจำนวน เพิ่มใหม่เข้าไป  1 เช่น 30+1=31
}

// Value receiver – ทำงานกับสำเนา
func (p Person) Greet() string {
    return "Hello, " + p.Name
}

func sum1(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}
// เรียกใช้: sum(1,2,3,4)
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return // naked return
}

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

// add returns the sum of two integers
func add(x, y int) int {
    return x + y
}
// addFloat is a generic version that works with any numeric type (requires Go 1.18+)
func addGeneric[T int | float64](x, y T) T {
    return x + y
}

func sum2(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

var fn func(int) int = func(x int) int { return x * 2 }
// ส่งฟังก์ชันเป็น argument
func apply(fn func(int) int, val int) int {
    return fn(val)
}
func main() {
	numbers := []int{2, 4, 6,10,8}
	s4 := sum2(numbers...)   
	fmt.Println("s4:", s4)

    square := func(n int) int { return n * n }
    fmt.Println(square(7)) // 49
    result := apply(func(x int) int { return x + 10 }, 9) // 19
    fmt.Println(result)  
}
