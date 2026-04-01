# The Complete Go Programming Guide: From Basics to Production-Ready Applications

**โดย คงนคร จันทะคุณ**  
kongnakornjantakun@gmail.com  
*Version 4.0 – April 2026*

---

## 📖 บทนำ

ในยุคที่ซอฟต์แวร์มีความซับซ้อนมากขึ้นเรื่อย ๆ ภาษาโปรแกรมมิ่งที่เรียบง่าย มีประสิทธิภาพสูง และสามารถจัดการกับการทำงานพร้อมกันได้ดี กลายเป็นสิ่งที่นักพัฒนาต้องการอย่างยิ่ง ภาษา Go (หรือ Golang) ถือกำเนิดขึ้นจากความต้องการของ engineers ที่ Google ซึ่งเผชิญกับความท้าทายในการพัฒนาและบำรุงรักษาระบบขนาดใหญ่ที่มีการทำงานพร้อมกันสูง พวกเขาต้องการภาษาใหม่ที่ผสมผสานความรวดเร็วในการทำงานของภาษา C ความง่ายของภาษา Python และความสามารถในการจัดการ concurrency ที่ดีขึ้น

คู่มือเล่มนี้เกิดจากความตั้งใจที่จะรวบรวมองค์ความรู้เกี่ยวกับภาษา Go ตั้งแต่ระดับพื้นฐานจนถึงระดับมืออาชีพ ครอบคลุมทั้งไวยากรณ์พื้นฐาน การจัดการโปรเจกต์ด้วย Go Modules การทดสอบหน่วย การทำงานพร้อมกัน (concurrency) ไปจนถึงการออกแบบสถาปัตยกรรมระดับ Production และการประยุกต์ใช้ Domain-Driven Design (DDD) ร่วมกับ Go รวมถึงการเชื่อมต่อกับระบบภายนอกที่ใช้ในโลกแห่งความจริง เช่น Redis, RabbitMQ, MQTT, InfluxDB, WebSocket, SMS, LINE Notify และ Discord

คู่มือนี้ถูกออกแบบให้เป็น **ทั้งตำราเรียนและคู่มืออ้างอิง** โดยเน้นให้ผู้อ่านสามารถนำไปประยุกต์ใช้ได้ทันที ตั้งแต่การติดตั้ง การเขียนโปรแกรมพื้นฐาน ไปจนถึงการออกแบบสถาปัตยกรรมแบบ Clean Architecture และ Domain‑Driven Design (DDD) รวมถึงการเชื่อมต่อกับระบบภายนอกที่พบได้บ่อยในโลกแห่งความจริง

### วัตถุประสงค์
- ให้ผู้อ่านเข้าใจภาษา Go อย่างลึกซึ้ง ตั้งแต่ไวยากรณ์จนถึง concurrency
- เสนอแนวทางการจัดโครงสร้างโปรเจกต์สำหรับการผลิตจริง
- นำเสนอเทคนิคการทดสอบหน่วย (Unit Test) และการวัดประสิทธิภาพ
- แนะนำรูปแบบสถาปัตยกรรม Clean Architecture + DDD + CQRS
- สอนการผสาน Redis, RabbitMQ, MQTT, InfluxDB, WebSocket, SMS, LINE Notify, Discord
- จัดเตรียมเทมเพลตและ checklist ที่ช่วยให้ทีมทำงานเป็นระบบ
- ประกอบด้วยแผนภาพ (Mermaid) สำหรับอธิบายโครงสร้างและกระบวนการทำงาน
- ให้โค้ดตัวอย่างที่สามารถนำไปรันทดสอบได้จริง

### กลุ่มเป้าหมาย
- **ผู้เริ่มต้น** ที่ต้องการเรียนรู้ภาษา Go ตั้งแต่ศูนย์
- **นักพัฒนาที่เปลี่ยนภาษา** จากภาษาอื่นมาสู่ Go
- **นักพัฒนาที่ต้องการยกระดับ** สู่การเป็น Go Developer มืออาชีพ
- **สถาปนิกซอฟต์แวร์** ที่สนใจการออกแบบระบบด้วย Go

### วิธีการอ่าน
- หากยังไม่เคยเขียน Go มาก่อน ให้เริ่มจาก **ภาคที่ 1–3** เพื่อทำความเข้าใจพื้นฐาน
- หากต้องการออกแบบแอปพลิเคชันทันที ให้ข้ามไป **ภาคที่ 7–8** เพื่อศึกษา Clean Architecture และ DDD
- หากต้องการเชื่อมต่อกับระบบอื่น (ฐานข้อมูล time‑series, message queue, IoT) ให้ดู **ภาคที่ 9**

---

## 🧭 สารบัญ

### ภาคที่ 1: ปฐมบทกับการเขียนโปรแกรม
- บทที่ 1: ความรู้เบื้องต้นเกี่ยวกับการเขียนโปรแกรมคอมพิวเตอร์
- บทที่ 2: รู้จักกับภาษา Go
- บทที่ 3: พื้นฐานการใช้งาน Terminal
- บทที่ 4: เตรียมสภาพแวดล้อมสำหรับพัฒนา
- บทที่ 5: สร้างแอปพลิเคชันแรกของคุณ

### ภาคที่ 2: พื้นฐานภาษาและโครงสร้างข้อมูล
- บทที่ 6: ระบบเลขฐานสองและฐานสิบ
- บทที่ 7: เลขฐานสิบหก, ฐานแปด, ASCII, UTF8, Unicode และ Runes
- บทที่ 8: ตัวแปร, ค่าคงที่ และชนิดข้อมูลพื้นฐาน
- บทที่ 9: คำสั่งควบคุมการทำงาน
- บทที่ 10: ฟังก์ชัน
- บทที่ 11: แพคเกจและการนำเข้า
- บทที่ 12: การเริ่มต้นทำงานของแพคเกจ
- บทที่ 13: การสร้างชนิดข้อมูลใหม่ (Types)
- บทที่ 14: เมธอด (Methods)
- บทที่ 15: พอยน์เตอร์ (Pointer)
- บทที่ 16: อินเทอร์เฟซ (Interfaces)

### ภาคที่ 3: การจัดการโปรเจกต์และโครงสร้างข้อมูลขั้นสูง
- บทที่ 17: Go Modules - การจัดการโปรเจกต์สมัยใหม่
- บทที่ 18: Go Module Proxies
- บทที่ 19: การทดสอบหน่วย (Unit Tests)
- บทที่ 20: อาเรย์ (Arrays)
- บทที่ 21: สไลซ์ (Slices)
- บทที่ 22: แมพ (Maps)
- บทที่ 23: การจัดการข้อผิดพลาด (Errors)

### ภาคที่ 4: การพัฒนาแอปพลิเคชันเชิงปฏิบัติ
- บทที่ 24: ฟังก์ชันนิรนาม (Anonymous functions) และ Closure
- บทที่ 25: การจัดการข้อมูล JSON และ XML
- บทที่ 26: พื้นฐานการสร้าง HTTP Server
- บทที่ 27: Enum, Iota และ Bitmask
- บทที่ 28: วันที่และเวลา
- บทที่ 29: การจัดเก็บข้อมูล: ไฟล์และฐานข้อมูล
- บทที่ 30: การทำงานพร้อมกัน (Concurrency)
- บทที่ 31: การบันทึกเหตุการณ์ (Logging)
- บทที่ 32: เทมเพลต (Templates)
- บทที่ 33: การจัดการค่า Configuration

### ภาคที่ 5: สู่การเป็นนักพัฒนา Go มืออาชีพ
- บทที่ 34: การวัดประสิทธิภาพ (Benchmarks)
- บทที่ 35: สร้าง HTTP Client
- บทที่ 36: การวิเคราะห์โปรไฟล์ (Program Profiling)
- บทที่ 37: การจัดการ Context
- บทที่ 38: Generics - การเขียนโค้ดแบบยืดหยุ่น
- บทที่ 39: Go กับกระบวนทัศน์ OOP?
- บทที่ 40: การอัปเกรดหรือดาวน์เกรดเวอร์ชัน Go
- บทที่ 41: คำแนะนำในการออกแบบโค้ดที่ดี
- บทที่ 42: ชีทสรุป (Cheatsheet)

### ภาคที่ 6: เครื่องมือและไลบรารียอดนิยม
- บทที่ 43: chi, viper, cobra, zap และเครื่องมือสำคัญ
- บทที่ 44: GORM – ORM ทรงพลังสำหรับ Go
- บทที่ 45: การส่งอีเมลด้วย gomail และ hermes

### ภาคที่ 7: การออกแบบสถาปัตยกรรมและ Workflow
- บทที่ 46: Clean Architecture และโครงสร้างโปรเจกต์
- บทที่ 47: Blueprint สำหรับโปรเจกต์ Go ระดับ Production
- บทที่ 48: การออกแบบ Workflow และ Task Management

### ภาคที่ 8: Domain-Driven Design (DDD) กับ Go
- บทที่ 49: หลักการ DDD และการนำไปใช้ใน Go
- บทที่ 50: Aggregates, Event Storming และ CQRS
- บทที่ 51: การออกแบบบริการด้วย Go-DDD

### ภาคที่ 9: การผสานระบบภายนอกและคุณลักษณะเสริม
- บทที่ 52: Redis สำหรับ Cache และ Message Queue
- บทที่ 53: RabbitMQ – Message Broker มาตรฐานองค์กร
- บทที่ 54: MQTT สำหรับ IoT และระบบเรียลไทม์
- บทที่ 55: InfluxDB – Time‑Series Database
- บทที่ 56: WebSocket และ Socket.IO
- บทที่ 57: การส่ง SMS และ LINE Notify
- บทที่ 58: Discord Webhook สำหรับแจ้งเตือน

### ภาคที่ 10: เทมเพลต กระบวนการพัฒนา และตัวอย่างโค้ด
- บทที่ 59: ตัวอย่างโค้ดครบวงจร (Full‑stack Example)
- บทที่ 60: Task List Template
- บทที่ 61: Checklist Template
- บทที่ 62: แผนภาพการทำงาน (Workflow Diagram)
- บทที่ 63: mop Config – การจัดการ Configuration

---

## ภาคที่ 1: ปฐมบทกับการเขียนโปรแกรม

### บทที่ 1: ความรู้เบื้องต้นเกี่ยวกับการเขียนโปรแกรมคอมพิวเตอร์

#### 1.1 การเขียนโปรแกรมคืออะไร?
การเขียนโปรแกรม (Programming) คือกระบวนการสร้างชุดคำสั่งที่ใช้ควบคุมการทำงานของคอมพิวเตอร์ให้ทำงานตามที่เราต้องการ โดยใช้ภาษาเฉพาะที่คอมพิวเตอร์สามารถเข้าใจได้ ภาษาที่มนุษย์ใช้เขียนเรียกว่า "ภาษาคอมพิวเตอร์ระดับสูง" (High-Level Language) เช่น Go, Python, Java ซึ่งจากนั้นจะถูกแปลงเป็นภาษาเครื่อง (Machine Language) ที่เป็นเลขฐานสอง (0 และ 1) ที่ซีพียูสามารถประมวลผลได้

#### 1.2 โครงสร้างพื้นฐานของโปรแกรม
โปรแกรมคอมพิวเตอร์โดยทั่วไปประกอบด้วย:
- **ข้อมูล (Data)** : ตัวเลข, ข้อความ, รายการต่างๆ
- **การประมวลผล (Processing)** : การดำเนินการกับข้อมูล เช่น การคำนวณ การเปรียบเทียบ
- **การควบคุมการทำงาน (Control Flow)** : การตัดสินใจ (if-else), การวนซ้ำ (loop)
- **การจัดเก็บ (Storage)** : หน่วยความจำ, ไฟล์, ฐานข้อมูล
- **อินพุต/เอาท์พุต (I/O)** : การรับข้อมูลจากผู้ใช้ หรือแสดงผล

#### 1.3 ตัวแปลภาษาและคอมไพเลอร์
- **คอมไพเลอร์ (Compiler)** : แปลงซอร์สโค้ดทั้งหมดเป็นไฟล์ได้ก่อนรัน (เช่น Go, C)
- **อินเทอร์พรีเตอร์ (Interpreter)** : แปลงและรันทีละคำสั่ง (เช่น Python, JavaScript)

Go เป็นภาษาแบบคอมไพล์ (compiled) ซึ่งมีข้อดีคือทำงานเร็วและสร้างไฟล์ binary ที่รันได้ทันทีโดยไม่ต้องพึ่งพาสิ่งแวดล้อมอื่น (ยกเว้นระบบปฏิบัติการ)

#### 1.4 กระบวนทัศน์การเขียนโปรแกรม
Go รองรับการเขียนโปรแกรมแบบ:
- **Procedural** : ใช้ฟังก์ชันและลำดับขั้นตอน
- **Concurrent** : ทำงานพร้อมกันด้วย goroutine
- **Functional** (บางส่วน) : ฟังก์ชันเป็น first-class citizen
- **ไม่ใช่ OOP แบบคลาสสิก** : ใช้ struct และ interface แทน inheritance

#### 1.5 ขั้นตอนการพัฒนาโปรแกรม
1. เขียนซอร์สโค้ด (.go)
2. คอมไพล์ (go build)
3. ทดสอบรัน (go run)
4. แก้ไขข้อผิดพลาด (debug)
5. จัดการแพคเกจ (go mod)
6. ทดสอบหน่วย (go test)

---

### บทที่ 2: รู้จักกับภาษา Go

#### 2.1 ประวัติความเป็นมา
ภาษา Go (หรือ Golang) ถูกพัฒนาโดย Google เริ่มต้นในปี 2007 โดย Robert Griesemer, Rob Pike, และ Ken Thompson เปิดตัวเป็นโอเพนซอร์สในปี 2009 จุดประสงค์เพื่อแก้ปัญหาที่เกิดขึ้นในภาษา C++ และ Java ในระบบขนาดใหญ่ของ Google เช่น การคอมไพล์ที่ช้า, การจัดการการทำงานพร้อมกันที่ซับซ้อน, และความยุ่งยากในการบำรุงรักษา

#### 2.2 จุดเด่นของภาษา Go
- **เรียบง่ายและอ่านง่าย** : ไวยากรณ์กระชับ ไม่มีฟีเจอร์ที่ซับซ้อนเกินจำเป็น
- **คอมไพล์เร็ว** : สามารถคอมไพล์โปรเจกต์ขนาดใหญ่ได้ในไม่กี่วินาที
- **การจัดการหน่วยความจำอัตโนมัติ** : มี garbage collector ที่มีประสิทธิภาพ
- **Concurrency ระดับภาษา** : goroutine และ channel ทำให้เขียนโปรแกรม concurrent ได้ง่าย
- **Static typing** : ตรวจสอบชนิดข้อมูลตั้งแต่ตอนคอมไพล์ ป้องกันข้อผิดพลาด
- **เครื่องมือที่ครบครัน** : go fmt, go test, go mod, go vet, go doc
- **สามารถคอมไพล์ข้ามแพลตฟอร์ม** (cross-compile) ไปยัง Windows, Linux, macOS, ARM เป็นต้น

#### 2.3 โครงสร้างภาษา Go
- ไม่มี class แต่ใช้ struct และ method
- ไม่มี inheritance แต่ใช้ composition และ interface
- ไม่มี exception handling แต่ใช้ error return value
- มี garbage collection
- มี pointer แต่ไม่มี pointer arithmetic

#### 2.4 ตัวอย่าง Hello World ใน Go
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

#### 2.5 ใครใช้ Go บ้าง?
- **Google** : ระบบ backend, Kubernetes, Docker
- **Uber** : ระบบการจับคู่การเดินทาง
- **Dropbox** : เปลี่ยนจาก Python มาเป็น Go สำหรับระบบจัดเก็บไฟล์
- **Netflix** : ส่วนของ proxy และ caching
- **Cloudflare** : เครื่องมือโครงสร้างพื้นฐาน

---

### บทที่ 3: พื้นฐานการใช้งาน Terminal

#### 3.1 Terminal คืออะไร?
Terminal (หรือ command line, console) เป็นเครื่องมือที่ให้เราสั่งงานคอมพิวเตอร์ผ่านข้อความ แทนการใช้ GUI การพัฒนา Go มักใช้ terminal ในการรันคำสั่งต่างๆ เช่น go build, go run, go test, git เป็นต้น

#### 3.2 คำสั่งพื้นฐาน (Unix/Linux/macOS)
- `pwd` : แสดงไดเรกทอรีปัจจุบัน
- `ls` : แสดงรายการไฟล์ (ใช้ `ls -la` แสดงรายละเอียด)
- `cd <path>` : เปลี่ยนไดเรกทอรี
- `mkdir <name>` : สร้างโฟลเดอร์
- `touch <file>` : สร้างไฟล์
- `rm <file>` : ลบไฟล์ (ใช้ `rm -rf` ลบโฟลเดอร์)
- `cp <source> <dest>` : คัดลอกไฟล์
- `mv <source> <dest>` : ย้ายหรือเปลี่ยนชื่อ
- `cat <file>` : แสดงเนื้อหาไฟล์
- `echo <text>` : แสดงข้อความ
- `grep <pattern> <file>` : ค้นหาข้อความ

#### 3.3 คำสั่งพื้นฐานสำหรับ Windows (Command Prompt หรือ PowerShell)
- `cd` : เปลี่ยนไดเรกทอรี
- `dir` : แสดงรายการไฟล์
- `mkdir` : สร้างโฟลเดอร์
- `del` : ลบไฟล์
- `copy` : คัดลอก
- `move` : ย้าย
- `type` : แสดงเนื้อหาไฟล์

#### 3.4 การตั้งค่า PATH
หลังจากติดตั้ง Go แล้ว เราต้องให้ terminal สามารถเรียกใช้คำสั่ง `go` ได้ โดยเพิ่มไดเรกทอรีของ Go (เช่น `/usr/local/go/bin` หรือ `C:\Go\bin`) ลงในตัวแปรสภาพแวดล้อม PATH

#### 3.5 การใช้ go command
- `go version` : ตรวจสอบเวอร์ชัน Go
- `go env` : แสดงตัวแปรสภาพแวดล้อมของ Go
- `go run <file>` : คอมไพล์และรันโปรแกรม
- `go build` : คอมไพล์เป็น binary
- `go mod init` : เริ่มต้น go module
- `go get` : ดาวน์โหลดแพคเกจ (deprecated ใช้ `go install` หรือ `go mod download`)
- `go test` : รัน test

---

### บทที่ 4: เตรียมสภาพแวดล้อมสำหรับพัฒนา

#### 4.1 การติดตั้ง Go
1. ดาวน์โหลดจาก https://go.dev/dl/
2. เลือกเวอร์ชันที่เหมาะสมกับ OS ของคุณ
3. ติดตั้งตามขั้นตอน
4. ตรวจสอบการติดตั้ง: เปิด terminal แล้วพิมพ์ `go version`

#### 4.2 ตัวแปรสภาพแวดล้อมที่สำคัญ
- `GOROOT` : ตำแหน่งที่ติดตั้ง Go (มักตั้งค่าให้อัตโนมัติ)
- `GOPATH` : ตำแหน่ง workspace ของ Go (เลิกใช้แล้ว ถ้าใช้ modules)
- `GOBIN` : ตำแหน่งที่ติดตั้ง binaries ที่ `go install`
- `GOOS`/`GOARCH` : ใช้ cross-compile

#### 4.3 การเลือก Editor/IDE
- **Visual Studio Code** : มี extension "Go" โดย官方 ให้ linting, autocomplete, debugging
- **GoLand** : IDE เฉพาะของ JetBrains
- **Vim/Neovim** : ใช้ plugin vim-go
- **Sublime Text** : มี GoSublime

#### 4.4 การตั้งค่า Go Modules
Go Modules เป็นระบบจัดการ dependencies เริ่มใช้ตั้งแต่ Go 1.11 โดยค่าเริ่มต้นใช้งานได้ตั้งแต่ Go 1.16 เป็นต้นไป
```bash
mkdir myproject
cd myproject
go mod init example.com/myproject
```
ไฟล์ `go.mod` จะถูกสร้างขึ้น

#### 4.5 workspace structure (แบบเดิมกับแบบใหม่)
- **แบบ GOPATH** (ก่อน Go 1.11): ต้องวางโค้ดใน `$GOPATH/src`
- **แบบ Modules** (ปัจจุบัน): สามารถสร้างโปรเจกต์ได้ทุกที่ โดยมี go.mod กำกับ

#### 4.6 การติดตั้งเครื่องมือเสริม
- **gopls** : language server (ติดตั้งโดยอัตโนมัติใน VS Code)
- **dlv** : delve debugger (`go install github.com/go-delve/delve/cmd/dlv@latest`)
- **golangci-lint** : linter (`go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest`)

---

### บทที่ 5: สร้างแอปพลิเคชันแรกของคุณ

#### 5.1 สร้างโปรเจกต์
```bash
mkdir hello
cd hello
go mod init hello
```

#### 5.2 เขียนโค้ด
สร้างไฟล์ `main.go`:
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

#### 5.3 รันโปรแกรม
```bash
go run main.go
```
หรือคอมไพล์แล้วรัน:
```bash
go build -o hello
./hello   # (Linux/macOS) หรือ hello.exe (Windows)
```

#### 5.4 เพิ่มฟังก์ชัน
```go
package main

import "fmt"

func greet(name string) string {
    return "Hello, " + name
}

func main() {
    message := greet("Gopher")
    fmt.Println(message)
}
```
ทดสอบ `go run main.go` ควรได้ "Hello, Gopher"

#### 5.5 การอ่านค่าจาก command line
```go
package main

import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Please provide a name")
        return
    }
    name := os.Args[1]
    fmt.Printf("Hello, %s!\n", name)
}
```
รัน: `go run main.go Somchai`

#### 5.6 การจัดรูปแบบโค้ด
ใช้คำสั่ง `go fmt` เพื่อจัดรูปแบบโค้ดให้เป็นมาตรฐาน:
```bash
go fmt ./...
```

#### 5.7 ข้อผิดพลาดที่พบบ่อย
- ใช้ตัวแปรแต่ไม่ได้ใช้ -> Go จะคอมไพล์ไม่ผ่าน (ยกเว้นตัวแปร _)
- import แพคเกจแต่ไม่ได้ใช้ -> ต้องลบออกหรือใช้ _
- ใส่เครื่องหมาย `,` หรือ `;` ไม่ถูกต้อง

---

## ภาคที่ 2: พื้นฐานภาษาและโครงสร้างข้อมูล

*(หมายเหตุ: เนื้อหาในภาคที่ 2–5 ได้ถูกนำมาจาก BOOK.md ซึ่งครอบคลุมพื้นฐานภาษา Go อย่างละเอียด เนื่องจากพื้นที่จำกัด จึงขอนำเสนอเฉพาะสารบัญและเนื้อหาสำคัญบางส่วนเพื่อให้เห็นภาพรวม)*

- บทที่ 6: ระบบเลขฐานสองและฐานสิบ
- บทที่ 7: เลขฐานสิบหก, ฐานแปด, ASCII, UTF8, Unicode และ Runes
- บทที่ 8: ตัวแปร, ค่าคงที่ และชนิดข้อมูลพื้นฐาน
- บทที่ 9: คำสั่งควบคุมการทำงาน
- บทที่ 10: ฟังก์ชัน
- บทที่ 11: แพคเกจและการนำเข้า
- บทที่ 12: การเริ่มต้นทำงานของแพคเกจ
- บทที่ 13: การสร้างชนิดข้อมูลใหม่ (Types)
- บทที่ 14: เมธอด (Methods)
- บทที่ 15: พอยน์เตอร์ (Pointer)
- บทที่ 16: อินเทอร์เฟซ (Interfaces)

---

## ภาคที่ 3: การจัดการโปรเจกต์และโครงสร้างข้อมูลขั้นสูง

- บทที่ 17: Go Modules - การจัดการโปรเจกต์สมัยใหม่
- บทที่ 18: Go Module Proxies
- บทที่ 19: การทดสอบหน่วย (Unit Tests)
- บทที่ 20: อาเรย์ (Arrays)
- บทที่ 21: สไลซ์ (Slices)
- บทที่ 22: แมพ (Maps)
- บทที่ 23: การจัดการข้อผิดพลาด (Errors)

---

## ภาคที่ 4: การพัฒนาแอปพลิเคชันเชิงปฏิบัติ

- บทที่ 24: ฟังก์ชันนิรนาม (Anonymous functions) และ Closure
- บทที่ 25: การจัดการข้อมูล JSON และ XML
- บทที่ 26: พื้นฐานการสร้าง HTTP Server
- บทที่ 27: Enum, Iota และ Bitmask
- บทที่ 28: วันที่และเวลา
- บทที่ 29: การจัดเก็บข้อมูล: ไฟล์และฐานข้อมูล
- บทที่ 30: การทำงานพร้อมกัน (Concurrency)
- บทที่ 31: การบันทึกเหตุการณ์ (Logging)
- บทที่ 32: เทมเพลต (Templates)
- บทที่ 33: การจัดการค่า Configuration

---

## ภาคที่ 5: สู่การเป็นนักพัฒนา Go มืออาชีพ

- บทที่ 34: การวัดประสิทธิภาพ (Benchmarks)
- บทที่ 35: สร้าง HTTP Client
- บทที่ 36: การวิเคราะห์โปรไฟล์ (Program Profiling)
- บทที่ 37: การจัดการ Context
- บทที่ 38: Generics - การเขียนโค้ดแบบยืดหยุ่น
- บทที่ 39: Go กับกระบวนทัศน์ OOP?
- บทที่ 40: การอัปเกรดหรือดาวน์เกรดเวอร์ชัน Go
- บทที่ 41: คำแนะนำในการออกแบบโค้ดที่ดี
- บทที่ 42: ชีทสรุป (Cheatsheet)

---

## ภาคที่ 6: เครื่องมือและไลบรารียอดนิยม

### บทที่ 43: chi, viper, cobra, zap และเครื่องมือสำคัญ

#### 43.1 chi – เราเตอร์และมิดเดิลแวร์

[chi](https://github.com/go-chi/chi) เป็น lightweight router ที่มีประสิทธิภาพสูง รองรับ middleware และเข้ากันได้กับ net/http มาตรฐาน

**การติดตั้ง**
```bash
go get github.com/go-chi/chi/v5
```

**ตัวอย่างพื้นฐาน**
```go
package main

import (
    "net/http"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
)

func main() {
    r := chi.NewRouter()
    
    // ใช้ middleware พื้นฐาน
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)
    
    // กำหนด routes
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello World"))
    })
    
    // route พร้อมพารามิเตอร์
    r.Get("/users/{id}", getUser)
    
    // group routes
    r.Route("/api", func(r chi.Router) {
        r.Get("/users", listUsers)
        r.Post("/users", createUser)
        
        // sub-router พร้อม middleware เฉพาะ
        r.Route("/admin", func(r chi.Router) {
            r.Use(adminOnly)
            r.Get("/dashboard", adminDashboard)
        })
    })
    
    http.ListenAndServe(":8080", r)
}

func getUser(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    w.Write([]byte("User ID: " + id))
}
```

**มิดเดิลแวร์ที่ใช้บ่อย**
- `middleware.Logger` – บันทึก request
- `middleware.Recoverer` – จับ panic
- `middleware.Timeout` – กำหนด timeout
- `middleware.Compress` – บีบอัด response
- `middleware.RealIP` – ดึง IP จริงจาก proxy

#### 43.2 viper – การจัดการ configuration

[viper](https://github.com/spf13/viper) รองรับหลายรูปแบบ (JSON, YAML, ENV, flags) และสามารถโหลดจากไฟล์, environment variables, หรือ remote system

**การติดตั้ง**
```bash
go get github.com/spf13/viper
```

**ตัวอย่างการใช้งาน**
```go
package config

import (
    "log"
    "github.com/spf13/viper"
)

type Config struct {
    Server   ServerConfig
    Database DatabaseConfig
    Redis    RedisConfig
    JWT      JWTConfig
}

type ServerConfig struct {
    Port int
    Mode string
}

type DatabaseConfig struct {
    Host     string
    Port     int
    User     string
    Password string
    Name     string
}

func LoadConfig() (*Config, error) {
    viper.SetConfigName("config")      // ชื่อไฟล์ (ไม่รวมนามสกุล)
    viper.SetConfigType("yaml")        // yaml, json, toml, etc.
    viper.AddConfigPath(".")           // path ที่ค้นหา
    viper.AddConfigPath("/etc/app/")
    viper.AutomaticEnv()               // อ่านจาก environment variables
    
    // กำหนดค่า default
    viper.SetDefault("server.port", 8080)
    viper.SetDefault("server.mode", "debug")
    
    // อ่านไฟล์
    if err := viper.ReadInConfig(); err != nil {
        log.Printf("Config file not found: %v", err)
    }
    
    var cfg Config
    if err := viper.Unmarshal(&cfg); err != nil {
        return nil, err
    }
    
    return &cfg, nil
}
```

**ไฟล์ config.yaml ตัวอย่าง**
```yaml
server:
  port: 8080
  mode: release

database:
  host: localhost
  port: 3306
  user: root
  password: secret
  name: mydb

redis:
  addr: localhost:6379
  password: ""
  db: 0

jwt:
  secret: "your-secret-key"
  access_expiry: 15m
  refresh_expiry: 7d
```

#### 43.3 cobra – การสร้าง CLI

[cobra](https://github.com/spf13/cobra) เป็นไลบรารีสำหรับสร้าง command-line application รองรับ commands, flags, และ subcommands

**การติดตั้ง**
```bash
go get -u github.com/spf13/cobra/cobra
```

**การเริ่มต้นโปรเจกต์**
```bash
cobra init --pkg-name mycli
```

**ตัวอย่างการเพิ่ม command**
```go
// cmd/root.go
package cmd

import (
    "fmt"
    "os"
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "mycli",
    Short: "My CLI application",
    Long:  "A sample CLI built with cobra",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Hello from CLI")
    },
}

// cmd/serve.go
var serveCmd = &cobra.Command{
    Use:   "serve",
    Short: "Start the server",
    Run: func(cmd *cobra.Command, args []string) {
        port, _ := cmd.Flags().GetInt("port")
        fmt.Printf("Starting server on port %d\n", port)
    },
}

func init() {
    serveCmd.Flags().IntP("port", "p", 8080, "port to listen on")
    rootCmd.AddCommand(serveCmd)
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}
```

**การใช้งาน**
```bash
go run main.go serve --port 3000
```

#### 43.4 gorm – ORM

[gorm](https://gorm.io) เป็น ORM ที่มีฟีเจอร์ครบถ้วน: associations, hooks, preloading, transactions, etc. (รายละเอียดในบทที่ 44)

#### 43.5 validator – การตรวจสอบข้อมูล

[validator](https://github.com/go-playground/validator) ใช้ struct tags ในการกำหนด validation rules

**การติดตั้ง**
```bash
go get github.com/go-playground/validator/v10
```

**ตัวอย่าง**
```go
package main

import (
    "fmt"
    "github.com/go-playground/validator/v10"
)

type RegisterRequest struct {
    Name     string `validate:"required,min=3,max=50"`
    Email    string `validate:"required,email"`
    Password string `validate:"required,min=8"`
    Age      int    `validate:"gte=18,lte=99"`
}

func main() {
    validate := validator.New()
    
    req := RegisterRequest{
        Name:     "Jo",
        Email:    "invalid",
        Password: "short",
        Age:      16,
    }
    
    err := validate.Struct(req)
    if err != nil {
        for _, err := range err.(validator.ValidationErrors) {
            fmt.Printf("Field %s failed on tag %s\n", err.Field(), err.Tag())
        }
    }
}
```

**Tags ที่ใช้บ่อย**
- `required` – ต้องมีค่า
- `email` – รูปแบบ email
- `min`, `max` – ขนาดต่ำสุด/สูงสุด
- `gte`, `lte` – มากกว่าหรือเท่ากับ, น้อยกว่าหรือเท่ากับ
- `oneof=red blue` – ต้องเป็นหนึ่งในค่าที่กำหนด
- `uuid` – ต้องเป็น UUID
- `url` – ต้องเป็น URL

#### 43.6 jwt – การจัดการ JWT

[jwt-go](https://github.com/golang-jwt/jwt) เป็นไลบรารีมาตรฐานสำหรับ JWT

**การติดตั้ง**
```bash
go get github.com/golang-jwt/jwt/v5
```

**ตัวอย่างการสร้างและตรวจสอบ token**
```go
package jwtutil

import (
    "time"
    "github.com/golang-jwt/jwt/v5"
)

type Claims struct {
    UserID uint `json:"user_id"`
    jwt.RegisteredClaims
}

var secretKey = []byte("your-secret-key")

func GenerateAccessToken(userID uint) (string, error) {
    claims := Claims{
        UserID: userID,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
        },
    }
    
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(secretKey)
}

func ValidateToken(tokenString string) (*Claims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return secretKey, nil
    })
    
    if err != nil {
        return nil, err
    }
    
    if claims, ok := token.Claims.(*Claims); ok && token.Valid {
        return claims, nil
    }
    
    return nil, jwt.ErrInvalidKey
}
```

#### 43.7 zap – structured logging

[zap](https://github.com/uber-go/zap) เป็น logger ที่มีความเร็วสูงและรองรับ structured logging

**การติดตั้ง**
```bash
go get go.uber.org/zap
```

**ตัวอย่าง**
```go
package logger

import (
    "go.uber.org/zap"
)

var Log *zap.Logger

func InitLogger(mode string) {
    var err error
    if mode == "production" {
        Log, err = zap.NewProduction()
    } else {
        Log, err = zap.NewDevelopment()
    }
    if err != nil {
        panic(err)
    }
    defer Log.Sync()
}

func main() {
    InitLogger("development")
    defer Log.Sync()
    
    Log.Info("User logged in",
        zap.String("user", "john"),
        zap.Int("id", 123),
        zap.Duration("duration", time.Second*2),
    )
    
    Log.Error("Database error",
        zap.Error(err),
        zap.String("query", "SELECT * FROM users"),
    )
}
```

#### 43.8 gomail – การส่งอีเมล

[gomail](https://github.com/go-gomail/gomail) เป็นไลบรารีที่ใช้งานง่ายสำหรับ SMTP

**การติดตั้ง**
```bash
go get gopkg.in/gomail.v2
```

**ตัวอย่าง**
```go
package mail

import (
    "gopkg.in/gomail.v2"
)

type Mailer struct {
    dialer *gomail.Dialer
    from   string
}

func NewMailer(host string, port int, user, pass, from string) *Mailer {
    return &Mailer{
        dialer: gomail.NewDialer(host, port, user, pass),
        from:   from,
    }
}

func (m *Mailer) Send(to, subject, body string) error {
    msg := gomail.NewMessage()
    msg.SetHeader("From", m.from)
    msg.SetHeader("To", to)
    msg.SetHeader("Subject", subject)
    msg.SetBody("text/html", body)
    
    return m.dialer.DialAndSend(msg)
}
```

#### 43.9 hermes – สร้าง HTML email ที่สวยงาม

[hermes](https://github.com/matcornic/hermes) ใช้สร้าง email template แบบ responsive

**การติดตั้ง**
```bash
go get github.com/matcornic/hermes/v2
```

**ตัวอย่าง**
```go
package email

import (
    "github.com/matcornic/hermes/v2"
)

type EmailGenerator struct {
    h hermes.Hermes
}

func NewEmailGenerator(appURL, appName string) *EmailGenerator {
    h := hermes.Hermes{
        Product: hermes.Product{
            Name: appName,
            Link: appURL,
            Logo: appURL + "/logo.png",
        },
    }
    return &EmailGenerator{h: h}
}

func (g *EmailGenerator) WelcomeEmail(name, verifyURL string) (string, error) {
    email := hermes.Email{
        Body: hermes.Body{
            Name: name,
            Intros: []string{
                "Welcome to our platform!",
            },
            Actions: []hermes.Action{
                {
                    Instructions: "Please click below to verify your email address:",
                    Button: hermes.Button{
                        Color: "#22BC66",
                        Text:  "Verify Email",
                        Link:  verifyURL,
                    },
                },
            },
            Outros: []string{
                "If you didn't sign up, you can ignore this email.",
            },
        },
    }
    
    return g.h.GenerateHTML(email)
}
```

#### 43.10 air – hot-reload

[air](https://github.com/cosmtrek/air) ใช้สำหรับ reload อัตโนมัติเมื่อไฟล์เปลี่ยนแปลง เหมาะสำหรับการพัฒนา

**การติดตั้ง**
```bash
go install github.com/cosmtrek/air@latest
```

**การใช้งาน**
สร้างไฟล์ `.air.toml` ใน root ของโปรเจกต์ (หรือใช้ default) แล้วรัน:
```bash
air
```

**ตัวอย่าง .air.toml**
```toml
root = "."
tmp_dir = "tmp"

[build]
  cmd = "go build -o ./tmp/main ."
  bin = "./tmp/main"
  include_ext = ["go", "tpl", "tmpl", "html"]
  exclude_dir = ["assets", "tmp", "vendor"]
  delay = 1000
  stop_on_error = true
  send_interrupt = false
  kill_delay = 500
```

---

### บทที่ 44: GORM – ORM ทรงพลังสำหรับ Go

#### 44.1 บทนำ

GORM (Go Object Relational Mapping) เป็น ORM ที่ได้รับความนิยมสูงสุดในภาษา Go ช่วยลดความซับซ้อนในการจัดการฐานข้อมูล ด้วยการแมป struct กับตารางโดยอัตโนมัติ พร้อมฟังก์ชัน CRUD ที่ใช้งานง่าย การจัดการ transaction แบบมีระบบ และฟีเจอร์ขั้นสูงอย่าง query cache, queue processor

#### 44.2 การติดตั้งและการตั้งค่าพื้นฐาน

```go
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres // หรือ driver อื่นตามที่ใช้
```

**ตัวอย่างการเชื่อมต่อ PostgreSQL**

```go
package main

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
)

func main() {
    dsn := "host=localhost user=gorm password=gorm dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Bangkok"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect database")
    }

    // ใช้ db ต่อไป
}
```

#### 44.3 กำหนด Model (Entity)

```go
type User struct {
    ID        uint           `gorm:"primaryKey"`
    Name      string         `gorm:"size:100;not null"`
    Email     string         `gorm:"uniqueIndex;size:100;not null"`
    Age       int            `gorm:"default:0"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`
}
```

#### 44.4 CRUD Operations

**Create**

```go
user := User{Name: "สมชาย", Email: "somchai@example.com", Age: 30}
result := db.Create(&user) // สร้าง record ใหม่
fmt.Println(user.ID)       // คืนค่า ID ที่ถูกสร้าง
fmt.Println(result.Error)  // error ถ้ามี
```

**Read**

```go
// ดึง record แรกที่ตรงเงื่อนไข
var user User
db.First(&user, 1)                 // by primary key
db.First(&user, "email = ?", "somchai@example.com")

// ดึงทั้งหมด
var users []User
db.Find(&users)

// พร้อมเงื่อนไข
db.Where("age > ?", 20).Find(&users)
db.Where(&User{Name: "สมชาย"}).Find(&users)
```

**Update**

```go
// อัปเดต single column
db.Model(&user).Update("Name", "สมชาย ใหม่")

// อัปเดตหลาย columns ด้วย struct (ไม่สนใจ zero values)
db.Model(&user).Updates(User{Name: "สมชาย ใหม่", Age: 31})

// อัปเดตหลาย columns ด้วย map
db.Model(&user).Updates(map[string]interface{}{"name": "สมชาย ใหม่", "age": 31})
```

**Delete**

```go
// soft delete (ถ้ามี gorm.DeletedAt)
db.Delete(&user, 1)

// hard delete
db.Unscoped().Delete(&user, 1)
```

#### 44.5 SessionFactory

Session factory เป็นรูปแบบการสร้าง `*gorm.DB` ที่มี configuration คงที่ (เช่น logging, connection pool) และสามารถสร้าง session ใหม่สำหรับแต่ละ request หรือ transaction

**ตัวอย่าง session factory**

```go
package db

import (
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
    "log"
    "time"
)

type SessionFactory struct {
    db *gorm.DB
}

func NewSessionFactory(dsn string) (*SessionFactory, error) {
    // ตั้งค่า logger และ connection pool
    newLogger := logger.New(
        log.New(os.Stdout, "\r\n", log.LstdFlags),
        logger.Config{
            SlowThreshold: time.Second,
            LogLevel:      logger.Info,
            Colorful:      true,
        },
    )

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
        Logger: newLogger,
        NowFunc: func() time.Time { return time.Now().UTC() },
    })
    if err != nil {
        return nil, err
    }

    // ตั้งค่า connection pool
    sqlDB, _ := db.DB()
    sqlDB.SetMaxIdleConns(10)
    sqlDB.SetMaxOpenConns(100)
    sqlDB.SetConnMaxLifetime(time.Hour)

    return &SessionFactory{db: db}, nil
}

// NewSession คืนค่า session ใหม่สำหรับการทำ transaction หรือ query แบบแยก context
func (sf *SessionFactory) NewSession() *gorm.DB {
    return sf.db.Session(&gorm.Session{})
}
```

**การใช้งาน**

```go
factory, _ := db.NewSessionFactory(dsn)

// สร้าง session ใหม่สำหรับ request นี้
session := factory.NewSession()
var user User
session.First(&user, 1)

// เมื่อต้องการ transaction
session.Transaction(func(tx *gorm.DB) error {
    // ใช้ tx แทน session
    return nil
})
```

#### 44.6 การใช้ GORM Transaction เพื่อ Rollback

GORM มีฟังก์ชัน `db.Transaction` ที่ช่วยให้เราสามารถรวมหลายคำสั่ง SQL ไว้ใน transaction เดียวกันได้ โดยหากฟังก์ชันที่ส่งเข้าไปคืนค่า `error` GORM จะทำการ rollback โดยอัตโนมัติ ถ้าคืน `nil` จะ commit

**ตัวอย่างการใช้งาน Transaction**

```go
func PlaceOrder(db *gorm.DB, userID uint, items []CartItem) error {
    // เริ่ม transaction
    return db.Transaction(func(tx *gorm.DB) error {
        // 1. คำนวณราคารวม และตรวจสอบสต็อกพร้อม lock
        var total float64
        for _, item := range items {
            var stock Stock
            // Lock แถว stock เพื่อป้องกัน race condition
            if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
                Where("product_id = ?", item.ProductID).
                First(&stock).Error; err != nil {
                return err // สินค้าไม่มีในระบบ
            }
            if stock.Quantity < item.Quantity {
                return errors.New("สินค้าไม่พอ")
            }
            // หักสต็อก (จะบันทึกภายหลัง)
            stock.Quantity -= item.Quantity
            if err := tx.Save(&stock).Error; err != nil {
                return err
            }
            // คำนวณราคารวม
            total += getPrice(item.ProductID) * float64(item.Quantity)
        }

        // 2. สร้าง order
        order := Order{UserID: userID, Total: total}
        if err := tx.Create(&order).Error; err != nil {
            return err
        }

        // 3. สร้าง receipt
        receipt := Receipt{OrderID: order.ID, Amount: total}
        if err := tx.Create(&receipt).Error; err != nil {
            return err
        }

        // ทุกอย่างสำเร็จ -> commit อัตโนมัติ
        return nil
    })
}
```

#### 44.7 Query Cache

GORM เองไม่มี query cache ในตัว แต่สามารถใช้ plugin หรือจัดการเองผ่าน Redis หรือ memory cache

**ตัวอย่างการใช้ Redis cache กับ GORM**

```go
import (
    "context"
    "encoding/json"
    "github.com/go-redis/redis/v8"
    "gorm.io/gorm"
    "time"
)

type CachedDB struct {
    db    *gorm.DB
    cache *redis.Client
}

func (c *CachedDB) FirstWithCache(dest interface{}, conds ...interface{}) error {
    // สร้าง cache key จากเงื่อนไข
    key := fmt.Sprintf("query:%v", conds)

    // พยายามอ่านจาก cache
    val, err := c.cache.Get(context.Background(), key).Result()
    if err == nil {
        // พบใน cache
        return json.Unmarshal([]byte(val), dest)
    }

    // ไม่พบใน cache, query ฐานข้อมูล
    if err := c.db.First(dest, conds...).Error; err != nil {
        return err
    }

    // บันทึกผลลัพธ์ลง cache (serialize)
    data, _ := json.Marshal(dest)
    c.cache.Set(context.Background(), key, data, 5*time.Minute)
    return nil
}
```

---

### บทที่ 45: การส่งอีเมลด้วย gomail และ hermes

*เนื้อหาดูได้ในบทที่ 43 (gomail) และ 43.9 (hermes)*

---

## ภาคที่ 7: การออกแบบสถาปัตยกรรมและ Workflow

### บทที่ 46: Clean Architecture และโครงสร้างโปรเจกต์

#### 46.1 โครงสร้างโปรเจกต์แบบ Clean Architecture

โครงสร้างที่แบ่งเป็น 3 ชั้นหลัก:
- **Delivery** – รับและส่งข้อมูล (HTTP handlers, gRPC, CLI)
- **Usecase** – business logic (interactors)
- **Repository** – การเข้าถึงข้อมูล (database, external API)

นอกจากนี้ยังมี **Models** (Entities) ที่ใช้ร่วมกันทุกชั้น

**โครงสร้างโฟลเดอร์ตัวอย่าง**
```
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── config/            # การตั้งค่า
│   ├── delivery/
│   │   ├── http/          # HTTP handlers
│   │   │   ├── handler.go
│   │   │   ├── middleware.go
│   │   │   └── routes.go
│   │   └── cli/           # (optional) CLI commands
│   ├── models/            # entities / DTOs
│   ├── repository/        # implementations
│   │   ├── user_repo.go
│   │   ├── user_repo_mock.go (สำหรับ test)
│   │   └── redis/         # redis implementation
│   └── usecase/           # business logic
│       ├── user_usecase.go
│       └── auth_usecase.go
├── pkg/                   # reusable packages
│   ├── jwt/
│   ├── mail/
│   └── redis/
├── go.mod
└── config.yaml
```

#### 46.2 Models (Entities)

**models/user.go**
```go
package models

import "time"

type User struct {
    ID          uint      `json:"id"`
    Name        string    `json:"name"`
    Email       string    `json:"email"`
    Password    string    `json:"-"`          // ไม่ส่งกลับใน JSON
    IsVerified  bool      `json:"is_verified"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

// สำหรับการ register
type RegisterRequest struct {
    Name     string `json:"name" validate:"required,min=3"`
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,min=8"`
}

// สำหรับ login
type LoginRequest struct {
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required"`
}
```

#### 46.3 Repository Interface

**internal/repository/user_repo.go**
```go
package repository

import (
    "context"
    "your-project/internal/models"
)

type UserRepository interface {
    Create(ctx context.Context, user *models.User) error
    GetByID(ctx context.Context, id uint) (*models.User, error)
    GetByEmail(ctx context.Context, email string) (*models.User, error)
    Update(ctx context.Context, user *models.User) error
    Delete(ctx context.Context, id uint) error
}
```

**Implementation ด้วย GORM (internal/repository/user_repo_gorm.go)**
```go
package repository

import (
    "context"
    "gorm.io/gorm"
    "your-project/internal/models"
)

type userRepositoryGorm struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepositoryGorm{db: db}
}

func (r *userRepositoryGorm) Create(ctx context.Context, user *models.User) error {
    return r.db.WithContext(ctx).Create(user).Error
}

func (r *userRepositoryGorm) GetByID(ctx context.Context, id uint) (*models.User, error) {
    var user models.User
    err := r.db.WithContext(ctx).First(&user, id).Error
    if err != nil {
        return nil, err
    }
    return &user, nil
}
// ... methods อื่นๆ
```

#### 46.4 Usecase

**internal/usecase/user_usecase.go**
```go
package usecase

import (
    "context"
    "errors"
    "golang.org/x/crypto/bcrypt"
    "your-project/internal/models"
    "your-project/internal/repository"
)

type UserUsecase interface {
    Register(ctx context.Context, req *models.RegisterRequest) (*models.User, error)
    GetUserByID(ctx context.Context, id uint) (*models.User, error)
}

type userUsecase struct {
    userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
    return &userUsecase{userRepo: userRepo}
}

func (u *userUsecase) Register(ctx context.Context, req *models.RegisterRequest) (*models.User, error) {
    // ตรวจสอบว่ามี email ซ้ำหรือไม่
    existing, _ := u.userRepo.GetByEmail(ctx, req.Email)
    if existing != nil {
        return nil, errors.New("email already registered")
    }
    
    // Hash password
    hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }
    
    user := &models.User{
        Name:       req.Name,
        Email:      req.Email,
        Password:   string(hashed),
        IsVerified: false,
    }
    
    if err := u.userRepo.Create(ctx, user); err != nil {
        return nil, err
    }
    
    return user, nil
}
```

#### 46.5 Delivery (HTTP)

**internal/delivery/http/handler.go**
```go
package http

import (
    "encoding/json"
    "net/http"
    "your-project/internal/usecase"
    "your-project/internal/models"
    "github.com/go-chi/chi/v5"
    "github.com/go-playground/validator/v10"
)

type UserHandler struct {
    userUsecase usecase.UserUsecase
    validate    *validator.Validate
}

func NewUserHandler(userUsecase usecase.UserUsecase) *UserHandler {
    return &UserHandler{
        userUsecase: userUsecase,
        validate:    validator.New(),
    }
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
    var req models.RegisterRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }
    
    if err := h.validate.Struct(req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    user, err := h.userUsecase.Register(r.Context(), &req)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}
```

**internal/delivery/http/routes.go**
```go
package http

import (
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
)

func SetupRoutes(userHandler *UserHandler, authHandler *AuthHandler) *chi.Mux {
    r := chi.NewRouter()
    
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)
    
    r.Post("/api/register", userHandler.Register)
    r.Post("/api/login", authHandler.Login)
    
    // Protected routes
    r.Group(func(r chi.Router) {
        r.Use(authHandler.AuthMiddleware)
        r.Get("/api/users/{id}", userHandler.GetUser)
        // ...
    })
    
    return r
}
```

---

### บทที่ 47: Blueprint สำหรับโปรเจกต์ Go ระดับ Production

#### 47.1 โครงสร้างโฟลเดอร์หลัก

```
project-root/
├── cmd/                 # จุดเริ่มต้นของโปรแกรม (executable)
├── internal/            # โค้ดเฉพาะของแอปพลิเคชัน (ไม่ถูก import จากภายนอก)
├── pkg/                 # โค้ดที่นำกลับไปใช้ได้ (reusable) ในโปรเจกต์อื่น
├── api/                 # ไฟล์ที่เกี่ยวข้องกับ API (เช่น Swagger)
├── configs/             # ไฟล์ configuration
├── deploy/              # Docker, Kubernetes files
├── migrations/          # SQL migration files
├── scripts/             # สคริปต์ช่วยพัฒนา
└── test/                # integration tests
```

#### 47.2 ชั้นการทำงานภายใน `internal/`

**apps/** – จุดรวม dependencies และกำหนด routing

- `bootstrap/injection/` – ใช้สำหรับ wire dependencies (repository, service, handler) ด้วย dependency injection แบบ manual
- `router/v1/` – กำหนด routes สำหรับ API version 1 แยก public/protected

**core/** – ชั้นโดเมนของแอปพลิเคชัน (แบ่งตาม domain module)

แต่ละ module (เช่น `auth`, `user`, `iot`) มีโครงสร้างย่อย:

- `entity/` – entity หลัก (struct พร้อม behavior) และ value objects
- `repository/` – interface สำหรับ repository
- `service/` – interface สำหรับ service และการ implement business logic
- `dto/` – data transfer objects (request/response)
- `model/` – model ที่ใช้กับฐานข้อมูล (optional)
- `handler/` – HTTP handlers (รับ request, เรียก service, ส่ง response)
- `routes.go` – ลงทะเบียน routes ของ module นี้

**platform/** – ชั้น infrastructure

- `config/` – โหลด configuration (viper)
- `db/` – การเชื่อมต่อ PostgreSQL, Redis
- `cache/` – ฟังก์ชันช่วยสำหรับ Redis cache
- `queue/` – message queue (Redis pub/sub) + worker pool + dead letter queue
- `logger/` – structured logger (slog) + middleware

**transport/** – ชั้น delivery

- `middleware/` – middleware ต่างๆ (CORS, rate limit, auth, logging, recovery, security headers)
- `httpx/` – utilities สำหรับ response, request, validation
- `utils/` – helper functions (context, pagination)

#### 47.3 การเพิ่มโมดูลใหม่ (Feature)

สมมติต้องการเพิ่มโมดูล `product`:

1. สร้างโครงสร้างใน `internal/core/product/`:
   - `entity/product.go` – entity และ behavior
   - `repository/repository.go` – interface สำหรับ repository
   - `service/service.go` – interface สำหรับ service + implementation
   - `dto/product_dto.go` – request/response structs
   - `handler/product_handler.go` – HTTP handlers
   - `routes.go` – routes สำหรับ product

2. เพิ่ม repository implementation ใน `internal/platform/db/postgres/`
3. เพิ่ม service implementation ใน `internal/core/product/service/service_impl.go`
4. เพิ่ม handler ใน `internal/core/product/handler/`
5. ลงทะเบียน dependencies ใน `internal/apps/app/bootstrap/injection/`
6. เพิ่ม routes ใน `internal/apps/app/router/v1/protected_routes.go`
7. อย่าลืมเพิ่ม migrations ถ้ามีการเปลี่ยนแปลง schema

---

### บทที่ 48: การออกแบบ Workflow และ Task Management

#### 48.1 Workflow การพัฒนา Feature ใหม่

1. **Analyze** – ทำความเข้าใจ requirement, ระบุ domain models, use cases
2. **Design** – ออกแบบ entities, value objects, repository interface, service interface
3. **Implement Domain** – เขียน entity, repository interface, service interface ใน `internal/core/<module>`
4. **Implement Infrastructure** – เขียน repository implementation (GORM), cache, queue (ถ้าจำเป็น)
5. **Implement Service** – เขียน business logic ใน service implementation
6. **Implement Handler** – เขียน HTTP handlers, ตรวจสอบ input validation, ใช้ service
7. **Register Routes** – ลงทะเบียน routes ใน router
8. **Test** – เขียน unit tests (domain, service), integration tests (handler)
9. **Document** – อัปเดต API docs (Swagger) ถ้ามี

#### 48.2 Task List Template

**Phase 1: Domain Design**
- [ ] ระบุ domain model (entity, value objects)
- [ ] กำหนด invariants (business rules)
- [ ] ระบุ use cases (methods in service)
- [ ] กำหนด events (ถ้ามี)
- [ ] ออกแบบ repository interface (methods)
- [ ] ออกแบบ DTOs (request/response)

**Phase 2: Implementation**
- [ ] สร้าง entity struct และ behavior methods
- [ ] สร้าง repository interface
- [ ] สร้าง service interface
- [ ] สร้าง DTO structs
- [ ] เขียน unit tests สำหรับ entity
- [ ] เขียน unit tests สำหรับ service (mock repository)

**Phase 3: Infrastructure**
- [ ] สร้าง repository implementation (GORM)
- [ ] สร้าง migration file (ถ้ามี)
- [ ] ตั้งค่า Redis cache (ถ้าจำเป็น)
- [ ] ตั้งค่า message queue (ถ้าจำเป็น)
- [ ] ทดสอบ repository ด้วย integration test

**Phase 4: Delivery**
- [ ] สร้าง HTTP handlers
- [ ] เพิ่ม input validation (go-playground/validator)
- [ ] สร้าง routes
- [ ] ลงทะเบียน dependencies ใน injection
- [ ] ทดสอบ handler ด้วย httptest

**Phase 5: Integration & Documentation**
- [ ] ทดสอบ end-to-end ด้วย curl/Postman
- [ ] อัปเดต Swagger docs (ถ้ามี)
- [ ] อัปเดต README (ถ้าจำเป็น)
- [ ] รัน linter (`golangci-lint run`) และแก้ไข warnings

**Phase 6: Review & Deploy**
- [ ] Code review
- [ ] ตรวจสอบ performance (ถ้ามีการ query มาก)
- [ ] รัน test coverage (`go test -cover`), ควร > 80%
- [ ] รัน race detector (`go test -race`)
- [ ] Deploy to staging
- [ ] ทดสอบใน staging
- [ ] Deploy to production

#### 48.3 Checklist Template

**Code Quality Checklist**
- [ ] All exported functions have comments
- [ ] No unused imports or variables (go vet)
- [ ] Code formatted with go fmt
- [ ] Error handling is explicit (no ignored errors)
- [ ] No use of panic in library code (only in main/init)
- [ ] Context is passed as first parameter
- [ ] Interfaces are small and focused
- [ ] No global state except configuration

**Security Checklist**
- [ ] Passwords hashed with bcrypt
- [ ] JWT secret loaded from environment, not hardcoded
- [ ] Refresh tokens stored in Redis, not in DB
- [ ] Access token short-lived (≤15min)
- [ ] CORS configured properly (allow only trusted origins)
- [ ] Input validation on all endpoints
- [ ] SQL injection prevented by using parameterized queries (GORM)
- [ ] No sensitive data in logs
- [ ] HTTPS enforced in production
- [ ] Rate limiting on auth endpoints

**Performance Checklist**
- [ ] Database indexes created on frequently queried columns
- [ ] User data cached in Redis
- [ ] Connection pools configured for DB and Redis
- [ ] Use of goroutines for non-blocking tasks (email sending)
- [ ] Avoid N+1 queries (use Preload in GORM)
- [ ] Benchmarks for critical paths

**Testing Checklist**
- [ ] Unit tests cover business logic (usecase)
- [ ] Repository tests with testcontainers or in-memory DB
- [ ] HTTP handler tests with httptest
- [ ] Mock external dependencies (Redis, Mailer)
- [ ] Race condition tests with `-race` flag
- [ ] Test coverage >80%

**Deployment Checklist**
- [ ] Configurable via environment variables
- [ ] Graceful shutdown (wait for existing requests)
- [ ] Health check endpoint
- [ ] Logging to stdout (for container)
- [ ] Docker image built with non-root user
- [ ] Secrets not baked into image
- [ ] Database migration runs automatically on startup (or separate step)
- [ ] Readiness and liveness probes configured
- [ ] Monitoring (Prometheus metrics) exposed

---

## ภาคที่ 8: Domain-Driven Design (DDD) กับ Go

### บทที่ 49: หลักการ DDD และการนำไปใช้ใน Go

#### 49.1 Domain-Driven Design (DDD) คืออะไร?

Domain-Driven Design เป็นแนวทางการออกแบบซอฟต์แวร์ที่เน้นการสร้างโมเดลที่สะท้อนความรู้ความเข้าใจ (domain knowledge) อย่างแท้จริง โดยมีหลักการสำคัญคือการทำให้ซอฟต์แวร์สอดคล้องกับความต้องการทางธุรกิจผ่านการร่วมมือกันระหว่างนักพัฒนาและผู้เชี่ยวชาญในโดเมน (domain experts)

#### 49.2 หลักการสำคัญของ DDD

1. **Ubiquitous Language (ภาษาร่วม)**
   - สร้างภาษากลางที่ใช้ร่วมกันระหว่างนักพัฒนาและผู้เชี่ยวชาญโดเมน
   - ใช้ศัพท์เดียวกันในโค้ด, การสนทนา, และเอกสาร
   - ลดความเข้าใจผิดและเพิ่มความสอดคล้อง

2. **Bounded Context (บริบทที่จำกัด)**
   - แบ่งโดเมนขนาดใหญ่ออกเป็นบริทย่อยที่มีขอบเขตชัดเจน
   - แต่ละ Bounded Context มีโมเดลของตัวเองและภาษาร่วมของตัวเอง
   - ลดความซับซ้อนและความขัดแย้งของโมเดล

3. **Entities และ Value Objects**
   - **Entity**: วัตถุที่มีเอกลักษณ์ (identity) และสามารถเปลี่ยนแปลงได้ (mutable) เช่น `User`, `Order`
   - **Value Object**: วัตถุที่ไม่มีเอกลักษณ์ในตัวเอง ถูกกำหนดโดยคุณสมบัติ (immutable) เช่น `Address`, `Money`

4. **Aggregates**
   - กลุ่มของ Entities และ Value Objects ที่ถูกจัดการเป็นหน่วยเดียวกัน
   - มี Aggregate Root (รูท) เป็นตัวควบคุมความสอดคล้องของข้อมูล
   - เช่น `Order` (aggregate root) ประกอบด้วย `OrderItem` (entity) และ `Address` (value object)

5. **Domain Events**
   - เหตุการณ์สำคัญในโดเมนที่เกิดขึ้น เช่น `OrderPlaced`, `UserRegistered`
   - ใช้ในการสื่อสารระหว่าง aggregates หรือระหว่าง bounded contexts

6. **Repositories**
   - ให้ abstraction ในการเข้าถึง aggregate roots
   - ซ่อนรายละเอียดของแหล่งข้อมูล (database, cache)

7. **Domain Services**
   - ใช้สำหรับ logic ที่ไม่เหมาะจะอยู่ใน entity หรือ value object
   - เช่น `TransferService` ที่โอนเงินระหว่างบัญชี

#### 49.3 การนำ DDD ไปใช้กับ Go: ข้อแนะนำเฉพาะภาษา

1. **จัดโครงสร้างโปรเจกต์ตามโมดูล**
   ```
   /cmd
     /myapp
       main.go
   /internal
     /domain
       /order
         order.go (entity, value objects)
         repository.go (interface)
         events.go
     /application
       order_service.go
     /infrastructure
       /persistence
         order_repo_mysql.go
       /bus
         event_bus_kafka.go
     /presentation
       /http
         order_handler.go
   /pkg
     /shared
       errors.go, uuid.go, etc.
   ```

2. **ใช้ interface เพื่อ Dependency Inversion**
   - Application layer รับ domain interface
   - Infrastructure ถูก inject ผ่าน constructor
   - ใช้ `wire` (Google Wire) หรือ manual DI สำหรับการประกอบ dependencies

3. **จัดการ Transaction**
   - นิยมใช้ `Unit of Work` pattern: application service เริ่ม transaction ผ่าน interface

4. **Domain Events**
   - ใช้ channel หรือ event bus ภายใน memory ก่อน แล้วค่อยเพิ่ม infrastructure

5. **Value Objects กับ immutability**
   - ใช้ struct พร้อม private fields และ constructor functions
   - เปรียบเทียบด้วย `==` หรือ implement `Equals` method

---

### บทที่ 50: Aggregates, Event Storming และ CQRS

#### 50.1 Aggregate (กลุ่มวัตถุที่มีความสอดคล้อง)

**Aggregate** คือกลุ่มของวัตถุ (Entities + Value Objects) ที่ถูกยึดเข้าด้วยกันโดย **Aggregate Root** (รูท) ซึ่งเป็นตัวเดียวที่อนุญาตให้เข้าถึงหรือแก้ไขวัตถุอื่นภายในกลุ่มจากภายนอก การออกแบบ Aggregate ช่วยรักษาความถูกต้องของข้อมูล (invariants) และลดความซับซ้อนในการจัดการธุรกรรม

**หลักการสำคัญ:**
- **Aggregate Root** มี identity และเป็นจุดเดียวที่ภายนอกเข้าถึงได้
- การเปลี่ยนแปลงใด ๆ ภายใน Aggregate ต้องทำผ่าน Root เท่านั้น
- ภายใน Aggregate เดียวกันต้องมีความสอดคล้องกันในทันที (consistency boundary)
- ระหว่าง Aggregates ควรใช้ **eventual consistency** ผ่าน Domain Events

**ตัวอย่างใน Go:**

```go
// aggregate root: Order
type Order struct {
    id      OrderID
    status  OrderStatus
    items   []OrderItem   // value object
    total   Money
}

func (o *Order) AddItem(product Product, quantity int) error {
    if o.status != Draft {
        return errors.New("cannot add item to non-draft order")
    }
    // invariant: total must not exceed limit
    newTotal := o.total.Add(product.Price.Mul(quantity))
    if newTotal.GreaterThan(MaxOrderTotal) {
        return errors.New("order total exceeds limit")
    }
    o.items = append(o.items, NewOrderItem(product, quantity))
    o.total = newTotal
    o.addDomainEvent(OrderItemAdded{OrderID: o.id, ProductID: product.ID})
    return nil
}

// repository interface รับเฉพาะ aggregate root
type OrderRepository interface {
    Save(order *Order) error
    FindByID(id OrderID) (*Order, error)
}
```

#### 50.2 Event Storming (เทคนิคค้นพบโดเมน)

**Event Storming** เป็นเวิร์กช็อปแบบมีส่วนร่วมที่ช่วยให้ทีม (นักพัฒนา, นักธุรกิจ, ผู้เชี่ยวชาญโดเมน) ระบุและเข้าใจโดเมนผ่านเหตุการณ์สำคัญที่เกิดขึ้นในระบบ โดยใช้โน้ตสีต่าง ๆ บนกระดาน

**สัญลักษณ์ทั่วไป:**
- **สีส้ม** – Domain Events (สิ่งที่เกิดขึ้นแล้ว) เช่น `OrderPlaced`, `PaymentReceived`
- **สีน้ำเงิน** – Commands (คำสั่งที่ทำให้เกิดเหตุการณ์) เช่น `PlaceOrder`, `RefundPayment`
- **สีเหลือง** – Aggregates (กลุ่มข้อมูลที่ถูกคำสั่งเรียก) เช่น `Order`, `Customer`
- **สีม่วง** – External Systems / Policies (ระบบภายนอกหรือกฎ)
- **สีเขียว** – Read Models / Views (ข้อมูลที่แสดงผล)

**ขั้นตอนคร่าว ๆ:**
1. ระบุ Domain Events (อดีต) โดยเรียงตามลำดับเวลา
2. ระบุ Commands ที่ทำให้เกิดเหตุการณ์นั้น
3. จับคู่ Command กับ Aggregate (ผู้รับผิดชอบ)
4. เพิ่ม Policies / Rules และ External Systems
5. ระบุ Read Models ที่จำเป็นต่อการแสดงผล

Event Storming นำไปสู่การกำหนด **Bounded Context** และ **Aggregates** ที่ชัดเจน ก่อนเริ่มเขียนโค้ด

#### 50.3 CQRS (Command Query Responsibility Segregation)

CQRS แยกโมเดลการ **เขียน** (Command) และ **อ่าน** (Query) ออกจากกัน ทำให้สามารถปรับแต่งให้เหมาะสมกับแต่ละฝั่งได้ เช่น ใช้ฐานข้อมูลแบบ Event Sourcing สำหรับ Command และฐานข้อมูลแบบ Materialized View สำหรับ Query

**ใน Go สามารถจัดโครงสร้างได้ดังนี้:**

##### แยก Command และ Query Models
```go
// Command models (เขียน)
type PlaceOrderCommand struct {
    OrderID string
    Items   []OrderItemDTO
}

// Query models (อ่าน)
type OrderView struct {
    OrderID    string
    Status     string
    TotalPrice float64
    Items      []OrderItemView
}
```

##### Command Handlers
```go
type PlaceOrderHandler struct {
    repo      domain.OrderRepository
    eventBus  domain.EventBus
}

func (h *PlaceOrderHandler) Handle(ctx context.Context, cmd PlaceOrderCommand) error {
    order, err := domain.NewOrder(cmd.OrderID)
    if err != nil {
        return err
    }
    for _, item := range cmd.Items {
        order.AddItem(item.ProductID, item.Quantity)
    }
    if err := h.repo.Save(order); err != nil {
        return err
    }
    // Publish events for projection
    for _, event := range order.Events() {
        h.eventBus.Publish(event)
    }
    return nil
}
```

##### Query Handlers (อ่านจาก Read Database)
```go
type OrderQueryHandler struct {
    db *sql.DB // หรือ ORM
}

func (h *OrderQueryHandler) GetOrder(ctx context.Context, orderID string) (*OrderView, error) {
    var view OrderView
    err := h.db.QueryRowContext(ctx, "SELECT ... FROM order_views WHERE id = ?", orderID).Scan(&view)
    return &view, err
}
```

##### Projection (สร้าง Read Model จาก Events)
```go
type OrderProjection struct {
    db *sql.DB
}

func (p *OrderProjection) HandleEvent(event domain.DomainEvent) {
    switch e := event.(type) {
    case OrderPlaced:
        p.db.Exec("INSERT INTO order_views (id, status, total) VALUES (?, ?, ?)", e.OrderID, "Placed", e.Total)
    case OrderItemAdded:
        p.db.Exec("INSERT INTO order_items_view ...")
    }
}
```

#### 50.4 การนำ CQRS ไปใช้ใน Go อย่างมีประสิทธิภาพ

- **ใช้ Interface ในการแยก**: Command handlers, Query handlers, Projection handlers แต่ละตัวเป็น struct ที่ implements interface ต่างกัน ทำให้ test และ replace ได้ง่าย
- **Event Store**: ใน Go สามารถใช้ database เช่น PostgreSQL พร้อม JSONB เก็บ events
- **Read Database**: อาจใช้ฐานข้อมูลแยก (SQL, NoSQL) หรือ cache เช่น Redis สำหรับการ query ที่รวดเร็ว
- **Concurrency**: Goroutine + channel ใช้ในการประมวลผล projection แบบ asynchronous

**ข้อควรระวัง:**
- CQRS เพิ่มความซับซ้อน เหมาะสำหรับระบบที่ต้องการความยืดหยุ่นสูง (microservices, high scalability)
- ไม่จำเป็นต้องใช้ Event Sourcing เสมอไป CQRS สามารถแยกโมเดลอ่าน-เขียนโดยใช้ฐานข้อมูลเดียวกันได้ (แต่แยกตารางหรือ schema)
- การจัดการ eventual consistency ต้องออกแบบ UX ให้เหมาะสม

---

### บทที่ 51: การออกแบบบริการด้วย Go-DDD

#### 51.1 โครงสร้างโปรเจกต์แบบ DDD เต็มรูปแบบ

```
project/
├── cmd/
│   └── api/
│       └── main.go                 # entry point
├── internal/
│   ├── domain/                     # ชั้นโดเมน (core business logic)
│   │   ├── user/
│   │   │   ├── entity.go           # User entity
│   │   │   ├── value_objects.go    # Email, Password, etc.
│   │   │   ├── repository.go       # interface
│   │   │   ├── service.go          # domain services
│   │   │   └── events.go           # domain events
│   │   ├── order/
│   │   │   └── ...
│   │   └── shared/                 # shared value objects
│   ├── application/                # ชั้นแอปพลิเคชัน (use cases)
│   │   ├── user/
│   │   │   ├── register.go         # use case
│   │   │   ├── login.go
│   │   │   └── dto.go              # input/output DTOs
│   │   └── order/
│   │       └── ...
│   ├── infrastructure/             # ชั้นโครงสร้างพื้นฐาน
│   │   ├── persistence/
│   │   │   ├── gorm/
│   │   │   │   ├── user_repo.go    # implementation
│   │   │   │   └── models.go       # GORM models
│   │   │   └── redis/
│   │   ├── mail/
│   │   └── bus/                    # event bus
│   └── interfaces/                 # ชั้นติดต่อกับภายนอก
│       ├── http/
│       │   ├── handlers/
│       │   ├── middleware/
│       │   └── routes.go
│       └── cli/
├── pkg/                            # reusable packages
└── go.mod
```

#### 51.2 การสร้าง Domain Layer

**1. Entities (domain/user/entity.go)**
```go
package user

import (
    "time"
    "github.com/google/uuid"
)

type User struct {
    id          uuid.UUID
    email       Email          // value object
    password    Password       // value object
    name        string
    isVerified  bool
    createdAt   time.Time
    updatedAt   time.Time
    events      []DomainEvent
}

// Constructor
func NewUser(email, password, name string) (*User, error) {
    emailVO, err := NewEmail(email)
    if err != nil {
        return nil, err
    }
    passwordVO, err := NewPassword(password)
    if err != nil {
        return nil, err
    }
    
    user := &User{
        id:         uuid.New(),
        email:      *emailVO,
        password:   *passwordVO,
        name:       name,
        isVerified: false,
        createdAt:  time.Now(),
        updatedAt:  time.Now(),
        events:     []DomainEvent{},
    }
    
    user.addDomainEvent(NewUserRegisteredEvent(user.id, user.email.String()))
    
    return user, nil
}

// Getters (exported)
func (u *User) ID() uuid.UUID      { return u.id }
func (u *User) Email() Email       { return u.email }
func (u *User) Name() string       { return u.name }
func (u *User) IsVerified() bool   { return u.isVerified }
func (u *User) Events() []DomainEvent { return u.events }
func (u *User) ClearEvents()          { u.events = nil }

// Behavior methods
func (u *User) Verify() {
    u.isVerified = true
    u.updatedAt = time.Now()
    u.addDomainEvent(NewUserVerifiedEvent(u.id))
}

func (u *User) ChangePassword(old, new string) error {
    if err := u.password.Compare(old); err != nil {
        return ErrInvalidPassword
    }
    newPassword, err := NewPassword(new)
    if err != nil {
        return err
    }
    u.password = *newPassword
    u.updatedAt = time.Now()
    return nil
}

func (u *User) addDomainEvent(event DomainEvent) {
    u.events = append(u.events, event)
}
```

**2. Value Objects (domain/user/value_objects.go)**
```go
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
    // validate email format
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

**3. Domain Events (domain/user/events.go)**
```go
package user

import (
    "time"
    "github.com/google/uuid"
)

type DomainEvent interface {
    OccurredAt() time.Time
}

type UserRegisteredEvent struct {
    UserID    uuid.UUID `json:"user_id"`
    Email     string    `json:"email"`
    Occurred  time.Time `json:"occurred_at"`
}

func (e UserRegisteredEvent) OccurredAt() time.Time { return e.Occurred }

func NewUserRegisteredEvent(userID uuid.UUID, email string) UserRegisteredEvent {
    return UserRegisteredEvent{
        UserID:   userID,
        Email:    email,
        Occurred: time.Now(),
    }
}

type UserVerifiedEvent struct {
    UserID   uuid.UUID `json:"user_id"`
    Occurred time.Time `json:"occurred_at"`
}

func NewUserVerifiedEvent(userID uuid.UUID) UserVerifiedEvent {
    return UserVerifiedEvent{
        UserID:   userID,
        Occurred: time.Now(),
    }
}
```

#### 51.3 การสร้าง Application Layer

**Use Case (application/user/register.go)**
```go
package user

import (
    "context"
    "your-project/internal/domain/user"
    "your-project/internal/infrastructure/bus"
)

type RegisterUseCase struct {
    userRepo user.Repository
    eventBus *bus.EventBus
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

func NewRegisterUseCase(repo user.Repository, eventBus *bus.EventBus) *RegisterUseCase {
    return &RegisterUseCase{
        userRepo: repo,
        eventBus: eventBus,
    }
}

func (uc *RegisterUseCase) Execute(ctx context.Context, input RegisterInput) (*RegisterOutput, error) {
    // 1. ตรวจสอบว่า email ซ้ำหรือไม่
    emailVO, _ := user.NewEmail(input.Email)
    existing, _ := uc.userRepo.FindByEmail(ctx, *emailVO)
    if existing != nil {
        return nil, ErrEmailAlreadyExists
    }
    
    // 2. สร้าง User entity
    newUser, err := user.NewUser(input.Email, input.Password, input.Name)
    if err != nil {
        return nil, err
    }
    
    // 3. บันทึกผ่าน repository
    if err := uc.userRepo.Save(ctx, newUser); err != nil {
        return nil, err
    }
    
    // 4. Dispatch domain events
    for _, event := range newUser.Events() {
        uc.eventBus.Publish(event)
    }
    newUser.ClearEvents()
    
    // 5. ส่ง output
    return &RegisterOutput{
        ID:    newUser.ID().String(),
        Email: newUser.Email().String(),
        Name:  newUser.Name(),
    }, nil
}
```

#### 51.4 การสร้าง Infrastructure Layer

**Event Bus (infrastructure/bus/event_bus.go)**
```go
package bus

import (
    "context"
    "sync"
    "your-project/internal/domain/user"
)

type EventHandler func(context.Context, user.DomainEvent) error

type EventBus struct {
    handlers map[string][]EventHandler
    mu       sync.RWMutex
}

func NewEventBus() *EventBus {
    return &EventBus{
        handlers: make(map[string][]EventHandler),
    }
}

func (b *EventBus) Subscribe(eventName string, handler EventHandler) {
    b.mu.Lock()
    defer b.mu.Unlock()
    b.handlers[eventName] = append(b.handlers[eventName], handler)
}

func (b *EventBus) Publish(event user.DomainEvent) {
    b.mu.RLock()
    handlers := b.handlers[eventName(event)]
    b.mu.RUnlock()
    
    for _, h := range handlers {
        go h(context.Background(), event) // async
    }
}

func eventName(event user.DomainEvent) string {
    switch event.(type) {
    case user.UserRegisteredEvent:
        return "UserRegistered"
    case user.UserVerifiedEvent:
        return "UserVerified"
    default:
        return "Unknown"
    }
}
```

#### 51.5 สรุปประโยชน์ของการใช้ DDD ร่วมกับ Go

- **ความชัดเจนของโดเมน**: โค้ดสะท้อนภาษาธุรกิจ (Ubiquitous Language)
- **การแยกหน้าที่**: แต่ละ layer มีความรับผิดชอบชัดเจน ลด coupling
- **ทดสอบง่าย**: domain layer ไม่ขึ้นกับ infrastructure สามารถ unit test ด้วย mock
- **ปรับเปลี่ยน infrastructure ได้**: เปลี่ยนฐานข้อมูลหรือ event bus โดยไม่กระทบโดเมน
- **Go เหมาะสม**: struct, interface, package system ช่วยให้จัดระเบียบตาม bounded context ได้ดี และการทำงาน concurrency ผ่าน goroutine ช่วยให้จัดการ domain events ได้มีประสิทธิภาพ

---

## ภาคที่ 9: การผสานระบบภายนอกและคุณลักษณะเสริม

### บทที่ 52: Redis สำหรับ Cache และ Message Queue

#### 52.1 บทนำ

Redis (Remote Dictionary Server) เป็นฐานข้อมูลแบบ in-memory ที่รวดเร็วสูง นิยมใช้สำหรับ caching, session storage, message queue, และ real-time analytics ในบทนี้เราจะสำรวจการใช้งาน Redis กับ Go ผ่านไลบรารี `go-redis/redis` ในรูปแบบ cache และ message queue พร้อมตัวอย่างการนำไปใช้จริง

#### 52.2 การติดตั้ง Redis และ go-redis

**ติดตั้ง Redis** (ผ่าน Docker เพื่อความสะดวก)
```bash
docker run -d --name redis -p 6379:6379 redis:alpine
```

**ติดตั้ง go-redis**
```bash
go get github.com/redis/go-redis/v9
```

#### 52.3 การเชื่อมต่อและตั้งค่าเบื้องต้น

```go
package main

import (
    "context"
    "fmt"
    "github.com/redis/go-redis/v9"
)

func main() {
    // สร้าง Redis client
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    ctx := context.Background()

    // ทดสอบการเชื่อมต่อ
    pong, err := rdb.Ping(ctx).Result()
    if err != nil {
        fmt.Println("Error connecting to Redis:", err)
        return
    }
    fmt.Println("Connected to Redis:", pong)
}
```

#### 52.4 Redis Cache – การใช้งานพื้นฐาน

**Set/Get**
```go
// เก็บค่า
err := rdb.Set(ctx, "key", "value", 10*time.Minute).Err()

// ดึงค่า
val, err := rdb.Get(ctx, "key").Result()
if err == redis.Nil {
    fmt.Println("key does not exist")
} else if err != nil {
    panic(err)
} else {
    fmt.Println("key:", val)
}
```

**Hash – สำหรับเก็บ object**
```go
// เก็บ hash
err := rdb.HSet(ctx, "user:123", map[string]interface{}{
    "name":  "John",
    "email": "john@example.com",
}).Err()

// ดึง hash
user, err := rdb.HGetAll(ctx, "user:123").Result()
```

**Set – สำหรับ unique items**
```go
// เพิ่มสมาชิก
rdb.SAdd(ctx, "tags", "golang", "redis", "cache")

// ดึงสมาชิกทั้งหมด
members, err := rdb.SMembers(ctx, "tags").Result()
```

**Expiration และ TTL**
```go
// ตั้งเวลา expire
rdb.Expire(ctx, "key", 1*time.Hour)

// ดู TTL ที่เหลือ
ttl, err := rdb.TTL(ctx, "key").Result()
```

#### 52.5 Redis Cache – การใช้งานระดับ Production

**Cache-aside pattern (Lazy loading)**
```go
func GetUser(ctx context.Context, rdb *redis.Client, db *gorm.DB, id uint) (*User, error) {
    key := fmt.Sprintf("user:%d", id)
    
    // 1. พยายามอ่านจาก cache
    val, err := rdb.Get(ctx, key).Result()
    if err == nil {
        var user User
        if err := json.Unmarshal([]byte(val), &user); err == nil {
            return &user, nil
        }
    }
    
    // 2. Cache miss: query database
    var user User
    if err := db.First(&user, id).Error; err != nil {
        return nil, err
    }
    
    // 3. เก็บลง cache (async)
    go func() {
        data, _ := json.Marshal(user)
        rdb.Set(context.Background(), key, data, 10*time.Minute)
    }()
    
    return &user, nil
}
```

**Cache invalidation**
```go
func UpdateUser(ctx context.Context, rdb *redis.Client, db *gorm.DB, user *User) error {
    // 1. อัปเดตฐานข้อมูล
    if err := db.Save(user).Error; err != nil {
        return err
    }
    
    // 2. ลบ cache
    key := fmt.Sprintf("user:%d", user.ID)
    if err := rdb.Del(ctx, key).Err(); err != nil {
        // log error but not fail
    }
    
    return nil
}
```

**Cache warming (preload cache)**
```go
func WarmupUserCache(ctx context.Context, rdb *redis.Client, db *gorm.DB) {
    var users []User
    db.Find(&users)
    for _, u := range users {
        key := fmt.Sprintf("user:%d", u.ID)
        data, _ := json.Marshal(u)
        rdb.Set(ctx, key, data, 10*time.Minute)
    }
}
```

#### 52.6 Redis Message Queue (Pub/Sub)

Redis มีระบบ publish/subscribe ที่ใช้สำหรับส่งข้อความระหว่าง services แบบ asynchronous

**Publisher**
```go
// ส่งข้อความไปยัง channel
err := rdb.Publish(ctx, "order_events", `{"event":"order_created","order_id":123}`).Err()
```

**Subscriber**
```go
// สมัครรับข้อความ
pubsub := rdb.Subscribe(ctx, "order_events")
defer pubsub.Close()

// รอรับข้อความ
ch := pubsub.Channel()
for msg := range ch {
    fmt.Println("Received:", msg.Payload)
    // ประมวลผลข้อความ
}
```

#### 52.7 Redis Queue with List (Producer-Consumer)

**Producer (LPush)**
```go
// ส่งงานเข้ารายการ
job := map[string]interface{}{"task": "send_email", "user_id": 123}
data, _ := json.Marshal(job)
rdb.LPush(ctx, "email_queue", data)
```

**Consumer (BRPop) – blocking pop**
```go
// รับงานจาก queue (block จนกว่าจะมีงาน)
result, err := rdb.BRPop(ctx, 0, "email_queue").Result()
if err != nil {
    // handle error
}
data := result[1] // ค่าที่ได้เป็น string

var job EmailJob
json.Unmarshal([]byte(data), &job)
// ประมวลผล job
```

#### 52.8 Worker Pool with Redis Queue

```go
type Worker struct {
    id      int
    rdb     *redis.Client
    queue   string
    handler func(string) error
    stop    chan struct{}
}

func (w *Worker) Start(ctx context.Context) {
    for {
        select {
        case <-w.stop:
            return
        default:
            // Blocking pop
            result, err := w.rdb.BRPop(ctx, 0, w.queue).Result()
            if err != nil {
                continue
            }
            data := result[1]
            if err := w.handler(data); err != nil {
                // ถ้าผิดพลาด อาจ push กลับไป dead letter queue
                w.rdb.LPush(ctx, "dead_letter_queue", data)
            }
        }
    }
}

func (w *Worker) Stop() {
    close(w.stop)
}
```

#### 52.9 Dead Letter Queue

Dead Letter Queue (DLQ) ใช้สำหรับเก็บข้อความที่ไม่สามารถประมวลผลได้หลังจากลองหลายครั้ง

```go
type JobWithRetry struct {
    Data      string
    RetryCount int
}

func processWithRetry(rdb *redis.Client, data string, maxRetries int) error {
    err := process(data) // ลองประมวลผล
    if err == nil {
        return nil
    }
    
    // ถ้า error ให้ดู retry count
    var job JobWithRetry
    json.Unmarshal([]byte(data), &job)
    if job.RetryCount < maxRetries {
        job.RetryCount++
        newData, _ := json.Marshal(job)
        // ส่งกลับไปยัง queue เพื่อลองใหม่
        rdb.LPush(context.Background(), "main_queue", newData)
    } else {
        // ส่งไป DLQ
        rdb.LPush(context.Background(), "dead_letter_queue", data)
        // แจ้งเตือนผู้ดูแล
    }
    return err
}
```

---

### บทที่ 53: RabbitMQ – Message Broker มาตรฐานองค์กร

#### 53.1 บทนำ

RabbitMQ เป็น message broker ที่ใช้โปรโตคอล AMQP (Advanced Message Queuing Protocol) ถูกออกแบบมาเพื่อการสื่อสารระหว่าง services ในระบบขนาดใหญ่ มีความน่าเชื่อถือสูง รองรับ routing ที่ซับซ้อน และสามารถจัดการกับปริมาณข้อความสูงได้ดี ในบทนี้เราจะแนะนำการใช้งาน RabbitMQ ร่วมกับ Go ผ่านไลบรารี `rabbitmq/amqp091-go`

#### 53.2 การติดตั้ง RabbitMQ และไลบรารี

**ติดตั้ง RabbitMQ ผ่าน Docker**
```bash
docker run -d --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:management
```
- Port 5672: สำหรับ AMQP protocol
- Port 15672: สำหรับ management UI (http://localhost:15672, user: guest, pass: guest)

**ติดตั้ง Go client**
```bash
go get github.com/rabbitmq/amqp091-go
```

#### 53.3 การเชื่อมต่อเบื้องต้น

```go
package main

import (
    "log"
    "github.com/rabbitmq/amqp091-go"
)

func main() {
    // เชื่อมต่อ RabbitMQ
    conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
    if err != nil {
        log.Fatal("Failed to connect to RabbitMQ:", err)
    }
    defer conn.Close()

    // สร้าง channel
    ch, err := conn.Channel()
    if err != nil {
        log.Fatal("Failed to open channel:", err)
    }
    defer ch.Close()

    // ประกาศ queue
    q, err := ch.QueueDeclare(
        "hello", // name
        false,   // durable
        false,   // delete when unused
        false,   // exclusive
        false,   // no-wait
        nil,     // arguments
    )
    if err != nil {
        log.Fatal("Failed to declare queue:", err)
    }

    log.Println("Connected and queue declared")
}
```

#### 53.4 Producer – ส่งข้อความ

```go
// ส่งข้อความไปยัง queue
body := "Hello World!"
err = ch.Publish(
    "",     // exchange
    q.Name, // routing key (queue name)
    false,  // mandatory
    false,  // immediate
    amqp091.Publishing{
        ContentType: "text/plain",
        Body:        []byte(body),
    })
if err != nil {
    log.Fatal("Failed to publish message:", err)
}
log.Printf("Sent %s", body)
```

#### 53.5 Consumer – รับข้อความ

```go
// สร้าง consumer
msgs, err := ch.Consume(
    q.Name, // queue
    "",     // consumer
    true,   // auto-ack
    false,  // exclusive
    false,  // no-local
    false,  // no-wait
    nil,    // args
)
if err != nil {
    log.Fatal("Failed to register consumer:", err)
}

// รับข้อความแบบ continuous
go func() {
    for d := range msgs {
        log.Printf("Received a message: %s", d.Body)
        // ประมวลผลข้อความ
    }
}()

// ปิดโปรแกรมด้วย ctrl+c
select {}
```

#### 53.6 Work Queue – การกระจายงาน

ในรูปแบบ work queue งานจะถูกกระจายไปยัง consumer หลายตัวแบบ round-robin

**Producer (ส่งงาน)**
```go
// ประกาศ durable queue (รอด survive restart)
q, err := ch.QueueDeclare(
    "task_queue",
    true,  // durable
    false, // delete when unused
    false, // exclusive
    false, // no-wait
    nil,
)

// ส่งงานโดยตั้งค่า persistent
err = ch.Publish(
    "",
    q.Name,
    false,
    false,
    amqp091.Publishing{
        DeliveryMode: amqp091.Persistent, // ทำให้ message อยู่รอดแม้ RabbitMQ restart
        ContentType:  "text/plain",
        Body:         []byte(task),
    })
```

**Consumer (รับงาน)**
```go
// ตั้งค่า prefetch count เพื่อไม่ให้ consumer รับงานเยอะเกินไป
err = ch.Qos(
    1,     // prefetch count
    0,     // prefetch size
    false, // global
)

msgs, err := ch.Consume(
    q.Name,
    "",
    false, // auto-ack (ต้อง ack เอง)
    false,
    false,
    false,
    nil,
)

for d := range msgs {
    log.Printf("Received a message: %s", d.Body)
    // ทำงาน
    time.Sleep(time.Second) // simulate work
    
    // ack เพื่อบอก RabbitMQ ว่าทำงานสำเร็จ
    d.Ack(false)
}
```

#### 53.7 Publish/Subscribe (Fanout Exchange)

ใช้เมื่อต้องการส่งข้อความไปยังหลาย consumer ผ่าน exchange แบบ fanout

**Producer (ส่งข้อความไปยัง exchange)**
```go
// ประกาศ exchange
err = ch.ExchangeDeclare(
    "logs",   // name
    "fanout", // type
    true,     // durable
    false,    // auto-deleted
    false,    // internal
    false,    // no-wait
    nil,      // arguments
)

// ส่งข้อความไปยัง exchange (ไม่ต้องระบุ routing key)
err = ch.Publish(
    "logs", // exchange
    "",     // routing key
    false,
    false,
    amqp091.Publishing{
        ContentType: "text/plain",
        Body:        []byte("Hello World!"),
    })
```

**Consumer (รับข้อความจาก exchange)**
```go
// ประกาศ exchange (ต้องตรงกัน)
ch.ExchangeDeclare("logs", "fanout", true, false, false, false, nil)

// สร้าง queue ชั่วคราว (ไม่คงทน)
q, err := ch.QueueDeclare(
    "",    // name (auto-generated)
    false, // durable
    false, // delete when unused
    true,  // exclusive (เฉพาะ connection นี้)
    false, // no-wait
    nil,
)

// Bind queue กับ exchange
err = ch.QueueBind(
    q.Name, // queue name
    "",     // routing key (ignored for fanout)
    "logs", // exchange
    false,
    nil,
)

// consume
msgs, _ := ch.Consume(q.Name, "", true, false, false, false, nil)
```

#### 53.8 Routing (Direct Exchange)

ใช้ direct exchange เพื่อส่งข้อความไปยัง queue ที่มี routing key ตรงกัน

**Producer**
```go
// ประกาศ direct exchange
ch.ExchangeDeclare("direct_logs", "direct", true, false, false, false, nil)

severity := "error" // routing key
err = ch.Publish(
    "direct_logs",
    severity,
    false,
    false,
    amqp091.Publishing{
        ContentType: "text/plain",
        Body:        []byte("Error message"),
    })
```

**Consumer (รับเฉพาะ routing key ที่สนใจ)**
```go
// ประกาศ exchange
ch.ExchangeDeclare("direct_logs", "direct", true, false, false, false, nil)

// สร้าง queue ชั่วคราว
q, _ := ch.QueueDeclare("", false, false, true, false, nil)

// Bind queue กับ routing keys ที่ต้องการ
for _, s := range []string{"error", "warning"} {
    ch.QueueBind(q.Name, s, "direct_logs", false, nil)
}

// consume
msgs, _ := ch.Consume(q.Name, "", true, false, false, false, nil)
```

#### 53.9 Topic Exchange – การจับคู่แบบ wildcard

Topic exchange ใช้ routing key ที่มี pattern เช่น `*.error` หรือ `order.*`

**Producer**
```go
// ประกาศ topic exchange
ch.ExchangeDeclare("topic_logs", "topic", true, false, false, false, nil)

// ส่งข้อความพร้อม routing key
ch.Publish("topic_logs", "order.created", false, false, ...)
ch.Publish("topic_logs", "order.paid", false, false, ...)
ch.Publish("topic_logs", "order.shipped", false, false, ...)
```

**Consumer (รับ pattern)**
```go
// Bind queue กับ routing pattern
ch.QueueBind(q.Name, "order.*", "topic_logs", false, nil) // รับ order.created, order.paid, order.shipped
ch.QueueBind(q.Name, "*.error", "topic_logs", false, nil) // รับทุกเหตุการณ์ที่มี .error
```

#### 53.10 RPC (Remote Procedure Call)

RabbitMQ สามารถใช้ทำ RPC โดยส่ง request ไปยัง queue และรอ response กลับผ่าน reply queue

**Client**
```go
// สร้าง reply queue ชั่วคราว
replyQueue, _ := ch.QueueDeclare("", false, false, true, false, nil)

// ตั้งค่า correlationId และ replyTo
corrId := uuid.New().String()
err = ch.Publish(
    "",
    "rpc_queue",
    false,
    false,
    amqp091.Publishing{
        ContentType:   "text/plain",
        CorrelationId: corrId,
        ReplyTo:       replyQueue.Name,
        Body:          []byte("ping"),
    })

// รอรับ response
msgs, _ := ch.Consume(replyQueue.Name, "", true, false, false, false, nil)
for d := range msgs {
    if d.CorrelationId == corrId {
        fmt.Println("Response:", string(d.Body))
        break
    }
}
```

**Server**
```go
// ประกาศ queue สำหรับ RPC
q, _ := ch.QueueDeclare("rpc_queue", false, false, false, false, nil)

msgs, _ := ch.Consume(q.Name, "", false, false, false, false, nil)

for d := range msgs {
    // ประมวลผลคำขอ
    response := "pong"
    
    // ส่ง response กลับไปยัง reply queue
    ch.Publish(
        "",
        d.ReplyTo,
        false,
        false,
        amqp091.Publishing{
            ContentType:   "text/plain",
            CorrelationId: d.CorrelationId,
            Body:          []byte(response),
        })
    d.Ack(false)
}
```

---

### บทที่ 54: MQTT สำหรับ IoT และระบบเรียลไทม์

#### 54.1 บทนำ

MQTT (Message Queuing Telemetry Transport) เป็นโปรโตคอล lightweight สำหรับการสื่อสารแบบ publish/subscribe เหมาะสำหรับ IoT, mobile apps, และระบบที่ต้องการการส่งข้อมูลแบบ real-time ด้วยแบนด์วิดท์ต่ำ ใน Go มีไลบรารียอดนิยมเช่น `eclipse/paho.mqtt.golang`

#### 54.2 การติดตั้งและเชื่อมต่อ

**ติดตั้งไลบรารี**
```bash
go get github.com/eclipse/paho.mqtt.golang
```

**เชื่อมต่อกับ MQTT broker** (เช่น Mosquitto)
```go
package main

import (
    "fmt"
    "time"
    mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
    opts := mqtt.NewClientOptions()
    opts.AddBroker("tcp://localhost:1883")
    opts.SetClientID("go_mqtt_client")
    opts.SetUsername("user") // ถ้าต้องการ authentication
    opts.SetPassword("pass")
    opts.SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {
        fmt.Printf("Received: %s from topic: %s\n", msg.Payload(), msg.Topic())
    })

    client := mqtt.NewClient(opts)
    if token := client.Connect(); token.Wait() && token.Error() != nil {
        panic(token.Error())
    }
    defer client.Disconnect(250)

    fmt.Println("Connected to MQTT broker")
    // ... do something
}
```

#### 54.3 Publish และ Subscribe

**Publish**
```go
// publish ข้อความ
token := client.Publish("sensors/temperature", 0, false, "25.6")
token.Wait()
```

**Subscribe**
```go
// subscribe topic
token := client.Subscribe("sensors/#", 0, nil) // wildcard # รับทุก subtopic
token.Wait()
```

**Subscribe ด้วย handler**
```go
// กำหนด handler เฉพาะ topic
handler := func(client mqtt.Client, msg mqtt.Message) {
    fmt.Printf("Temperature: %s\n", msg.Payload())
}
token := client.Subscribe("sensors/temperature", 0, handler)
token.Wait()
```

#### 54.4 QoS (Quality of Service) Levels

- **QoS 0** – At most once (fire and forget)
- **QoS 1** – At least once (รับประกันถึง broker แต่ซ้ำได้)
- **QoS 2** – Exactly once (รับประกันไม่ซ้ำ)

```go
// ตัวอย่างใช้ QoS 1
token := client.Publish("important/topic", 1, false, "important message")
token.Wait()
```

#### 54.5 Retained Messages

Retained message คือข้อความที่ broker จะเก็บไว้และส่งให้ subscriber ใหม่ทันทีเมื่อ subscribe

```go
// publish with retained flag
token := client.Publish("device/status", 0, true, "online")
token.Wait()
```

#### 54.6 Last Will and Testament (LWT)

LWT คือข้อความที่ broker จะส่งเมื่อ client disconnect อย่างไม่ปกติ

```go
opts.SetWill("device/status", "offline", 0, true) // กำหนด LWT
```

#### 54.7 ตัวอย่าง: IoT Sensor Data Collector

```go
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "time"
    mqtt "github.com/eclipse/paho.mqtt.golang"
)

type SensorData struct {
    DeviceID  string    `json:"device_id"`
    Temp      float64   `json:"temperature"`
    Humidity  float64   `json:"humidity"`
    Timestamp time.Time `json:"timestamp"`
}

func main() {
    opts := mqtt.NewClientOptions()
    opts.AddBroker("tcp://localhost:1883")
    opts.SetClientID("data_collector")
    opts.SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {
        var data SensorData
        if err := json.Unmarshal(msg.Payload(), &data); err != nil {
            log.Printf("Invalid JSON: %v", err)
            return
        }
        fmt.Printf("Received from %s: temp=%.2f, humidity=%.2f\n", 
            data.DeviceID, data.Temp, data.Humidity)
        // บันทึกลง InfluxDB หรือ database อื่น
    })

    client := mqtt.NewClient(opts)
    if token := client.Connect(); token.Wait() && token.Error() != nil {
        log.Fatal(token.Error())
    }

    // Subscribe topic ทั้งหมดของ sensor
    if token := client.Subscribe("sensors/+/data", 0, nil); token.Wait() && token.Error() != nil {
        log.Fatal(token.Error())
    }

    // ปิดโปรแกรมด้วย ctrl+c
    select {}
}
```

---

### บทที่ 55: InfluxDB – Time‑Series Database

#### 55.1 บทนำ

InfluxDB เป็นฐานข้อมูล time-series ที่ออกแบบมาเพื่อจัดเก็บและสืบค้นข้อมูลที่มี timestamp เช่น ข้อมูลเซ็นเซอร์, metrics, logs มีประสิทธิภาพสูงและมี query language เฉพาะ (Flux หรือ InfluxQL) ใน Go มี client library อย่าง `influxdata/influxdb-client-go`

#### 55.2 การติดตั้งและเชื่อมต่อ

**ติดตั้ง InfluxDB (Docker)**
```bash
docker run -d --name influxdb -p 8086:8086 influxdb:latest
```

**ติดตั้ง Go client**
```bash
go get github.com/influxdata/influxdb-client-go/v2
```

**เชื่อมต่อ**
```go
package main

import (
    "fmt"
    "time"
    influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func main() {
    // สร้าง client
    client := influxdb2.NewClient("http://localhost:8086", "my-token")
    defer client.Close()

    // ใช้ write API
    writeAPI := client.WriteAPI("my-org", "my-bucket")

    // สร้าง point
    p := influxdb2.NewPoint("temperature",
        map[string]string{"device": "sensor1"},
        map[string]interface{}{"value": 25.6},
        time.Now())
    writeAPI.WritePoint(p)

    // flush
    writeAPI.Flush()

    fmt.Println("Data written")
}
```

#### 55.3 การเขียนข้อมูล (Write)

**เขียนข้อมูลทีละหลายจุด**
```go
writeAPI := client.WriteAPI("my-org", "my-bucket")
for i := 0; i < 10; i++ {
    p := influxdb2.NewPoint("cpu",
        map[string]string{"host": "server1"},
        map[string]interface{}{"usage": 50 + i},
        time.Now())
    writeAPI.WritePoint(p)
    time.Sleep(1 * time.Second)
}
writeAPI.Flush()
```

**ใช้แบบ batch ด้วย WriteAPIBlocking**
```go
writeAPI := client.WriteAPIBlocking("my-org", "my-bucket")
p := influxdb2.NewPoint("temperature",
    map[string]string{"device": "sensor2"},
    map[string]interface{}{"value": 30.2},
    time.Now())
err := writeAPI.WritePoint(context.Background(), p)
```

#### 55.4 การอ่านข้อมูล (Query) ด้วย Flux

```go
// สร้าง query API
queryAPI := client.QueryAPI("my-org")

// Flux query
flux := `from(bucket:"my-bucket")
  |> range(start: -1h)
  |> filter(fn: (r) => r._measurement == "temperature")
  |> filter(fn: (r) => r.device == "sensor1")
  |> aggregateWindow(every: 1m, fn: mean)`

// Execute query
result, err := queryAPI.Query(context.Background(), flux)
if err != nil {
    panic(err)
}

// Iterate over results
for result.Next() {
    record := result.Record()
    fmt.Printf("Time: %s, Value: %v\n", record.Time(), record.Value())
}
if result.Err() != nil {
    panic(result.Err())
}
```

#### 55.5 การใช้ InfluxQL (ภาษา SQL-like)

```go
query := `SELECT mean(value) FROM temperature WHERE time > now() - 1h GROUP BY time(5m)`
result, err := queryAPI.Query(context.Background(), query)
// ... iterate
```

#### 55.6 การจัดการ Tags และ Fields

- **Tags**: ใช้สำหรับ indexing (เช่น device ID, location) – ต้องเป็น string
- **Fields**: ใช้สำหรับค่าที่วัด (เช่น temperature, humidity) – 可以是数值

```go
p := influxdb2.NewPoint("sensor_data",
    map[string]string{
        "device_id": "sensor1",
        "location":  "room_101",
    },
    map[string]interface{}{
        "temperature": 25.6,
        "humidity":    60,
        "battery":     4.2,
    },
    time.Now())
```

---

### บทที่ 56: WebSocket และ Socket.IO

#### 56.1 WebSocket พื้นฐาน

WebSocket เป็นโปรโตคอลที่ช่วยให้ client-server สื่อสารแบบ full-duplex เหนือ TCP เหมาะสำหรับ real-time applications เช่น chat, live notifications, gaming ใน Go เราใช้ไลบรารี `gorilla/websocket`

**การติดตั้ง**
```bash
go get github.com/gorilla/websocket
```

**Server ตัวอย่าง**
```go
package main

import (
    "log"
    "net/http"
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool { return true },
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Print("upgrade failed:", err)
        return
    }
    defer conn.Close()

    for {
        // อ่านข้อความจาก client
        msgType, msg, err := conn.ReadMessage()
        if err != nil {
            break
        }
        log.Printf("Received: %s", msg)

        // ส่งข้อความกลับ
        err = conn.WriteMessage(msgType, []byte("Echo: "+string(msg)))
        if err != nil {
            break
        }
    }
}

func main() {
    http.HandleFunc("/ws", wsHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

**Client (JavaScript)**
```javascript
const socket = new WebSocket('ws://localhost:8080/ws');
socket.onmessage = (event) => console.log('Received:', event.data);
socket.send('Hello');
```

#### 56.2 WebSocket Hub – กระจายข้อความ

ใช้ pattern hub เพื่อจัดการ connections หลายตัวและ broadcast ข้อความ

```go
type Client struct {
    conn *websocket.Conn
    send chan []byte
    hub  *Hub
}

type Hub struct {
    clients    map[*Client]bool
    broadcast  chan []byte
    register   chan *Client
    unregister chan *Client
}

func NewHub() *Hub {
    return &Hub{
        clients:    make(map[*Client]bool),
        broadcast:  make(chan []byte),
        register:   make(chan *Client),
        unregister: make(chan *Client),
    }
}

func (h *Hub) Run() {
    for {
        select {
        case client := <-h.register:
            h.clients[client] = true
        case client := <-h.unregister:
            if _, ok := h.clients[client]; ok {
                delete(h.clients, client)
                close(client.send)
            }
        case message := <-h.broadcast:
            for client := range h.clients {
                select {
                case client.send <- message:
                default:
                    close(client.send)
                    delete(h.clients, client)
                }
            }
        }
    }
}

func (c *Client) readPump() {
    defer func() {
        c.hub.unregister <- c
        c.conn.Close()
    }()
    for {
        _, msg, err := c.conn.ReadMessage()
        if err != nil {
            break
        }
        c.hub.broadcast <- msg
    }
}

func (c *Client) writePump() {
    defer c.conn.Close()
    for {
        select {
        case msg, ok := <-c.send:
            if !ok {
                c.conn.WriteMessage(websocket.CloseMessage, []byte{})
                return
            }
            if err := c.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
                return
            }
        }
    }
}
```

#### 56.3 Socket.IO – การใช้งานที่ง่ายขึ้น

Socket.IO เป็น wrapper เหนือ WebSocket ที่ให้คุณสมบัติเพิ่มเติม เช่น fallback, rooms, auto-reconnect ใน Go มีไลบรารี `googollee/go-socket.io`

**การติดตั้ง**
```bash
go get github.com/googollee/go-socket.io
```

**Server ตัวอย่าง**
```go
package main

import (
    "log"
    socketio "github.com/googollee/go-socket.io"
    "net/http"
)

func main() {
    server := socketio.NewServer(nil)

    server.OnConnect("/", func(s socketio.Conn) error {
        s.SetContext("")
        log.Println("Connected:", s.ID())
        return nil
    })

    server.OnEvent("/", "chat message", func(s socketio.Conn, msg string) {
        log.Println("Message:", msg)
        s.Emit("chat message", "Echo: "+msg)
    })

    server.OnDisconnect("/", func(s socketio.Conn, reason string) {
        log.Println("Disconnected:", reason)
    })

    go server.Serve()
    defer server.Close()

    http.Handle("/socket.io/", server)
    http.Handle("/", http.FileServer(http.Dir("./static")))
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

**Client (JavaScript)**
```html
<script src="/socket.io/socket.io.js"></script>
<script>
  const socket = io();
  socket.on('chat message', (msg) => console.log(msg));
  socket.emit('chat message', 'Hello from client');
</script>
```

---

### บทที่ 57: การส่ง SMS และ LINE Notify

#### 57.1 การส่ง SMS ผ่าน Twilio

Twilio เป็นบริการ SMS API ยอดนิยม

**การติดตั้ง**
```bash
go get github.com/twilio/twilio-go
```

**ตัวอย่างการส่ง SMS**
```go
package main

import (
    "fmt"
    "github.com/twilio/twilio-go"
    twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

func main() {
    client := twilio.NewRestClientWithParams(twilio.ClientParams{
        Username: "ACxxxxxxxxxxxxxx",
        Password: "yyyyyyyyyyyyyyyy",
    })

    params := &twilioApi.CreateMessageParams{}
    params.SetTo("+66812345678")
    params.SetFrom("+1234567890")
    params.SetBody("Hello from Go!")

    resp, err := client.Api.CreateMessage(params)
    if err != nil {
        fmt.Println("Error sending SMS:", err)
    } else {
        fmt.Println("SMS SID:", *resp.Sid)
    }
}
```

#### 57.2 การส่ง SMS ผ่านบริการไทย (SMS360)

บริการในไทย เช่น SMS360, ThaiBulkSMS ฯลฯ มักใช้ HTTP API

```go
func sendSMS360(phone, message string) error {
    url := "https://api.sms360.com/send"
    data := url.Values{}
    data.Set("api_key", "your-api-key")
    data.Set("phone", phone)
    data.Set("message", message)

    resp, err := http.PostForm(url, data)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("failed to send SMS: %s", resp.Status)
    }
    return nil
}
```

#### 57.3 LINE Notify – การแจ้งเตือนผ่าน LINE

LINE Notify ให้บริการแจ้งเตือนฟรี โดยต้องสร้าง token ที่ https://notify-bot.line.me/

**ส่งข้อความผ่าน LINE Notify API**
```go
func sendLineNotify(message string) error {
    url := "https://notify-api.line.me/api/notify"
    data := url.Values{}
    data.Set("message", message)

    req, err := http.NewRequest("POST", url, strings.NewReader(data.Encode()))
    if err != nil {
        return err
    }
    req.Header.Set("Authorization", "Bearer YOUR_TOKEN")
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("LINE Notify failed: %s", resp.Status)
    }
    return nil
}
```

**ส่งรูปภาพ (image)**
```go
// สามารถส่งภาพโดยใช้ multipart form
func sendLineNotifyWithImage(message, imagePath string) error {
    var b bytes.Buffer
    w := multipart.NewWriter(&b)
    w.WriteField("message", message)
    file, _ := os.Open(imagePath)
    defer file.Close()
    part, _ := w.CreateFormFile("imageFile", imagePath)
    io.Copy(part, file)
    w.Close()

    req, _ := http.NewRequest("POST", "https://notify-api.line.me/api/notify", &b)
    req.Header.Set("Authorization", "Bearer YOUR_TOKEN")
    req.Header.Set("Content-Type", w.FormDataContentType())

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}
```

---

### บทที่ 58: Discord Webhook สำหรับแจ้งเตือน

#### 58.1 Discord Webhook คืออะไร

Discord Webhook เป็นวิธีส่งข้อความอัตโนมัติไปยัง channel ใน Discord โดยไม่ต้องใช้ bot token เหมาะสำหรับการแจ้งเตือนจากระบบ (เช่น deployment status, error logs)

#### 58.2 การสร้าง Webhook ใน Discord

1. เปิด Discord, ไปที่ server ของคุณ
2. คลิกขวาที่ channel -> Edit Channel -> Integrations -> Webhooks -> Create Webhook
3. คัดลอก URL (https://discord.com/api/webhooks/...)

#### 58.3 ส่งข้อความผ่าน Go

**โครงสร้าง payload**
```go
type DiscordWebhook struct {
    Content   string         `json:"content,omitempty"`
    Username  string         `json:"username,omitempty"`
    AvatarURL string         `json:"avatar_url,omitempty"`
    Embeds    []DiscordEmbed `json:"embeds,omitempty"`
}

type DiscordEmbed struct {
    Title       string `json:"title,omitempty"`
    Description string `json:"description,omitempty"`
    Color       int    `json:"color,omitempty"` // hexadecimal color (e.g., 0x00ff00)
    Fields      []struct {
        Name   string `json:"name"`
        Value  string `json:"value"`
        Inline bool   `json:"inline,omitempty"`
    } `json:"fields,omitempty"`
    Timestamp string `json:"timestamp,omitempty"`
}
```

**ฟังก์ชันส่งข้อความ**
```go
func sendDiscordWebhook(webhookURL string, content string) error {
    payload := DiscordWebhook{Content: content}
    data, _ := json.Marshal(payload)

    resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(data))
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
        return fmt.Errorf("discord webhook returned %s", resp.Status)
    }
    return nil
}
```

**ส่งแบบ embed (สวยงาม)**
```go
embed := DiscordEmbed{
    Title:       "🚀 Deployment",
    Description: "New version v1.2.3 deployed successfully",
    Color:       0x00ff00,
    Fields: []struct{
        Name   string `json:"name"`
        Value  string `json:"value"`
        Inline bool   `json:"inline,omitempty"`
    }{
        {Name: "Environment", Value: "Production", Inline: true},
        {Name: "Version", Value: "v1.2.3", Inline: true},
        {Name: "Deployed By", Value: "CI/CD", Inline: false},
    },
    Timestamp: time.Now().Format(time.RFC3339),
}
payload := DiscordWebhook{Embeds: []DiscordEmbed{embed}}
```

---

## ภาคที่ 10: เทมเพลต กระบวนการพัฒนา และตัวอย่างโค้ด

### บทที่ 59: ตัวอย่างโค้ดครบวงจร (Full‑stack Example)

*(หมายเหตุ: ตัวอย่างโค้ดเต็มรูปแบบจะถูกแทรกในฉบับสมบูรณ์ แต่เนื่องจากข้อจำกัดของพื้นที่ ขอสรุปเป็นโครงสร้างหลัก)*

- โครงสร้างโปรเจกต์ตาม Clean Architecture
- การลงทะเบียนผู้ใช้ด้วย JWT และ refresh token ใน Redis
- การส่งอีเมลยืนยันด้วย gomail + hermes
- การ query ข้อมูลพร้อม cache ใน Redis
- การใช้ GORM transaction สำหรับการสร้าง order
- การใช้ RabbitMQ สำหรับ async email sending
- การแจ้งเตือนผ่าน Discord Webhook เมื่อเกิด error

### บทที่ 60: Task List Template (แสดงไว้ในบทที่ 48)

### บทที่ 61: Checklist Template (แสดงไว้ในบทที่ 48)

### บทที่ 62: แผนภาพการทำงาน (Workflow Diagram)

แผนภาพ Mermaid ที่แทรกในบทต่างๆ สามารถนำไปใช้ในการอธิบายระบบ

**ตัวอย่าง: Clean Architecture Workflow**
```mermaid
graph TD
    A[HTTP Request] --> B[Handler]
    B --> C[Usecase]
    C --> D[Repository Interface]
    D --> E[Database Implementation]
    C --> F[Domain Events]
    F --> G[Event Bus]
    G --> H[Other Services]
    C --> I[Response]
    I --> B
```

**ตัวอย่าง: JWT Authentication Flow**
```mermaid
sequenceDiagram
    participant Client
    participant API
    participant DB
    participant Redis

    Client->>API: POST /login (email, password)
    API->>DB: Validate credentials
    DB-->>API: User data
    API->>API: Generate access & refresh tokens
    API->>Redis: Store refresh token
    API-->>Client: Tokens
```

### บทที่ 63: mop Config – การจัดการ Configuration

**ไฟล์ config/config.yaml**
```yaml
server:
  port: 8080
  mode: release
  read_timeout: 15s
  write_timeout: 15s

database:
  host: localhost
  port: 5432
  user: postgres
  password: postgres
  name: userdb
  sslmode: disable
  max_open_conns: 25
  max_idle_conns: 25
  conn_max_lifetime: 5m

redis:
  addr: localhost:6379
  password: ""
  db: 0
  pool_size: 10
  ttl: 10m

jwt:
  secret: "your-secret-key-change-in-production"
  access_expiry: 15m
  refresh_expiry: 168h  # 7 days

smtp:
  host: smtp.gmail.com
  port: 587
  username: your-email@gmail.com
  password: your-app-password
  from: your-email@gmail.com

log:
  level: info
  format: json
  output: stdout

rabbitmq:
  url: amqp://guest:guest@localhost:5672/
  exchange: events

mqtt:
  broker: tcp://localhost:1883
  client_id: go_mqtt_client

influxdb:
  url: http://localhost:8086
  token: my-token
  org: my-org
  bucket: my-bucket
```

**การโหลด config ด้วย viper (Go)**
```go
package config

import (
    "github.com/spf13/viper"
)

type Config struct {
    Server   ServerConfig
    Database DatabaseConfig
    Redis    RedisConfig
    JWT      JWTConfig
    SMTP     SMTPConfig
    Log      LogConfig
    RabbitMQ RabbitMQConfig
    MQTT     MQTTConfig
    InfluxDB InfluxDBConfig
}

func LoadConfig() (*Config, error) {
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(".")
    viper.AutomaticEnv()

    if err := viper.ReadInConfig(); err != nil {
        return nil, err
    }

    var cfg Config
    if err := viper.Unmarshal(&cfg); err != nil {
        return nil, err
    }
    return &cfg, nil
}
```

---

เราจะเพิ่มคำอธิบายใต้แผนภาพภาษาไทยสำหรับแต่ละส่วน เพื่อให้ผู้อ่านเข้าใจการไหลของข้อมูลได้ง่ายขึ้น โดยใช้แผนภาพ Mermaid ที่นำเสนอในคำตอบก่อนหน้านี้ พร้อมคำอธิบายภาษาไทยใต้ภาพ

---

## ภาพรวมการไหลของข้อมูล (Overview Data Flow)

```mermaid
graph TD
    subgraph Part1[Part 1: Fundamentals]
        A1[Source Code] --> A2[Compiler]
        A2 --> A3[Executable]
        A3 --> A4[Run & Debug]
    end

    subgraph Part2[Part 2: Basic Language & Data Structures]
        B1[Variables & Types] --> B2[Control Flow]
        B2 --> B3[Functions]
        B3 --> B4[Structs & Interfaces]
    end

    subgraph Part3[Part 3: Project Management & Advanced Data Structures]
        C1[Go Modules] --> C2[Tests]
        C2 --> C3[Arrays/Slices/Maps]
        C3 --> C4[Error Handling]
    end

    subgraph Part4[Part 4: Practical Application Development]
        D1[HTTP Server] --> D2[JSON/XML]
        D2 --> D3[Concurrency]
        D3 --> D4[Logging & Config]
    end

    subgraph Part5[Part 5: Professional Go Development]
        E1[Benchmarks] --> E2[Profiling]
        E2 --> E3[Context]
        E3 --> E4[Generics]
    end

    subgraph Part6[Part 6: Popular Tools & Libraries]
        F1[chi, viper, cobra, zap] --> F2[GORM]
        F2 --> F3[gomail & hermes]
    end

    subgraph Part7[Part 7: Architecture Design & Workflow]
        G1[Clean Architecture] --> G2[Blueprint]
        G2 --> G3[Workflow & Tasks]
    end

    subgraph Part8[Part 8: Domain-Driven Design]
        H1[DDD Principles] --> H2[Aggregates & Events]
        H2 --> H3[CQRS & Services]
    end

    subgraph Part9[Part 9: External Systems Integration]
        I1[Redis] --> I2[RabbitMQ]
        I2 --> I3[MQTT]
        I3 --> I4[InfluxDB]
        I4 --> I5[WebSocket]
        I5 --> I6[SMS/LINE/Discord]
    end

    subgraph Part10[Part 10: Templates & Examples]
        J1[Full‑stack Example] --> J2[Task List]
        J2 --> J3[Checklist]
        J3 --> J4[Workflow Diagram]
        J4 --> J5[mop Config]
    end

    A4 --> B1
    C4 --> D1
    D4 --> E1
    E4 --> F1
    F3 --> G1
    G3 --> H1
    H3 --> I1
    I6 --> J1
```

**คำอธิบาย:**  
แผนภาพนี้แสดงความสัมพันธ์ของเนื้อหาทั้ง 10 ภาค โดยเริ่มจากภาคที่ 1 (พื้นฐานการเขียนโปรแกรม) ไหลไปสู่ภาคที่ 2 (โครงสร้างข้อมูลพื้นฐาน) จากนั้นต่อเนื่องไปจนถึงภาคที่ 10 (เทมเพลตและตัวอย่างโค้ด) แต่ละภาคส่งต่อข้อมูลและแนวคิดไปยังภาคถัดไป ทำให้เห็นภาพรวมว่าความรู้แต่ละส่วนเชื่อมโยงกันอย่างไร

---

## ภาคที่ 1: ปฐมบทกับการเขียนโปรแกรม (บทที่ 1–5)

**แผนภาพ: จากแนวคิดสู่โปรแกรมแรก**

```mermaid
graph LR
    subgraph User[ผู้ใช้]
        U1[เขียนโค้ด]
    end

    subgraph Terminal[เทอร์มินัล]
        T1[go run] --> T2[คอมไพเลอร์]
        T2 --> T3[ไบนารี]
        T3 --> T4[ผลลัพธ์]
    end

    U1 -->|main.go| T1
    T4 -->|แสดงผล| U1
```

**คำอธิบาย:**  
ผู้ใช้เขียนโค้ด (ไฟล์ `.go`) แล้วสั่ง `go run` ซึ่งจะส่งให้คอมไพเลอร์แปลงเป็นไบนารีและทำงานทันที ผลลัพธ์แสดงบนเทอร์มินัล ผู้ใช้เห็นผลและสามารถปรับปรุงโค้ดต่อไป วงจรนี้แสดงการพัฒนาโปรแกรมแรกด้วย Go

---
# ภาคที่ 2: พื้นฐานภาษาและโครงสร้างข้อมูล (บทที่ 6–16)

## แผนภาพแสดงความสัมพันธ์ของเนื้อหา

```mermaid
flowchart TD
    subgraph DataRepresentation["1. การแทนข้อมูล (บทที่ 6-7)"]
        direction TB
        A1["เลขฐานสอง (Binary)<br/>0 และ 1<br/>หน่วยเก็บข้อมูลพื้นฐาน"]
        A2["เลขฐานสิบ (Decimal)<br/>ที่มนุษย์ใช้"]
        A3["เลขฐานสิบหก (Hex)<br/>0x0-0xF<br/>แสดงข้อมูลขนาดใหญ่"]
        A4["เลขฐานแปด (Octal)<br/>0o0-0o7<br/>สิทธิ์ไฟล์ Unix"]
        A5["ASCII / UTF-8 / Unicode<br/>การแทนข้อความ<br/>Rune: ตัวอักษรเดี่ยว"]
        
        A1 --> A3
        A2 --> A3
        A3 --> A5
        A4 --> A5
    end

    subgraph DataStorage["2. การจัดเก็บข้อมูล (บทที่ 8-9)"]
        direction TB
        B1["ตัวแปร (Variables)<br/>var, :=<br/>เก็บค่าที่เปลี่ยนได้"]
        B2["ค่าคงที่ (Constants)<br/>const, iota<br/>ค่าที่ไม่เปลี่ยนแปลง"]
        B3["ชนิดข้อมูลพื้นฐาน<br/>int, float, string, bool<br/>byte, rune"]
        B4["โครงสร้างควบคุม<br/>if-else, for, switch<br/>ควบคุมการทำงาน"]
        
        B1 --> B4
        B2 --> B4
        B3 --> B4
    end

    subgraph CodeOrganization["3. การจัดระเบียบโค้ด (บทที่ 10-12)"]
        direction TB
        C1["ฟังก์ชัน (Functions)<br/>func, return<br/>multiple returns<br/>defer, panic/recover"]
        C2["แพคเกจ (Packages)<br/>package, import<br/>exported/unexported"]
        C3["การเริ่มต้นแพคเกจ<br/>init()<br/>ลำดับการทำงาน"]
        
        C1 --> C2
        C2 --> C3
    end

    subgraph AdvancedTypes["4. ชนิดข้อมูลขั้นสูง (บทที่ 13-16)"]
        direction TB
        D1["การสร้างชนิดใหม่<br/>type, struct<br/>type alias"]
        D2["เมธอด (Methods)<br/>func (t T) method()<br/>value receiver<br/>pointer receiver"]
        D3["พอยน์เตอร์ (Pointers)<br/>*, &<br/>การอ้างอิงที่อยู่<br/>ไม่มี pointer arithmetic"]
        D4["อินเทอร์เฟซ (Interfaces)<br/>interface{}<br/>implicit implementation<br/>polymorphism"]
        
        D1 --> D2
        D2 --> D3
        D3 --> D4
    end

    subgraph Flow["ลำดับการเรียนรู้"]
        direction LR
        F1[เข้าใจการแทนข้อมูล] --> F2[รู้จักการจัดเก็บ] --> F3[จัดระเบียบโค้ด] --> F4[สร้างโครงสร้างขั้นสูง]
    end

    DataRepresentation --> DataStorage
    DataStorage --> CodeOrganization
    CodeOrganization --> AdvancedTypes
    Flow -.-> DataRepresentation
```

---

## คำอธิบายภาษาไทย (แบบละเอียด)

ภาคที่ 2 ครอบคลุมบทที่ 6–16 โดยมีเนื้อหาแบ่งเป็น 4 ช่วงหลัก ตามแผนภาพด้านบน ซึ่งแสดงลำดับการเรียนรู้ที่ต่อเนื่องกัน

---

### 1. การแทนข้อมูล (บทที่ 6–7)

**เป้าหมาย:** เข้าใจว่าคอมพิวเตอร์เก็บข้อมูลอย่างไร

#### บทที่ 6: ระบบเลขฐานสองและฐานสิบ

คอมพิวเตอร์ทำงานด้วยไฟฟ้า รู้จักแค่สถานะ "เปิด" (1) และ "ปิด" (0) ดังนั้นข้อมูลทุกชนิดจึงถูกแปลงเป็นเลขฐานสอง (Binary)

| ฐาน | คำอธิบาย | ตัวอย่าง |
|-----|---------|---------|
| **ฐานสอง (Binary)** | ตัวเลข 0 และ 1 เท่านั้น<br/>ใช้แสดงข้อมูลระดับเครื่อง | `1010₂` = 10₁₀ |
| **ฐานสิบ (Decimal)** | ตัวเลข 0-9 ตามที่มนุษย์ใช้ | `10₁₀` |
| **การแปลง** | ฐานสอง → ฐานสิบ: นำแต่ละหลักคูณ 2^(ตำแหน่ง)<br/>ฐานสิบ → ฐานสอง: หาร 2 ซ้ำๆ เอาเศษ | `1010₂ = 1×8 + 0×4 + 1×2 + 0×1 = 10` |

**Bit, Byte, และหน่วยความจำ**
- **Bit**: หน่วยย่อยที่สุด (0 หรือ 1)
- **Byte**: 8 bits = 1 byte (ใช้แทนตัวอักษร 1 ตัวใน ASCII)
- **หน่วยที่ใหญ่กว่า**: KB (1024 bytes), MB (1024 KB), GB (1024 MB)

#### บทที่ 7: เลขฐานสิบหก, ฐานแปด, ASCII, UTF-8, Unicode, Rune

| หัวข้อ | คำอธิบาย | ตัวอย่าง |
|-------|---------|---------|
| **เลขฐานสิบหก (Hexadecimal)** | ใช้เลข 0-9 และตัวอักษร A-F (10-15)<br/>นิยมเขียนนำหน้าด้วย `0x`<br/>1 hex digit = 4 bits (ครึ่ง byte) | `0xFF` = 255₁₀<br/>`0x1A` = 26₁₀ |
| **เลขฐานแปด (Octal)** | ใช้เลข 0-7<br/>นิยมเขียนนำหน้าด้วย `0o`<br/>ใช้กำหนดสิทธิ์ไฟล์ใน Unix | `0o755` = รหัสสิทธิ์ไฟล์ |
| **ASCII** | ตารางรหัส 7 bits (0-127)<br/>ครอบคลุมตัวอักษรอังกฤษ, ตัวเลข, เครื่องหมาย | `A` = 65, `a` = 97, `0` = 48 |
| **Unicode** | มาตรฐานสากลสำหรับตัวอักษรทุกภาษา<br/>มีมากกว่า 140,000 ตัวอักษร | U+0E01 = "ก" (ไทย)<br/>U+4E2D = "中" (จีน) |
| **UTF-8** | การเข้ารหัส Unicode แบบ variable-length<br/>ใช้ 1-4 bytes ต่อตัวอักษร<br/>เข้ากันได้กับ ASCII | ภาษาอังกฤษ: 1 byte<br/>ภาษาไทย: 3 bytes |
| **Rune** | ชนิดข้อมูลใน Go สำหรับแทน Unicode code point<br/>`rune` = `int32` alias<br/>ใช้ `'a'` (single quote) | `'ก'` = U+0E01 (3585)<br/>`'A'` = 65 |

**ตัวอย่างใน Go**
```go
package main

import "fmt"

func main() {
    // ตัวเลขฐานต่างๆ
    binary := 0b1010   // 10
    octal  := 0o12     // 10
    hex    := 0xA      // 10
    
    fmt.Println(binary, octal, hex)
    
    // Rune และ UTF-8
    text := "Hello, โลก"
    fmt.Println("Length in bytes:", len(text))        // 11 bytes
    fmt.Println("Number of runes:", len([]rune(text))) // 8 runes
    
    // แสดงแต่ละ rune
    for i, r := range text {
        fmt.Printf("Position %d: %c (U+%04X)\n", i, r, r)
    }
}
```

**ความสำคัญของ rune ใน Go**
- `string` ใน Go เป็น sequence ของ bytes (UTF-8)
- การวนลูปด้วย `for range` จะได้ rune อัตโนมัติ
- ใช้ `[]rune(str)` เพื่อจัดการทีละตัวอักษร

---

### 2. การจัดเก็บข้อมูล (บทที่ 8–9)

**เป้าหมาย:** เรียนรู้การประกาศตัวแปรและควบคุมการทำงาน

#### บทที่ 8: ตัวแปร, ค่าคงที่, และชนิดข้อมูลพื้นฐาน

**ตัวแปร (Variables)**
```go
// วิธีประกาศ
var name string = "Go"     // ระบุชนิด
var version = 1.23          // type inference
lang := "Golang"            // short declaration (ใช้ภายในฟังก์ชัน)

// ประกาศหลายตัว
var x, y int = 10, 20
a, b := "hello", true

// zero value (ค่าเริ่มต้น)
var i int       // 0
var s string    // "" (empty string)
var b bool      // false
var p *int      // nil
```

**ค่าคงที่ (Constants)**
```go
const Pi = 3.14159
const (
    StatusOK    = 200
    StatusNotFound = 404
)

// iota: ตัวนับอัตโนมัติ (เริ่มที่ 0)
const (
    Sunday = iota     // 0
    Monday            // 1
    Tuesday           // 2
    // ... ไปเรื่อยๆ
)
```

**ชนิดข้อมูลพื้นฐาน**

| ชนิด | ขนาด (bit) | ช่วงค่า | ตัวอย่าง |
|------|-----------|---------|---------|
| `bool` | 1 | true/false | `var ok bool = true` |
| `int` | 32/64 | ขึ้นอยู่กับระบบ | `var age int = 30` |
| `int8` | 8 | -128 ถึง 127 | `var score int8 = 100` |
| `int16` | 16 | -32768 ถึง 32767 | - |
| `int32` (rune) | 32 | -2³¹ ถึง 2³¹-1 | `var char rune = 'A'` |
| `int64` | 64 | -9×10¹⁸ ถึง 9×10¹⁸ | - |
| `uint` | 32/64 | 0 ถึง 2^n-1 | `var count uint = 100` |
| `float32` | 32 | ~1.4×10⁻⁴⁵ ถึง 3.4×10³⁸ | `var pi float32 = 3.14` |
| `float64` | 64 | ~4.9×10⁻³²⁴ ถึง 1.8×10³⁰⁸ | `var price float64 = 99.99` |
| `string` | - | sequence of bytes | `var name string = "Go"` |
| `byte` | 8 | alias of uint8 | `var b byte = 'A'` |
| `rune` | 32 | alias of int32 | `var r rune = 'ก'` |

#### บทที่ 9: คำสั่งควบคุมการทำงาน

**if-else**
```go
// if พื้นฐาน
if x > 0 {
    fmt.Println("positive")
} else if x < 0 {
    fmt.Println("negative")
} else {
    fmt.Println("zero")
}

// if พร้อม short statement
if err := doSomething(); err != nil {
    fmt.Println("Error:", err)
    return
}
```

**for loop** (Go มีแค่ `for` ไม่มี `while`)
```go
// แบบคลาสสิก (for i := 0; i < 10; i++)
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// while-style (for condition {})
count := 0
for count < 10 {
    count++
}

// infinite loop
for {
    // ใช้ break ออก
}

// range (iterating over slices, maps, strings)
numbers := []int{1, 2, 3}
for index, value := range numbers {
    fmt.Printf("index=%d, value=%d\n", index, value)
}

// range บน string ได้ rune
for i, r := range "Go" {
    fmt.Printf("%d: %c\n", i, r)
}
```

**switch** (ไม่ต้องใช้ `break`)
```go
// switch กับค่า
switch day := time.Now().Weekday(); day {
case time.Saturday, time.Sunday:
    fmt.Println("Weekend")
default:
    fmt.Println("Weekday")
}

// switch แบบไม่มี expression (ใช้แทน if-else chain)
score := 85
switch {
case score >= 90:
    fmt.Println("A")
case score >= 80:
    fmt.Println("B")
case score >= 70:
    fmt.Println("C")
default:
    fmt.Println("F")
}

// fallthrough (ไปตรวจ case ถัดไป)
switch 2 {
case 1:
    fmt.Println("1")
case 2:
    fmt.Println("2")
    fallthrough
case 3:
    fmt.Println("3")  // จะพิมพ์ "2" และ "3"
}
```

---

### 3. การจัดระเบียบโค้ด (บทที่ 10–12)

**เป้าหมาย:** เขียนโค้ดที่เป็นระเบียบ แบ่งเป็นฟังก์ชันและแพคเกจ

#### บทที่ 10: ฟังก์ชัน (Functions)

```go
// ฟังก์ชันพื้นฐาน
func add(x int, y int) int {
    return x + y
}

// การประกาศแบบสั้น
func add(x, y int) int {
    return x + y
}

// คืนค่าหลายค่า
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// named return values
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return  // naked return
}

// variadic parameters (รับพารามิเตอร์จำนวนไม่แน่นอน)
func sum(numbers ...int) int {
    total := 0
    for _, n := range numbers {
        total += n
    }
    return total
}
// เรียกใช้: sum(1, 2, 3) หรือ sum([]int{1,2,3}...)

// defer: ทำงานก่อนออกจากฟังก์ชัน
func readFile(filename string) error {
    f, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer f.Close()  // จะถูกเรียกเมื่อฟังก์ชันจบ
    
    // อ่านไฟล์...
    return nil
}

// panic และ recover
func safeDivide(a, b int) (result int) {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
            result = 0
        }
    }()
    
    if b == 0 {
        panic("division by zero")
    }
    return a / b
}
```

#### บทที่ 11: แพคเกจและการนำเข้า (Packages and Imports)

**หลักการ**
- โค้ด Go จัดกลุ่มเป็นแพคเกจ (package)
- `package main` คือ executable program
- แพคเกจอื่นๆ คือ reusable libraries

**การตั้งชื่อและ visibility**
- ตัวพิมพ์ใหญ่: **exported** (สามารถเข้าถึงจากภายนอกแพคเกจได้)
- ตัวพิมพ์เล็ก: **unexported** (ใช้ภายในแพคเกจเท่านั้น)

**ตัวอย่างโครงสร้างโปรเจกต์**
```
myproject/
├── go.mod
├── main.go
└── math/
    └── math.go
```

**math/math.go**
```go
package math

// Exported function (ขึ้นต้นด้วยตัวพิมพ์ใหญ่)
func Add(a, b int) int {
    return a + b
}

// Unexported function (ใช้ภายในแพคเกจ)
func subtract(a, b int) int {
    return a - b
}
```

**main.go**
```go
package main

import (
    "fmt"
    "myproject/math"  // import แพคเกจ custom
)

func main() {
    result := math.Add(5, 3)
    fmt.Println(result)  // 8
    
    // math.subtract(5, 3)  // ERROR: ไม่สามารถเข้าถึงได้
}
```

#### บทที่ 12: การเริ่มต้นทำงานของแพคเกจ (Package Initialization)

**ลำดับการทำงาน**
1. ตัวแปรระดับแพคเกจถูกกำหนดค่า
2. ฟังก์ชัน `init()` ถูกเรียก (เรียงตามลำดับในไฟล์ และข้ามแพคเกจ)
3. ฟังก์ชัน `main()` ถูกเรียก (เฉพาะแพคเกจ main)

```go
package main

import "fmt"

var globalVar = initGlobal()

func initGlobal() string {
    fmt.Println("1. Initializing global variable")
    return "ready"
}

func init() {
    fmt.Println("2. First init function")
}

func init() {
    fmt.Println("3. Second init function")
}

func main() {
    fmt.Println("4. Main function")
}

// ผลลัพธ์:
// 1. Initializing global variable
// 2. First init function
// 3. Second init function
// 4. Main function
```

**การใช้ init() เพื่อตั้งค่า**
```go
package database

import (
    "database/sql"
    _ "github.com/lib/pq"  // blank import: เรียกเฉพาะ init()
)

var DB *sql.DB

func init() {
    // ตั้งค่า connection pool
    var err error
    DB, err = sql.Open("postgres", "connection string")
    if err != nil {
        panic(err)
    }
}
```

---

### 4. ชนิดข้อมูลขั้นสูง (บทที่ 13–16)

**เป้าหมาย:** สร้างโครงสร้างข้อมูลที่ซับซ้อนและเขียนโค้ดที่ยืดหยุ่น

#### บทที่ 13: การสร้างชนิดข้อมูลใหม่ (Types)

```go
// type alias
type MyInt int
type UserID string

// struct (โครงสร้างข้อมูล)
type Person struct {
    Name    string
    Age     int
    Address string
}

// struct พร้อม tags (ใช้กับ JSON, validation)
type User struct {
    ID        uint   `json:"id"`
    Email     string `json:"email" validate:"required,email"`
    Password  string `json:"-"`  // ไม่ถูกส่งใน JSON
    CreatedAt time.Time `json:"created_at"`
}

// embedded struct
type Employee struct {
    Person      // embedded field (ไม่ต้องมีชื่อ)
    Position string
    Salary   float64
}

func main() {
    var id MyInt = 100
    var uid UserID = "user-123"
    
    // สร้าง struct
    p1 := Person{Name: "Alice", Age: 30}
    p2 := Person{
        Name: "Bob",
        Age:  25,
    }
    
    // เข้าถึง field
    fmt.Println(p1.Name)
    p1.Age = 31
    
    // embedded struct
    emp := Employee{
        Person:   Person{Name: "Charlie", Age: 35},
        Position: "Developer",
        Salary:   75000,
    }
    fmt.Println(emp.Name)  // เข้าถึง field จาก Person โดยตรง
}
```

#### บทที่ 14: เมธอด (Methods)

เมธอดคือฟังก์ชันที่ผูกกับ type (receiver)

```go
// กำหนด type
type Rectangle struct {
    Width, Height float64
}

// Value receiver (รับสำเนา)
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Pointer receiver (รับ reference)
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}

// ใช้ pointer receiver เมื่อต้องการแก้ไข struct
func (r *Rectangle) SetWidth(w float64) {
    r.Width = w
}

// เมธอดบนชนิดอื่นๆ (ไม่ใช่แค่ struct)
type Counter int

func (c *Counter) Increment() {
    *c++
}

func (c Counter) Value() int {
    return int(c)
}

func main() {
    rect := Rectangle{Width: 10, Height: 5}
    
    // value receiver (ไม่แก้ไขต้นฉบับ)
    area := rect.Area()  // 50
    
    // pointer receiver (แก้ไขต้นฉบับ)
    rect.Scale(2)  // rect เปลี่ยนเป็น Width=20, Height=10
    
    // counter
    var cnt Counter = 5
    cnt.Increment()
    fmt.Println(cnt.Value())  // 6
}
```

**เมื่อใช้ value receiver vs pointer receiver**

| Value Receiver | Pointer Receiver |
|----------------|------------------|
| รับสำเนา ไม่แก้ไขต้นฉบับ | รับ reference แก้ไขต้นฉบับได้ |
| เหมาะกับ struct เล็ก | เหมาะกับ struct ใหญ่ |
| เรียกใช้กับ value หรือ pointer ก็ได้ | เรียกใช้กับ value หรือ pointer ก็ได้ |
| ไม่เปลี่ยนค่าในเมธอด | เปลี่ยนค่าในเมธอดได้ |

#### บทที่ 15: พอยน์เตอร์ (Pointers)

พอยน์เตอร์คือตัวแปรที่เก็บ **ที่อยู่หน่วยความจำ** (memory address)

```go
func main() {
    x := 10
    p := &x  // p ชี้ไปที่ x (เก็บ address)
    
    fmt.Println(x)   // 10
    fmt.Println(p)   // 0xc0000120a0 (address)
    fmt.Println(*p)  // 10 (dereference: อ่านค่า)
    
    // เปลี่ยนค่าผ่าน pointer
    *p = 20
    fmt.Println(x)   // 20 (เปลี่ยนแล้ว)
    
    // zero value ของ pointer คือ nil
    var q *int
    fmt.Println(q)  // nil
    
    // สร้าง pointer ใหม่ด้วย new
    r := new(int)
    *r = 100
    fmt.Println(*r)  // 100
}

// ฟังก์ชันที่รับ pointer
func increment(p *int) {
    *p++
}

// ฟังก์ชันที่คืน pointer
func newInt(value int) *int {
    return &value
}

// เปรียบเทียบ: pass by value vs pass by pointer
func byValue(x int) {
    x = 100  // ไม่มีผลนอกฟังก์ชัน
}

func byPointer(x *int) {
    *x = 100  // มีผลนอกฟังก์ชัน
}
```

**ข้อควรรู้เกี่ยวกับ pointer ใน Go**
- Go มี pointer แต่ **ไม่มี pointer arithmetic** (ต่างจาก C)
- ใช้ `&` เพื่อ get address, `*` เพื่อ dereference
- `nil` pointer dereference ทำให้เกิด panic
- ใช้ pointer กับ struct ขนาดใหญ่เพื่อประหยัดหน่วยความจำ

#### บทที่ 16: อินเทอร์เฟซ (Interfaces)

Interface กำหนด **พฤติกรรม** (method set) ที่ type ต้องมี

```go
// ประกาศ interface
type Speaker interface {
    Speak() string
}

type Animal interface {
    Speak() string
    Move() string
}

// type ที่ implement interface
type Dog struct {
    Name string
}

func (d Dog) Speak() string {
    return "Woof!"
}

func (d Dog) Move() string {
    return "Running on 4 legs"
}

type Cat struct {
    Name string
}

func (c Cat) Speak() string {
    return "Meow!"
}

func (c Cat) Move() string {
    return "Walking gracefully"
}

// ฟังก์ชันที่รับ interface
func MakeSound(s Speaker) {
    fmt.Println(s.Speak())
}

func main() {
    dog := Dog{Name: "Max"}
    cat := Cat{Name: "Luna"}
    
    // Dog และ Cat implement Speaker
    MakeSound(dog)  // Woof!
    MakeSound(cat)  // Meow!
    
    // empty interface (interface{}) สามารถเก็บค่าใดก็ได้
    var anything interface{}
    anything = 42
    anything = "hello"
    anything = Dog{Name: "Buddy"}
    
    // type assertion (ตรวจสอบชนิด)
    val, ok := anything.(Dog)
    if ok {
        fmt.Println(val.Speak())
    }
    
    // type switch
    switch v := anything.(type) {
    case int:
        fmt.Println("int:", v)
    case string:
        fmt.Println("string:", v)
    case Dog:
        fmt.Println("Dog:", v.Name)
    default:
        fmt.Println("unknown type")
    }
}
```

**Interface Composition** (การรวม interface)
```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

// ประกอบ interface
type ReadWriter interface {
    Reader
    Writer
}
```

**ข้อดีของ interface**
- **Polymorphism**: เขียนโค้ดที่ทำงานกับ type ต่างๆ ที่มีพฤติกรรมเดียวกัน
- **Decoupling**: แยก implementation ออกจาก contract
- **Testability**: สร้าง mock ได้ง่าย

**Interface ใน Go ต่างจากภาษา OOP อื่น**
- Go interface **implement implicitly** (ไม่ต้องประกาศ implements)
- Interface เป็น **ค่าที่มี 2 ส่วน**: dynamic type + dynamic value
- Interface สามารถเป็น `nil` (ถ้า type เป็น nil และ value เป็น nil)

---

## สรุปภาคที่ 2: ความสัมพันธ์ของเนื้อหา

```mermaid
flowchart LR
    subgraph Knowledge["องค์ความรู้"]
        K1[การแทนข้อมูล<br/>บทที่ 6-7] --> K2[การจัดเก็บ<br/>บทที่ 8-9]
        K2 --> K3[การจัดระเบียบ<br/>บทที่ 10-12]
        K3 --> K4[ชนิดข้อมูลขั้นสูง<br/>บทที่ 13-16]
    end
    
    subgraph Skills["ทักษะที่ได้"]
        S1[เข้าใจการทำงาน<br/>ระดับบิต/ไบต์]
        S2[เขียนโปรแกรม<br/>แบบมีเงื่อนไข]
        S3[แบ่งโค้ดเป็น<br/>ฟังก์ชัน/แพคเกจ]
        S4[สร้าง struct,<br/>methods, interfaces]
    end
    
    K1 --> S1
    K2 --> S2
    K3 --> S3
    K4 --> S4
```

**สิ่งที่ได้เรียนรู้**
1. **บทที่ 6-7**: รู้ว่าคอมพิวเตอร์จัดเก็บข้อมูลอย่างไร (Binary, Hex, UTF-8, Rune)
2. **บทที่ 8-9**: ประกาศตัวแปรและควบคุมการทำงานด้วย if, for, switch
3. **บทที่ 10-12**: จัดระเบียบโค้ดด้วยฟังก์ชันและแพคเกจ
4. **บทที่ 13-16**: สร้าง struct, กำหนดเมธอด, ใช้ pointer, และออกแบบ interface

พื้นฐานเหล่านี้จะนำไปใช้ใน **ภาคที่ 3** เพื่อจัดการโปรเจกต์, สร้าง test, และใช้ slice/map ในการจัดการข้อมูลจำนวนมาก
## ภาคที่ 3: การจัดการโปรเจกต์และโครงสร้างข้อมูลขั้นสูง (บทที่ 17–23)

**แผนภาพ: การจัดการโปรเจกต์ → ทดสอบ → คอลเลกชัน → ข้อผิดพลาด**

```mermaid
graph TD
    subgraph Modules[โมดูล]
        M1[go mod init] --> M2[go.mod]
        M2 --> M3[go get / tidy]
        M3 --> M4[go.sum]
    end

    subgraph Testing[การทดสอบ]
        T1[_test.go] --> T2[go test]
        T2 --> T3[รายงาน coverage]
    end

    subgraph Collections[คอลเลกชัน]
        C1[อาเรย์] --> C2[สไลซ์]
        C2 --> C3[แมพ]
    end

    subgraph Errors[ข้อผิดพลาด]
        E1[errors.New] --> E2[fmt.Errorf]
        E2 --> E3[errors.Is / As]
    end

    M4 --> T1
    T3 --> C1
    C3 --> E1
```

**คำอธิบาย:**  
เริ่มจากสร้างโมดูลด้วย `go mod init` ได้ไฟล์ `go.mod` และ `go.sum` จากนั้นเขียน test (`_test.go`) แล้วรัน `go test` เพื่อวัด coverage ข้อมูลที่ถูกต้องผ่านการทดสอบจะถูกนำไปใช้กับคอลเลกชัน (array, slice, map) และในที่สุดการจัดการข้อผิดพลาดจะเกิดขึ้นเมื่อพบ error

---
# ภาคที่ 3: การจัดการโปรเจกต์และโครงสร้างข้อมูลขั้นสูง (บทที่ 17–23)

## แผนภาพแสดงความสัมพันธ์ของเนื้อหา

```mermaid
flowchart TD
    subgraph ProjectManagement["1. การจัดการโปรเจกต์ (บทที่ 17-18)"]
        direction TB
        A1["Go Modules (บทที่ 17)<br/>go.mod, go.sum<br/>module, require, replace<br/>go mod init/tidy/vendor"]
        A2["Go Module Proxies (บทที่ 18)<br/>GOPROXY, GOSUMDB<br/>GOPRIVATE, Athens<br/>ความปลอดภัยและความเร็ว"]
        
        A1 --> A2
    end

    subgraph Testing["2. การทดสอบ (บทที่ 19)"]
        direction TB
        B1["Unit Tests<br/>_test.go, TestXxx(t *testing.T)<br/>t.Run(), table-driven tests<br/>t.Helper(), t.Parallel()"]
        B2["Test Coverage<br/>go test -cover<br/>go test -coverprofile<br/>go tool cover -html"]
        B3["Test Doubles<br/>Mock, Stub, Fake<br/>interface-based mocking"]
        
        B1 --> B2
        B2 --> B3
    end

    subgraph DataStructures["3. โครงสร้างข้อมูลขั้นสูง (บทที่ 20-22)"]
        direction TB
        C1["Arrays (บทที่ 20)<br/>[n]T, fixed size<br/>value type, compile-time size<br/>rarely used directly"]
        C2["Slices (บทที่ 21)<br/>[]T, dynamic size<br/>len, cap, append, copy<br/>backed by array, reference type"]
        C3["Maps (บทที่ 22)<br/>map[K]V, hash table<br/>make(), delete, comma ok idiom<br/>unordered, not safe for concurrent"]
        
        C1 --> C2
        C2 --> C3
    end

    subgraph ErrorHandling["4. การจัดการข้อผิดพลาด (บทที่ 23)"]
        direction TB
        D1["Error Interface<br/>error interface<br/>errors.New, fmt.Errorf"]
        D2["Error Wrapping<br/>%w, errors.Unwrap<br/>errors.Is, errors.As<br/>error chains"]
        D3["Best Practices<br/>check errors early<br/>don't ignore errors<br/>custom error types"]
        
        D1 --> D2
        D2 --> D3
    end

    subgraph Flow["ลำดับการเรียนรู้"]
        direction LR
        F1[จัดการโปรเจกต์] --> F2[เขียนทดสอบ] --> F3[ใช้โครงสร้างข้อมูล] --> F4[จัดการข้อผิดพลาด]
    end

    ProjectManagement --> Testing
    Testing --> DataStructures
    DataStructures --> ErrorHandling
    Flow -.-> ProjectManagement
```

---

## คำอธิบายภาษาไทย (แบบละเอียด)

ภาคที่ 3 ครอบคลุมบทที่ 17–23 โดยมีเนื้อหาแบ่งเป็น 4 ช่วงหลัก ตามแผนภาพด้านบน ซึ่งแสดงลำดับการเรียนรู้ที่ต่อเนื่องกัน

---

## 1. การจัดการโปรเจกต์ (บทที่ 17–18)

### บทที่ 17: Go Modules – การจัดการโปรเจกต์สมัยใหม่

**เป้าหมาย:** จัดการ dependencies (แพคเกจภายนอก) และ versioning

#### Go Modules คืออะไร?

Go Modules เป็นระบบ dependency management อย่างเป็นทางการของ Go เริ่มตั้งแต่ Go 1.11 (stable ตั้งแต่ 1.16) แทนที่ GOPATH workspace แบบเดิม

#### โครงสร้างไฟล์

```
myproject/
├── go.mod          # ไฟล์หลัก ระบุ module name และ dependencies
├── go.sum          # checksum ของ dependencies (ความปลอดภัย)
├── main.go
└── internal/       # โค้ดเฉพาะของโปรเจกต์
```

#### คำสั่งพื้นฐาน

```bash
# เริ่มต้นโปรเจกต์ใหม่
go mod init github.com/username/myproject

# เพิ่ม dependency อัตโนมัติ (เมื่อ import ในโค้ด)
go mod tidy

# ดาวน์โหลด dependencies
go mod download

# แสดง dependency graph
go mod graph

# แก้ไข go.mod โดยตรง
go mod edit -require=github.com/gorilla/mux@v1.8.0

# ลบ unused dependencies
go mod tidy -v

# สร้าง vendor folder (copy dependencies เข้าโปรเจกต์)
go mod vendor
```

#### ตัวอย่าง go.mod

```go
module github.com/username/myapp

go 1.21

require (
    github.com/gin-gonic/gin v1.9.1
    github.com/go-sql-driver/mysql v1.7.1
    gorm.io/gorm v1.25.5
)

require (
    github.com/bytedance/sonic v1.9.1 // indirect
    github.com/jinzhu/inflection v1.0.0 // indirect
)

replace github.com/old/package => github.com/new/package v1.0.0

exclude github.com/broken/package v1.2.3
```

**คำอธิบายแต่ละส่วน:**
- `module`: ชื่อของโปรเจกต์ (ปกติเป็น repository URL)
- `go`: เวอร์ชัน Go ที่ใช้
- `require`: dependencies ที่โปรเจกต์ต้องการ
- `indirect`: dependency ที่ถูกนำเข้ามาทางอ้อม
- `replace`: ใช้แทนที่แพคเกจด้วยเวอร์ชันหรือ source อื่น
- `exclude`: ห้ามใช้เวอร์ชันที่ระบุ

#### ตัวอย่างการใช้งานจริง

**1. สร้างโปรเจกต์ใหม่**

```bash
# สร้างโฟลเดอร์
mkdir my-rest-api
cd my-rest-api

# เริ่มต้น Go module
go mod init github.com/mycompany/my-rest-api

# สร้างไฟล์ main.go
cat > main.go << 'EOF'
package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "pong"})
    })
    r.Run() // listen and serve on 0.0.0.0:8080
}
EOF

# รัน go mod tidy จะดึง gin มาให้อัตโนมัติ
go mod tidy

# ดูผลลัพธ์ go.mod
cat go.mod
```

**ผลลัพธ์ go.mod:**
```
module github.com/mycompany/my-rest-api

go 1.21

require github.com/gin-gonic/gin v1.9.1

require (
    github.com/bytedance/sonic v1.9.1 // indirect
    github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311 // indirect
    github.com/gabriel-vasile/mimetype v1.4.2 // indirect
    // ... indirect dependencies อื่นๆ
)
```

**2. การอัปเกรด dependency**

```bash
# ดูเวอร์ชันล่าสุด
go list -u -m all

# อัปเกรดแพคเกจเฉพาะ
go get -u github.com/gin-gonic/gin

# อัปเกรดเป็นเวอร์ชันเฉพาะ
go get github.com/gin-gonic/gin@v1.9.0

# อัปเกรดทั้งหมด
go get -u ./...
```

**3. การใช้ replace สำหรับ local development**

```go
// go.mod
module github.com/mycompany/myapp

go 1.21

require (
    github.com/mycompany/mylib v1.0.0
)

// ใช้ local path แทน (ตอนพัฒนา)
replace github.com/mycompany/mylib => ../mylib
```

**4. Semantic Import Versioning (SIV)**

Go ใช้ major version ใน import path เมื่อ major version > 1

```go
// v1 (ไม่ต้องระบุ /v1)
import "github.com/user/pkg"

// v2 ขึ้นไป ต้องระบุใน path
import "github.com/user/pkg/v2"
import "github.com/user/pkg/v3"
```

---

### บทที่ 18: Go Module Proxies

**เป้าหมาย:** เพิ่มความเร็วและความปลอดภัยในการดาวน์โหลด dependencies

#### Go Module Proxy คืออะไร?

Go Module Proxy เป็นเซิร์ฟเวอร์ที่เก็บ cache ของ Go modules ช่วยให้:
- ดาวน์โหลด dependencies ได้เร็วขึ้น (cache ในประเทศ)
- ทำงานในองค์กรที่มี firewall
- ป้องกันการลบ module ออกจาก GitHub (sum.golang.org)
- ตรวจสอบความถูกต้องของ module

#### ตัวแปรสภาพแวดล้อมที่สำคัญ

```bash
# ตั้งค่า GOPROXY (ใช้ proxy หลายตัว)
export GOPROXY=https://proxy.golang.org,direct

# ใช้ private proxy
export GOPROXY=https://goproxy.io,https://proxy.golang.org,direct

# ตั้งค่า GOSUMDB (ตรวจสอบ checksum)
export GOSUMDB=sum.golang.org

# ตั้งค่า GOPRIVATE (private repositories)
export GOPRIVATE=github.com/mycompany/*,gitlab.com/private/*

# ปิด checksum สำหรับ private repo
export GONOSUMDB=github.com/mycompany/*
```

#### ตัวอย่างการใช้งานจริง

**1. ตั้งค่า proxy สำหรับองค์กร**

```bash
# .bashrc หรือ .zshrc
export GOPROXY=https://goproxy.company.com,https://proxy.golang.org,direct
export GOPRIVATE=github.com/mycompany/*
export GONOSUMDB=github.com/mycompany/*
export GOSUMDB=sum.golang.org
```

**2. ใช้ Athens (self-hosted proxy)**

Athens เป็น Go module proxy ที่สามารถติดตั้งเองได้

```yaml
# docker-compose.yml
version: '3'
services:
  athens:
    image: gomods/athens:latest
    ports:
      - "3000:3000"
    environment:
      - ATHENS_STORAGE_TYPE=disk
      - ATHENS_DISK_STORAGE_ROOT=/var/lib/athens
    volumes:
      - ./athens-storage:/var/lib/athens
```

```bash
# เริ่มต้น Athens
docker-compose up -d

# ตั้งค่า GOPROXY
export GOPROXY=http://localhost:3000
```

**3. การทำงานของ GOPROXY**

```
1. go build/run/test
   ↓
2. ตรวจสอบ GOPROXY
   ↓
3. proxy.golang.org มี cache ไหม?
   ├── มี → ดาวน์โหลดจาก proxy
   └── ไม่มี → ดึงจาก VCS (GitHub) → เก็บ cache
   ↓
4. ตรวจสอบ GOSUMDB
   ↓
5. ดาวน์โหลด module
```

**4. ปัญหาที่พบบ่อยและ解决方法**

```bash
# ปัญหา: checksum mismatch สำหรับ private repo
# วิธีแก้: ตั้ง GONOSUMDB
export GONOSUMDB=github.com/private/*

# ปัญหา: proxy ไม่มี module (404)
# วิธีแก้: ใช้ direct
export GOPROXY=https://proxy.golang.org,direct

# ปัญหา: module ถูกลบจาก GitHub
# วิธีแก้: ใช้ proxy ที่มี cache (proxy.golang.org มี cache)
# หรือใช้ vendor
go mod vendor
```

---

## 2. การทดสอบหน่วย (บทที่ 19)

**เป้าหมาย:** เขียน test เพื่อให้โค้ดมีความมั่นใจและบำรุงรักษาง่าย

### บทที่ 19: Unit Tests

#### โครงสร้างการทดสอบ

```go
// calculator/calculator.go
package calculator

func Add(a, b int) int {
    return a + b
}

func Divide(a, b int) (int, error) {
    if b == 0 {
        return 0, ErrDivisionByZero
    }
    return a / b, nil
}
```

```go
// calculator/calculator_test.go
package calculator

import (
    "testing"
)

// 1. Test function (ต้องขึ้นต้นด้วย Test)
func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5
    
    if result != expected {
        t.Errorf("Add(2,3) = %d; want %d", result, expected)
    }
}

// 2. Table-driven tests (แนะนำ)
func TestDivide(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
        hasError bool
    }{
        {"positive numbers", 10, 2, 5, false},
        {"negative numbers", -10, 2, -5, false},
        {"division by zero", 10, 0, 0, true},
        {"zero divided", 0, 5, 0, false},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result, err := Divide(tt.a, tt.b)
            
            if tt.hasError {
                if err == nil {
                    t.Errorf("expected error but got none")
                }
                return
            }
            
            if err != nil {
                t.Errorf("unexpected error: %v", err)
            }
            
            if result != tt.expected {
                t.Errorf("Divide(%d,%d) = %d; want %d", 
                    tt.a, tt.b, result, tt.expected)
            }
        })
    }
}

// 3. Helper function
func testHelper(t *testing.T, msg string) {
    t.Helper() // บรรทัดนี้บอกว่าเป็น helper (error จะแสดง caller)
    if msg == "" {
        t.Fatal("message cannot be empty")
    }
}
```

#### ตัวอย่างการใช้งานจริง

**1. การทดสอบ HTTP Handler**

```go
// handler/user_handler.go
package handler

import (
    "encoding/json"
    "net/http"
)

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

type UserHandler struct {
    userService UserService
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
    // รับ ID จาก URL
    id := r.URL.Query().Get("id")
    if id == "" {
        http.Error(w, "id required", http.StatusBadRequest)
        return
    }
    
    user, err := h.userService.GetUser(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}
```

```go
// handler/user_handler_test.go
package handler

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
)

// Mock UserService
type mockUserService struct {
    getUserFunc func(id string) (*User, error)
}

func (m *mockUserService) GetUser(id string) (*User, error) {
    return m.getUserFunc(id)
}

func TestGetUser(t *testing.T) {
    tests := []struct {
        name           string
        userID         string
        mockService    func() UserService
        expectedStatus int
        expectedUser   *User
    }{
        {
            name:   "success",
            userID: "1",
            mockService: func() UserService {
                return &mockUserService{
                    getUserFunc: func(id string) (*User, error) {
                        return &User{ID: 1, Name: "John"}, nil
                    },
                }
            },
            expectedStatus: http.StatusOK,
            expectedUser:   &User{ID: 1, Name: "John"},
        },
        {
            name:   "user not found",
            userID: "999",
            mockService: func() UserService {
                return &mockUserService{
                    getUserFunc: func(id string) (*User, error) {
                        return nil, ErrUserNotFound
                    },
                }
            },
            expectedStatus: http.StatusNotFound,
            expectedUser:   nil,
        },
        {
            name:           "missing id",
            userID:         "",
            mockService:    func() UserService { return &mockUserService{} },
            expectedStatus: http.StatusBadRequest,
            expectedUser:   nil,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // สร้าง handler
            handler := &UserHandler{
                userService: tt.mockService(),
            }
            
            // สร้าง request
            req := httptest.NewRequest("GET", "/user?id="+tt.userID, nil)
            w := httptest.NewRecorder()
            
            // เรียก handler
            handler.GetUser(w, req)
            
            // ตรวจสอบ status code
            if w.Code != tt.expectedStatus {
                t.Errorf("status = %d; want %d", w.Code, tt.expectedStatus)
            }
            
            // ตรวจสอบ response body
            if tt.expectedUser != nil {
                var user User
                json.NewDecoder(w.Body).Decode(&user)
                if user.ID != tt.expectedUser.ID {
                    t.Errorf("user.ID = %d; want %d", user.ID, tt.expectedUser.ID)
                }
            }
        })
    }
}
```

**2. การใช้ Subtests และ Parallel**

```go
func TestDatabaseQueries(t *testing.T) {
    // setup database (once)
    db := setupTestDB(t)
    defer db.Close()
    
    tests := []struct {
        name string
        query string
        expected int
    }{
        {"count users", "SELECT COUNT(*) FROM users", 10},
        {"count orders", "SELECT COUNT(*) FROM orders", 25},
    }
    
    for _, tt := range tests {
        tt := tt // capture range variable (Go 1.21+ ไม่ต้องแล้ว)
        t.Run(tt.name, func(t *testing.T) {
            t.Parallel() // รัน parallel
            
            var count int
            err := db.QueryRow(tt.query).Scan(&count)
            if err != nil {
                t.Fatal(err)
            }
            if count != tt.expected {
                t.Errorf("count = %d; want %d", count, tt.expected)
            }
        })
    }
}
```

**3. Test Coverage**

```bash
# ดู coverage ใน terminal
go test -cover

# สร้าง coverage profile
go test -coverprofile=coverage.out

# ดู coverage ใน browser
go tool cover -html=coverage.out

# ดูเฉพาะบาง package
go test -cover ./internal/...

# ตั้งค่า threshold (ใน CI)
go test -cover -coverprofile=coverage.out ./...
go tool cover -func=coverage.out | grep total
# total: (statements) 85.2%
```

**4. Test Main (Setup/Teardown ทั้ง package)**

```go
// package_test.go
package mypackage

import (
    "os"
    "testing"
)

// TestMain รันก่อน test ทั้งหมด
func TestMain(m *testing.M) {
    // setup (รันก่อน test)
    setup()
    
    // รัน test
    code := m.Run()
    
    // teardown (รันหลัง test)
    teardown()
    
    os.Exit(code)
}

func setup() {
    // ตั้งค่า database connection, mock server, etc.
}

func teardown() {
    // clean up
}
```

**5. การใช้ testdata folder**

```
myproject/
├── internal/
│   └── parser/
│       ├── parser.go
│       ├── parser_test.go
│       └── testdata/           # ไฟล์สำหรับ test
│           ├── valid.json
│           ├── invalid.json
│           └── sample.csv
```

```go
func TestParseJSON(t *testing.T) {
    // อ่านไฟล์จาก testdata
    data, err := os.ReadFile("testdata/valid.json")
    if err != nil {
        t.Fatal(err)
    }
    
    result, err := ParseJSON(data)
    if err != nil {
        t.Fatal(err)
    }
    
    // assert...
}
```

---

## 3. โครงสร้างข้อมูลขั้นสูง (บทที่ 20–22)

### บทที่ 20: Arrays (อาเรย์)

**เป้าหมาย:** เข้าใจโครงสร้างข้อมูลแบบ fixed-size

#### ลักษณะสำคัญ

- **Fixed size**: ขนาดคงที่ กำหนดตอน compile time
- **Value type**: การ assign หรือส่งเป็น argument จะ copy ทั้ง array
- **Contiguous memory**: เก็บข้อมูลต่อเนื่องในหน่วยความจำ

#### ตัวอย่างการใช้งาน

```go
package main

import "fmt"

func main() {
    // 1. ประกาศ array
    var arr1 [5]int                    // [0 0 0 0 0]
    arr2 := [3]string{"a", "b", "c"}   // [a b c]
    arr3 := [...]int{1, 2, 3}          // compiler นับขนาดอัตโนมัติ
    
    fmt.Println(arr1, arr2, arr3)
    
    // 2. เข้าถึงและแก้ไข
    arr1[0] = 10
    arr1[1] = 20
    fmt.Println(arr1[0])  // 10
    
    // 3. array เป็น value type (copy)
    original := [3]int{1, 2, 3}
    copy := original
    copy[0] = 100
    fmt.Println(original) // [1 2 3] (ไม่เปลี่ยน)
    fmt.Println(copy)     // [100 2 3]
    
    // 4. วนลูป array
    nums := [5]int{10, 20, 30, 40, 50}
    
    // แบบ index
    for i := 0; i < len(nums); i++ {
        fmt.Println(nums[i])
    }
    
    // แบบ range
    for index, value := range nums {
        fmt.Printf("index=%d, value=%d\n", index, value)
    }
    
    // 5. ส่ง array เข้าฟังก์ชัน (copy)
    func modifyArray(arr [3]int) {
        arr[0] = 999
    }
    
    data := [3]int{1, 2, 3}
    modifyArray(data)
    fmt.Println(data) // [1 2 3] (ไม่เปลี่ยน)
    
    // 6. ส่ง pointer ของ array (rarely used)
    func modifyArrayPtr(arr *[3]int) {
        arr[0] = 999
    }
    
    modifyArrayPtr(&data)
    fmt.Println(data) // [999 2 3]
}
```

#### ข้อควรจำ

- **Arrays ใช้ไม่บ่อยใน Go** เพราะขนาดคงที่และ copy ข้อมูล
- **ส่วนใหญ่ใช้ Slices แทน** (บทถัดไป)
- ใช้ array เมื่อต้องการ:
  - ขนาดคงที่แน่นอน (เช่น 7 วันในสัปดาห์)
  - ประสิทธิภาพสูงมาก (stack allocation)
  - ทำงานกับระบบ low-level

---

### บทที่ 21: Slices (สไลซ์)

**เป้าหมาย:** เข้าใจโครงสร้างข้อมูล dynamic ที่ใช้บ่อยที่สุดใน Go

#### Slices คืออะไร?

Slice คือ **view** ที่ยืดหยุ่นของ array ประกอบด้วย 3 ส่วน:
- **pointer** → ชี้ไปยัง underlying array
- **len** → ความยาวปัจจุบัน
- **cap** → ความจุทั้งหมด

```
Slice Structure:
┌─────────────────────────────────┐
│ Slice Header                    │
├─────────────────────────────────┤
│ ptr ──────┐                     │
│ len = 3   │                     │
│ cap = 5   │                     │
└───────────┼─────────────────────┘
            ↓
     ┌─────┬─────┬─────┬─────┬─────┐
     │ 10  │ 20  │ 30  │ 40  │ 50  │  ← underlying array
     └─────┴─────┴─────┴─────┴─────┘
            ↑           ↑
            ptr         ptr+cap
```

#### ตัวอย่างการใช้งาน

```go
package main

import "fmt"

func main() {
    // 1. สร้าง slice
    var s1 []int                    // nil slice (len=0, cap=0)
    s2 := []int{1, 2, 3}           // literal
    s3 := make([]int, 5)           // length=5, capacity=5
    s4 := make([]int, 5, 10)       // length=5, capacity=10
    
    fmt.Println(s1, s2, s3, s4)
    
    // 2. append (เพิ่มสมาชิก)
    numbers := []int{1, 2, 3}
    numbers = append(numbers, 4)     // [1,2,3,4]
    numbers = append(numbers, 5, 6)  // [1,2,3,4,5,6]
    
    // append slice กับ slice
    more := []int{7, 8, 9}
    numbers = append(numbers, more...)  // ... ใช้ unpack slice
    
    fmt.Println(numbers)  // [1,2,3,4,5,6,7,8,9]
    
    // 3. len และ cap
    slice := make([]int, 3, 5)
    fmt.Printf("len=%d, cap=%d\n", len(slice), cap(slice))  // len=3, cap=5
    
    slice = append(slice, 100)
    fmt.Printf("len=%d, cap=%d\n", len(slice), cap(slice))  // len=4, cap=5
    
    slice = append(slice, 200, 300)
    fmt.Printf("len=%d, cap=%d\n", len(slice), cap(slice))  // len=6, cap=10 (double)
    
    // 4. Slicing (สร้าง slice ใหม่จาก slice)
    original := []int{1, 2, 3, 4, 5}
    sub := original[1:4]   // [2,3,4]
    fmt.Println(sub)
    
    // sub ชี้ไปที่ underlying array เดียวกับ original
    sub[0] = 999
    fmt.Println(original)  // [1,999,3,4,5] (เปลี่ยนด้วย!)
    
    // 5. copy slice (deep copy)
    src := []int{1, 2, 3}
    dst := make([]int, len(src))
    copy(dst, src)
    dst[0] = 999
    fmt.Println(src)  // [1,2,3] (ไม่เปลี่ยน)
    fmt.Println(dst)  // [999,2,3]
    
    // 6. การลบ element
    // ลบ element ที่ index 2
    s := []int{1, 2, 3, 4, 5}
    index := 2
    s = append(s[:index], s[index+1:]...)
    fmt.Println(s)  // [1,2,4,5]
    
    // 7. การ insert element
    s = []int{1, 2, 4, 5}
    s = append(s[:2], append([]int{3}, s[2:]...)...)
    fmt.Println(s)  // [1,2,3,4,5]
}

// 8. ฟังก์ชันที่รับและคืน slice
func filter(numbers []int, predicate func(int) bool) []int {
    result := make([]int, 0, len(numbers))
    for _, n := range numbers {
        if predicate(n) {
            result = append(result, n)
        }
    }
    return result
}
```

#### ตัวอย่างการใช้งานจริง

**1. Pagination (การแบ่งหน้า)**

```go
type PaginatedResult struct {
    Data       []User
    Total      int
    Page       int
    PerPage    int
    TotalPages int
}

func GetUsers(page, perPage int) (*PaginatedResult, error) {
    // ดึงข้อมูลทั้งหมด
    allUsers, err := userRepo.FindAll()
    if err != nil {
        return nil, err
    }
    
    total := len(allUsers)
    totalPages := (total + perPage - 1) / perPage
    
    // คำนวณ start และ end index
    start := (page - 1) * perPage
    end := start + perPage
    
    if start >= total {
        return &PaginatedResult{
            Data:       []User{},
            Total:      total,
            Page:       page,
            PerPage:    perPage,
            TotalPages: totalPages,
        }, nil
    }
    
    if end > total {
        end = total
    }
    
    // slice ข้อมูล
    return &PaginatedResult{
        Data:       allUsers[start:end],
        Total:      total,
        Page:       page,
        PerPage:    perPage,
        TotalPages: totalPages,
    }, nil
}
```

**2. Stack (LIFO) ด้วย Slice**

```go
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
    if len(s.items) == 0 {
        var zero T
        return zero, false
    }
    
    lastIndex := len(s.items) - 1
    item := s.items[lastIndex]
    s.items = s.items[:lastIndex]
    return item, true
}

func (s *Stack[T]) Peek() (T, bool) {
    if len(s.items) == 0 {
        var zero T
        return zero, false
    }
    return s.items[len(s.items)-1], true
}
```

**3. Queue (FIFO) ด้วย Slice**

```go
type Queue[T any] struct {
    items []T
}

func (q *Queue[T]) Enqueue(item T) {
    q.items = append(q.items, item)
}

func (q *Queue[T]) Dequeue() (T, bool) {
    if len(q.items) == 0 {
        var zero T
        return zero, false
    }
    
    item := q.items[0]
    q.items = q.items[1:]
    return item, true
}
```

**4. Performance Tip: Preallocate capacity**

```go
// ไม่ดี: ทำให้เกิดการ reallocation บ่อย
func badBuild() []int {
    var result []int
    for i := 0; i < 1000000; i++ {
        result = append(result, i)  // reallocate หลายครั้ง
    }
    return result
}

// ดี: preallocate capacity
func goodBuild() []int {
    result := make([]int, 0, 1000000)  // allocate ครั้งเดียว
    for i := 0; i < 1000000; i++ {
        result = append(result, i)
    }
    return result
}
```

---

### บทที่ 22: Maps (แมพ)

**เป้าหมาย:** เข้าใจโครงสร้างข้อมูล key-value ที่ใช้บ่อย

#### Maps คืออะไร?

Map เป็น **hash table** ที่เก็บ key-value pairs
- Key ต้องเป็น comparable type (==, !=)
- Value สามารถเป็นชนิดใดก็ได้
- **ไม่ปลอดภัยต่อ concurrent access** (ต้องใช้ mutex หรือ sync.Map)

#### ตัวอย่างการใช้งาน

```go
package main

import "fmt"

func main() {
    // 1. สร้าง map
    var m1 map[string]int           // nil map (อ่านได้ แต่เขียนไม่ได้)
    m2 := make(map[string]int)      // empty map
    m3 := map[string]int{           // map literal
        "apple":  5,
        "banana": 3,
    }
    
    // nil map จะ panic ถ้าเขียน
    // m1["key"] = 1  // panic!
    
    // 2. เพิ่ม/แก้ไข
    m2["apple"] = 10
    m2["banana"] = 7
    m2["orange"] = 4
    
    fmt.Println(m2)  // map[apple:10 banana:7 orange:4]
    
    // 3. อ่านค่า
    value := m2["apple"]           // 10
    value2 := m2["grape"]          // 0 (zero value, ไม่มี key นี้)
    
    // 4. comma ok idiom (ตรวจสอบว่ามี key หรือไม่)
    if val, ok := m2["banana"]; ok {
        fmt.Printf("banana exists with value %d\n", val)
    }
    
    if _, ok := m2["grape"]; !ok {
        fmt.Println("grape does not exist")
    }
    
    // 5. ลบ key
    delete(m2, "orange")
    fmt.Println(m2)  // map[apple:10 banana:7]
    
    // 6. วนลูป map (order is random)
    for key, value := range m2 {
        fmt.Printf("%s: %d\n", key, value)
    }
    
    // 7. วนลูปเฉพาะ key
    for key := range m2 {
        fmt.Println(key)
    }
    
    // 8. map ของ struct
    type User struct {
        Name string
        Age  int
    }
    
    users := map[int]User{
        1: {Name: "Alice", Age: 30},
        2: {Name: "Bob", Age: 25},
    }
    
    // 9. map ของ map (nested)
    matrix := map[string]map[string]int{
        "row1": {"col1": 1, "col2": 2},
        "row2": {"col1": 3, "col2": 4},
    }
    
    // 10. map กับ slice เป็น value
    tags := map[string][]string{
        "golang": {"backend", "fast", "concurrent"},
        "python": {"ai", "ml", "easy"},
    }
}
```

#### ตัวอย่างการใช้งานจริง

**1. Cache (LRU cache แบบง่าย)**

```go
type Cache struct {
    store map[string]interface{}
    mu    sync.RWMutex
    ttl   time.Duration
}

func NewCache(ttl time.Duration) *Cache {
    c := &Cache{
        store: make(map[string]interface{}),
        ttl:   ttl,
    }
    
    // start cleanup goroutine
    go c.cleanup()
    
    return c
}

func (c *Cache) Set(key string, value interface{}) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.store[key] = value
}

func (c *Cache) Get(key string) (interface{}, bool) {
    c.mu.RLock()
    defer c.mu.RUnlock()
    val, ok := c.store[key]
    return val, ok
}

func (c *Cache) cleanup() {
    ticker := time.NewTicker(c.ttl)
    for range ticker.C {
        c.mu.Lock()
        // ลบ expired keys (สมมติมีเวลา expire)
        for k := range c.store {
            delete(c.store, k)
        }
        c.mu.Unlock()
    }
}
```

**2. Group By (การจัดกลุ่มข้อมูล)**

```go
type Order struct {
    ID       int
    UserID   int
    Amount   float64
    Status   string
}

func GroupOrdersByUser(orders []Order) map[int][]Order {
    result := make(map[int][]Order)
    
    for _, order := range orders {
        result[order.UserID] = append(result[order.UserID], order)
    }
    
    return result
}

func GroupOrdersByStatus(orders []Order) map[string][]Order {
    result := make(map[string][]Order)
    
    for _, order := range orders {
        result[order.Status] = append(result[order.Status], order)
    }
    
    return result
}
```

**3. Set (ใช้ map แทน set)**

```go
type Set[T comparable] struct {
    items map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
    return &Set[T]{
        items: make(map[T]struct{}),
    }
}

func (s *Set[T]) Add(item T) {
    s.items[item] = struct{}{}
}

func (s *Set[T]) Remove(item T) {
    delete(s.items, item)
}

func (s *Set[T]) Contains(item T) bool {
    _, ok := s.items[item]
    return ok
}

func (s *Set[T]) Size() int {
    return len(s.items)
}

func (s *Set[T]) List() []T {
    result := make([]T, 0, len(s.items))
    for item := range s.items {
        result = append(result, item)
    }
    return result
}

// ใช้งาน
func main() {
    tags := NewSet[string]()
    tags.Add("golang")
    tags.Add("programming")
    tags.Add("golang")  // duplicate (ไม่มีผล)
    
    fmt.Println(tags.Contains("golang"))   // true
    fmt.Println(tags.Size())               // 2
    fmt.Println(tags.List())               // [golang programming]
}
```

**4. Counter (นับความถี่)**

```go
func WordCount(text string) map[string]int {
    words := strings.Fields(strings.ToLower(text))
    counter := make(map[string]int)
    
    for _, word := range words {
        // ลบ punctuation
        word = strings.Trim(word, ".,!?;:")
        counter[word]++
    }
    
    return counter
}

// ใช้งาน
text := "Go is great. Go is fast. Go is concurrent!"
counts := WordCount(text)
// map[go:3 is:3 great:1 fast:1 concurrent:1]
```

**5. Graph (กราฟ) ด้วย map**

```go
type Graph struct {
    nodes map[string][]string  // adjacency list
}

func NewGraph() *Graph {
    return &Graph{
        nodes: make(map[string][]string),
    }
}

func (g *Graph) AddEdge(from, to string) {
    g.nodes[from] = append(g.nodes[from], to)
    g.nodes[to] = append(g.nodes[to], from)  // undirected
}

func (g *Graph) BFS(start, target string) []string {
    visited := make(map[string]bool)
    queue := []string{start}
    parent := make(map[string]string)
    visited[start] = true
    
    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]
        
        if current == target {
            // reconstruct path
            path := []string{}
            for node := target; node != ""; node = parent[node] {
                path = append([]string{node}, path...)
            }
            return path
        }
        
        for _, neighbor := range g.nodes[current] {
            if !visited[neighbor] {
                visited[neighbor] = true
                parent[neighbor] = current
                queue = append(queue, neighbor)
            }
        }
    }
    
    return nil
}
```

**6. Map of functions**

```go
// Strategy pattern ด้วย map
type Operation func(a, b int) int

var operations = map[string]Operation{
    "add":      func(a, b int) int { return a + b },
    "subtract": func(a, b int) int { return a - b },
    "multiply": func(a, b int) int { return a * b },
    "divide":   func(a, b int) int { return a / b },
}

func calculate(op string, a, b int) (int, error) {
    if fn, ok := operations[op]; ok {
        return fn(a, b), nil
    }
    return 0, fmt.Errorf("unknown operation: %s", op)
}
```

---

## 4. การจัดการข้อผิดพลาด (บทที่ 23)

**เป้าหมาย:** เขียนโค้ดที่จัดการ error ได้อย่างถูกต้องและมีประสิทธิภาพ

### บทที่ 23: Errors

#### Error Interface

```go
// error interface (built-in)
type error interface {
    Error() string
}
```

#### ตัวอย่างการใช้งาน

```go
package main

import (
    "errors"
    "fmt"
)

// 1. สร้าง error ด้วย errors.New
var ErrNotFound = errors.New("not found")
var ErrInvalidInput = errors.New("invalid input")

// 2. สร้าง error ด้วย fmt.Errorf
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero: %f / %f", a, b)
    }
    return a / b, nil
}

// 3. Custom error type
type ValidationError struct {
    Field   string
    Value   interface{}
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation failed for %s (value: %v): %s", 
        e.Field, e.Value, e.Message)
}

// 4. ฟังก์ชันที่ใช้ custom error
func validateUser(user User) error {
    if user.Name == "" {
        return &ValidationError{
            Field:   "name",
            Value:   user.Name,
            Message: "name is required",
        }
    }
    if user.Age < 0 {
        return &ValidationError{
            Field:   "age",
            Value:   user.Age,
            Message: "age must be positive",
        }
    }
    return nil
}
```

#### Error Wrapping (Go 1.13+)

```go
package main

import (
    "errors"
    "fmt"
)

// ฟังก์ชันที่ wrap error
func readConfig(path string) error {
    data, err := os.ReadFile(path)
    if err != nil {
        // Wrap error ด้วย context
        return fmt.Errorf("read config file %s: %w", path, err)
    }
    // process data...
    return nil
}

func process() error {
    if err := readConfig("/etc/app/config.yaml"); err != nil {
        // Wrap เพิ่ม layer
        return fmt.Errorf("process failed: %w", err)
    }
    return nil
}

func main() {
    err := process()
    if err != nil {
        // ตรวจสอบด้วย errors.Is
        if errors.Is(err, os.ErrNotExist) {
            fmt.Println("config file does not exist")
        }
        
        // ตรวจสอบ type ด้วย errors.As
        var pathErr *os.PathError
        if errors.As(err, &pathErr) {
            fmt.Printf("path error: %s\n", pathErr.Path)
        }
        
        // ดู error chain
        fmt.Printf("full error: %+v\n", err)
    }
}
```

#### ตัวอย่างการใช้งานจริง

**1. การจัดการ error ใน HTTP handler**

```go
// domain errors
var (
    ErrUserNotFound     = errors.New("user not found")
    ErrInvalidPassword  = errors.New("invalid password")
    ErrEmailExists      = errors.New("email already exists")
)

// service layer
type UserService struct {
    repo UserRepository
}

func (s *UserService) Login(email, password string) (*User, error) {
    user, err := s.repo.FindByEmail(email)
    if err != nil {
        if errors.Is(err, ErrUserNotFound) {
            return nil, fmt.Errorf("login failed: %w", ErrUserNotFound)
        }
        return nil, fmt.Errorf("database error: %w", err)
    }
    
    if !user.VerifyPassword(password) {
        return nil, fmt.Errorf("login failed: %w", ErrInvalidPassword)
    }
    
    return user, nil
}

// handler layer
type UserHandler struct {
    service UserService
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
    var req LoginRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "invalid request", http.StatusBadRequest)
        return
    }
    
    user, err := h.service.Login(req.Email, req.Password)
    if err != nil {
        // แยก error type
        switch {
        case errors.Is(err, ErrUserNotFound):
            http.Error(w, "user not found", http.StatusNotFound)
        case errors.Is(err, ErrInvalidPassword):
            http.Error(w, "invalid password", http.StatusUnauthorized)
        default:
            // log internal error
            log.Printf("login error: %v", err)
            http.Error(w, "internal server error", http.StatusInternalServerError)
        }
        return
    }
    
    json.NewEncoder(w).Encode(user)
}
```

**2. Retry pattern with error handling**

```go
func Retry(maxRetries int, fn func() error) error {
    var err error
    
    for i := 0; i < maxRetries; i++ {
        err = fn()
        if err == nil {
            return nil
        }
        
        // ตรวจสอบว่า error ที่เกิดขึ้น retry ได้หรือไม่
        var retryable *RetryableError
        if errors.As(err, &retryable) {
            delay := time.Duration(1<<i) * time.Second
            log.Printf("retry %d after %v: %v", i+1, delay, err)
            time.Sleep(delay)
            continue
        }
        
        // error ไม่สามารถ retry ได้
        return err
    }
    
    return fmt.Errorf("max retries exceeded: %w", err)
}

type RetryableError struct {
    Err error
}

func (e *RetryableError) Error() string {
    return fmt.Sprintf("retryable: %v", e.Err)
}

func (e *RetryableError) Unwrap() error {
    return e.Err
}

// ใช้งาน
err := Retry(3, func() error {
    resp, err := http.Get("https://api.example.com/data")
    if err != nil {
        return &RetryableError{Err: err}  // network error, retry
    }
    defer resp.Body.Close()
    
    if resp.StatusCode != http.StatusOK {
        if resp.StatusCode >= 500 {
            return &RetryableError{Err: fmt.Errorf("server error: %d", resp.StatusCode)}
        }
        return fmt.Errorf("client error: %d", resp.StatusCode)  // ไม่ retry
    }
    
    return nil
})
```

**3. Error grouping and handling**

```go
type ErrorGroup struct {
    errors []error
    mu     sync.Mutex
}

func (g *ErrorGroup) Add(err error) {
    g.mu.Lock()
    defer g.mu.Unlock()
    g.errors = append(g.errors, err)
}

func (g *ErrorGroup) Error() string {
    if len(g.errors) == 0 {
        return ""
    }
    
    messages := make([]string, len(g.errors))
    for i, err := range g.errors {
        messages[i] = err.Error()
    }
    return strings.Join(messages, "; ")
}

func (g *ErrorGroup) HasErrors() bool {
    g.mu.Lock()
    defer g.mu.Unlock()
    return len(g.errors) > 0
}

// ใช้งาน
func ValidateUser(user User) error {
    var errGroup ErrorGroup
    
    if user.Name == "" {
        errGroup.Add(errors.New("name is required"))
    }
    
    if !strings.Contains(user.Email, "@") {
        errGroup.Add(errors.New("invalid email"))
    }
    
    if user.Age < 18 {
        errGroup.Add(errors.New("age must be at least 18"))
    }
    
    if errGroup.HasErrors() {
        return &errGroup
    }
    
    return nil
}
```

**4. Panic recovery (เมื่อจำเป็น)**

```go
// สำหรับ service ที่ต้องรันตลอดเวลา
func RunWithRecovery(fn func()) {
    defer func() {
        if r := recover(); r != nil {
            // log panic
            log.Printf("panic recovered: %v\n%s", r, debug.Stack())
            
            // send alert
            sendAlert(fmt.Sprintf("panic: %v", r))
        }
    }()
    
    fn()
}

// ใช้กับ goroutine
go func() {
    RunWithRecovery(func() {
        // risky operation
        doSomething()
    })
}()
```

**5. Best Practices สรุป**

```go
// ✅ DO: ตรวจสอบ error ทันที
data, err := readFile()
if err != nil {
    return fmt.Errorf("read file: %w", err)
}

// ❌ DON'T: ignore errors
data, _ := readFile()

// ✅ DO: ให้ context แก่ error
return fmt.Errorf("failed to process user %d: %w", userID, err)

// ❌ DON'T: return plain error without context
return err

// ✅ DO: ใช้ custom error types สำหรับ business logic
var ErrInsufficientFunds = errors.New("insufficient funds")

// ✅ DO: ใช้ errors.Is และ errors.As ในการตรวจสอบ
if errors.Is(err, ErrNotFound) {
    // handle not found
}

var valErr *ValidationError
if errors.As(err, &valErr) {
    fmt.Printf("field: %s\n", valErr.Field)
}
```

---

## สรุปภาคที่ 3: ความสัมพันธ์ของเนื้อหา

```mermaid
flowchart LR
    subgraph Skills["ทักษะที่ได้"]
        S1[สร้างและจัดการ<br/>โปรเจกต์ Go]
        S2[เขียน Unit Tests<br/>และวัด coverage]
        S3[ใช้ Slices และ Maps<br/>จัดการข้อมูล]
        S4[จัดการ error<br/>อย่างมืออาชีพ]
    end
    
    subgraph Applications["การประยุกต์ใช้"]
        A1[โปรเจกต์ที่<br/>มีโครงสร้าง]
        A2[โค้ดที่มั่นใจ<br/>ผ่านการทดสอบ]
        A3[โครงสร้างข้อมูล<br/>ที่มีประสิทธิภาพ]
        A4[ระบบที่เสถียร<br/>debug ได้ง่าย]
    end
    
    Skills --> Applications
```

**สิ่งที่ได้เรียนรู้**

1. **บทที่ 17-18 (Go Modules)**: 
   - จัดการ dependencies ด้วย go.mod
   - ใช้ module proxy เพื่อความเร็วและความปลอดภัย
   - จัดการ version และ private modules

2. **บทที่ 19 (Testing)**:
   - เขียน table-driven tests
   - ใช้ mock และ stub
   - วัด test coverage

3. **บทที่ 20-22 (Data Structures)**:
   - Array: fixed-size, value type
   - Slice: dynamic, reference type, ใช้บ่อยที่สุด
   - Map: key-value, hash table

4. **บทที่ 23 (Error Handling)**:
   - error interface และ error wrapping
   - custom error types
   - errors.Is และ errors.As
   - retry pattern และ error grouping

พื้นฐานเหล่านี้จะนำไปใช้ใน **ภาคที่ 4** เพื่อพัฒนาแอปพลิเคชันเชิงปฏิบัติ เช่น HTTP server, concurrency, JSON handling, และการบันทึกข้อมูล


## ภาคที่ 5: สู่การเป็นนักพัฒนา Go มืออาชีพ (บทที่ 34–42)
 ```mermaid
flowchart TD
    subgraph Performance[การวัดและปรับปรุงประสิทธิภาพ]
        A[บทที่ 34: Benchmark] --> B[go test -bench]
        B --> C[บทที่ 36: Profiling<br/>pprof]
        C --> D[ระบุจุดช้า]
        D --> E[บทที่ 38: Generics<br/>ลด duplication]
        E --> F[บทที่ 37: Context<br/>จัดการ timeout/cancel]
        F --> G[บทที่ 35: HTTP Client<br/>ปรับแต่ง connection pool]
    end

    subgraph CodeDesign[การออกแบบโค้ด]
        H[บทที่ 39: OOP ใน Go<br/>struct + interface + composition]
        I[บทที่ 41: คำแนะนำ<br/>idioms, naming, error handling]
        J[บทที่ 40: การอัปเกรด/ดาวน์เกรด<br/>Go version]
    end

    subgraph Reference[เอกสารอ้างอิง]
        K[บทที่ 42: Cheatsheet]
    end

    Performance --> CodeDesign
    CodeDesign --> Reference
```

**คำอธิบายภาษาไทย (แบบละเอียด):**

ภาคที่ 5 (บทที่ 34–42) มุ่งเน้นยกระดับผู้พัฒนาไปสู่ระดับมืออาชีพ ครอบคลุมการวัดประสิทธิภาพ, การปรับแต่งโค้ด, การใช้เครื่องมือขั้นสูง, และแนวทางการออกแบบที่ดี โดยมีลำดับการเรียนรู้ดังนี้

### 1. การวัดและปรับปรุงประสิทธิภาพ (Performance)

- **บทที่ 34: Benchmark**  
  ผู้เขียน Benchmark ในไฟล์ `_test.go` โดยใช้ฟังก์ชัน `BenchmarkXxx(b *testing.B)` และรันด้วย `go test -bench`. ช่วยให้รู้ว่าโค้ดส่วนไหนทำงานช้าหรือใช้หน่วยความจำมาก

- **บทที่ 36: Profiling**  
  ใช้ `pprof` (built‑in) เพื่อเก็บและวิเคราะห์ข้อมูล CPU, memory, goroutine ผ่าน `go tool pprof` หรือเปิด endpoint `http://localhost:6060/debug/pprof` เมื่อใช้ `net/http/pprof`. ผล profiling ช่วยระบุจุด bottleneck ได้แม่นยำ

- **บทที่ 38: Generics**  
  ตั้งแต่ Go 1.18 เป็นต้นมา Generics ช่วยลดการเขียนโค้ดซ้ำซ้อนสำหรับชนิดข้อมูลที่คล้ายกัน โดยใช้ type parameters เช่น `func Map[T any](s []T, f func(T) T) []T`. ทำให้โค้ดปลอดภัยและลด dependency กับ interface{}

- **บทที่ 37: Context**  
  `context.Context` ใช้ในการส่งค่า deadline, cancellation, และ request‑scoped values ผ่านขอบเขตของ API เป็นมาตรฐานสำหรับการจัดการ timeout และการยกเลิกงาน (โดยเฉพาะใน HTTP server และ goroutine) ช่วยให้ระบบตอบสนองได้ดีและไม่เกิด goroutine ค้าง

- **บทที่ 35: HTTP Client**  
  การสร้าง HTTP client แบบมืออาชีพต้องปรับแต่ง `http.Transport` เช่น กำหนด `MaxIdleConns`, `IdleConnTimeout`, และ `Timeout` เพื่อเพิ่มประสิทธิภาพการเรียก API ภายนอก

### 2. การออกแบบโค้ด (Code Design)

- **บทที่ 39: OOP ใน Go**  
  Go ไม่มี class แต่ใช้ **struct + method** และ **interface** เพื่อสร้าง abstraction การสืบทอดทำผ่าน **composition** (embedding) แทน inheritance แนวทางนี้ลด coupling และเพิ่มความยืดหยุ่น

- **บทที่ 41: คำแนะนำในการออกแบบโค้ดที่ดี**  
  รวบรวม best practices เช่น:
  - **Keep it simple**: หลีกเลี่ยงความซับซ้อนเกินจำเป็น
  - **Accept interfaces, return structs**: ฟังก์ชันควรรับ interface, คืน concrete type
  - **Use zero values**: ใช้ zero value ให้เป็นประโยชน์
  - **Error handling**: ตรวจสอบ error ทันที, ห่อ error ด้วยข้อมูลเพิ่มเติม
  - **Naming**: ชื่อแพคเกจสั้น, ชื่อฟังก์ชันและตัวแปรสื่อความหมาย

- **บทที่ 40: การอัปเกรดหรือดาวน์เกรดเวอร์ชัน Go**  
  การจัดการเวอร์ชัน Go ในโปรเจกต์: ใช้ `go mod edit -go=1.x`, อัปเดต `go.mod`, ใช้ `go fix` สำหรับโค้ดที่ต้องปรับให้เข้ากับเวอร์ชันใหม่, หรือใช้ `gvm` (Go Version Manager) สำหรับการสลับเวอร์ชันในการพัฒนา

### 3. เอกสารอ้างอิง (Reference)

- **บทที่ 42: Cheatsheet**  
  สรุปคำสั่ง Go ที่ใช้บ่อย, ชนิดข้อมูล, การประกาศตัวแปร, โครงสร้างควบคุม, slice, map, channel, การทำงานกับไฟล์, JSON, HTTP และ context เป็นตารางอ้างอิงที่ช่วยให้พัฒนาได้รวดเร็ว

---

### แผนภาพสรุป

แผนภาพด้านบนแสดงลำดับการทำงาน:
1. เริ่มจาก **Benchmark** → **Profiling** → ระบุจุดช้า → แก้ด้วย **Generics** หรือ **Context** และปรับแต่ง **HTTP Client**
2. หลังจากปรับปรุงประสิทธิภาพแล้ว ศึกษาการออกแบบโค้ดให้สอดคล้องกับแนวทางของ Go (OOP แบบ Go, idioms, version management)
3. ปิดท้ายด้วย **Cheatsheet** สำหรับการทบทวนและใช้งานจริง

การเรียนรู้ในภาคนี้จะเปลี่ยนนักพัฒนาจากผู้ใช้ภาษาให้เป็นผู้ที่เข้าใจกลไกภายใน สามารถวิเคราะห์และปรับปรุงระบบให้มีประสิทธิภาพสูง พร้อมเขียนโค้ดที่บำรุงรักษาง่ายและยืดหยุ่นตามความต้องการของธุรกิจ
---

## ภาคที่ 6: เครื่องมือและไลบรารียอดนิยม (บทที่ 43–45)

**แผนภาพ: Router → Config → CLI → Logger → ORM → Email**

```mermaid
graph LR
    subgraph Router[เราเตอร์]
        R1[chi] --> R2[มิดเดิลแวร์]
    end

    subgraph Config[การตั้งค่า]
        V1[viper] --> V2[config.yaml]
        V2 --> R1
    end

    subgraph CLI[บรรทัดคำสั่ง]
        C1[cobra] --> C2[คำสั่ง]
        C2 --> V1
    end

    subgraph Logger[การบันทึก]
        Z1[zap] --> Z2[บันทึกแบบมีโครงสร้าง]
    end

    subgraph ORM[ORM]
        G1[GORM] --> G2[การจัดการฐานข้อมูล]
    end

    subgraph Email[อีเมล]
        M1[gomail] --> M2[hermes]
        M2 --> M3[HTML Email]
    end

    R2 --> Z1
    R2 --> G1
    G1 --> M1
```

**คำอธิบาย:**  
การกำหนดค่า (viper) จะถูกโหลดจากไฟล์ `config.yaml` และใช้ใน router (chi) ส่วน cobra ใช้สร้าง CLI ที่อาจเรียกใช้ viper หรือ chi ได้ ระหว่างการทำงาน router จะเรียก logger (zap) และ ORM (GORM) ซึ่ง GORM สามารถเชื่อมต่อกับ gomail เพื่อส่งอีเมลผ่าน hermes ที่สร้าง HTML template

---

## ภาคที่ 7: การออกแบบสถาปัตยกรรมและ Workflow (บทที่ 46–48)
เราจะนำเสนอแผนภาพใหม่สำหรับภาคที่ 7 (บทที่ 46–48) พร้อมคำอธิบายภาษาไทยแบบละเอียด โดยเน้นให้เห็นภาพรวมของสถาปัตยกรรม Clean Architecture, Blueprint สำหรับโปรเจกต์ระดับ Production, และการจัดการ Workflow รวมถึง Task Management

---

## ภาคที่ 7: การออกแบบสถาปัตยกรรมและ Workflow (บทที่ 46–48)

### แผนภาพหลัก: Clean Architecture + Production Blueprint + Workflow

```mermaid
flowchart TB
    subgraph Delivery[Delivery Layer]
        D1[HTTP Handlers<br/>chi / gin]
        D2[Middleware<br/>Auth, Logger, CORS]
        D3[DTOs & Validation]
    end

    subgraph Usecase[Usecase Layer]
        U1[Business Logic<br/>Use Cases]
        U2[Domain Services]
        U3[Repository Interfaces]
    end

    subgraph Repository[Repository Layer]
        R1[Repository Implementations<br/>GORM / SQL]
        R2[Cache<br/>Redis]
        R3[External API Clients]
    end

    subgraph Infrastructure[Infrastructure Layer]
        I1[Database<br/>PostgreSQL]
        I2[Cache Store<br/>Redis]
        I3[Message Broker<br/>RabbitMQ]
    end

    subgraph Blueprint[Production Blueprint]
        B1[cmd/ / internal/ / pkg/]
        B2[configs/ / migrations/]
        B3[deploy/ docker/ k8s]
        B4[scripts/ / test/]
    end

    subgraph Workflow[Workflow & Task Management]
        W1[Task List Template]
        W2[Checklist Template]
        W3[Development Workflow]
    end

    Delivery --> Usecase
    Usecase --> Repository
    Repository --> Infrastructure

    Blueprint --> Delivery
    Blueprint --> Usecase
    Blueprint --> Repository
    Blueprint --> Infrastructure

    Workflow --> Blueprint
    Workflow --> Delivery
```

---

### คำอธิบายภาษาไทย (แบบละเอียด)

ภาคที่ 7 ประกอบด้วย 3 บทหลัก ได้แก่:

- **บทที่ 46: Clean Architecture และโครงสร้างโปรเจกต์**  
- **บทที่ 47: Blueprint สำหรับโปรเจกต์ Go ระดับ Production**  
- **บทที่ 48: การออกแบบ Workflow และ Task Management**

ทั้งสามบทนี้เชื่อมโยงกันเพื่อให้ผู้อ่านสามารถนำไปสร้างแอปพลิเคชันที่มีโครงสร้างแข็งแรง พร้อมขยาย และมีกระบวนการพัฒนาที่มีระเบียบ

#### 1. Clean Architecture (บทที่ 46)

สถาปัตยกรรม Clean Architecture แบ่งชั้นอย่างชัดเจนเพื่อแยกความรับผิดชอบ (Separation of Concerns) ทำให้โค้ดทดสอบง่าย เปลี่ยนแปลงโครงสร้างพื้นฐานได้โดยไม่กระทบตรรกะธุรกิจ

**ชั้น Delivery**  
- รับ HTTP Request ผ่าน router (chi)  
- ใช้ middleware จัดการเรื่อง logging, authentication, CORS, panic recovery  
- แปลง request body เป็น DTO และตรวจสอบข้อมูล (validation)  
- ส่ง DTO ไปยัง Usecase และแปลงผลลัพธ์กลับเป็น HTTP Response

**ชั้น Usecase**  
- เก็บตรรกะธุรกิจหลัก (business logic) ของแอปพลิเคชัน  
- แต่ละ use case (เช่น RegisterUser, PlaceOrder) จะเรียกใช้ repository interface เพื่อดึงหรือบันทึกข้อมูล  
- ใช้ domain services สำหรับกฎที่ซับซ้อน  
- สนับสนุนการส่ง domain events (ถ้ามี)

**ชั้น Repository**  
- กำหนด interface สำหรับการเข้าถึงข้อมูล (ไม่ขึ้นกับเทคโนโลยี)  
- การ implement (GORM, SQL, Redis) อยู่ในชั้น infrastructure แต่ถูก inject ผ่าน interface  
- ช่วยให้สามารถเปลี่ยนฐานข้อมูลหรือ cache ได้โดยไม่ต้องแก้ use case

**ชั้น Infrastructure**  
- เป็นการ implement จริงของ repository (GORM, Redis client, HTTP client)  
- จัดการ connection pool, migration, และการเชื่อมต่อกับระบบภายนอก

**ข้อดีของ Clean Architecture**  
- ทดสอบ unit ได้ง่าย (mock repository)  
- เปลี่ยนฐานข้อมูลหรือ cache ได้โดยไม่กระทบ business logic  
- โค้ดอ่านง่ายและบำรุงรักษา

---

#### 2. Blueprint สำหรับโปรเจกต์ Go ระดับ Production (บทที่ 47)

Blueprint นี้เป็นโครงสร้างโฟลเดอร์และไฟล์ที่ผ่านการพิสูจน์แล้วจากโปรเจกต์จริง โดยแยกตามหลัก Clean Architecture และเพิ่มส่วนประกอบที่จำเป็นสำหรับการผลิตจริง

**โครงสร้างหลัก**
- **cmd/** : จุดเริ่มต้นของแอปพลิเคชัน (main.go)  
- **internal/** : โค้ดเฉพาะแอปพลิเคชัน ไม่ให้ import จากภายนอก  
  - **core/** : domain models, repositories interfaces, services  
  - **application/** : use cases, DTOs  
  - **infrastructure/** : การ implement repository, cache, queue, logger  
  - **interfaces/** : HTTP handlers, middleware, routes  
- **pkg/** : โค้ดที่ reusable (jwt, mail, redis client)  
- **configs/** : ไฟล์คอนฟิก (YAML)  
- **migrations/** : SQL migration scripts  
- **deploy/** : Dockerfile, docker-compose, Kubernetes manifests  
- **scripts/** : สคริปต์ช่วย (migrate, seed)  
- **test/** : integration tests

**ฟีเจอร์พร้อมใช้ใน blueprint**  
- Authentication (JWT access/refresh token) เก็บ refresh token ใน Redis  
- Rate limiting ตาม IP  
- Health check endpoints (/health, /ready, /live)  
- Caching ด้วย Redis  
- Message queue (Redis pub/sub หรือ RabbitMQ)  
- Transaction management  
- Structured logging (slog หรือ zap)  
- Graceful shutdown

---

#### 3. การออกแบบ Workflow และ Task Management (บทที่ 48)

บทนี้มุ่งเน้นกระบวนการพัฒนาที่เป็นระบบ โดยมีเครื่องมือช่วยในการวางแผนและควบคุมคุณภาพ

**Workflow การพัฒนา Feature ใหม่**
1. **Analyze** – ทำความเข้าใจ requirement, domain models, use cases  
2. **Design** – ออกแบบ entities, value objects, repository interface, service interface  
3. **Implement Domain** – เขียน entity, repository interface, service interface ใน `internal/core/`  
4. **Implement Infrastructure** – เขียน repository implementation (GORM), cache, queue  
5. **Implement Service** – เขียน business logic ใน service implementation  
6. **Implement Handler** – เขียน HTTP handlers, ตรวจสอบ input validation, เรียก service  
7. **Register Routes** – ลงทะเบียน routes ใน router  
8. **Test** – unit tests (domain, service), integration tests (handler)  
9. **Document** – อัปเดต API docs, README

**Task List Template**  
เป็นเทมเพลตสำหรับการติดตามงานในแต่ละ phase ประกอบด้วยรายการย่อยที่ต้องทำ (checklist) เช่น  
- Phase 1: Domain Design (ระบุ entity, กำหนด invariants, ออกแบบ repository interface)  
- Phase 2: Implementation (สร้าง entity struct, เขียน unit tests)  
- Phase 3: Infrastructure (สร้าง repository implementation, migration)  
- Phase 4: Delivery (สร้าง HTTP handlers, routes, dependency injection)  
- Phase 5: Integration (end-to-end test, linter)  
- Phase 6: Review & Deploy (code review, coverage, race detector, deploy)

**Checklist Template**  
เป็นรายการตรวจสอบคุณภาพด้านต่าง ๆ เช่น  
- **Code Quality**: comment, go fmt, error handling, no panic in libs  
- **Security**: bcrypt for password, JWT secret from env, CORS, input validation  
- **Performance**: indexes, caching, connection pools, avoid N+1 queries  
- **Testing**: unit test coverage >80%, race detector passed, mock external dependencies  
- **Deployment**: graceful shutdown, health checks, logging to stdout, docker multi-stage

**ประโยชน์**  
- ลดความเสี่ยงในการลืมขั้นตอนสำคัญ  
- สร้างมาตรฐานเดียวกันในทีม  
- ช่วยให้การ code review มีประสิทธิภาพ  
- ทำให้การ deploy ราบรื่น

---

### แผนภาพเสริม: ความสัมพันธ์ระหว่างบท

```mermaid
graph LR
    subgraph Ch46[บทที่ 46: Clean Architecture]
        A[Delivery] --> B[Usecase]
        B --> C[Repository]
        C --> D[Infrastructure]
    end

    subgraph Ch47[บทที่ 47: Production Blueprint]
        E[cmd/ internal/ pkg/]
        F[configs/ migrations/]
        G[deploy/ scripts/]
    end

    subgraph Ch48[บทที่ 48: Workflow & Task Management]
        H[Task List Template]
        I[Checklist Template]
        J[Development Workflow]
    end

    Ch46 -->|นำไปใช้| Ch47
    Ch47 -->|กำหนดโครงสร้าง| Ch48
    Ch48 -->|ควบคุมกระบวนการ| Ch46
```

**คำอธิบายเพิ่มเติม:**  
- Clean Architecture (บท 46) ให้หลักการออกแบบ  
- Production Blueprint (บท 47) นำหลักการนั้นมาเป็นโครงสร้างไฟล์และส่วนประกอบสำเร็จรูป  
- Workflow & Task Management (บท 48) ใช้โครงสร้างนั้นเพื่อกำหนดกระบวนการพัฒนาและการตรวจสอบคุณภาพ

---

แผนภาพและคำอธิบายนี้จะช่วยให้ผู้อ่านเห็นภาพรวมของภาคที่ 7 ได้อย่างชัดเจน สามารถนำไปปรับใช้กับโปรเจกต์จริงได้ทันที
---

## ภาคที่ 8: Domain-Driven Design (DDD) (บทที่ 49–51)

**แผนภาพ: Domain → Application → Infrastructure**

```mermaid
graph LR
    subgraph Domain[ชั้นโดเมน]
        A[Entities & Value Objects] --> B[Aggregates]
        B --> C[Domain Events]
        C --> D[Repository Interface]
    end

    subgraph Application[ชั้นแอปพลิเคชัน]
        E[Use Cases] --> F[DTOs]
        F --> G[Service Orchestration]
    end

    subgraph Infrastructure[ชั้นโครงสร้างพื้นฐาน]
        H[Repository Implementation] --> I[Event Bus]
        I --> J[Message Broker]
    end

    A --> E
    D --> H
    C --> I
```

**คำอธิบาย:**  
ชั้นโดเมน (Entities, Value Objects, Aggregates) กำหนดพฤติกรรมทางธุรกิจและสร้าง Domain Events ชั้นแอปพลิเคชันใช้ Use Cases ในการจัดลำดับการทำงาน โดยรับ DTO และเรียกใช้ Service Orchestration ชั้นโครงสร้างพื้นฐาน Implement Repository และ Event Bus เพื่อติดต่อกับฐานข้อมูลหรือ Message Broker

---

## ภาคที่ 9: การผสานระบบภายนอก (บทที่ 52–58)

**แผนภาพ: Redis, RabbitMQ, MQTT, InfluxDB, WebSocket, Notifications**
 ```mermaid
flowchart TB
    subgraph DataSources[แหล่งข้อมูลภายนอก]
        direction LR
        IoT[IoT Sensors] -->|MQTT| MQTT
        User[ผู้ใช้] -->|HTTP Request| API
    end

    subgraph Core[ระบบหลัก Go]
        API[HTTP API / WebSocket]
        Redis[(Redis)]
        Rabbit[(RabbitMQ)]
        Influx[(InfluxDB)]
    end

    subgraph Processing[การประมวลผล]
        Workers[Worker Pool]
        Cache[Cache Layer]
        Stream[Event Stream]
    end

    subgraph Notifications[การแจ้งเตือน]
        SMS[SMS Gateway]
        LINE[LINE Notify]
        Discord[Discord Webhook]
    end

    subgraph Clients[ไคลเอนต์]
        Browser[เว็บเบราว์เซอร์]
        Mobile[แอปมือถือ]
    end

    %% การไหลของข้อมูล
    MQTT -->|Telemetry| Influx
    API -->|Read/Write| Redis
    API -->|Publish Event| Rabbit
    Rabbit -->|Consume| Workers
    Workers -->|Store Metrics| Influx
    Workers -->|Update Cache| Redis
    Influx -->|Query| API
    Redis -->|Cached Data| API

    %% Real-time
    API -->|WebSocket| Browser
    API -->|WebSocket| Mobile

    %% Notifications trigger
    Workers -->|Alert| SMS
    Workers -->|Alert| LINE
    Workers -->|Alert| Discord
    Influx -->|Threshold Exceeded| Workers
```

**คำอธิบายภาษาไทย:**

แผนภาพนี้แสดงการทำงานร่วมกันของระบบภายนอกทั้ง 7 ชนิดที่กล่าวถึงในบทที่ 52–58 โดยแบ่งเป็น 5 กลุ่มหลัก:

1. **แหล่งข้อมูลภายนอก**  
   - **IoT Sensors**: ส่งข้อมูลผ่าน MQTT ไปยัง InfluxDB โดยตรง  
   - **ผู้ใช้**: ส่ง HTTP request ผ่าน API ที่เขียนด้วย Go

2. **ระบบหลัก Go**  
   - **HTTP API / WebSocket**: รับคำขอจากผู้ใช้และให้บริการ real‑time ผ่าน WebSocket  
   - **Redis**: ใช้เป็น cache (เก็บข้อมูลที่เรียกบ่อย) และ message queue (Pub/Sub)  
   - **RabbitMQ**: รับ event จาก API และกระจายไปยัง workers  
   - **InfluxDB**: เก็บข้อมูลอนุกรมเวลาจาก IoT และผลลัพธ์จากการประมวลผล

3. **การประมวลผล**  
   - **Worker Pool**: ดึงงานจาก RabbitMQ, ประมวลผล (เช่น คำนวณค่าเฉลี่ย, ตรวจจับเหตุการณ์), จากนั้นอัปเดต cache และ InfluxDB  
   - **Cache Layer**: ข้อมูลที่ถูกเรียกบ่อยจะถูกเก็บใน Redis เพื่อลดภาระฐานข้อมูล  
   - **Event Stream**: ใช้ Redis Pub/Sub หรือ RabbitMQ เพื่อส่ง event ระหว่าง component

4. **การแจ้งเตือน**  
   - เมื่อ worker ตรวจพบเงื่อนไข (เช่น ค่าเกินเกณฑ์) จะส่งการแจ้งเตือนผ่าน **SMS**, **LINE Notify**, และ **Discord Webhook** ตามช่องทางที่กำหนด

5. **ไคลเอนต์**  
   - เบราว์เซอร์และแอปมือถือรับข้อมูล real‑time ผ่าน WebSocket และแสดงผลพร้อมรับการแจ้งเตือน

**การไหลของข้อมูลโดยสรุป:**
- ข้อมูลจาก IoT → InfluxDB (จัดเก็บ)
- ผู้ใช้ → API → Redis (cache) / RabbitMQ (event) → Workers → InfluxDB / Redis
- InfluxDB → API → WebSocket → ไคลเอนต์ (แสดงผล)
- Workers → SMS/LINE/Discord (แจ้งเตือน)

โครงสร้างนี้ช่วยให้ระบบมีความยืดหยุ่น รองรับการขยายขนาด และสามารถตอบสนองต่อเหตุการณ์แบบเรียลไทม์ได้อย่างมีประสิทธิภาพ
---

## ภาคที่ 10: เทมเพลต กระบวนการพัฒนา และตัวอย่างโค้ด (บทที่ 59–63)

**แผนภาพ: Example Application → Tasks → Checklist → Diagrams → Config**

```mermaid
graph LR
    subgraph Example[ตัวอย่าง]
        E1[Full‑stack Example] --> E2[โครงสร้างโค้ด]
    end

    subgraph Processes[กระบวนการ]
        P1[Task List] --> P2[Checklist]
        P2 --> P3[Workflow Diagram]
    end

    subgraph Config[การตั้งค่า]
        C1[mop Config] --> C2[YAML/ENV]
        C2 --> E1
    end

    E2 --> P1
    P3 --> C1
```

**คำอธิบาย:**  
จากตัวอย่างโปรเจกต์ครบวงจร (Full‑stack Example) จะได้โครงสร้างโค้ดที่สามารถนำไปใช้เป็นต้นแบบ โครงสร้างนี้ช่วยให้ทีมสร้าง Task List และ Checklist เพื่อติดตามงาน พร้อมทั้ง Workflow Diagram ที่อธิบายขั้นตอนการทำงาน สุดท้าย mop Config จัดการค่า configuration (YAML/ENV) เพื่อให้แอปพลิเคชันทำงานได้ในหลายสภาพแวดล้อม

---

แผนภาพเหล่านี้สามารถนำไปแทรกในหนังสือหรือเอกสารการสอน เพื่อช่วยให้ผู้อ่านเห็นภาพรวมและลำดับการไหลของข้อมูลในแต่ละหัวข้อได้ดียิ่งขึ้น



## 📚 บทสรุป

คู่มือนี้ครอบคลุมเนื้อหาตั้งแต่พื้นฐานภาษา Go ไปจนถึงการออกแบบสถาปัตยกรรมระดับองค์กรและการผสานระบบภายนอกที่ใช้ในโลกแห่งความจริง โดยมุ่งเน้นให้ผู้อ่านสามารถนำไปประยุกต์ใช้ได้ทันที

**หวังว่าคู่มือนี้จะเป็นประโยชน์สำหรับผู้ที่ต้องการเริ่มต้นและพัฒนาทักษะการเขียนโปรแกรมด้วย Go อย่างจริงจัง ขอให้สนุกกับการเขียนโปรแกรม!**

---

**ผู้เขียน:** คงนคร จันทะคุณ  
**อีเมล:** kongnakornjantakun@gmail.com  
**วันที่:** เมษายน 2026