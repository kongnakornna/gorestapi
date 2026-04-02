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

func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

func main() {
	//     p := Person{Name: "Alice", Age: 30}
	//     p.HaveBirthday()
	//     fmt.Println(p.Age) // 31 (เปลี่ยนแปลงจริง)
		
	//     fmt.Println(p.Greet()) // Hello, Alice
	// 	a, b := split(10)
	// 	fmt.Println(a, b) // ผลลัพธ์: 4 6
	// 	fmt.Println(sum(1,2,3,4)) 

	// 	result, err := divide(10, 2)
	// 	if err != nil {
	// 		fmt.Println("Error:", err)
	// 	} else {
	// 		fmt.Println("Result:", result)
	// 	}
	
	//  // ----- Using add with ints -----
	//     total := 10          // total is int
	//     // Correct: use = for reassignment, not :=
	//     total = add(5, 3)    // total becomes 8
	//     fmt.Println("total (int):", total)

	//     // ----- Using add with float64 (explicit conversion) -----
	//     var f float64
	//     // Error previously: cannot use add(5,3) (int) as float64
	//     // Fix: convert the int result to float64
	//     f = float64(add(5, 3))
	//     fmt.Println("f (float64):", f)

	//     // ----- Alternative: using generic function with floats directly -----
	//     f2 := addGeneric(2.5, 3.7) // f2 is float64
	//     fmt.Println("f2 (generic float):", f2)

	//     // ----- If you need to mix int and float, convert explicitly -----
	//     var mixed float64
	//     mixed = float64(add(5, 3)) + 2.5
	//     fmt.Println("mixed:", mixed)

	//     // ----- Short variable declaration with a new variable -----
	//     sum := add(10, 20) // new variable, okay
	//     fmt.Println("sum:", sum)

	s1 := sum(1, 2, 3, 4)       // 10
	s2 := sum()                 // 0 (ไม่มีอาร์กิวเมนต์)
	s3 := sum(5, 10)            // 15
 
	numbers := []int{2, 4, 6}
	s4 := sum(numbers...)   
	fmt.Println("s4:", s4)
}