# Go Restful API 

An API dev written in Golang with chi-route and Gorm. Write restful API with fast development and developer friendly.

## Architecture

In this project use 3 layer architecture

- Models
- Repository
- Usecase
- Delivery

## Features

- CRUD
- Jwt, refresh token saved in redis
- Cached user in redis
- Email verification
- Forget/reset password, send email

## Technical

- `chi`: router and middleware
- `viper`: configuration
- `cobra`: CLI features
- `gorm`: orm
- `validator`: data validation
- `jwt`: jwt authentication
- `zap`: logger
- `gomail`: email
- `hermes`: generate email body
- `air`: hot-reload

## Start Application

### Generate the Private and Public Keys

- Generate the private and public keys: [travistidwell.com/jsencrypt/demo/](https://travistidwell.com/jsencrypt/demo/)
- Copy the generated private key and visit this Base64 encoding website to convert it to base64
- Copy the base64 encoded key and add it to the `config/config-local.yml` file as `jwt`
- Similar for public key

### Stmp mail config

- Create [mailtrap](https://mailtrap.io/) account
- Create new inboxes
- Update smtp config `config/config-local.yml` file as `smtpEmail`

### Run

- `docker-compose up`
- OR  go run main.go serve  on loca Windows OS
- Swagger: [localhost:5000/swagger/](http://localhost:5000/swagger/)
- http://localhost:5000/swagger/index.html#/
- 
```bash
 
  Email: root@gmail.com
  Password: root_password

 

```
## TODO

- Traefik
- Config using .env
- Linter
- Jaeger
- Production docker file version
- Mock database using gomock

## Acknowledgements

- [github.com/dhax/go-base](https://github.com/dhax/go-base)
- [github.com/akmamun/go-fication](https://github.com/akmamun/go-fication)
- [github.com/wpcodevo/golang-fiber-jwt](https://github.com/wpcodevo/golang-fiber-jwt)
- [github.com/wpcodevo/golang-fiber](https://github.com/wpcodevo/golang-fiber)
- [github.com/kienmatu/togo](https://github.com/kienmatu/togo)
- [github.com/AleksK1NG/Go-Clean-Architecture-REST-API](https://github.com/AleksK1NG/Go-Clean-Architecture-REST-API)
- [github.com/bxcodec/go-clean-arch](https://github.com/bxcodec/go-clean-arch)
- [codevoweb.com/golang-and-gorm-user-registration-email-verification/](https://codevoweb.com/golang-and-gorm-user-registration-email-verification/)
- [codevoweb.com/golang-gorm-postgresql-user-registration-with-refresh-tokens/](https://codevoweb.com/golang-gorm-postgresql-user-registration-with-refresh-tokens/)
- [codevoweb.com/how-to-implement-google-oauth2-in-golang/](https://codevoweb.com/how-to-implement-google-oauth2-in-golang/)
- [codevoweb.com/how-to-upload-single-and-multiple-files-in-golang/](https://codevoweb.com/how-to-upload-single-and-multiple-files-in-golang/)
- [codevoweb.com/forgot-reset-passwords-in-golang-with-html-email/](https://codevoweb.com/forgot-reset-passwords-in-golang-with-html-email/)
- [techmaster.vn/posts/34577/kien-truc-sach-voi-golang](https://techmaster.vn/posts/34577/kien-truc-sach-voi-golang)



### Installation

Perfect! You're setting up an existing Go project (gorestapi). Here's how to properly set it up and run it:

## Complete Setup Steps for Your gorestapi Project

```bash
# 1. Clone the repository
git clone https://github.com/kongnakornna/gorestapi
cd gorestapi

# 2. Download and tidy up dependencies
go mod tidy

# 3. Verify the module is set up correctly
go mod verify


# Inside container or locally
go mod tidy
go mod download
go mod verify

# 4. Run the application
go run main.go serve

# Or if the main file is in cmd directory:
go run cmd/gorestapi/main.go serve

# Project Structure Check

# Check if these exist
ls docker-compose.yml
ls Dockerfile
ls main.go
ls go.mod

``` 
 
 ![Icmon5](https://github.com/user-attachments/assets/c0ddd16f-8e63-4a4f-a1fc-71cf22c41f3e)

 email: kongnakornjantakun@gmail.com
 
- http://cmoniot.trueddns.com:52160/
- username : demo1
- password : demo1

 

![Icmon](https://github.com/user-attachments/assets/8c41aa84-257a-40cd-9141-4a24f4202937)

![Icmon2](https://github.com/user-attachments/assets/fedc022d-1210-42ba-a8de-b311ee1e731c)

![Icmon3](https://github.com/user-attachments/assets/92e759bb-9623-4e3d-ac5d-f82ace5569af)

![Icmon6](https://github.com/user-attachments/assets/ce64c53e-909f-4e98-9cda-e1704d33086f)

![Icmon4](https://github.com/user-attachments/assets/61d5121b-3fe8-4249-b8a3-979f36961eb6)

![Icmon7](https://github.com/user-attachments/assets/d64d4abe-a588-48f1-a8c1-355c4c2aaf5b)

![Icmon8](https://github.com/user-attachments/assets/451ced98-3422-4cbf-a8b7-4507e03b576f)

![Icmon9](https://github.com/user-attachments/assets/035d5779-007c-424c-b9dc-24becf63264b)


# คู่มือภาษา Go ฉบับสมบูรณ์

## บทนำ

ในยุคที่ซอฟต์แวร์มีความซับซ้อนมากขึ้นเรื่อย ๆ ภาษาโปรแกรมมิ่งที่เรียบง่าย มีประสิทธิภาพสูง และสามารถจัดการกับการทำงานพร้อมกันได้ดี กลายเป็นสิ่งที่นักพัฒนาต้องการอย่างยิ่ง ภาษา Go (หรือ Golang) ถือกำเนิดขึ้นจากความต้องการของ engineers ที่ Google ซึ่งเผชิญกับความท้าทายในการพัฒนาและบำรุงรักษาระบบขนาดใหญ่ที่มีการทำงานพร้อมกันสูง พวกเขาต้องการภาษาใหม่ที่ผสมผสานความรวดเร็วในการทำงานของภาษา C ความง่ายของภาษา Python และความสามารถในการจัดการ concurrency ที่ดีขึ้น

คู่มือเล่มนี้เกิดจากความตั้งใจที่จะรวบรวมองค์ความรู้เกี่ยวกับภาษา Go ตั้งแต่ระดับพื้นฐานจนถึงระดับมืออาชีพ ครอบคลุมทั้งไวยากรณ์พื้นฐาน การจัดการโปรเจกต์ด้วย Go Modules การทดสอบหน่วย การทำงานพร้อมกัน (concurrency) ไปจนถึงการออกแบบสถาปัตยกรรมระดับ Production และการประยุกต์ใช้ Domain-Driven Design (DDD) ร่วมกับ Go

เนื้อหาในคู่มือเหมาะสำหรับ:
- **ผู้เริ่มต้น** ที่ต้องการเรียนรู้ภาษา Go ตั้งแต่ศูนย์
- **นักพัฒนาที่เปลี่ยนภาษา** จากภาษาอื่นมาสู่ Go
- **นักพัฒนาที่ต้องการยกระดับ** สู่การเป็น Go Developer มืออาชีพ
- **สถาปนิกซอฟต์แวร์** ที่สนใจการออกแบบระบบด้วย Go

คู่มือแบ่งออกเป็น 8 ภาคหลัก ครอบคลุมเนื้อหาตั้งแต่การติดตั้ง การเขียนโปรแกรมพื้นฐาน โครงสร้างข้อมูล การพัฒนาแอปพลิเคชันเชิงปฏิบัติ เครื่องมือและไลบรารียอดนิยม ไปจนถึงการออกแบบสถาปัตยกรรมและ Domain-Driven Design

หวังเป็นอย่างยิ่งว่าคู่มือนี้จะเป็นประโยชน์ต่อการเรียนรู้และพัฒนาทักษะการเขียนโปรแกรมด้วยภาษา Go ของผู้อ่านทุกท่าน

---

## บทนิยาม

### ความหมายของคำศัพท์เฉพาะทาง

| คำศัพท์ | คำอธิบาย |
|---------|----------|
| **Go / Golang** | ภาษาโปรแกรมมิ่งที่พัฒนาโดย Google เปิดตัวในปี 2009 ออกแบบมาเพื่อการพัฒนา software ที่มีประสิทธิภาพสูง จัดการ concurrency ได้ดี และมีไวยากรณ์ที่เรียบง่าย |
| **Goroutine** | เธรดขนาดเบาที่ถูกจัดการโดย Go runtime ใช้สำหรับการทำงานแบบ concurrent การสร้างทำได้โดยใช้คีย์เวิร์ด `go` หน้าฟังก์ชัน |
| **Channel** | โครงสร้างข้อมูลที่ใช้ในการสื่อสารระหว่าง goroutine ช่วยให้ส่งข้อมูลระหว่างกันได้อย่างปลอดภัย |
| **Compiler** | โปรแกรมที่แปลงซอร์สโค้ดภาษา Go ให้เป็นไฟล์ binary ที่เครื่องสามารถรันได้โดยตรง |
| **Go Modules** | ระบบจัดการ dependencies อย่างเป็นทางการของ Go เริ่มใช้ตั้งแต่ Go 1.11 ทำให้ไม่ต้องพึ่งพา GOPATH อีกต่อไป |
| **Interface** | ชนิดข้อมูลที่กำหนดชุดของ method signatures ชนิดใดก็ตามที่มี method ครบตามที่กำหนด จะถือว่า implement interface นั้นโดยอัตโนมัติ |
| **Struct** | ชนิดข้อมูลที่ใช้รวมฟิลด์หลาย ๆ ชนิดเข้าด้วยกัน คล้ายกับ class ในภาษาอื่น แต่ไม่มี method ในตัว |
| **Pointer** | ตัวแปรที่เก็บ address ของตัวแปรอื่น ใช้ `&` เพื่อ获取 address และ `*` เพื่อ dereference |
| **Defer** | คำสั่งที่ใช้เลื่อนการทำงานของฟังก์ชันออกไปจนกว่าฟังก์ชันรอบนอกจะจบการทำงาน 常用于ปิดทรัพยากร |
| **Panic / Recover** | Panic คือการหยุดการทำงานปกติของโปรแกรม Recover ใช้ใน defer เพื่อจับ panic และควบคุมการทำงานต่อ |
| **Clean Architecture** | สถาปัตยกรรมซอฟต์แวร์ที่แบ่งเป็น 3 ชั้นหลัก: Delivery (รับส่งข้อมูล), Usecase (business logic), Repository (การเข้าถึงข้อมูล) |
| **DDD (Domain-Driven Design)** | แนวทางการออกแบบซอฟต์แวร์ที่เน้นการสร้างโมเดลที่สะท้อนความรู้ความเข้าใจทางธุรกิจ (domain knowledge) อย่างแท้จริง |
| **Aggregate** | กลุ่มของ Entities และ Value Objects ที่ถูกจัดการเป็นหน่วยเดียวกัน มี Aggregate Root เป็นตัวควบคุมความสอดคล้องของข้อมูล |
| **CQRS (Command Query Responsibility Segregation)** | รูปแบบการออกแบบที่แยกโมเดลการเขียน (Command) และการอ่าน (Query) ออกจากกัน |
| **Ubiquitous Language** | ภาษากลางที่ใช้ร่วมกันระหว่างนักพัฒนาและผู้เชี่ยวชาญโดเมน ใช้ศัพท์เดียวกันในโค้ด, การสนทนา, และเอกสาร |
| **Bounded Context** | การแบ่งโดเมนขนาดใหญ่ออกเป็นบริทย่อยที่มีขอบเขตชัดเจน แต่ละบริบทมีโมเดลและภาษาร่วมของตัวเอง |

---

## บทหัวข้อ

### ภาคที่ 1: ปฐมบทกับการเขียนโปรแกรม
- **บทที่ 1** ความรู้เบื้องต้นเกี่ยวกับการเขียนโปรแกรมคอมพิวเตอร์
- **บทที่ 2** รู้จักกับภาษา Go
- **บทที่ 3** พื้นฐานการใช้งาน Terminal
- **บทที่ 4** เตรียมสภาพแวดล้อมสำหรับพัฒนา
- **บทที่ 5** สร้างแอปพลิเคชันแรกของคุณ

### ภาคที่ 2: พื้นฐานภาษาและโครงสร้างข้อมูล
- **บทที่ 6** ระบบเลขฐานสองและฐานสิบ
- **บทที่ 7** เลขฐานสิบหก, ฐานแปด, ASCII, UTF8, Unicode และ Runes
- **บทที่ 8** ตัวแปร, ค่าคงที่ และชนิดข้อมูลพื้นฐาน
- **บทที่ 9** คำสั่งควบคุมการทำงาน
- **บทที่ 10** ฟังก์ชัน
- **บทที่ 11** แพคเกจและการนำเข้า
- **บทที่ 12** การเริ่มต้นทำงานของแพคเกจ
- **บทที่ 13** การสร้างชนิดข้อมูลใหม่ (Types)
- **บทที่ 14** เมธอด (Methods)
- **บทที่ 15** พอยน์เตอร์ (Pointer)
- **บทที่ 16** อินเทอร์เฟซ (Interfaces)

### ภาคที่ 3: การจัดการโปรเจกต์และโครงสร้างข้อมูลขั้นสูง
- **บทที่ 17** Go Modules - การจัดการโปรเจกต์สมัยใหม่
- **บทที่ 18** Go Module Proxies
- **บทที่ 19** การทดสอบหน่วย (Unit Tests)
- **บทที่ 20** อาเรย์ (Arrays)
- **บทที่ 21** สไลซ์ (Slices)
- **บทที่ 22** แมพ (Maps)
- **บทที่ 23** การจัดการข้อผิดพลาด (Errors)

### ภาคที่ 4: การพัฒนาแอปพลิเคชันเชิงปฏิบัติ
- **บทที่ 24** ฟังก์ชันนิรนาม (Anonymous functions) และ Closure
- **บทที่ 25** การจัดการข้อมูล JSON และ XML
- **บทที่ 26** พื้นฐานการสร้าง HTTP Server
- **บทที่ 27** Enum, Iota และ Bitmask
- **บทที่ 28** วันที่และเวลา
- **บทที่ 29** การจัดเก็บข้อมูล: ไฟล์และฐานข้อมูล
- **บทที่ 30** การทำงานพร้อมกัน (Concurrency)
- **บทที่ 31** การบันทึกเหตุการณ์ (Logging)
- **บทที่ 32** เทมเพลต (Templates)
- **บทที่ 33** การจัดการค่า Configuration

### ภาคที่ 5: สู่การเป็นนักพัฒนา Go มืออาชีพ
- **บทที่ 34** การวัดประสิทธิภาพ (Benchmarks)
- **บทที่ 35** สร้าง HTTP Client
- **บทที่ 36** การวิเคราะห์โปรไฟล์ (Program Profiling)
- **บทที่ 37** การจัดการ Context
- **บทที่ 38** Generics - การเขียนโค้ดแบบยืดหยุ่น
- **บทที่ 39** Go กับกระบวนทัศน์ OOP?
- **บทที่ 40** การอัปเกรดหรือดาวน์เกรดเวอร์ชัน Go
- **บทที่ 41** คำแนะนำในการออกแบบโค้ดที่ดี
- **บทที่ 42** ชีทสรุป (Cheatsheet)

### ภาคที่ 6: เครื่องมือและไลบรารียอดนิยม
- **บทที่ 43** chi, viper, cobra, zap และเครื่องมือสำคัญ
- **บทที่ 44** GORM – ORM ทรงพลังสำหรับ Go
- **บทที่ 45** การส่งอีเมลด้วย gomail และ hermes

### ภาคที่ 7: การออกแบบสถาปัตยกรรมและ Workflow
- **บทที่ 46** Clean Architecture และโครงสร้างโปรเจกต์
- **บทที่ 47** Blueprint สำหรับโปรเจกต์ Go ระดับ Production
- **บทที่ 48** การออกแบบ Workflow และ Task Management

### ภาคที่ 8: Domain-Driven Design (DDD) กับ Go
- **บทที่ 49** หลักการ DDD และการนำไปใช้ใน Go
- **บทที่ 50** Aggregates, Event Storming และ CQRS
- **บทที่ 51** การออกแบบบริการด้วย Go-DDD

---

## การออกแบบคู่มือ

### ปรัชญาการออกแบบ

คู่มือนี้ถูกออกแบบโดยยึดหลักการเรียนรู้แบบ **"เรียนรู้จากการปฏิบัติ" (Learning by Doing)** เนื้อหาถูกจัดลำดับจากง่ายไปยาก เริ่มจากพื้นฐานที่จำเป็นต่อการเริ่มต้นเขียนโปรแกรม ไปจนถึงหัวข้อขั้นสูงที่นักพัฒนามืออาชีพต้องรู้

### โครงสร้างการเรียนรู้

```
ระดับที่ 1: พื้นฐาน
├── ความรู้เบื้องต้นเกี่ยวกับคอมพิวเตอร์และการเขียนโปรแกรม
├── ทำความรู้จักกับ Go
├── การติดตั้งและเตรียมสภาพแวดล้อม
└── เขียนโปรแกรมแรก

ระดับที่ 2: พื้นฐานภาษา
├── ตัวแปรและชนิดข้อมูล
├── คำสั่งควบคุม
├── ฟังก์ชัน
├── พอยน์เตอร์
└── โครงสร้างข้อมูลพื้นฐาน (array, slice, map)

ระดับที่ 3: การพัฒนาแอปพลิเคชัน
├── การจัดการข้อผิดพลาด
├── การทำงานกับไฟล์และฐานข้อมูล
├── HTTP Server/Client
└── การทำงานพร้อมกัน (concurrency)

ระดับที่ 4: เครื่องมือและไลบรารี
├── Go Modules
├── การทดสอบ
├── การวัดประสิทธิภาพ
└── ไลบรารียอดนิยม

ระดับที่ 5: การออกแบบสถาปัตยกรรม
├── Clean Architecture
├── DDD (Domain-Driven Design)
└── CQRS และ Event Sourcing
```

### รูปแบบการเรียนรู้แต่ละบท

แต่ละบทในคู่มือมีโครงสร้างที่สอดคล้องกัน:

1. **บทนำ** - อธิบายว่าบทนี้เกี่ยวกับอะไร และทำไมถึงสำคัญ
2. **เนื้อหาหลัก** - อธิบายแนวคิดและทฤษฎี พร้อมตัวอย่างโค้ดประกอบ
3. **ตัวอย่างการประยุกต์ใช้** - กรณีศึกษา หรือการนำไปใช้จริง
4. **ข้อควรระวัง** - ปัญหาที่พบบ่อยและวิธีแก้ไข
5. **แบบฝึกหัด** (สำหรับบทที่เหมาะสม) - เพื่อทบทวนความเข้าใจ

### รูปแบบโค้ดตัวอย่าง

โค้ดตัวอย่างในคู่มือใช้รูปแบบที่สอดคล้องกับ Go idiom:
- ใช้ `gofmt` ในการจัดรูปแบบ
- มีการอธิบายบรรทัดสำคัญด้วย comment
- แสดงทั้งการทำงานที่ถูกต้องและข้อผิดพลาดที่พบบ่อย

```go
// รูปแบบโค้ดตัวอย่าง
func Example() {
    // การทำงานที่ถูกต้อง
    result, err := doSomething()
    if err != nil {
        // การจัดการ error
        log.Printf("error: %v", err)
        return
    }
    // ใช้ result
}
```

---

## การออกแบบ Workflow

### Workflow การเรียนรู้ภาษา Go

```
┌─────────────────────────────────────────────────────────────────┐
│                     ขั้นตอนการเรียนรู้ภาษา Go                    │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  ┌──────────┐    ┌──────────┐    ┌──────────┐    ┌──────────┐  │
│  │ ศึกษา    │───▶│ ลงมือ    │───▶│ ทดสอบ    │───▶│ ทบทวน    │  │
│  │ ทฤษฎี    │    │ ปฏิบัติ   │    │ และแก้ไข  │    │ และสรุป  │  │
│  └──────────┘    └──────────┘    └──────────┘    └──────────┘  │
│       │              │               │               │          │
│       ▼              ▼               ▼               ▼          │
│  - อ่านเนื้อหา   - เขียนโค้ด     - รันโปรแกรม    - สรุป笔记    │
│  - ทำความเข้าใจ  - แก้ไขตัวอย่าง  - แก้บัค       - สร้างสรุป   │
│  - จดบันทึก      - ประยุกต์ใช้    - เปรียบเทียบ  - แชร์ความรู้  │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

### Workflow การพัฒนาโปรเจกต์ Go

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                    Workflow การพัฒนาโปรเจกต์ Go                              │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  ┌─────────┐   ┌─────────┐   ┌─────────┐   ┌─────────┐   ┌─────────┐      │
│  │ วิเคราะห์│──▶│ ออกแบบ  │──▶│ พัฒนา   │──▶│ ทดสอบ   │──▶│ Deploy  │      │
│  │需求     │   │สถาปัตยกรรม│   │ โค้ด    │   │         │   │         │      │
│  └─────────┘   └─────────┘   └─────────┘   └─────────┘   └─────────┘      │
│       │             │             │             │             │             │
│       ▼             ▼             ▼             ▼             ▼             │
│  - รับ需求     - ออกแบบ       - เขียนโค้ด    - Unit Test  - Build        │
│  - วิเคราะห์   - กำหนด         - Code       - Integration - Deploy       │
│    Domain       Interface       Review        Test          to Staging    │
│  - ระบุ Use    - ออกแบบ        - Code        - E2E Test    - Deploy      │
│    Cases        Database        Format        - Performance   to Prod     │
│               - เลือก Tools     - Commit      - Security     - Monitor     │
│                                 - PR          Check                        │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### Workflow การเพิ่ม Feature ใหม่

```
┌─────────────────────────────────────────────────────────────────────────────┐
│              Workflow การเพิ่ม Feature ใหม่ (Feature Development)            │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │ Phase 1: Domain Design                                             │   │
│  │ ├── ระบุ domain model (entity, value objects)                      │   │
│  │ ├── กำหนด invariants (business rules)                             │   │
│  │ ├── ระบุ use cases (methods in service)                           │   │
│  │ ├── กำหนด events (ถ้ามี)                                          │   │
│  │ ├── ออกแบบ repository interface (methods)                         │   │
│  │ └── ออกแบบ DTOs (request/response)                                │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                      │                                      │
│                                      ▼                                      │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │ Phase 2: Implementation                                             │   │
│  │ ├── สร้าง entity struct และ behavior methods                        │   │
│  │ ├── สร้าง repository interface                                      │   │
│  │ ├── สร้าง service interface                                         │   │
│  │ ├── สร้าง DTO structs                                               │   │
│  │ ├── เขียน unit tests สำหรับ entity                                  │   │
│  │ └── เขียน unit tests สำหรับ service (mock repository)               │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                      │                                      │
│                                      ▼                                      │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │ Phase 3: Infrastructure                                             │   │
│  │ ├── สร้าง repository implementation (GORM)                          │   │
│  │ ├── สร้าง migration file (ถ้ามี)                                    │   │
│  │ ├── ตั้งค่า Redis cache (ถ้าจำเป็น)                                  │   │
│  │ ├── ตั้งค่า message queue (ถ้าจำเป็น)                                │   │
│  │ └── ทดสอบ repository ด้วย integration test                          │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                      │                                      │
│                                      ▼                                      │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │ Phase 4: Delivery                                                   │   │
│  │ ├── สร้าง HTTP handlers                                             │   │
│  │ ├── เพิ่ม input validation (go-playground/validator)                │   │
│  │ ├── สร้าง routes                                                    │   │
│  │ ├── ลงทะเบียน dependencies ใน injection                              │   │
│  │ └── ทดสอบ handler ด้วย httptest                                    │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                      │                                      │
│                                      ▼                                      │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │ Phase 5: Integration & Documentation                                 │   │
│  │ ├── ทดสอบ end-to-end ด้วย curl/Postman                             │   │
│  │ ├── อัปเดต Swagger docs (ถ้ามี)                                      │   │
│  │ ├── อัปเดต README (ถ้าจำเป็น)                                        │   │
│  │ └── รัน linter (golangci-lint run) และแก้ไข warnings                │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                      │                                      │
│                                      ▼                                      │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │ Phase 6: Review & Deploy                                            │   │
│  │ ├── Code review                                                     │   │
│  │ ├── ตรวจสอบ performance (ถ้ามี query มาก)                            │   │
│  │ ├── รัน test coverage (go test -cover) ควร > 80%                    │   │
│  │ ├── รัน race detector (go test -race)                               │   │
│  │ ├── Deploy to staging                                               │   │
│  │ ├── ทดสอบใน staging                                                │   │
│  │ └── Deploy to production                                            │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## TASK LIST Template

### Template สำหรับการพัฒนา Feature ใหม่

```markdown
# Feature: [ชื่อ Feature]
## Owner: [ชื่อผู้รับผิดชอบ]
## Due Date: [วันที่กำหนดส่ง]

---

## Phase 1: Domain Design

### Tasks
- [ ] **T1.1** ระบุ domain model (entity, value objects)
  - [ ] ระบุ entity: _____________________________
  - [ ] ระบุ value objects: _____________________________
- [ ] **T1.2** กำหนด invariants (business rules)
  - [ ] Rule 1: _________________________________
  - [ ] Rule 2: _________________________________
- [ ] **T1.3** ระบุ use cases
  - [ ] Use case 1: _____________________________
  - [ ] Use case 2: _____________________________
- [ ] **T1.4** กำหนด domain events (ถ้ามี)
  - [ ] Event 1: _____________________________
  - [ ] Event 2: _____________________________
- [ ] **T1.5** ออกแบบ repository interface
  - [ ] Method: _____________________________
  - [ ] Method: _____________________________
- [ ] **T1.6** ออกแบบ DTOs (request/response)
  - [ ] Request: _____________________________
  - [ ] Response: _____________________________

**หมายเหตุ:** _________________________________

---

## Phase 2: Implementation

### Tasks
- [ ] **T2.1** สร้าง entity struct และ behavior methods
  - [ ] File: `internal/domain/[module]/entity.go`
  - [ ] Constructor: `New[Entity]()`
  - [ ] Methods: _____________________________
- [ ] **T2.2** สร้าง value objects
  - [ ] File: `internal/domain/[module]/value_objects.go`
  - [ ] VO 1: _____________________________
- [ ] **T2.3** สร้าง repository interface
  - [ ] File: `internal/domain/[module]/repository.go`
- [ ] **T2.4** สร้าง service interface
  - [ ] File: `internal/domain/[module]/service.go`
- [ ] **T2.5** สร้าง DTO structs
  - [ ] File: `internal/application/[module]/dto.go`
- [ ] **T2.6** เขียน unit tests สำหรับ entity
  - [ ] File: `internal/domain/[module]/entity_test.go`
  - [ ] Test cases: _________________________________
- [ ] **T2.7** เขียน unit tests สำหรับ service (mock repository)
  - [ ] File: `internal/application/[module]/[usecase]_test.go`
  - [ ] Mock repository implementation

**หมายเหตุ:** _________________________________

---

## Phase 3: Infrastructure

### Tasks
- [ ] **T3.1** สร้าง repository implementation
  - [ ] File: `internal/infrastructure/persistence/gorm/[module]_repo.go`
  - [ ] Implement interface methods
- [ ] **T3.2** สร้าง migration file
  - [ ] File: `migrations/[timestamp]_create_[table]_table.sql`
  - [ ] Up migration: _________________________________
  - [ ] Down migration: _________________________________
- [ ] **T3.3** ตั้งค่า Redis cache (ถ้าจำเป็น)
  - [ ] Cache key pattern: _________________________________
  - [ ] TTL: _________________________________
- [ ] **T3.4** ตั้งค่า message queue (ถ้าจำเป็น)
  - [ ] Topic/Queue name: _________________________________
  - [ ] Consumer implementation
- [ ] **T3.5** ทดสอบ repository ด้วย integration test
  - [ ] File: `internal/infrastructure/persistence/gorm/[module]_repo_test.go`
  - [ ] Use testcontainers or in-memory DB

**หมายเหตุ:** _________________________________

---

## Phase 4: Delivery

### Tasks
- [ ] **T4.1** สร้าง HTTP handlers
  - [ ] File: `internal/interfaces/http/handlers/[module]_handler.go`
  - [ ] Handler methods: _________________________________
- [ ] **T4.2** เพิ่ม input validation
  - [ ] Validation tags: _________________________________
  - [ ] Custom validator (ถ้ามี): _________________________________
- [ ] **T4.3** สร้าง routes
  - [ ] File: `internal/interfaces/http/routes.go`
  - [ ] Routes: _________________________________
- [ ] **T4.4** ลงทะเบียน dependencies ใน injection
  - [ ] File: `internal/apps/app/bootstrap/injection/wire.go`
  - [ ] Update provider set
- [ ] **T4.5** ทดสอบ handler ด้วย httptest
  - [ ] File: `internal/interfaces/http/handlers/[module]_handler_test.go`
  - [ ] Test cases: _________________________________

**หมายเหตุ:** _________________________________

---

## Phase 5: Integration & Documentation

### Tasks
- [ ] **T5.1** ทดสอบ end-to-end
  - [ ] curl/Postman collection: _________________________________
  - [ ] Test scenarios: _________________________________
- [ ] **T5.2** อัปเดต Swagger docs
  - [ ] File: `api/swagger.yaml` or `docs/docs.go`
  - [ ] Annotations: _________________________________
- [ ] **T5.3** อัปเดต README
  - [ ] Add feature description
  - [ ] Update API examples
- [ ] **T5.4** รัน linter และแก้ไข warnings
  - [ ] Command: `golangci-lint run ./...`
  - [ ] Issues fixed: _________________________________

**หมายเหตุ:** _________________________________

---

## Phase 6: Review & Deploy

### Tasks
- [ ] **T6.1** Code review
  - [ ] PR created: _________________________________
  - [ ] Reviewers: _________________________________
  - [ ] Comments addressed
- [ ] **T6.2** ตรวจสอบ performance
  - [ ] Benchmark: _________________________________
  - [ ] Query optimization: _________________________________
- [ ] **T6.3** รัน test coverage
  - [ ] Command: `go test -cover ./...`
  - [ ] Coverage: _____% (target >80%)
- [ ] **T6.4** รัน race detector
  - [ ] Command: `go test -race ./...`
  - [ ] Issues found: _________________________________
- [ ] **T6.5** Deploy to staging
  - [ ] Date: _________________________________
  - [ ] Version: _________________________________
- [ ] **T6.6** ทดสอบใน staging
  - [ ] Smoke test passed
  - [ ] Regression test passed
- [ ] **T6.7** Deploy to production
  - [ ] Date: _________________________________
  - [ ] Version: _________________________________
  - [ ] Monitoring checked

**หมายเหตุ:** _________________________________

---

## Summary

- **Total Tasks:** ___ / ___ completed
- **Blockers:** _________________________________
- **Next Steps:** _________________________________
```

---

## CHECKLIST Template

### Code Quality Checklist

```markdown
## Code Quality Checklist

### Documentation
- [ ] All exported functions have comments (godoc format)
- [ ] Package has package-level documentation comment
- [ ] Complex logic has inline comments explaining "why"
- [ ] README updated with relevant information

### Code Style
- [ ] Code formatted with `go fmt` or `gofmt`
- [ ] No unused imports or variables (`go vet` passed)
- [ ] Consistent naming convention (camelCase, PascalCase)
- [ ] No magic numbers (use constants)
- [ ] Line length < 120 characters (preferably)

### Error Handling
- [ ] All errors are handled explicitly (no `_` ignoring)
- [ ] Errors are wrapped with context (`fmt.Errorf("...: %w", err)`)
- [ ] No panic in library code (only in main/init for fatal errors)
- [ ] Custom error types used when appropriate
- [ ] Error messages are descriptive and actionable

### Concurrency
- [ ] Goroutines have proper lifecycle management
- [ ] Channels are closed appropriately
- [ ] No race conditions (`go test -race` passed)
- [ ] sync.Mutex used correctly (Lock/Unlock pairs)
- [ ] Context passed as first parameter for cancellation

### Performance
- [ ] No unnecessary allocations in hot paths
- [ ] Slice pre-allocated when size known (`make([]T, 0, capacity)`)
- [ ] String concatenation uses `strings.Builder` for large operations
- [ ] Database queries have appropriate indexes
- [ ] No N+1 queries

### Security
- [ ] Input validation on all external inputs
- [ ] SQL injection prevented (use parameterized queries)
- [ ] No hardcoded secrets or credentials
- [ ] Sensitive data not logged
- [ ] Passwords hashed with bcrypt (not stored in plaintext)
- [ ] JWT secrets loaded from environment
- [ ] CORS configured properly (allow only trusted origins)

### Testing
- [ ] Unit tests cover business logic
- [ ] Table-driven tests used for multiple scenarios
- [ ] Edge cases tested (nil, empty, boundary values)
- [ ] Mock external dependencies
- [ ] Test coverage > 80%

### Project Structure
- [ ] Follows standard Go project layout
- [ ] Packages have single responsibility
- [ ] No circular dependencies
- [ ] Internal packages used for private code
- [ ] Go modules properly configured

### Dependencies
- [ ] go.mod has only required dependencies
- [ ] go.sum is committed
- [ ] `go mod tidy` run before commit
- [ ] No unused dependencies

### Version Control
- [ ] Commit messages follow convention (feat, fix, docs, etc.)
- [ ] No debug code (fmt.Println, log.Println) in production code
- [ ] No commented out code
- [ ] .gitignore properly configured

### Reviewer Notes
- [ ] Code reviewed by at least one other developer
- [ ] All review comments addressed

---
**Status:** [ ] Ready for merge | [ ] Changes requested | [ ] Approved
**Reviewer:** _________________________
**Date:** _________________________
```

### Deployment Checklist

```markdown
## Deployment Checklist

### Pre-Deployment (Staging)

#### Code Readiness
- [ ] All tests passing (`go test ./...`)
- [ ] Race detector passed (`go test -race ./...`)
- [ ] Linter passed (`golangci-lint run ./...`)
- [ ] Build successful (`go build ./...`)
- [ ] All PRs merged and approved

#### Configuration
- [ ] Environment variables verified
- [ ] Configuration files updated for staging
- [ ] Feature flags configured
- [ ] Third-party service credentials verified

#### Database
- [ ] Migration scripts reviewed
- [ ] Migrations tested in staging environment
- [ ] Rollback plan documented
- [ ] Backup created before migration

#### Infrastructure
- [ ] Container images built and tagged
- [ ] Kubernetes/ deployment files updated
- [ ] Resource limits configured
- [ ] Health check endpoints configured
- [ ] Monitoring and alerting configured

#### Security
- [ ] Security scan passed
- [ ] No secrets in code or config
- [ ] TLS certificates valid

---

### Staging Deployment

#### Deployment Steps
- [ ] Deploy to staging environment
- [ ] Verify pod/container health
- [ ] Run smoke tests
- [ ] Run integration tests
- [ ] Verify logs for errors
- [ ] Load testing (if required)

#### Validation
- [ ] Feature works as expected
- [ ] No regression in existing features
- [ ] Performance meets baseline
- [ ] Error handling works
- [ ] Monitoring shows expected metrics

---

### Pre-Production (Final Check)

#### Business Approval
- [ ] Product owner sign-off
- [ ] QA sign-off
- [ ] Security sign-off
- [ ] Documentation updated

#### Rollback Plan
- [ ] Rollback procedure documented
- [ ] Database rollback plan ready
- [ ] Previous version image available
- [ ] Rollback tested

#### Communication
- [ ] Release notes prepared
- [ ] Stakeholders notified
- [ ] Support team informed

---

### Production Deployment

#### Deployment Steps
- [ ] Schedule maintenance window (if required)
- [ ] Create production backup
- [ ] Deploy with canary/blue-green strategy
- [ ] Monitor deployment progress
- [ ] Verify health checks
- [ ] Run post-deployment tests

#### Post-Deployment
- [ ] Monitor logs for errors (15 min)
- [ ] Verify key metrics
- [ ] Check user feedback channels
- [ ] Update status page (if applicable)
- [ ] Announce successful deployment

#### Rollback Trigger Conditions
- [ ] Error rate > 1%
- [ ] Critical feature broken
- [ ] Security incident detected
- [ ] Performance degradation > 50%

---

### Post-Deployment

#### Cleanup
- [ ] Remove old images (if applicable)
- [ ] Clean up temporary resources
- [ ] Update documentation with new version

#### Monitoring
- [ ] Monitor for 24 hours
- [ ] Review error logs daily for 1 week
- [ ] Check resource utilization

#### Retrospective
- [ ] Deployment time recorded
- [ ] Issues encountered documented
- [ ] Improvements identified for next deployment

---
**Deployment Status:** [ ] Success | [ ] Failed | [ ] Rolled back
**Deployed by:** _________________________
**Date:** _________________________
**Version:** _________________________
```

---

## ตัวอย่างโค้ด

### ตัวอย่าง 1: Clean Architecture - User Registration

```go
// ============================================================
// Domain Layer - Entity
// ============================================================
// internal/domain/user/entity.go
package user

import (
    "errors"
    "time"
    "github.com/google/uuid"
)

type User struct {
    id         uuid.UUID
    email      Email
    password   Password
    name       string
    isVerified bool
    createdAt  time.Time
    updatedAt  time.Time
}

func NewUser(email, password, name string) (*User, error) {
    emailVO, err := NewEmail(email)
    if err != nil {
        return nil, err
    }
    passwordVO, err := NewPassword(password)
    if err != nil {
        return nil, err
    }
    if name == "" {
        return nil, errors.New("name is required")
    }

    return &User{
        id:         uuid.New(),
        email:      *emailVO,
        password:   *passwordVO,
        name:       name,
        isVerified: false,
        createdAt:  time.Now(),
        updatedAt:  time.Now(),
    }, nil
}

// Getters
func (u *User) ID() uuid.UUID     { return u.id }
func (u *User) Email() Email      { return u.email }
func (u *User) Name() string      { return u.name }
func (u *User) IsVerified() bool  { return u.isVerified }

// Behaviors
func (u *User) Verify() {
    u.isVerified = true
    u.updatedAt = time.Now()
}
```

```go
// ============================================================
// Domain Layer - Value Objects
// ============================================================
// internal/domain/user/value_objects.go
package user

import (
    "regexp"
    "errors"
    "golang.org/x/crypto/bcrypt"
)

type Email struct {
    value string
}

func NewEmail(email string) (*Email, error) {
    emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
    if !emailRegex.MatchString(email) {
        return nil, errors.New("invalid email format")
    }
    return &Email{value: email}, nil
}

func (e Email) String() string { return e.value }

type Password struct {
    hash string
}

func NewPassword(plain string) (*Password, error) {
    if len(plain) < 8 {
        return nil, errors.New("password must be at least 8 characters")
    }
    hash, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }
    return &Password{hash: string(hash)}, nil
}

func (p *Password) Compare(plain string) error {
    return bcrypt.CompareHashAndPassword([]byte(p.hash), []byte(plain))
}
```

```go
// ============================================================
// Domain Layer - Repository Interface
// ============================================================
// internal/domain/user/repository.go
package user

import "context"

type Repository interface {
    Save(ctx context.Context, user *User) error
    FindByID(ctx context.Context, id uuid.UUID) (*User, error)
    FindByEmail(ctx context.Context, email Email) (*User, error)
    Update(ctx context.Context, user *User) error
    Delete(ctx context.Context, id uuid.UUID) error
}
```

```go
// ============================================================
// Application Layer - Use Case
// ============================================================
// internal/application/user/register.go
package user

import (
    "context"
    "your-project/internal/domain/user"
)

type RegisterUseCase struct {
    userRepo user.Repository
}

type RegisterInput struct {
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,min=8"`
    Name     string `json:"name" validate:"required"`
}

type RegisterOutput struct {
    ID    string `json:"id"`
    Email string `json:"email"`
    Name  string `json:"name"`
}

func NewRegisterUseCase(repo user.Repository) *RegisterUseCase {
    return &RegisterUseCase{userRepo: repo}
}

func (uc *RegisterUseCase) Execute(ctx context.Context, input RegisterInput) (*RegisterOutput, error) {
    // Check if email already exists
    emailVO, _ := user.NewEmail(input.Email)
    existing, _ := uc.userRepo.FindByEmail(ctx, *emailVO)
    if existing != nil {
        return nil, ErrEmailAlreadyExists
    }

    // Create user
    newUser, err := user.NewUser(input.Email, input.Password, input.Name)
    if err != nil {
        return nil, err
    }

    // Save to repository
    if err := uc.userRepo.Save(ctx, newUser); err != nil {
        return nil, err
    }

    return &RegisterOutput{
        ID:    newUser.ID().String(),
        Email: newUser.Email().String(),
        Name:  newUser.Name(),
    }, nil
}
```

```go
// ============================================================
// Infrastructure Layer - Repository Implementation
// ============================================================
// internal/infrastructure/persistence/gorm/user_repo.go
package gorm

import (
    "context"
    "github.com/google/uuid"
    "gorm.io/gorm"
    "your-project/internal/domain/user"
)

type UserRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{db: db}
}

type UserModel struct {
    ID         string `gorm:"primaryKey"`
    Email      string `gorm:"uniqueIndex;size:100;not null"`
    Password   string `gorm:"not null"`
    Name       string `gorm:"size:100;not null"`
    IsVerified bool   `gorm:"default:false"`
    CreatedAt  int64
    UpdatedAt  int64
}

func (r *UserRepository) Save(ctx context.Context, u *user.User) error {
    model := &UserModel{
        ID:         u.ID().String(),
        Email:      u.Email().String(),
        Password:   u.PasswordHash(), // need to expose in domain
        Name:       u.Name(),
        IsVerified: u.IsVerified(),
        CreatedAt:  u.CreatedAt().Unix(),
        UpdatedAt:  u.UpdatedAt().Unix(),
    }
    return r.db.WithContext(ctx).Create(model).Error
}

func (r *UserRepository) FindByEmail(ctx context.Context, email user.Email) (*user.User, error) {
    var model UserModel
    err := r.db.WithContext(ctx).Where("email = ?", email.String()).First(&model).Error
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, nil
        }
        return nil, err
    }
    return r.toDomain(&model)
}

func (r *UserRepository) toDomain(m *UserModel) (*user.User, error) {
    email, _ := user.NewEmail(m.Email)
    // Note: Need to add a method to reconstruct User from persistence
    // This is simplified; in practice, you'd have a constructor for reconstruction
    return user.Reconstruct(m.ID, *email, m.Password, m.Name, m.IsVerified, m.CreatedAt, m.UpdatedAt)
}
```

```go
// ============================================================
// Interface Layer - HTTP Handler
// ============================================================
// internal/interfaces/http/handlers/user_handler.go
package handlers

import (
    "encoding/json"
    "net/http"
    "your-project/internal/application/user"
    "github.com/go-playground/validator/v10"
)

type UserHandler struct {
    registerUC *user.RegisterUseCase
    validate   *validator.Validate
}

func NewUserHandler(registerUC *user.RegisterUseCase) *UserHandler {
    return &UserHandler{
        registerUC: registerUC,
        validate:   validator.New(),
    }
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
    var req user.RegisterInput
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    if err := h.validate.Struct(req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    output, err := h.registerUC.Execute(r.Context(), req)
    if err != nil {
        switch err {
        case user.ErrEmailAlreadyExists:
            http.Error(w, "Email already registered", http.StatusConflict)
        default:
            http.Error(w, "Internal server error", http.StatusInternalServerError)
        }
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(output)
}
```

```go
// ============================================================
// Main Entry Point - Dependency Injection
// ============================================================
// cmd/api/main.go
package main

import (
    "log"
    "net/http"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "your-project/internal/application/user"
    "your-project/internal/domain/user"
    gormRepo "your-project/internal/infrastructure/persistence/gorm"
    "your-project/internal/interfaces/http/handlers"
    "your-project/internal/interfaces/http/routes"
)

func main() {
    // Connect to database
    dsn := "host=localhost user=postgres password=postgres dbname=myapp port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect database:", err)
    }

    // Auto migrate (for development only)
    db.AutoMigrate(&gormRepo.UserModel{})

    // Dependency Injection
    userRepo := gormRepo.NewUserRepository(db)
    registerUC := user.NewRegisterUseCase(userRepo)
    userHandler := handlers.NewUserHandler(registerUC)

    // Setup routes
    router := routes.SetupRoutes(userHandler)

    // Start server
    log.Println("Server starting on :8080")
    if err := http.ListenAndServe(":8080", router); err != nil {
        log.Fatal("server failed:", err)
    }
}
```

### ตัวอย่าง 2: Concurrency - Worker Pool Pattern

```go
// ============================================================
// Worker Pool Pattern for Processing Tasks Concurrently
// ============================================================
package main

import (
    "context"
    "fmt"
    "sync"
    "time"
)

// Job represents a unit of work
type Job struct {
    ID     int
    Payload string
}

// Result represents the outcome of processing a job
type Result struct {
    JobID int
    Output string
    Error error
}

// WorkerPool manages a pool of workers for concurrent job processing
type WorkerPool struct {
    numWorkers int
    jobQueue   chan Job
    resultQueue chan Result
    wg         sync.WaitGroup
    ctx        context.Context
    cancel     context.CancelFunc
}

// NewWorkerPool creates a new worker pool with specified number of workers
func NewWorkerPool(numWorkers int, queueSize int) *WorkerPool {
    ctx, cancel := context.WithCancel(context.Background())
    return &WorkerPool{
        numWorkers:  numWorkers,
        jobQueue:    make(chan Job, queueSize),
        resultQueue: make(chan Result, queueSize),
        ctx:         ctx,
        cancel:      cancel,
    }
}

// Start launches the worker pool
func (wp *WorkerPool) Start() {
    for i := 0; i < wp.numWorkers; i++ {
        wp.wg.Add(1)
        go wp.worker(i)
    }
}

// worker processes jobs from the queue
func (wp *WorkerPool) worker(id int) {
    defer wp.wg.Done()
    for {
        select {
        case <-wp.ctx.Done():
            fmt.Printf("Worker %d stopping\n", id)
            return
        case job, ok := <-wp.jobQueue:
            if !ok {
                return
            }
            result := wp.processJob(job)
            select {
            case wp.resultQueue <- result:
            case <-wp.ctx.Done():
                return
            }
        }
    }
}

// processJob handles a single job
func (wp *WorkerPool) processJob(job Job) Result {
    // Simulate processing time
    time.Sleep(100 * time.Millisecond)
    
    // Process the job
    output := fmt.Sprintf("Processed job %d with payload: %s", job.ID, job.Payload)
    
    return Result{
        JobID:  job.ID,
        Output: output,
        Error:  nil,
    }
}

// Submit adds a job to the queue
func (wp *WorkerPool) Submit(job Job) bool {
    select {
    case wp.jobQueue <- job:
        return true
    case <-wp.ctx.Done():
        return false
    default:
        // Queue is full, could return false or block
        return false
    }
}

// Results returns a channel for consuming results
func (wp *WorkerPool) Results() <-chan Result {
    return wp.resultQueue
}

// Stop gracefully shuts down the worker pool
func (wp *WorkerPool) Stop() {
    wp.cancel()
    close(wp.jobQueue)
    wp.wg.Wait()
    close(wp.resultQueue)
}

// Example usage with timeout and error handling
func main() {
    // Create worker pool with 5 workers
    pool := NewWorkerPool(5, 100)
    pool.Start()
    
    // Create context with timeout for the whole process
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    // Submit 50 jobs
    go func() {
        for i := 0; i < 50; i++ {
            job := Job{
                ID:     i,
                Payload: fmt.Sprintf("data-%d", i),
            }
            if !pool.Submit(job) {
                fmt.Printf("Failed to submit job %d\n", i)
            }
        }
    }()
    
    // Collect results with timeout
    results := pool.Results()
    processed := 0
    errors := 0
    
    for {
        select {
        case <-ctx.Done():
            fmt.Println("Timeout reached")
            pool.Stop()
            fmt.Printf("Processed: %d, Errors: %d\n", processed, errors)
            return
        case result, ok := <-results:
            if !ok {
                fmt.Println("All jobs processed")
                fmt.Printf("Processed: %d, Errors: %d\n", processed, errors)
                return
            }
            if result.Error != nil {
                errors++
                fmt.Printf("Error processing job %d: %v\n", result.JobID, result.Error)
            } else {
                processed++
                fmt.Printf("Result: %s\n", result.Output)
            }
        }
    }
}
```

### ตัวอย่าง 3: Generic Repository Pattern

```go
// ============================================================
// Generic Repository Pattern with Generics (Go 1.18+)
// ============================================================
package repository

import (
    "context"
    "gorm.io/gorm"
)

// Entity interface defines methods that all entities must implement
type Entity interface {
    GetID() string
}

// Repository is a generic repository interface
type Repository[T Entity] interface {
    Create(ctx context.Context, entity T) error
    GetByID(ctx context.Context, id string) (T, error)
    Update(ctx context.Context, entity T) error
    Delete(ctx context.Context, id string) error
    Find(ctx context.Context, query Query) ([]T, error)
}

// Query represents search criteria
type Query struct {
    Filters map[string]interface{}
    Limit   int
    Offset  int
    OrderBy string
}

// GormRepository is a generic GORM implementation
type GormRepository[T Entity] struct {
    db *gorm.DB
}

func NewGormRepository[T Entity](db *gorm.DB) *GormRepository[T] {
    return &GormRepository[T]{db: db}
}

func (r *GormRepository[T]) Create(ctx context.Context, entity T) error {
    return r.db.WithContext(ctx).Create(entity).Error
}

func (r *GormRepository[T]) GetByID(ctx context.Context, id string) (T, error) {
    var entity T
    err := r.db.WithContext(ctx).Where("id = ?", id).First(&entity).Error
    return entity, err
}

func (r *GormRepository[T]) Update(ctx context.Context, entity T) error {
    return r.db.WithContext(ctx).Save(entity).Error
}

func (r *GormRepository[T]) Delete(ctx context.Context, id string) error {
    var entity T
    return r.db.WithContext(ctx).Where("id = ?", id).Delete(&entity).Error
}

func (r *GormRepository[T]) Find(ctx context.Context, query Query) ([]T, error) {
    var entities []T
    db := r.db.WithContext(ctx)
    
    for field, value := range query.Filters {
        db = db.Where(field+" = ?", value)
    }
    
    if query.Limit > 0 {
        db = db.Limit(query.Limit)
    }
    if query.Offset > 0 {
        db = db.Offset(query.Offset)
    }
    if query.OrderBy != "" {
        db = db.Order(query.OrderBy)
    }
    
    err := db.Find(&entities).Error
    return entities, err
}

// ============================================================
// Usage Example
// ============================================================
type User struct {
    ID    string `gorm:"primaryKey"`
    Name  string
    Email string
}

func (u User) GetID() string { return u.ID }

func main() {
    db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    
    // Create generic repository for User
    userRepo := NewGormRepository[User](db)
    
    // Create user
    user := User{ID: "1", Name: "John", Email: "john@example.com"}
    userRepo.Create(context.Background(), user)
    
    // Find users
    users, _ := userRepo.Find(context.Background(), Query{
        Filters: map[string]interface{}{"name": "John"},
        Limit:   10,
    })
    
    // Get by ID
    found, _ := userRepo.GetByID(context.Background(), "1")
}
```

---

## สรุป

คู่มือภาษา Go ฉบับสมบูรณ์นี้ถูกออกแบบมาเพื่อให้ผู้อ่านสามารถเรียนรู้ภาษา Go ได้อย่างเป็นระบบ ตั้งแต่พื้นฐานไปจนถึงระดับมืออาชีพ โดยครอบคลุม:

1. **บทนำ** - อธิบายวัตถุประสงค์และกลุ่มเป้าหมายของคู่มือ
2. **บทนิยาม** - รวบรวมคำศัพท์เฉพาะทางที่ใช้ในคู่มือ
3. **บทหัวข้อ** - โครงสร้างเนื้อหาทั้ง 8 ภาค 51 บท
4. **การออกแบบคู่มือ** - ปรัชญา โครงสร้างการเรียนรู้ และรูปแบบเนื้อหา
5. **การออกแบบ Workflow** - ขั้นตอนการเรียนรู้และกระบวนการพัฒนาโปรเจกต์
6. **TASK LIST Template** - แม่แบบสำหรับการพัฒนา feature ใหม่
7. **CHECKLIST Template** - แม่แบบสำหรับตรวจสอบคุณภาพโค้ดและการ deploy
8. **ตัวอย่างโค้ด** - โค้ดตัวอย่างที่ใช้ในสถานการณ์จริง

ผู้อ่านสามารถใช้คู่มือนี้เป็นแนวทางในการเรียนรู้ภาษา Go และนำเทมเพลตและตัวอย่างโค้ดไปประยุกต์ใช้ในโปรเจกต์จริงได้ ขอให้สนุกกับการเขียนโปรแกรมด้วยภาษา Go!