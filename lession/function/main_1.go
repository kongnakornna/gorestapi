// package main

// import "fmt"

// func main() {

// 	  result := add(40, 2)
// 	  fmt.Println(result)

// 	// fmt.Println(sub(4, 2))

// 	// fmt.Println(mul(4, 2))
// 	// fmt.Println(div(4, 2))

// 	// fmt.Println(div(4, 0))

// 	// n := 10
// 	// double(n)
// 	// fmt.Println(n)

// 	// result := func(a, b int) int {
// 	// 	return a + b
// 	// }(1, 2)

// 	// fmt.Println(result)
// }

// func hello(name string) {
// 	fmt.Println("Hello", name)
// }

// func add(a, b int) int {
// 	return a + b
// }

// func sub(a, b int) int {
// 	return a - b
// }

// func mul(a, b int) int {
// 	return a * b
// }

// func div(a, b int) int {
// 	if b == 0 {
// 		fmt.Println("Error: Division by zero.")
// 		return 0
// 	}
// 	return a / b
// }

// func double(n int) {
// 	n *= 2
// }

/*********************************/

// package main

// import (
// 	"errors"
// 	"fmt"
// )

// func getPerson(id int) (string, int, error) {
//     if id == 1 {
//         return "สมชาย", 30, nil
//     }
//     return "", 0, errors.New("ไม่พบข้อมูล")
// }

// func main() {
//     name, age, err := getPerson(1)
//     if err != nil {
//         fmt.Println("Error:", err)
//         return
//     }
//     fmt.Printf("ชื่อ: %s, อายุ: %d\n", name, age)
// }

// package main

// import "fmt"

// func sum(nums ...int) int {
//     total := 0
//     for _, num := range nums {
//         total += num
//     }
//     return total
// }

// func main() {
//     fmt.Println(sum(1, 2, 3))        // 6
//     fmt.Println(sum(10, 20, 30, 40)) // 100

//     // ส่ง slice
//     numbers := []int{1, 2, 3, 4, 5}
//     fmt.Println(sum(numbers...))
// }

// package main

// import "fmt"

// func counter() func() int {
//     count := 0
//     return func() int {
//         count++
//         return count
//     }
// }

// func main() {
//     c1 := counter()
//     fmt.Println(c1()) // 1
//     fmt.Println(c1()) // 2
//     fmt.Println(c1()) // 3

//     c2 := counter()
//     fmt.Println(c2()) // 1 (แยก instance)
// }

// package main

// import "fmt"
// func main() {
//     sum := applyOperation(5, 3, func(a, b int) int {
//         return a + b
//     })
//     fmt.Println("Sum:", sum)

//     multiply := getOperation("multiply")
//     fmt.Println("multiply:", multiply(4, 5))

// 	 add := getOperation("add")
//     fmt.Println("add:", add(4, 5))
// }

// // รับฟังก์ชันเป็นพารามิเตอร์
// func applyOperation(x, y int, operation func(int, int) int) int {
//     return operation(x, y)
// }

// // คืนค่าฟังก์ชัน
// func getOperation(op string) func(int, int) int {
//     switch op {
//     case "add":
//         return func(a, b int) int { return a + b }
//     case "multiply":
//         return func(a, b int) int { return a * b }
//     default:
//         return nil
//     }
// }

// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"strings"
// 	// local package
// 	// external package
// )

// func main() {
//     fmt.Println(rand.Intn(100))
//     fmt.Println(strings.ToUpper("hello"))
// }

// package main

// import "fmt"

// type Person struct {
//     Name    string
//     Age     int
//     Address string
// }

// // Struct with tags (สำหรับ JSON)
// type User struct {
//     ID    int    `json:"id"`
//     Name  string `json:"name"`
//     Email string `json:"email"`
// }

// func main() {
//     p1 := Person{"สมชาย", 30, "กรุงเทพฯ"}
//     p2 := Person{Name: "สมหญิง", Age: 25}

//     fmt.Println(p1.Name)
//     p2.Age = 126
// 	fmt.Println(p2)
// }

// package main

// import "fmt"

// type Animal struct {
//     Name string
// }

// func (a Animal) Speak() {
//     fmt.Println(a.Name, "makes sound")
// }

// type Dog struct {
//     Animal  // Embedding
//     Breed string
// }

// func (d Dog) Speak() {
//     fmt.Println(d.Name, "barks!")
// }

// func main() {
//     dog := Dog{
//         Animal: Animal{Name: "Max"},
//         Breed:  "Golden",
//     }

//     dog.Speak()           // Override
//     dog.Animal.Speak()    // Call parent
// }

// package main

// import "fmt"

// type Counter struct {
//     count int
// }

// // Value receiver (ไม่เปลี่ยนค่าเดิม)
// func (c Counter) Value() int {
//     return c.count
// }

// // Pointer receiver (เปลี่ยนค่าเดิม)
// func (c *Counter) Increment() {
//     c.count++
// }

// func main() {
//     c := Counter{count: 100}

//     c.Increment()
//     fmt.Println(c.count)  // 101
// }

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Product struct {
//     ID       int     `json:"id"`
//     Name     string  `json:"name"`
//     Price    float64 `json:"price"`
//     InStock  bool    `json:"in_stock"`
// }

// func main() {
//     // struct -> JSON
//     p := Product{ID: 1, Name: "Laptop", Price: 25000.50, InStock: true}
//     jsonData, _ := json.Marshal(p)
//     fmt.Println(string(jsonData))

//     // JSON -> struct
//     jsonString := `{"id":2,"name":"Mouse","price":500,"in_stock":false}`
//     var product Product
//     json.Unmarshal([]byte(jsonString), &product)
//     fmt.Printf("%+v\n", product)
// }

// package main

// import "fmt"

// func main() {
//     x := 10          // ตัวแปรปกติ เก็บค่า 10
//     p := &x          // พอยน์เตอร์ เก็บที่อยู่ของ x

//     fmt.Println(x)   // 10 (ค่าของตัวแปร x)
//     fmt.Println(&x)  // 0xc0000140a0 (ที่อยู่ของ x)
//     fmt.Println(p)   // 0xc0000140a0 (ที่อยู่เดียวกัน)
//     fmt.Println(*p)  // 10 (ค่าที่พอยน์เตอร์ชี้ไป)

// 	 // เปลี่ยนค่าผ่านพอยน์เตอร์
//     *p = 20
//     fmt.Println("New x:", x)  // 20

// }

// package main

// import "fmt"

// type Rectangle struct {
//     Width, Height float64
// }

// // Value receiver (ไม่เปลี่ยน struct เดิม)
// func (r Rectangle) Area() float64 {
//     return r.Width * r.Height
// }

// // Pointer receiver (เปลี่ยน struct เดิม)
// func (r *Rectangle) Scale(factor float64) {
//     r.Width *= factor
//     r.Height *= factor
// }

// func main() {
//     rect := Rectangle{Width: 10, Height: 5}
//     rect.Scale(2)
//     fmt.Println(rect)  // {20, 10}
// }

// package main

// import "fmt"

// func printAnything(v interface{}) {
//     fmt.Printf("Value: %v, Type: %T\n", v, v)
// }

// func main() {
//     var anything interface{}
//     var na interface{}

//     anything = 42
//     printAnything(anything)

//     anything = "hello"
//     printAnything(anything)

//      fmt.Println(anything)

// 	na = 42.22
// 	printAnything(na)
//     fmt.Println(na)

//     // Go 1.18+ use 'any' alias
//     var x any = "hello"
//     fmt.Println(x)
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
//     var i interface{} = "hello"

//     // Type assertion แบบ safe
//     if s, ok := i.(string); ok {
//         fmt.Println("String:", s)
//     } else {
//         fmt.Println("Not a string")
//     }

//     // อีกตัวอย่าง
//     if num, ok := i.(int); ok {
//         fmt.Println("Integer:", num)
//     } else {
//         fmt.Println("Not an integer")  // จะเข้ามาทางนี้
//     }

//     // ตัวอย่างกับ type ต่างๆ
//     checkType(42)
//     checkType(113.14)
//     checkType("golang")
//     checkType(true)
// }

// func checkType(v interface{}) {
//     switch v := v.(type) {
//     case int:
//         fmt.Printf("Integer: %d\n", v)
//     case float32:
//         fmt.Printf("float32: %f\n", v)  // ✅ ใช้ %f สำหรับ float
//     case float64:
//         fmt.Printf("float64: %f\n", v)  // ✅ ใช้ %f สำหรับ float
//     case string:
//         fmt.Printf("String length: %d, value: %s\n", len(v), v)
//     case bool:
//         fmt.Printf("Boolean: %t\n", v)
//     default:
//         fmt.Printf("Unknown type: %T\n", v)
//     }
// }

// package main

// import "fmt"

// // Repository interface
// type UserRepository interface {
//     GetUser(id int) (string, error)
// }

// // Service depends on interface
// type UserService struct {
//     repo UserRepository
// }

// func NewUserService(repo UserRepository) *UserService {
//     return &UserService{repo: repo}
// }

// // Concrete implementation
// type InMemoryUserRepo struct {
//     users map[int]string
// }

// func (r *InMemoryUserRepo) GetUser(id int) (string, error) {
//     if name, exists := r.users[id]; exists {
//         return name, nil
//     }
//     return "", fmt.Errorf("user not found")
// }

// func main() {
//     repo := &InMemoryUserRepo{users: map[int]string{1: "John"}}
//     service := NewUserService(repo)
//     name, _ := service.repo.GetUser(1)
//     fmt.Println("User:", name)
// }

// package main

// import (
// 	"errors"
// 	"fmt"
// )

// func divide(a, b float64) (float64, error) {
//     if a == 0 || b == 0 {
//         return 0, errors.New("a or b division by zero")
//     }

//     return a / b, nil
// }

// func main() {
//     result, err := divide(1066.5, 10)
//     if err != nil {
//         fmt.Println("Error:", err)
//         return
//     }
//     fmt.Println("Result:", result)

// 	 result1, err := divide(0, 0)
//     if err != nil {
//         fmt.Println("Error:", err)
//         return
//     }
//     fmt.Println("Result:", result1)
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// type ValidationError struct {
//     Field   string
//     Value   interface{}
//     Message string
//     Time    time.Time
// }

// func (e ValidationError) Error() string {
//     return fmt.Sprintf("Validation error at %s: field '%s' - %s",
//         e.Time.Format("2016-05-02 15:04:05"),
//         e.Field,
//         e.Message,
//     )
// }

// func validateAge(age int) error {
//     if age < 0 {
//         return ValidationError{
//             Field:   "age",
//             Value:   age,
//             Message: "age cannot be negative",
//             Time:    time.Now(),
//         }
//     }
//     return nil
// }

// func main() {
//     err := validateAge(-15)
//     if err != nil {
//         fmt.Println(err)
//     }
// }

package main

import "fmt"

func riskyOperation() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()
    
    panic("something went wrong")
}

func main() {
    riskyOperation()
    fmt.Println("Program continues after panic recovery")
}
