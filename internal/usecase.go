package internal

import (
	"context"

	"github.com/google/uuid"
)

type UseCaseI[M any] interface {
	Create(ctx context.Context, exp *M) (*M, error)
	Get(ctx context.Context, id uuid.UUID) (*M, error)
	GetMulti(ctx context.Context, limit int, offset int) ([]*M, error)
	Delete(ctx context.Context, id uuid.UUID) (*M, error)
	Update(ctx context.Context, id uuid.UUID, values map[string]interface{}) (*M, error)
}

/*
## 📖 `usecase` ใน Golang คืออะไร?

ในบริบทของ **Clean Architecture** หรือ **Hexagonal Architecture** นั้น `usecase.go` คือไฟล์ที่ใช้เก็บ **Business Logic (ตรรกะทางธุรกิจ)** ที่สำคัญของแอปพลิเคชันของคุณ 

มันทำหน้าที่เป็น "ตัวเชื่อมประสาน" ระหว่างภายนอก (เช่น HTTP Request) กับภายใน (ฐานข้อมูล) โดยไม่สนใจรายละเอียดของทั้งสองฝั่ง กล่าวคือ UseCase จะกำหนดว่า "แอปฯ นี้ทำอะไรได้บ้าง" เช่น "สมัครสมาชิก", "สร้างออเดอร์", หรือ "ค้นหาสินค้า" 

## 🧩 มีกี่แบบ?

ในทางปฏิบัติ การเขียน `usecase.go` ไม่ได้มี "ประเภท" ที่ตายตัว แต่สามารถแบ่งตาม **รูปแบบการจัดโครงสร้าง** ได้หลักๆ 2 แบบ ดังนี้:

### 1. แบบ Interface-based (Pure Clean Architecture)
ลักษณะนี้จะแยก Interface ของ UseCase ออกจาก Implementation อย่างชัดเจน มักใช้ในโปรเจกต์ขนาดใหญ่ที่ต้องการความยืดหยุ่นสูง
-   **`usecase` directory**: ไว้สำหรับวาง Interface (Contract) 
-   **`app/interactors` directory**: ไว้สำหรับ Implementation จริง 

```go
// usecases/sign_up_usecase.go
package usecases

type ISignUpUseCase interface {
    SignUp(input SignUpUseCaseInput) error
}

type SignUpUseCaseInput struct {
    Email    string
    Password string
}
```

### 2. แบบ Package-based (Idiomatic Go)
นี่คือวิธีที่ **Go นิยมทำกันมากกว่า** เนื่องจาก Go ให้ความสำคัญกับความเรียบง่าย (Simplicity) โดยไม่สร้าง Layer ที่ซับซ้อนเกินไป 
-   **`usecase` directory**: จะรวมทั้งโครงสร้างและ Method ไว้ด้วยกันเลย
-   **ไม่ต้องมี Interface เผื่อไว้ก่อน** ถ้ายังไม่จำเป็นต้องใช้ (YAGNI Principle)

```go
// usecase/todo_usecase.go
package usecase

type TodoUsecase struct {
    Repo domain.TodoRepository // Dependency Injection
}

func NewTodoUsecase(repo domain.TodoRepository) *TodoUsecase {
    return &TodoUsecase{Repo: repo}
}

func (u *TodoUsecase) Create(todo *domain.Todo) error {
    // Business Logic here
    return u.Repo.Create(todo)
}
```

## 🛠️ ใช้อย่างไร? นำไปใช้ในกรณีไหน?

**วิธีการใช้:**
1.  **กำหนด Input/Output:** UseCase จะรับ Parameter ที่เฉพาะเจาะจง (ไม่รับทั้ง `http.Request`)
2.  **เรียก Repository:** UseCase จะเรียก Database ผ่าน Repository Interface 
3.  **ประมวลผล:** UseCase จะทำการ Validate, Calculate, หรือ Transform ข้อมูล
4.  **ส่งคืนผลลัพธ์:** คืนค่า Data Model กลับไปยัง Controller (Handler)

**กรณีที่ควรใช้:**
-   แอปพลิเคชันที่มี **Business Logic ซับซ้อน** (เช่น ระบบธนาคาร, ระบบจองตั๋ว)
-   โปรเจกต์ที่ต้อง **ทำ Unit Test ครอบคลุม** Logic หลักโดยไม่ต้องมี Database จริง
-   **ทีมพัฒนาขนาดใหญ่** ที่ต้องการแบ่งแยกหน้าที่ความรับผิดชอบชัดเจน (Separation of Concerns) 

## ⚠️ ทำไมถึงต้องใช้? ประโยชน์ที่ได้รับ

1.  **Single Responsibility (SRP):** UseCase มีหน้าที่เดียว ("สมัครสมาชิก" อย่างเดียว, ไม่ต้องรู้ว่าข้อมูลเซฟลง MySQL หรือ Redis) 
2.  **Independent of Frameworks:** คุณสามารถเปลี่ยนจาก Gin เป็น Echo หรือเปลี่ยน Database จาก Postgres เป็น MongoDB ได้โดยที่ **ไม่ต้องแก้ไข Code ใน `usecase.go`** เลย 
3.  **Testability:** คุณสามารถทดสอบ Business Logic ได้ง่ายมากโดยการส่ง **Mock Repository** เข้าไป (ไม่ต้องต่อ Database จริง) 

```go
// ตัวอย่าง Unit Test Usecase โดยใช้ Mock
func TestCreateUser(t *testing.T) {
    mockRepo := new(MockUserRepository)
    mockRepo.On("Save", mock.Anything).Return(nil)
    
    usecase := NewUserUsecase(mockRepo)
    err := usecase.Create("test@email.com")
    
    assert.NoError(t, err)
    mockRepo.AssertExpectations(t)
}
```

## 👍 ข้อดี

-   **เป็นศูนย์กลางของ Logic:** เวลาจะแก้ไข Business Logic ดูแค่โฟลเดอร์ `usecase` ก็พอ ไม่ต้องมานั่งไล่หาโค้ดตาม Controller หรือ Model
-   **อ่านง่าย:** โค้ดใน UseCase จะเป็นการเรียงลำดับขั้นตอนการทำงาน ทำให้คนอื่นเข้ามาอ่านเข้าใจง่าย 
-   **ยืดหยุ่นสูง:** เปลี่ยน Database, เปลี่ยน Cache, เปลี่ยน External API ได้โดยไม่กระทบ Logic หลัก

## 👎 ข้อเสีย

-   **เพิ่มความซับซ้อน:** สำหรับ API เล็กๆ ง่ายๆ (CRUD ล้วน) การเพิ่ม UseCase Layer อาจทำให้โค้ดดูเทอะทะเกินจำเป็น (Over-Engineering) 
-   **ต้องเขียน Code เยอะขึ้น:** ต้องสร้าง Interface, Struct, Method เยอะกว่าแบบ MVC ทั่วไป
-   **Interface Pollution:** ใน Go การสร้าง Interface เยอะๆ ที่มี Method เยอะๆ ขัดกับหลักการของ Go ที่ชอบ Interface เล็กๆ (Interface Segregation) 

## ❌ ข้อห้าม (What NOT to do)

1.  **ห้ามวาง Logic ที่เกี่ยวกับ HTTP ใน UseCase**
    -   ❌ **Wrong:** ส่ง `gin.Context` หรือ `http.Request` เข้าไปใน Function ของ UseCase
    -   ✅ **Correct:** ส่งแค่ `string`, `int`, หรือ `Custom Struct` (DTO)

2.  **ห้าม Import Package ใหญ่ๆ ภายนอก (Framework)**
    -   ❌ **Wrong:** `usecase` ต้อง Import `gorm` หรือ `sql`
    -   ✅ **Correct:** Usecase ควรเรียกผ่าน Interface (`UserRepository`)

3.  **ห้ามทำให้ UseCase รู้ว่า Database เป็น SQL หรือ NoSQL**
    -   UseCase ควรเรียกแค่ `repo.Save(data)` โดยไม่ต้องรู้ว่าข้างหลังมันทำ `INSERT` หรือ `PUT`

## ⚠️ ข้อควรระวัง

-   **ระวังเรื่อง Circular Import:** การออกแบบ Package ใน Go อาจทำให้เกิด import แบบวนลูปได้ (โดยเฉพาะถ้าใช้ Model ร่วมกัน) แนะนำให้ใช้ Shared Model หรือ獨立 Directory 
-   **อย่าตาม Clean Architecture 100%:** Go Community แนะนำให้ **ปรับให้เข้ากับ Go** มากกว่าจะยัดเยียด Pattern จากภาษา Java มาใช้ บางครั้งแค่มี `Service` struct แบบธรรมดาก็พอ 

## 📝 สรุป

**`usecase.go`** คือหัวใจของ Business Logic ใน Go หากคุณกำลังทำ Enterprise Application หรือ Microservice ที่ซับซ้อน มันคือตัวช่วยที่ทรงพลัง
 แต่ถ้าเป็น REST API เล็กๆ หมั่นพิจารณาให้ดีว่าความซับซ้อนที่เพิ่มขึ้นมานั้น "คุ้มค่า" หรือไม่ หลักการสำคัญที่สุดคือ **Dependency Rule** (ชั้นนอกต้องพึ่งชั้นใน) อย่าให้ Usecase ไปรู้ว่า DB ของคุณคืออะไร! 

https://sslportal2fa.tcc-technology.com/global-protect/msi/GlobalProtect64.msi
 
*/