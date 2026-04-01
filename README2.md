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

**สิ่งที่ได้เรียนรู้**
1. **บทที่ 6-7**: รู้ว่าคอมพิวเตอร์จัดเก็บข้อมูลอย่างไร (Binary, Hex, UTF-8, Rune)
2. **บทที่ 8-9**: ประกาศตัวแปรและควบคุมการทำงานด้วย if, for, switch
3. **บทที่ 10-12**: จัดระเบียบโค้ดด้วยฟังก์ชันและแพคเกจ
4. **บทที่ 13-16**: สร้าง struct, กำหนดเมธอด, ใช้ pointer, และออกแบบ interface

พื้นฐานเหล่านี้จะนำไปใช้ใน **ภาคที่ 3** เพื่อจัดการโปรเจกต์, สร้าง test, และใช้ slice/map ในการจัดการข้อมูลจำนวนมาก

# ภาคที่ 4: การพัฒนาแอปพลิเคชันเชิงปฏิบัติ (บทที่ 24–33)

## แผนภาพแสดงความสัมพันธ์ของเนื้อหา

```mermaid
flowchart TB
    subgraph Input["1. การรับข้อมูล (Input Layer)"]
        direction TB
        A1["HTTP Server (บทที่ 26)<br/>net/http, chi, gin<br/>routing, middleware, handlers"]
        A2["JSON/XML Handling (บทที่ 25)<br/>encoding/json, xml<br/>Marshal/Unmarshal, tags"]
        A3["Enum/Iota/Bitmask (บทที่ 27)<br/>iota, bit operations<br/>status flags, permissions"]
        A4["Date & Time (บทที่ 28)<br/>time package, duration<br/>parsing, formatting, timezone"]
        
        A1 --> A2
        A2 --> A3
        A3 --> A4
    end

    subgraph Processing["2. การประมวลผล (Processing Layer)"]
        direction TB
        B1["Anonymous Functions & Closure (บทที่ 24)<br/>func literals, closures<br/>capturing variables, callbacks"]
        B2["Concurrency (บทที่ 30)<br/>goroutines, channels, select<br/>sync, worker pools, patterns"]
        B3["Error Handling (from บทที่ 23)<br/>error wrapping, custom errors<br/>panic recovery, retry"]
        
        B1 --> B2
        B2 --> B3
    end

    subgraph Storage["3. การจัดเก็บ (Storage Layer)"]
        direction TB
        C1["Files & Database (บทที่ 29)<br/>os, io, bufio<br/>SQL, migrations, transactions"]
        C2["Logging (บทที่ 31)<br/>slog, zap, logrus<br/>structured logging, levels"]
        
        C1 --> C2
    end

    subgraph Output["4. การส่งออก (Output Layer)"]
        direction TB
        D1["Templates (บทที่ 32)<br/>html/template, text/template<br/>layout, partials, functions"]
        D2["Configuration (บทที่ 33)<br/>viper, env, flags<br/>multiple sources, defaults"]
        
        D1 --> D2
    end

    Input --> Processing
    Processing --> Storage
    Storage --> Output

    subgraph Example["ตัวอย่างแอปพลิเคชัน"]
        E1["REST API Server<br/>+ Authentication<br/>+ Database<br/>+ Logging"]
        E2["Web Application<br/>+ Templates<br/>+ Session<br/>+ Configuration"]
    end

    Output --> E1
    Output --> E2
```

---

## คำอธิบายภาษาไทย (แบบละเอียด)

ภาคที่ 4 ครอบคลุมบทที่ 24–33 โดยมีเนื้อหาแบ่งเป็น 4 ช่วงหลัก ตามแผนภาพด้านบน ซึ่งแสดงการไหลของข้อมูลจาก Input → Processing → Storage → Output

---

## 1. การรับข้อมูล (Input Layer)

### บทที่ 24: ฟังก์ชันนิรนาม (Anonymous Functions) และ Closure

**เป้าหมาย:** เข้าใจการสร้างฟังก์ชันแบบไม่มีชื่อและการใช้งาน closure

#### Anonymous Functions

```go
package main

import "fmt"

func main() {
    // 1. ประกาศ anonymous function และเรียกใช้ทันที
    func() {
        fmt.Println("Hello from anonymous")
    }()
    
    // 2. ประกาศและเก็บไว้ในตัวแปร
    add := func(a, b int) int {
        return a + b
    }
    result := add(3, 5)
    fmt.Println(result) // 8
    
    // 3. ส่ง anonymous function เป็น argument
    numbers := []int{1, 2, 3, 4, 5}
    
    // filter function
    evens := filter(numbers, func(n int) bool {
        return n%2 == 0
    })
    fmt.Println(evens) // [2, 4]
    
    // 4. คืนค่าเป็น function
    multiplier := createMultiplier(3)
    fmt.Println(multiplier(5)) // 15
    fmt.Println(multiplier(10)) // 30
}

func filter(nums []int, predicate func(int) bool) []int {
    result := make([]int, 0)
    for _, n := range nums {
        if predicate(n) {
            result = append(result, n)
        }
    }
    return result
}

func createMultiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}
```

#### Closures (การปิดล้อมตัวแปร)

```go
package main

import "fmt"

func main() {
    // 1. Closure จับตัวแปรภายนอก
    counter := func() func() int {
        count := 0
        return func() int {
            count++
            return count
        }
    }()
    
    fmt.Println(counter()) // 1
    fmt.Println(counter()) // 2
    fmt.Println(counter()) // 3
    
    // 2. Closure ใน loop (ข้อควรระวัง!)
    funcs := make([]func(), 3)
    
    // ❌ Wrong: capture by reference
    for i := 0; i < 3; i++ {
        funcs[i] = func() {
            fmt.Println(i) // i is captured by reference
        }
    }
    
    for _, f := range funcs {
        f() // prints 3, 3, 3 (all the same!)
    }
    
    // ✅ Correct: capture by value
    funcs2 := make([]func(), 3)
    for i := 0; i < 3; i++ {
        i := i // create new variable
        funcs2[i] = func() {
            fmt.Println(i)
        }
    }
    
    for _, f := range funcs2 {
        f() // prints 0, 1, 2
    }
}
```

#### ตัวอย่างการใช้งานจริง

**1. Middleware Pattern**

```go
// HTTP middleware with closure
type Middleware func(http.Handler) http.Handler

// Logger middleware
func Logger(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        
        // สร้าง response wrapper เพื่อ capture status
        rw := &responseWriter{ResponseWriter: w, status: http.StatusOK}
        
        next.ServeHTTP(rw, r)
        
        duration := time.Since(start)
        log.Printf("%s %s %d %v", r.Method, r.URL.Path, rw.status, duration)
    })
}

// Auth middleware
func Auth(requiredRole string) Middleware {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            token := r.Header.Get("Authorization")
            
            user, err := validateToken(token)
            if err != nil {
                http.Error(w, "unauthorized", http.StatusUnauthorized)
                return
            }
            
            if user.Role != requiredRole {
                http.Error(w, "forbidden", http.StatusForbidden)
                return
            }
            
            // ใส่ user info ลงใน context
            ctx := context.WithValue(r.Context(), "user", user)
            next.ServeHTTP(w, r.WithContext(ctx))
        })
    }
}

// Apply middlewares
func applyMiddlewares(handler http.Handler, middlewares ...Middleware) http.Handler {
    for _, m := range middlewares {
        handler = m(handler)
    }
    return handler
}

// ใช้งาน
func main() {
    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, World!"))
    })
    
    // chain middlewares
    wrapped := applyMiddlewares(
        handler,
        Logger,
        Auth("admin"),
    )
    
    http.ListenAndServe(":8080", wrapped)
}
```

**2. Callback Pattern**

```go
// Event handler with callbacks
type EventHandler struct {
    onSuccess func(data interface{})
    onError   func(err error)
    onComplete func()
}

func (h *EventHandler) OnSuccess(fn func(interface{})) *EventHandler {
    h.onSuccess = fn
    return h
}

func (h *EventHandler) OnError(fn func(error)) *EventHandler {
    h.onError = fn
    return h
}

func (h *EventHandler) OnComplete(fn func()) *EventHandler {
    h.onComplete = fn
    return h
}

func (h *EventHandler) Execute(task func() (interface{}, error)) {
    go func() {
        defer func() {
            if h.onComplete != nil {
                h.onComplete()
            }
        }()
        
        result, err := task()
        if err != nil && h.onError != nil {
            h.onError(err)
            return
        }
        
        if h.onSuccess != nil {
            h.onSuccess(result)
        }
    }()
}

// ใช้งาน
func main() {
    handler := &EventHandler{}
    
    handler.
        OnSuccess(func(data interface{}) {
            fmt.Printf("Success: %v\n", data)
        }).
        OnError(func(err error) {
            fmt.Printf("Error: %v\n", err)
        }).
        OnComplete(func() {
            fmt.Println("Task completed")
        }).
        Execute(func() (interface{}, error) {
            // simulate work
            time.Sleep(2 * time.Second)
            return map[string]string{"status": "ok"}, nil
        })
    
    time.Sleep(3 * time.Second)
}
```

**3. Factory Pattern with Closure**

```go
// Database connection factory
type DBConfig struct {
    Host     string
    Port     int
    User     string
    Password string
    Database string
}

func NewDBConnection(config DBConfig) (*sql.DB, error) {
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        config.User, config.Password, config.Host, config.Port, config.Database)
    
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }
    
    db.SetMaxOpenConns(25)
    db.SetMaxIdleConns(25)
    db.SetConnMaxLifetime(5 * time.Minute)
    
    return db, nil
}

// Repository factory with closure
type RepositoryFactory struct {
    db *sql.DB
}

func (f *RepositoryFactory) NewUserRepository() *UserRepository {
    return &UserRepository{db: f.db}
}

func (f *RepositoryFactory) NewOrderRepository() *OrderRepository {
    return &OrderRepository{db: f.db}
}

// ใช้งาน
func main() {
    config := DBConfig{
        Host:     "localhost",
        Port:     3306,
        User:     "root",
        Password: "password",
        Database: "myapp",
    }
    
    db, _ := NewDBConnection(config)
    factory := &RepositoryFactory{db: db}
    
    userRepo := factory.NewUserRepository()
    orderRepo := factory.NewOrderRepository()
    
    // use repositories...
}
```

---

### บทที่ 25: การจัดการข้อมูล JSON และ XML

**เป้าหมาย:** แปลงข้อมูลระหว่าง Go struct กับ JSON/XML

#### JSON Handling

```go
package main

import (
    "encoding/json"
    "fmt"
    "time"
)

// struct with JSON tags
type User struct {
    ID        int       `json:"id"`
    Name      string    `json:"name"`
    Email     string    `json:"email,omitempty"`
    Password  string    `json:"-"`  // ไม่ถูก serialize
    CreatedAt time.Time `json:"created_at"`
    Tags      []string  `json:"tags,omitempty"`
    Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

// custom marshaling
type CustomTime struct {
    time.Time
}

func (ct CustomTime) MarshalJSON() ([]byte, error) {
    return []byte(fmt.Sprintf("\"%s\"", ct.Format("2006-01-02 15:04:05"))), nil
}

func (ct *CustomTime) UnmarshalJSON(data []byte) error {
    str := string(data)
    str = str[1 : len(str)-1] // remove quotes
    t, err := time.Parse("2006-01-02 15:04:05", str)
    if err != nil {
        return err
    }
    ct.Time = t
    return nil
}

func main() {
    // 1. Marshal (struct -> JSON)
    user := User{
        ID:        1,
        Name:      "John Doe",
        Email:     "john@example.com",
        CreatedAt: time.Now(),
        Tags:      []string{"golang", "developer"},
        Metadata: map[string]interface{}{
            "role": "admin",
            "active": true,
        },
    }
    
    jsonData, err := json.Marshal(user)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(jsonData))
    
    // 2. Marshal with indentation (pretty print)
    prettyJSON, _ := json.MarshalIndent(user, "", "  ")
    fmt.Println(string(prettyJSON))
    
    // 3. Unmarshal (JSON -> struct)
    jsonString := `{"id":2,"name":"Jane Doe","email":"jane@example.com"}`
    var user2 User
    err = json.Unmarshal([]byte(jsonString), &user2)
    if err != nil {
        panic(err)
    }
    fmt.Printf("%+v\n", user2)
    
    // 4. Unmarshal to map
    var data map[string]interface{}
    err = json.Unmarshal([]byte(jsonString), &data)
    fmt.Printf("name: %v\n", data["name"])
    
    // 5. Streaming JSON (for large data)
    decoder := json.NewDecoder(strings.NewReader(jsonString))
    var user3 User
    decoder.Decode(&user3)
    fmt.Printf("%+v\n", user3)
    
    // 6. Encoder for streaming output
    encoder := json.NewEncoder(os.Stdout)
    encoder.SetIndent("", "  ")
    encoder.Encode(user)
}
```

#### ตัวอย่างการใช้งานจริง

**1. API Response Wrapper**

```go
type APIResponse struct {
    Success bool        `json:"success"`
    Data    interface{} `json:"data,omitempty"`
    Error   *APIError   `json:"error,omitempty"`
    Meta    *APIMeta    `json:"meta,omitempty"`
}

type APIError struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
    Details string `json:"details,omitempty"`
}

type APIMeta struct {
    Page       int `json:"page"`
    PerPage    int `json:"per_page"`
    Total      int `json:"total"`
    TotalPages int `json:"total_pages"`
}

func JSONResponse(w http.ResponseWriter, status int, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    
    response := APIResponse{
        Success: status >= 200 && status < 300,
        Data:    data,
    }
    
    json.NewEncoder(w).Encode(response)
}

func JSONError(w http.ResponseWriter, status int, code int, message string) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    
    response := APIResponse{
        Success: false,
        Error: &APIError{
            Code:    code,
            Message: message,
        },
    }
    
    json.NewEncoder(w).Encode(response)
}

// ใช้งาน
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
    user, err := userService.GetUser(1)
    if err != nil {
        JSONError(w, http.StatusNotFound, 404, "User not found")
        return
    }
    
    JSONResponse(w, http.StatusOK, user)
}
```

**2. JSON Validation with Struct Tags**

```go
type RegisterRequest struct {
    Username string `json:"username" validate:"required,min=3,max=50"`
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,min=8"`
    Age      int    `json:"age" validate:"gte=18,lte=99"`
    Phone    string `json:"phone" validate:"omitempty,e164"`
}

func (r *RegisterRequest) Validate() error {
    validate := validator.New()
    return validate.Struct(r)
}

// ใช้งาน
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
    var req RegisterRequest
    
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        JSONError(w, http.StatusBadRequest, 400, "Invalid JSON")
        return
    }
    
    if err := req.Validate(); err != nil {
        JSONError(w, http.StatusBadRequest, 400, err.Error())
        return
    }
    
    // process registration...
}
```

**3. XML Handling**

```go
package main

import (
    "encoding/xml"
    "fmt"
)

// XML struct with tags
type Person struct {
    XMLName   xml.Name `xml:"person"`
    ID        int      `xml:"id,attr"`
    Name      string   `xml:"name"`
    Age       int      `xml:"age"`
    Address   Address  `xml:"address"`
    Phones    []Phone  `xml:"phones>phone"`
}

type Address struct {
    Street  string `xml:"street"`
    City    string `xml:"city"`
    Country string `xml:"country"`
}

type Phone struct {
    Type  string `xml:"type,attr"`
    Number string `xml:",chardata"`
}

func main() {
    // Marshal struct -> XML
    person := Person{
        ID:   1,
        Name: "John Doe",
        Age:  30,
        Address: Address{
            Street:  "123 Main St",
            City:    "Bangkok",
            Country: "Thailand",
        },
        Phones: []Phone{
            {Type: "home", Number: "02-123-4567"},
            {Type: "mobile", Number: "081-234-5678"},
        },
    }
    
    xmlData, _ := xml.MarshalIndent(person, "", "  ")
    fmt.Println(string(xmlData))
    
    // Unmarshal XML -> struct
    xmlString := `<?xml version="1.0" encoding="UTF-8"?>
    <person id="2">
        <name>Jane Doe</name>
        <age>25</age>
        <address>
            <street>456 Oak Ave</street>
            <city>Chiang Mai</city>
            <country>Thailand</country>
        </address>
        <phones>
            <phone type="mobile">082-987-6543</phone>
        </phones>
    </person>`
    
    var person2 Person
    xml.Unmarshal([]byte(xmlString), &person2)
    fmt.Printf("%+v\n", person2)
}
```

---

### บทที่ 26: พื้นฐานการสร้าง HTTP Server

**เป้าหมาย:** สร้าง HTTP server ที่มีประสิทธิภาพและปลอดภัย

#### Basic HTTP Server

```go
package main

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
)

func main() {
    // 1. Basic handler
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
    })
    
    // 2. Different HTTP methods
    http.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodGet:
            // get users
            w.Write([]byte("GET users"))
        case http.MethodPost:
            // create user
            w.Write([]byte("POST users"))
        case http.MethodPut:
            // update user
            w.Write([]byte("PUT users"))
        case http.MethodDelete:
            // delete user
            w.Write([]byte("DELETE users"))
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    })
    
    // 3. Server configuration
    server := &http.Server{
        Addr:         ":8080",
        ReadTimeout:  15 * time.Second,
        WriteTimeout: 15 * time.Second,
        IdleTimeout:  60 * time.Second,
        Handler:      nil, // use default ServeMux
    }
    
    // 4. Graceful shutdown
    go func() {
        log.Printf("Server starting on %s", server.Addr)
        if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("Server failed: %v", err)
        }
    }()
    
    // Wait for interrupt signal
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    log.Println("Shutting down server...")
    
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    
    if err := server.Shutdown(ctx); err != nil {
        log.Fatalf("Server forced to shutdown: %v", err)
    }
    
    log.Println("Server exited")
}
```

#### Using chi Router

```go
package main

import (
    "context"
    "encoding/json"
    "net/http"
    "time"
    
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "github.com/go-chi/cors"
    "github.com/go-chi/httprate"
)

func main() {
    r := chi.NewRouter()
    
    // Global middlewares
    r.Use(middleware.Logger)          // log requests
    r.Use(middleware.Recoverer)       // recover from panics
    r.Use(middleware.RealIP)          // get real IP
    r.Use(middleware.RequestID)       // add request ID
    r.Use(middleware.Timeout(60 * time.Second)) // timeout
    
    // CORS
    r.Use(cors.Handler(cors.Options{
        AllowedOrigins:   []string{"https://*", "http://*"},
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
        ExposedHeaders:   []string{"Link"},
        AllowCredentials: false,
        MaxAge:           300,
    }))
    
    // Rate limiting
    r.Use(httprate.LimitByIP(100, 1*time.Minute))
    
    // Routes
    r.Get("/", homeHandler)
    
    r.Route("/api/v1", func(r chi.Router) {
        r.Get("/health", healthHandler)
        
        r.Route("/users", func(r chi.Router) {
            r.Get("/", listUsersHandler)
            r.Post("/", createUserHandler)
            
            r.Route("/{userID}", func(r chi.Router) {
                r.Use(userContext)
                r.Get("/", getUserHandler)
                r.Put("/", updateUserHandler)
                r.Delete("/", deleteUserHandler)
            })
        })
    })
    
    // Static files
    r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
    
    // Start server
    http.ListenAndServe(":8080", r)
}

// Middleware to load user from database
func userContext(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        userID := chi.URLParam(r, "userID")
        
        // Load user from database
        user, err := loadUser(userID)
        if err != nil {
            http.Error(w, "User not found", http.StatusNotFound)
            return
        }
        
        ctx := context.WithValue(r.Context(), "user", user)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
    user := r.Context().Value("user").(*User)
    json.NewEncoder(w).Encode(user)
}
```

#### ตัวอย่างการใช้งานจริง: REST API Server

```go
// models/user.go
package models

import (
    "time"
    "golang.org/x/crypto/bcrypt"
)

type User struct {
    ID        int       `json:"id"`
    Username  string    `json:"username"`
    Email     string    `json:"email"`
    Password  string    `json:"-"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) HashPassword() error {
    hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    u.Password = string(hashed)
    return nil
}

func (u *User) CheckPassword(password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
    return err == nil
}

// handlers/user_handler.go
package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"
    
    "github.com/go-chi/chi/v5"
    "github.com/go-playground/validator/v10"
)

type UserHandler struct {
    userService UserService
    validate    *validator.Validate
}

func NewUserHandler(userService UserService) *UserHandler {
    return &UserHandler{
        userService: userService,
        validate:    validator.New(),
    }
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
    var req CreateUserRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        JSONError(w, http.StatusBadRequest, "Invalid request body")
        return
    }
    
    if err := h.validate.Struct(req); err != nil {
        JSONError(w, http.StatusBadRequest, err.Error())
        return
    }
    
    user, err := h.userService.CreateUser(r.Context(), &req)
    if err != nil {
        JSONError(w, http.StatusInternalServerError, err.Error())
        return
    }
    
    JSONResponse(w, http.StatusCreated, user)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
    idStr := chi.URLParam(r, "id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        JSONError(w, http.StatusBadRequest, "Invalid user ID")
        return
    }
    
    user, err := h.userService.GetUserByID(r.Context(), id)
    if err != nil {
        JSONError(w, http.StatusNotFound, "User not found")
        return
    }
    
    JSONResponse(w, http.StatusOK, user)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
    idStr := chi.URLParam(r, "id")
    id, _ := strconv.Atoi(idStr)
    
    var req UpdateUserRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        JSONError(w, http.StatusBadRequest, "Invalid request body")
        return
    }
    
    user, err := h.userService.UpdateUser(r.Context(), id, &req)
    if err != nil {
        JSONError(w, http.StatusInternalServerError, err.Error())
        return
    }
    
    JSONResponse(w, http.StatusOK, user)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
    idStr := chi.URLParam(r, "id")
    id, _ := strconv.Atoi(idStr)
    
    if err := h.userService.DeleteUser(r.Context(), id); err != nil {
        JSONError(w, http.StatusInternalServerError, err.Error())
        return
    }
    
    w.WriteHeader(http.StatusNoContent)
}

func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
    page, _ := strconv.Atoi(r.URL.Query().Get("page"))
    if page < 1 {
        page = 1
    }
    
    limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
    if limit < 1 || limit > 100 {
        limit = 20
    }
    
    users, total, err := h.userService.ListUsers(r.Context(), page, limit)
    if err != nil {
        JSONError(w, http.StatusInternalServerError, err.Error())
        return
    }
    
    response := &ListUsersResponse{
        Data:       users,
        Total:      total,
        Page:       page,
        PerPage:    limit,
        TotalPages: (total + limit - 1) / limit,
    }
    
    JSONResponse(w, http.StatusOK, response)
}
```

---

### บทที่ 27: Enum, Iota และ Bitmask

**เป้าหมาย:** ใช้ iota สำหรับสร้าง enum และ bitmask

#### Basic Iota

```go
package main

import "fmt"

// 1. Basic enum
type Weekday int

const (
    Sunday Weekday = iota  // 0
    Monday                 // 1
    Tuesday                // 2
    Wednesday              // 3
    Thursday               // 4
    Friday                 // 5
    Saturday               // 6
)

// 2. Skip values
const (
    _ = iota              // 0 (skip)
    KB = 1 << (10 * iota) // 1 << 10 = 1024
    MB                    // 1 << 20 = 1048576
    GB                    // 1 << 30 = 1073741824
    TB                    // 1 << 40 = 1099511627776
)

// 3. Enum with string representation
type Status int

const (
    StatusPending Status = iota
    StatusApproved
    StatusRejected
    StatusCancelled
)

func (s Status) String() string {
    return [...]string{"Pending", "Approved", "Rejected", "Cancelled"}[s]
}

// 4. Enum with JSON marshaling
type OrderStatus int

const (
    OrderStatusCreated OrderStatus = iota + 1  // start from 1
    OrderStatusPaid
    OrderStatusShipped
    OrderStatusDelivered
    OrderStatusCancelled
)

func (s OrderStatus) MarshalJSON() ([]byte, error) {
    return []byte(fmt.Sprintf(`"%s"`, s.String())), nil
}

func (s OrderStatus) String() string {
    return [...]string{"", "Created", "Paid", "Shipped", "Delivered", "Cancelled"}[s]
}

func main() {
    fmt.Println(Sunday)    // 0
    fmt.Println(Monday)    // 1
    fmt.Println(KB)        // 1024
    fmt.Println(StatusPending) // 0
    fmt.Println(StatusPending.String()) // "Pending"
}
```

#### Bitmask with Iota

```go
package main

import "fmt"

// Permission flags (bitmask)
type Permission int

const (
    PermRead   Permission = 1 << iota  // 1 (001)
    PermWrite                          // 2 (010)
    PermExecute                        // 4 (100)
    PermDelete                         // 8 (1000)
    PermAdmin   = PermRead | PermWrite | PermExecute | PermDelete  // 15 (1111)
)

func (p Permission) String() string {
    if p == 0 {
        return "None"
    }
    
    var flags []string
    if p&PermRead != 0 {
        flags = append(flags, "Read")
    }
    if p&PermWrite != 0 {
        flags = append(flags, "Write")
    }
    if p&PermExecute != 0 {
        flags = append(flags, "Execute")
    }
    if p&PermDelete != 0 {
        flags = append(flags, "Delete")
    }
    
    return strings.Join(flags, "|")
}

func (p *Permission) Add(perm Permission) {
    *p |= perm
}

func (p *Permission) Remove(perm Permission) {
    *p &^= perm
}

func (p Permission) Has(perm Permission) bool {
    return p&perm != 0
}

// ตัวอย่างการใช้งาน
type User struct {
    Name        string
    Permissions Permission
}

func main() {
    user := User{
        Name:        "john",
        Permissions: PermRead | PermWrite,
    }
    
    fmt.Println(user.Permissions) // Read|Write
    fmt.Println(user.Permissions.Has(PermRead))   // true
    fmt.Println(user.Permissions.Has(PermDelete)) // false
    
    user.Permissions.Add(PermDelete)
    fmt.Println(user.Permissions) // Read|Write|Delete
    
    user.Permissions.Remove(PermWrite)
    fmt.Println(user.Permissions) // Read|Delete
}
```

#### ตัวอย่างการใช้งานจริง

**1. Order Status Management**

```go
type OrderStatus int

const (
    OrderStatusDraft OrderStatus = iota + 1
    OrderStatusPending
    OrderStatusConfirmed
    OrderStatusProcessing
    OrderStatusShipped
    OrderStatusDelivered
    OrderStatusCancelled
    OrderStatusRefunded
)

var orderStatusTransitions = map[OrderStatus][]OrderStatus{
    OrderStatusDraft:      {OrderStatusPending, OrderStatusCancelled},
    OrderStatusPending:    {OrderStatusConfirmed, OrderStatusCancelled},
    OrderStatusConfirmed:  {OrderStatusProcessing, OrderStatusCancelled},
    OrderStatusProcessing: {OrderStatusShipped, OrderStatusCancelled},
    OrderStatusShipped:    {OrderStatusDelivered},
    OrderStatusDelivered:  {OrderStatusRefunded},
    OrderStatusCancelled:  {},
    OrderStatusRefunded:   {},
}

func (s OrderStatus) CanTransitionTo(target OrderStatus) bool {
    allowed := orderStatusTransitions[s]
    for _, allowedStatus := range allowed {
        if allowedStatus == target {
            return true
        }
    }
    return false
}

func (s OrderStatus) IsFinal() bool {
    return s == OrderStatusDelivered || 
           s == OrderStatusCancelled || 
           s == OrderStatusRefunded
}

func (s OrderStatus) String() string {
    return [...]string{
        "", "Draft", "Pending", "Confirmed", 
        "Processing", "Shipped", "Delivered", 
        "Cancelled", "Refunded",
    }[s]
}

// ใช้งาน
type Order struct {
    ID     int
    Status OrderStatus
}

func (o *Order) TransitionTo(newStatus OrderStatus) error {
    if !o.Status.CanTransitionTo(newStatus) {
        return fmt.Errorf("cannot transition from %s to %s", 
            o.Status, newStatus)
    }
    
    o.Status = newStatus
    return nil
}
```

**2. HTTP Status Code Group**

```go
type HTTPStatusGroup int

const (
    StatusGroupInfo HTTPStatusGroup = iota + 1
    StatusGroupSuccess
    StatusGroupRedirect
    StatusGroupClientError
    StatusGroupServerError
)

func GetStatusGroup(code int) HTTPStatusGroup {
    switch {
    case code >= 100 && code < 200:
        return StatusGroupInfo
    case code >= 200 && code < 300:
        return StatusGroupSuccess
    case code >= 300 && code < 400:
        return StatusGroupRedirect
    case code >= 400 && code < 500:
        return StatusGroupClientError
    case code >= 500 && code < 600:
        return StatusGroupServerError
    default:
        return 0
    }
}

func (g HTTPStatusGroup) String() string {
    return [...]string{
        "", "Informational", "Success", 
        "Redirection", "Client Error", "Server Error",
    }[g]
}
```

---

### บทที่ 28: วันที่และเวลา (Time)

**เป้าหมาย:** จัดการ date, time, duration, timezone อย่างถูกต้อง

#### Time Basics

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    // 1. Current time
    now := time.Now()
    fmt.Printf("Now: %v\n", now)
    fmt.Printf("Unix: %d\n", now.Unix())
    fmt.Printf("UnixNano: %d\n", now.UnixNano())
    
    // 2. Creating time
    t1 := time.Date(2024, time.December, 25, 10, 30, 0, 0, time.UTC)
    fmt.Println(t1)
    
    // 3. Parsing time
    t2, _ := time.Parse("2006-01-02 15:04:05", "2024-12-25 10:30:00")
    fmt.Println(t2)
    
    // 4. Formatting time
    fmt.Println(now.Format("2006-01-02 15:04:05"))
    fmt.Println(now.Format("Monday, 02 Jan 2006"))
    fmt.Println(now.Format(time.RFC3339))
    
    // 5. Time components
    fmt.Printf("Year: %d, Month: %d, Day: %d\n", now.Year(), now.Month(), now.Day())
    fmt.Printf("Hour: %d, Minute: %d, Second: %d\n", now.Hour(), now.Minute(), now.Second())
    fmt.Printf("Weekday: %s\n", now.Weekday())
    
    // 6. Time arithmetic
    tomorrow := now.Add(24 * time.Hour)
    yesterday := now.Add(-24 * time.Hour)
    oneWeekLater := now.AddDate(0, 0, 7)
    
    fmt.Printf("Tomorrow: %v\n", tomorrow)
    fmt.Printf("Yesterday: %v\n", yesterday)
    
    // 7. Time difference
    duration := tomorrow.Sub(now)
    fmt.Printf("Duration: %v\n", duration)
    fmt.Printf("Hours: %.2f\n", duration.Hours())
    fmt.Printf("Minutes: %.2f\n", duration.Minutes())
    
    // 8. Compare times
    fmt.Println(now.Before(tomorrow))  // true
    fmt.Println(now.After(yesterday))  // true
    fmt.Println(now.Equal(now))        // true
}
```

#### Timezone Handling

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    // 1. Load location
    loc, _ := time.LoadLocation("Asia/Bangkok")
    nyLoc, _ := time.LoadLocation("America/New_York")
    
    // 2. Create time with location
    t := time.Date(2024, 12, 25, 10, 30, 0, 0, loc)
    fmt.Println(t)
    
    // 3. Convert between timezones
    tUTC := t.UTC()
    tNY := t.In(nyLoc)
    
    fmt.Printf("Bangkok: %v\n", t)
    fmt.Printf("UTC: %v\n", tUTC)
    fmt.Printf("New York: %v\n", tNY)
    
    // 4. Parse with timezone
    t2, _ := time.ParseInLocation(
        "2006-01-02 15:04:05",
        "2024-12-25 10:30:00",
        loc,
    )
    fmt.Println(t2)
    
    // 5. Get current time in different timezone
    nowBKK := time.Now().In(loc)
    nowNY := time.Now().In(nyLoc)
    fmt.Printf("Bangkok now: %v\n", nowBKK)
    fmt.Printf("New York now: %v\n", nowNY)
}
```

#### Duration and Ticker

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    // 1. Duration parsing
    duration, _ := time.ParseDuration("1h30m")
    fmt.Println(duration) // 1h30m0s
    
    duration2, _ := time.ParseDuration("2h15m30s")
    fmt.Println(duration2.Hours())   // 2.2583333333333333
    fmt.Println(duration2.Minutes()) // 135.5
    
    // 2. Sleep
    fmt.Println("Sleeping...")
    time.Sleep(2 * time.Second)
    fmt.Println("Awake!")
    
    // 3. Timer (one-shot)
    timer := time.NewTimer(3 * time.Second)
    <-timer.C
    fmt.Println("Timer expired")
    
    // 4. Timer with stop
    timer2 := time.NewTimer(5 * time.Second)
    go func() {
        <-timer2.C
        fmt.Println("Timer expired")
    }()
    
    time.Sleep(2 * time.Second)
    if timer2.Stop() {
        fmt.Println("Timer stopped")
    }
    
    // 5. Ticker (repeating)
    ticker := time.NewTicker(1 * time.Second)
    done := make(chan bool)
    
    go func() {
        for {
            select {
            case t := <-ticker.C:
                fmt.Printf("Tick at %v\n", t)
            case <-done:
                ticker.Stop()
                return
            }
        }
    }()
    
    time.Sleep(5 * time.Second)
    done <- true
    fmt.Println("Ticker stopped")
    
    // 6. After (timeout pattern)
    timeout := time.After(3 * time.Second)
    select {
    case <-timeout:
        fmt.Println("Timeout!")
    }
    
    // 7. AfterFunc
    time.AfterFunc(2*time.Second, func() {
        fmt.Println("AfterFunc executed")
    })
    
    time.Sleep(3 * time.Second)
}
```

#### ตัวอย่างการใช้งานจริง

**1. Date Range Validation**

```go
type DateRange struct {
    Start time.Time
    End   time.Time
}

func (r *DateRange) Validate() error {
    if r.Start.IsZero() {
        return errors.New("start date is required")
    }
    if r.End.IsZero() {
        return errors.New("end date is required")
    }
    if r.End.Before(r.Start) {
        return errors.New("end date must be after start date")
    }
    return nil
}

func (r *DateRange) Contains(t time.Time) bool {
    return (t.Equal(r.Start) || t.After(r.Start)) && 
           (t.Equal(r.End) || t.Before(r.End))
}

func (r *DateRange) Overlaps(other *DateRange) bool {
    return r.Start.Before(other.End) && r.End.After(other.Start)
}

func (r *DateRange) Duration() time.Duration {
    return r.End.Sub(r.Start)
}

// ใช้งาน
func GetBookingsForDateRange(db *sql.DB, start, end time.Time) ([]Booking, error) {
    range := &DateRange{Start: start, End: end}
    if err := range.Validate(); err != nil {
        return nil, err
    }
    
    query := `SELECT id, user_id, start_time, end_time 
              FROM bookings 
              WHERE start_time >= ? AND end_time <= ?`
    
    rows, err := db.Query(query, start, end)
    // ... process rows
}
```

**2. Rate Limiter with Time Window**

```go
type RateLimiter struct {
    limit     int
    window    time.Duration
    requests  map[string][]time.Time
    mu        sync.RWMutex
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
    return &RateLimiter{
        limit:    limit,
        window:   window,
        requests: make(map[string][]time.Time),
    }
}

func (rl *RateLimiter) Allow(key string) bool {
    rl.mu.Lock()
    defer rl.mu.Unlock()
    
    now := time.Now()
    cutoff := now.Add(-rl.window)
    
    // เก็บเฉพาะ requests ที่อยู่ใน window
    requests := rl.requests[key]
    valid := []time.Time{}
    for _, t := range requests {
        if t.After(cutoff) {
            valid = append(valid, t)
        }
    }
    
    if len(valid) >= rl.limit {
        return false
    }
    
    valid = append(valid, now)
    rl.requests[key] = valid
    return true
}

func (rl *RateLimiter) Cleanup() {
    rl.mu.Lock()
    defer rl.mu.Unlock()
    
    cutoff := time.Now().Add(-rl.window)
    for key, requests := range rl.requests {
        valid := []time.Time{}
        for _, t := range requests {
            if t.After(cutoff) {
                valid = append(valid, t)
            }
        }
        if len(valid) == 0 {
            delete(rl.requests, key)
        } else {
            rl.requests[key] = valid
        }
    }
}

// ใช้งาน
func main() {
    limiter := NewRateLimiter(100, time.Minute)
    
    // Cleanup every hour
    ticker := time.NewTicker(time.Hour)
    go func() {
        for range ticker.C {
            limiter.Cleanup()
        }
    }()
    
    // Use in handler
    http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
        ip := r.RemoteAddr
        
        if !limiter.Allow(ip) {
            http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
            return
        }
        
        // process request
    })
}
```

**3. Schedule Job with Cron-like**

```go
type Scheduler struct {
    jobs   []*Job
    stopCh chan struct{}
}

type Job struct {
    Name     string
    Schedule string  // "0 9 * * *" = every day at 9am
    Task     func()
    LastRun  time.Time
    NextRun  time.Time
}

func NewScheduler() *Scheduler {
    return &Scheduler{
        jobs:   []*Job{},
        stopCh: make(chan struct{}),
    }
}

func (s *Scheduler) AddJob(job *Job) {
    s.jobs = append(s.jobs, job)
    s.calculateNextRun(job)
}

func (s *Scheduler) calculateNextRun(job *Job) {
    // Simplified cron parsing (use robfig/cron for production)
    now := time.Now()
    job.NextRun = now.Add(24 * time.Hour) // placeholder
}

func (s *Scheduler) Start() {
    ticker := time.NewTicker(1 * time.Minute)
    go func() {
        for {
            select {
            case <-ticker.C:
                s.checkJobs()
            case <-s.stopCh:
                ticker.Stop()
                return
            }
        }
    }()
}

func (s *Scheduler) checkJobs() {
    now := time.Now()
    for _, job := range s.jobs {
        if now.After(job.NextRun) {
            go job.Task()
            job.LastRun = now
            s.calculateNextRun(job)
        }
    }
}

func (s *Scheduler) Stop() {
    close(s.stopCh)
}
```

---

## 2. การประมวลผล (Processing Layer)

### บทที่ 30: การทำงานพร้อมกัน (Concurrency)

**เป้าหมาย:** ใช้ goroutines และ channels ในการเขียนโปรแกรม concurrent

#### Goroutines Basics

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    // 1. Start a goroutine
    go func() {
        fmt.Println("Hello from goroutine")
    }()
    
    // 2. Wait for goroutine (bad way)
    time.Sleep(100 * time.Millisecond)
    
    // 3. Use WaitGroup
    var wg sync.WaitGroup
    
    for i := 1; i <= 5; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            fmt.Printf("Worker %d done\n", id)
        }(i)
    }
    
    wg.Wait()
    fmt.Println("All workers done")
    
    // 4. Goroutine with panic recovery
    var wg2 sync.WaitGroup
    wg2.Add(1)
    
    go func() {
        defer wg2.Done()
        defer func() {
            if r := recover(); r != nil {
                fmt.Printf("Recovered from panic: %v\n", r)
            }
        }()
        
        panic("something went wrong")
    }()
    
    wg2.Wait()
}
```

#### Channels

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    // 1. Unbuffered channel
    ch := make(chan int)
    
    go func() {
        ch <- 42  // send
    }()
    
    value := <-ch  // receive
    fmt.Println(value)
    
    // 2. Buffered channel
    buffered := make(chan string, 3)
    buffered <- "one"
    buffered <- "two"
    buffered <- "three"
    // buffered <- "four" // would block
    
    fmt.Println(<-buffered)
    fmt.Println(<-buffered)
    fmt.Println(<-buffered)
    
    // 3. Channel with close
    ch2 := make(chan int, 5)
    
    go func() {
        for i := 0; i < 5; i++ {
            ch2 <- i
        }
        close(ch2)
    }()
    
    // Range over channel (auto stops when closed)
    for v := range ch2 {
        fmt.Println(v)
    }
    
    // 4. Select statement
    ch3 := make(chan string)
    ch4 := make(chan string)
    
    go func() {
        time.Sleep(1 * time.Second)
        ch3 <- "from channel 3"
    }()
    
    go func() {
        time.Sleep(2 * time.Second)
        ch4 <- "from channel 4"
    }()
    
    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-ch3:
            fmt.Println(msg1)
        case msg2 := <-ch4:
            fmt.Println(msg2)
        case <-time.After(3 * time.Second):
            fmt.Println("timeout")
        }
    }
    
    // 5. Non-blocking channel operations
    ch5 := make(chan int, 1)
    
    select {
    case ch5 <- 42:
        fmt.Println("Sent")
    default:
        fmt.Println("Channel is full")
    }
    
    select {
    case v := <-ch5:
        fmt.Printf("Received: %d\n", v)
    default:
        fmt.Println("Channel is empty")
    }
}
```

#### Worker Pool Pattern

```go
package main

import (
    "context"
    "fmt"
    "sync"
    "time"
)

type Job struct {
    ID    int
    Data  interface{}
}

type Result struct {
    JobID   int
    Output  interface{}
    Error   error
}

type WorkerPool struct {
    numWorkers int
    jobQueue   chan Job
    resultQueue chan Result
    wg         sync.WaitGroup
    ctx        context.Context
    cancel     context.CancelFunc
}

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

func (wp *WorkerPool) Start(workerFunc func(Job) Result) {
    // Start workers
    for i := 0; i < wp.numWorkers; i++ {
        wp.wg.Add(1)
        go wp.worker(i, workerFunc)
    }
    
    // Wait for all workers to finish in background
    go func() {
        wp.wg.Wait()
        close(wp.resultQueue)
    }()
}

func (wp *WorkerPool) worker(id int, workerFunc func(Job) Result) {
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
            result := workerFunc(job)
            select {
            case wp.resultQueue <- result:
            case <-wp.ctx.Done():
                return
            }
        }
    }
}

func (wp *WorkerPool) Submit(job Job) {
    select {
    case wp.jobQueue <- job:
    case <-wp.ctx.Done():
    }
}

func (wp *WorkerPool) Results() <-chan Result {
    return wp.resultQueue
}

func (wp *WorkerPool) Stop() {
    wp.cancel()
    close(wp.jobQueue)
    wp.wg.Wait()
}

// ตัวอย่างการใช้งาน
func main() {
    // Create worker pool with 3 workers
    pool := NewWorkerPool(3, 100)
    
    // Define worker function
    workerFunc := func(job Job) Result {
        // Simulate work
        time.Sleep(time.Duration(job.ID) * 100 * time.Millisecond)
        
        return Result{
            JobID:  job.ID,
            Output: fmt.Sprintf("Processed job %d", job.ID),
        }
    }
    
    pool.Start(workerFunc)
    
    // Submit 10 jobs
    for i := 1; i <= 10; i++ {
        pool.Submit(Job{ID: i, Data: fmt.Sprintf("data-%d", i)})
    }
    
    // Collect results
    for result := range pool.Results() {
        if result.Error != nil {
            fmt.Printf("Job %d failed: %v\n", result.JobID, result.Error)
        } else {
            fmt.Printf("Job %d: %s\n", result.JobID, result.Output)
        }
    }
    
    pool.Stop()
}
```

#### Pipeline Pattern

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    // Pipeline: numbers -> square -> filter -> sum
    numbers := generate(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
    
    squared := square(numbers)
    filtered := filter(squared, func(n int) bool {
        return n%2 == 0 // even numbers
    })
    
    sum := sum(filtered)
    
    fmt.Println("Sum of squares of even numbers:", <-sum)
}

func generate(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums {
            out <- n
        }
        close(out)
    }()
    return out
}

func square(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * n
        }
        close(out)
    }()
    return out
}

func filter(in <-chan int, predicate func(int) bool) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            if predicate(n) {
                out <- n
            }
        }
        close(out)
    }()
    return out
}

func sum(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        total := 0
        for n := range in {
            total += n
        }
        out <- total
        close(out)
    }()
    return out
}
```

#### Fan-Out / Fan-In Pattern

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    // Generate numbers
    numbers := generateNumbers(100)
    
    // Fan-out: create multiple workers
    worker1 := squareWorker(numbers)
    worker2 := squareWorker(numbers)
    worker3 := squareWorker(numbers)
    
    // Fan-in: merge results from all workers
    results := merge(worker1, worker2, worker3)
    
    // Collect results
    for result := range results {
        fmt.Println(result)
    }
}

func generateNumbers(n int) <-chan int {
    out := make(chan int)
    go func() {
        for i := 1; i <= n; i++ {
            out <- i
        }
        close(out)
    }()
    return out
}

func squareWorker(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            time.Sleep(10 * time.Millisecond) // simulate work
            out <- n * n
        }
        close(out)
    }()
    return out
}

func merge(channels ...<-chan int) <-chan int {
    var wg sync.WaitGroup
    out := make(chan int)
    
    // Start a goroutine for each input channel
    output := func(c <-chan int) {
        for n := range c {
            out <- n
        }
        wg.Done()
    }
    
    wg.Add(len(channels))
    for _, c := range channels {
        go output(c)
    }
    
    // Close output channel when all workers done
    go func() {
        wg.Wait()
        close(out)
    }()
    
    return out
}
```

---

### บทที่ 31: การบันทึกเหตุการณ์ (Logging)

**เป้าหมาย:** ใช้ structured logging สำหรับการ debug และ monitoring

#### Using slog (Go 1.21+)

```go
package main

import (
    "context"
    "log/slog"
    "os"
    "time"
)

func main() {
    // 1. Basic logging
    slog.Info("Application started")
    slog.Warn("Low memory", "usage", 85)
    slog.Error("Database connection failed", "error", "timeout")
    
    // 2. Structured logging with attributes
    slog.Info("User logged in",
        "user_id", 123,
        "username", "john_doe",
        "ip", "192.168.1.1",
    )
    
    // 3. Log with group
    slog.Info("Request processed",
        slog.Group("request",
            "method", "GET",
            "path", "/api/users",
            "duration_ms", 150,
        ),
        slog.Group("response",
            "status", 200,
            "size", 1024,
        ),
    )
    
    // 4. Log with context
    ctx := context.WithValue(context.Background(), "request_id", "abc-123")
    slog.InfoContext(ctx, "Processing request", "user_id", 123)
    
    // 5. Set log level
    opts := &slog.HandlerOptions{
        Level: slog.LevelDebug,
    }
    handler := slog.NewJSONHandler(os.Stdout, opts)
    logger := slog.New(handler)
    slog.SetDefault(logger)
    
    logger.Debug("Debug message", "detail", "only shown in debug mode")
    
    // 6. Custom handler with fields
    handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
        Level: slog.LevelInfo,
        ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
            // Rename timestamp to time
            if a.Key == slog.TimeKey {
                a.Key = "timestamp"
            }
            return a
        },
    })
    slog.SetDefault(slog.New(handler))
    
    slog.Info("With custom timestamp key", "value", 42)
}
```

#### Production Logging Setup

```go
package logger

import (
    "io"
    "log/slog"
    "os"
    "time"
)

type Config struct {
    Level      string `yaml:"level"`   // debug, info, warn, error
    Format     string `yaml:"format"`  // json, text
    OutputPath string `yaml:"output"`  // stdout, stderr, file path
}

func NewLogger(cfg Config) (*slog.Logger, error) {
    // Parse log level
    var level slog.Level
    switch cfg.Level {
    case "debug":
        level = slog.LevelDebug
    case "info":
        level = slog.LevelInfo
    case "warn":
        level = slog.LevelWarn
    case "error":
        level = slog.LevelError
    default:
        level = slog.LevelInfo
    }
    
    // Setup output
    var writer io.Writer
    switch cfg.OutputPath {
    case "stdout":
        writer = os.Stdout
    case "stderr":
        writer = os.Stderr
    default:
        file, err := os.OpenFile(cfg.OutputPath, 
            os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
        if err != nil {
            return nil, err
        }
        writer = file
    }
    
    // Create handler
    var handler slog.Handler
    opts := &slog.HandlerOptions{
        Level: level,
        ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
            // Add service name to all logs
            if a.Key == slog.TimeKey {
                a.Value = slog.StringValue(a.Value.Time().Format(time.RFC3339))
            }
            return a
        },
    }
    
    if cfg.Format == "json" {
        handler = slog.NewJSONHandler(writer, opts)
    } else {
        handler = slog.NewTextHandler(writer, opts)
    }
    
    return slog.New(handler), nil
}

// Middleware for HTTP logging
func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        
        // Capture response status
        rw := &responseWriter{ResponseWriter: w, status: http.StatusOK}
        
        next.ServeHTTP(rw, r)
        
        duration := time.Since(start)
        
        slog.Info("HTTP request",
            "method", r.Method,
            "path", r.URL.Path,
            "status", rw.status,
            "duration_ms", duration.Milliseconds(),
            "ip", r.RemoteAddr,
            "user_agent", r.UserAgent(),
        )
    })
}

type responseWriter struct {
    http.ResponseWriter
    status int
}

func (rw *responseWriter) WriteHeader(code int) {
    rw.status = code
    rw.ResponseWriter.WriteHeader(code)
}
```

#### Zap Logger (Alternative)

```go
package main

import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

func main() {
    // 1. Production logger
    logger, _ := zap.NewProduction()
    defer logger.Sync()
    
    logger.Info("Application started",
        zap.String("version", "1.0.0"),
        zap.Int("port", 8080),
    )
    
    logger.Error("Database error",
        zap.Error(fmt.Errorf("connection refused")),
        zap.String("host", "localhost"),
    )
    
    // 2. Development logger
    devLogger, _ := zap.NewDevelopment()
    devLogger.Debug("Debug message", zap.String("key", "value"))
    
    // 3. Custom logger
    config := zap.Config{
        Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
        Development: false,
        Encoding:    "json",
        EncoderConfig: zapcore.EncoderConfig{
            TimeKey:        "timestamp",
            LevelKey:       "level",
            NameKey:        "logger",
            CallerKey:      "caller",
            MessageKey:     "message",
            StacktraceKey:  "stacktrace",
            LineEnding:     zapcore.DefaultLineEnding,
            EncodeLevel:    zapcore.LowercaseLevelEncoder,
            EncodeTime:     zapcore.ISO8601TimeEncoder,
            EncodeDuration: zapcore.SecondsDurationEncoder,
            EncodeCaller:   zapcore.ShortCallerEncoder,
        },
        OutputPaths:      []string{"stdout", "/var/log/app.log"},
        ErrorOutputPaths: []string{"stderr"},
    }
    
    customLogger, _ := config.Build()
    defer customLogger.Sync()
    
    customLogger.Info("Custom logger",
        zap.String("service", "myapp"),
    )
    
    // 4. Sugared logger (less typed)
    sugar := logger.Sugar()
    sugar.Infow("User login",
        "user_id", 123,
        "username", "john",
    )
}
```

---

### บทที่ 32: เทมเพลต (Templates)

**เป้าหมาย:** สร้าง HTML และ text templates สำหรับ dynamic content

#### Text Templates

```go
package main

import (
    "bytes"
    "os"
    "text/template"
)

func main() {
    // 1. Basic template
    tmpl := template.Must(template.New("greeting").Parse("Hello, {{.Name}}!"))
    tmpl.Execute(os.Stdout, map[string]string{"Name": "World"})
    
    // 2. Template with conditionals
    const emailTmpl = `
    Subject: Welcome {{.Name}}
    
    {{if .IsVIP}}
    Dear VIP Customer {{.Name}},
    {{else}}
    Dear {{.Name}},
    {{end}}
    
    Thank you for joining our platform.
    {{if .Discount}}
    Use code {{.Discount.Code}} for {{.Discount.Percent}}% off!
    {{end}}
    
    Best regards,
    The Team
    `
    
    data := struct {
        Name     string
        IsVIP    bool
        Discount *Discount
    }{
        Name:  "John Doe",
        IsVIP: true,
        Discount: &Discount{
            Code:    "WELCOME20",
            Percent: 20,
        },
    }
    
    tmpl = template.Must(template.New("email").Parse(emailTmpl))
    tmpl.Execute(os.Stdout, data)
    
    // 3. Template with loops
    const listTmpl = `
    <ul>
    {{range .Items}}
        <li>{{.}}</li>
    {{else}}
        <li>No items found</li>
    {{end}}
    </ul>
    `
    
    tmpl = template.Must(template.New("list").Parse(listTmpl))
    tmpl.Execute(os.Stdout, map[string][]string{
        "Items": {"Apple", "Banana", "Orange"},
    })
    
    // 4. Template with custom functions
    funcMap := template.FuncMap{
        "upper": strings.ToUpper,
        "lower": strings.ToLower,
        "formatDate": func(t time.Time) string {
            return t.Format("2006-01-02")
        },
        "multiply": func(a, b int) int {
            return a * b
        },
    }
    
    const funcTmpl = `
    Uppercase: {{.Name | upper}}
    Lowercase: {{.Name | lower}}
    Date: {{.CreatedAt | formatDate}}
    Total: {{multiply .Price .Quantity}}
    `
    
    tmpl = template.Must(template.New("func").Funcs(funcMap).Parse(funcTmpl))
    tmpl.Execute(os.Stdout, struct {
        Name      string
        CreatedAt time.Time
        Price     int
        Quantity  int
    }{
        Name:      "John Doe",
        CreatedAt: time.Now(),
        Price:     100,
        Quantity:  3,
    })
}
```

#### HTML Templates

```go
package main

import (
    "html/template"
    "net/http"
)

type PageData struct {
    Title       string
    User        *User
    Products    []Product
    IsLoggedIn  bool
    Flash       string
    CSRFToken   string
}

func renderTemplate(w http.ResponseWriter, tmplName string, data interface{}) {
    // Parse templates with layout
    tmpl := template.Must(template.ParseFiles(
        "templates/layout.html",
        "templates/"+tmplName+".html",
    ))
    
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    tmpl.ExecuteTemplate(w, "layout", data)
}

// templates/layout.html
/*
<!DOCTYPE html>
<html>
<head>
    <title>{{.Title}} - My App</title>
    <link rel="stylesheet" href="/static/css/style.css">
</head>
<body>
    <header>
        {{template "navbar" .}}
    </header>
    
    <main>
        {{if .Flash}}
        <div class="alert">{{.Flash}}</div>
        {{end}}
        
        {{template "content" .}}
    </main>
    
    <footer>
        &copy; 2024 My App
    </footer>
    
    <script src="/static/js/app.js"></script>
</body>
</html>
*/

// templates/navbar.html
/*
{{define "navbar"}}
<nav>
    <a href="/">Home</a>
    {{if .IsLoggedIn}}
        <a href="/profile">Profile</a>
        <a href="/logout">Logout</a>
    {{else}}
        <a href="/login">Login</a>
        <a href="/register">Register</a>
    {{end}}
</nav>
{{end}}
*/

// templates/home.html
/*
{{define "content"}}
<h1>Welcome, {{.User.Name}}!</h1>

<div class="products">
    <h2>Our Products</h2>
    <div class="product-grid">
        {{range .Products}}
        <div class="product-card">
            <h3>{{.Name}}</h3>
            <p>{{.Description}}</p>
            <div class="price">${{.Price}}</div>
            <form method="POST" action="/cart/add">
                <input type="hidden" name="csrf_token" value="{{$.CSRFToken}}">
                <input type="hidden" name="product_id" value="{{.ID}}">
                <input type="number" name="quantity" value="1" min="1">
                <button type="submit">Add to Cart</button>
            </form>
        </div>
        {{end}}
    </div>
</div>
{{end}}
*/
```

#### Template Functions and Helpers

```go
package main

import (
    "html/template"
    "strings"
    "time"
)

// Custom template functions
var templateFuncs = template.FuncMap{
    // String helpers
    "upper":      strings.ToUpper,
    "lower":      strings.ToLower,
    "title":      strings.Title,
    "trim":       strings.TrimSpace,
    "truncate":   truncate,
    
    // Number helpers
    "formatNumber": formatNumber,
    "currency":     formatCurrency,
    
    // Time helpers
    "formatDate":   formatDate,
    "formatTime":   formatTime,
    "timeAgo":      timeAgo,
    
    // Array helpers
    "has":          contains,
    "join":         strings.Join,
    "first":        first,
    "last":         last,
    
    // HTML helpers
    "safeHTML":     func(s string) template.HTML { return template.HTML(s) },
    "safeJS":       func(s string) template.JS { return template.JS(s) },
}

func truncate(s string, length int) string {
    if len(s) <= length {
        return s
    }
    return s[:length] + "..."
}

func formatNumber(n int) string {
    return strconv.FormatInt(int64(n), 10)
}

func formatCurrency(amount float64) string {
    return fmt.Sprintf("$%.2f", amount)
}

func formatDate(t time.Time) string {
    return t.Format("2006-01-02")
}

func formatTime(t time.Time) string {
    return t.Format("15:04:05")
}

func timeAgo(t time.Time) string {
    duration := time.Since(t)
    
    switch {
    case duration < time.Minute:
        return "just now"
    case duration < time.Hour:
        return fmt.Sprintf("%d minutes ago", int(duration.Minutes()))
    case duration < 24*time.Hour:
        return fmt.Sprintf("%d hours ago", int(duration.Hours()))
    default:
        return fmt.Sprintf("%d days ago", int(duration.Hours()/24))
    }
}

func contains(slice []string, item string) bool {
    for _, s := range slice {
        if s == item {
            return true
        }
    }
    return false
}

func first(slice []interface{}) interface{} {
    if len(slice) == 0 {
        return nil
    }
    return slice[0]
}

func last(slice []interface{}) interface{} {
    if len(slice) == 0 {
        return nil
    }
    return slice[len(slice)-1]
}

// Setup template with functions
func setupTemplates() *template.Template {
    return template.Must(template.New("").
        Funcs(templateFuncs).
        ParseGlob("templates/*.html"))
}
```

---

### บทที่ 33: การจัดการค่า Configuration

**เป้าหมาย:** จัดการ configuration หลายแหล่ง (file, env, flags)

#### Using Viper

```go
package config

import (
    "fmt"
    "strings"
    "time"
    
    "github.com/spf13/viper"
)

type Config struct {
    Server   ServerConfig   `mapstructure:"server"`
    Database DatabaseConfig `mapstructure:"database"`
    Redis    RedisConfig    `mapstructure:"redis"`
    JWT      JWTConfig      `mapstructure:"jwt"`
    Log      LogConfig      `mapstructure:"log"`
    SMTP     SMTPConfig     `mapstructure:"smtp"`
}

type ServerConfig struct {
    Port         int           `mapstructure:"port"`
    Mode         string        `mapstructure:"mode"`
    ReadTimeout  time.Duration `mapstructure:"read_timeout"`
    WriteTimeout time.Duration `mapstructure:"write_timeout"`
}

type DatabaseConfig struct {
    Host            string        `mapstructure:"host"`
    Port            int           `mapstructure:"port"`
    User            string        `mapstructure:"user"`
    Password        string        `mapstructure:"password"`
    Name            string        `mapstructure:"name"`
    SSLMode         string        `mapstructure:"ssl_mode"`
    MaxOpenConns    int           `mapstructure:"max_open_conns"`
    MaxIdleConns    int           `mapstructure:"max_idle_conns"`
    ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"`
}

type RedisConfig struct {
    Addr     string `mapstructure:"addr"`
    Password string `mapstructure:"password"`
    DB       int    `mapstructure:"db"`
    PoolSize int    `mapstructure:"pool_size"`
}

type JWTConfig struct {
    Secret        string        `mapstructure:"secret"`
    AccessExpiry  time.Duration `mapstructure:"access_expiry"`
    RefreshExpiry time.Duration `mapstructure:"refresh_expiry"`
}

type LogConfig struct {
    Level  string `mapstructure:"level"`
    Format string `mapstructure:"format"`
    Output string `mapstructure:"output"`
}

type SMTPConfig struct {
    Host     string `mapstructure:"host"`
    Port     int    `mapstructure:"port"`
    Username string `mapstructure:"username"`
    Password string `mapstructure:"password"`
    From     string `mapstructure:"from"`
}

func LoadConfig(configPath string) (*Config, error) {
    // Set defaults
    viper.SetDefault("server.port", 8080)
    viper.SetDefault("server.mode", "debug")
    viper.SetDefault("server.read_timeout", "15s")
    viper.SetDefault("server.write_timeout", "15s")
    viper.SetDefault("log.level", "info")
    viper.SetDefault("log.format", "json")
    viper.SetDefault("log.output", "stdout")
    
    // Config file
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(configPath)
    viper.AddConfigPath(".")
    viper.AddConfigPath("/etc/app/")
    
    // Environment variables
    viper.AutomaticEnv()
    viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
    
    // Read config file
    if err := viper.ReadInConfig(); err != nil {
        if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
            return nil, fmt.Errorf("read config: %w", err)
        }
    }
    
    // Unmarshal
    var cfg Config
    if err := viper.Unmarshal(&cfg); err != nil {
        return nil, fmt.Errorf("unmarshal config: %w", err)
    }
    
    return &cfg, nil
}
```

#### Environment Variables

```go
package config

import (
    "os"
    "strconv"
    "time"
)

type Config struct {
    // Server
    ServerPort int
    ServerMode string
    
    // Database
    DBHost     string
    DBPort     int
    DBUser     string
    DBPassword string
    DBName     string
    
    // Redis
    RedisAddr     string
    RedisPassword string
    RedisDB       int
    
    // JWT
    JWTSecret        string
    JWTAccessExpiry  time.Duration
    JWTRefreshExpiry time.Duration
}

func LoadFromEnv() (*Config, error) {
    cfg := &Config{
        // Server defaults
        ServerPort: 8080,
        ServerMode: "debug",
        
        // Database defaults
        DBHost: "localhost",
        DBPort: 5432,
        
        // Redis defaults
        RedisAddr: "localhost:6379",
        RedisDB:   0,
    }
    
    // Server
    if port := os.Getenv("SERVER_PORT"); port != "" {
        if p, err := strconv.Atoi(port); err == nil {
            cfg.ServerPort = p
        }
    }
    
    if mode := os.Getenv("SERVER_MODE"); mode != "" {
        cfg.ServerMode = mode
    }
    
    // Database
    if host := os.Getenv("DB_HOST"); host != "" {
        cfg.DBHost = host
    }
    
    if port := os.Getenv("DB_PORT"); port != "" {
        if p, err := strconv.Atoi(port); err == nil {
            cfg.DBPort = p
        }
    }
    
    cfg.DBUser = os.Getenv("DB_USER")
    cfg.DBPassword = os.Getenv("DB_PASSWORD")
    cfg.DBName = os.Getenv("DB_NAME")
    
    // Redis
    if addr := os.Getenv("REDIS_ADDR"); addr != "" {
        cfg.RedisAddr = addr
    }
    
    cfg.RedisPassword = os.Getenv("REDIS_PASSWORD")
    
    if db := os.Getenv("REDIS_DB"); db != "" {
        if d, err := strconv.Atoi(db); err == nil {
            cfg.RedisDB = d
        }
    }
    
    // JWT
    cfg.JWTSecret = os.Getenv("JWT_SECRET")
    
    if expiry := os.Getenv("JWT_ACCESS_EXPIRY"); expiry != "" {
        if d, err := time.ParseDuration(expiry); err == nil {
            cfg.JWTAccessExpiry = d
        }
    } else {
        cfg.JWTAccessExpiry = 15 * time.Minute
    }
    
    if expiry := os.Getenv("JWT_REFRESH_EXPIRY"); expiry != "" {
        if d, err := time.ParseDuration(expiry); err == nil {
            cfg.JWTRefreshExpiry = d
        }
    } else {
        cfg.JWTRefreshExpiry = 7 * 24 * time.Hour
    }
    
    // Validation
    if cfg.JWTSecret == "" {
        return nil, fmt.Errorf("JWT_SECRET is required")
    }
    
    return cfg, nil
}
```

#### Command Line Flags

```go
package main

import (
    "flag"
    "fmt"
    "time"
)

type Config struct {
    // Server
    ServerPort int
    ServerMode string
    
    // Database
    DBHost     string
    DBPort     int
    DBUser     string
    DBPassword string
    DBName     string
    
    // Redis
    RedisAddr     string
    RedisPassword string
    RedisDB       int
    
    // JWT
    JWTSecret        string
    JWTAccessExpiry  time.Duration
    JWTRefreshExpiry time.Duration
}

func LoadFromFlags() *Config {
    cfg := &Config{}
    
    // Server flags
    flag.IntVar(&cfg.ServerPort, "port", 8080, "Server port")
    flag.StringVar(&cfg.ServerMode, "mode", "debug", "Server mode (debug/release)")
    
    // Database flags
    flag.StringVar(&cfg.DBHost, "db-host", "localhost", "Database host")
    flag.IntVar(&cfg.DBPort, "db-port", 5432, "Database port")
    flag.StringVar(&cfg.DBUser, "db-user", "", "Database user")
    flag.StringVar(&cfg.DBPassword, "db-password", "", "Database password")
    flag.StringVar(&cfg.DBName, "db-name", "", "Database name")
    
    // Redis flags
    flag.StringVar(&cfg.RedisAddr, "redis-addr", "localhost:6379", "Redis address")
    flag.StringVar(&cfg.RedisPassword, "redis-password", "", "Redis password")
    flag.IntVar(&cfg.RedisDB, "redis-db", 0, "Redis database")
    
    // JWT flags
    flag.StringVar(&cfg.JWTSecret, "jwt-secret", "", "JWT secret key")
    flag.DurationVar(&cfg.JWTAccessExpiry, "jwt-access-expiry", 15*time.Minute, "JWT access token expiry")
    flag.DurationVar(&cfg.JWTRefreshExpiry, "jwt-refresh-expiry", 7*24*time.Hour, "JWT refresh token expiry")
    
    flag.Parse()
    
    return cfg
}
```

#### Complete Configuration Example

```go
package config

import (
    "fmt"
    "strings"
    
    "github.com/spf13/viper"
)

// config.yaml
/*
server:
  port: 8080
  mode: release
  read_timeout: 15s
  write_timeout: 15s

database:
  host: localhost
  port: 5432
  user: postgres
  password: ${DB_PASSWORD}
  name: myapp
  ssl_mode: disable
  max_open_conns: 25
  max_idle_conns: 25
  conn_max_lifetime: 5m

redis:
  addr: localhost:6379
  password: ${REDIS_PASSWORD}
  db: 0
  pool_size: 10

jwt:
  secret: ${JWT_SECRET}
  access_expiry: 15m
  refresh_expiry: 168h

log:
  level: info
  format: json
  output: stdout

smtp:
  host: smtp.gmail.com
  port: 587
  username: ${SMTP_USERNAME}
  password: ${SMTP_PASSWORD}
  from: noreply@myapp.com
*/

type Config struct {
    Server   ServerConfig   `mapstructure:"server"`
    Database DatabaseConfig `mapstructure:"database"`
    Redis    RedisConfig    `mapstructure:"redis"`
    JWT      JWTConfig      `mapstructure:"jwt"`
    Log      LogConfig      `mapstructure:"log"`
    SMTP     SMTPConfig     `mapstructure:"smtp"`
}

func InitConfig() (*Config, error) {
    // Set config file
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(".")
    viper.AddConfigPath("./config")
    viper.AddConfigPath("/etc/app")
    
    // Environment variables
    viper.AutomaticEnv()
    viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
    
    // Read config file
    if err := viper.ReadInConfig(); err != nil {
        return nil, fmt.Errorf("failed to read config: %w", err)
    }
    
    // Expand environment variables in config
    for _, key := range viper.AllKeys() {
        val := viper.GetString(key)
        if strings.Contains(val, "${") {
            expanded := os.ExpandEnv(val)
            viper.Set(key, expanded)
        }
    }
    
    var cfg Config
    if err := viper.Unmarshal(&cfg); err != nil {
        return nil, fmt.Errorf("failed to unmarshal config: %w", err)
    }
    
    // Validate required fields
    if cfg.JWT.Secret == "" {
        return nil, fmt.Errorf("JWT secret is required")
    }
    
    return &cfg, nil
}
```

---

## สรุปภาคที่ 4: ตัวอย่างแอปพลิเคชันครบวงจร

```go
// main.go - Complete application example
package main

import (
    "context"
    "log/slog"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
    
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "github.com/go-chi/cors"
    
    "myapp/config"
    "myapp/handlers"
    "myapp/logger"
    "myapp/repository"
    "myapp/service"
)

func main() {
    // 1. Load configuration
    cfg, err := config.InitConfig()
    if err != nil {
        panic(err)
    }
    
    // 2. Setup logger
    log, err := logger.NewLogger(cfg.Log)
    if err != nil {
        panic(err)
    }
    slog.SetDefault(log)
    
    slog.Info("Application starting", "config", cfg)
    
    // 3. Connect to database
    db, err := repository.NewPostgresDB(cfg.Database)
    if err != nil {
        slog.Error("Failed to connect to database", "error", err)
        os.Exit(1)
    }
    defer db.Close()
    
    // 4. Connect to Redis
    redisClient, err := repository.NewRedisClient(cfg.Redis)
    if err != nil {
        slog.Error("Failed to connect to Redis", "error", err)
        os.Exit(1)
    }
    defer redisClient.Close()
    
    // 5. Initialize repositories
    userRepo := repository.NewUserRepository(db)
    sessionRepo := repository.NewSessionRepository(redisClient)
    
    // 6. Initialize services
    authService := service.NewAuthService(userRepo, sessionRepo, cfg.JWT)
    userService := service.NewUserService(userRepo)
    
    // 7. Initialize handlers
    authHandler := handlers.NewAuthHandler(authService)
    userHandler := handlers.NewUserHandler(userService)
    healthHandler := handlers.NewHealthHandler(db, redisClient)
    
    // 8. Setup router
    r := chi.NewRouter()
    
    // Global middlewares
    r.Use(middleware.RequestID)
    r.Use(middleware.RealIP)
    r.Use(logger.HTTPMiddleware)
    r.Use(middleware.Recoverer)
    r.Use(middleware.Timeout(60 * time.Second))
    r.Use(cors.Handler(cors.Options{
        AllowedOrigins:   []string{"*"},
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
        ExposedHeaders:   []string{"Link"},
        AllowCredentials: false,
        MaxAge:           300,
    }))
    
    // Health check
    r.Get("/health", healthHandler.Health)
    r.Get("/ready", healthHandler.Ready)
    r.Get("/live", healthHandler.Live)
    
    // Public routes
    r.Post("/api/v1/auth/register", authHandler.Register)
    r.Post("/api/v1/auth/login", authHandler.Login)
    r.Post("/api/v1/auth/refresh", authHandler.Refresh)
    
    // Protected routes
    r.Group(func(r chi.Router) {
        r.Use(authHandler.AuthMiddleware)
        
        r.Get("/api/v1/users/me", userHandler.GetMe)
        r.Put("/api/v1/users/me", userHandler.UpdateMe)
        r.Post("/api/v1/auth/logout", authHandler.Logout)
    })
    
    // 9. Create server
    server := &http.Server{
        Addr:         fmt.Sprintf(":%d", cfg.Server.Port),
        Handler:      r,
        ReadTimeout:  cfg.Server.ReadTimeout,
        WriteTimeout: cfg.Server.WriteTimeout,
        IdleTimeout:  120 * time.Second,
    }
    
    // 10. Start server
    go func() {
        slog.Info("Server starting", "port", cfg.Server.Port)
        if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            slog.Error("Server failed", "error", err)
            os.Exit(1)
        }
    }()
    
    // 11. Graceful shutdown
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    
    slog.Info("Shutting down server...")
    
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    
    if err := server.Shutdown(ctx); err != nil {
        slog.Error("Server forced to shutdown", "error", err)
    }
    
    slog.Info("Server exited")
}
```

---

**สิ่งที่ได้เรียนรู้ในภาคที่ 4:**

1. **การรับข้อมูล (Input)**: HTTP server, JSON/XML, Enum/Bitmask, Time handling
2. **การประมวลผล (Processing)**: Anonymous functions, Closures, Concurrency patterns
3. **การจัดเก็บ (Storage)**: Files, Database, Structured logging
4. **การส่งออก (Output)**: HTML/Template, Configuration management

พื้นฐานเหล่านี้จะนำไปใช้ใน **ภาคที่ 5** เพื่อพัฒนาสู่การเป็น Go Developer มืออาชีพ ด้วยการวัดประสิทธิภาพ, profiling, context management, และ generics

## ภาคที่ 5: สู่การเป็นนักพัฒนา Go มืออาชีพ (บทที่ 34–42)
 
## แผนภาพแสดงความสัมพันธ์ของเนื้อหา

```mermaid
flowchart TB
    subgraph Performance["1. การวัดและปรับปรุงประสิทธิภาพ (Performance Optimization)"]
        direction TB
        A1["Benchmark (บทที่ 34)<br/>go test -bench<br/>benchmem, benchtime<br/>benchstat, comparing results"]
        A2["Profiling (บทที่ 36)<br/>pprof (CPU, Memory, Goroutine, Block, Mutex)<br/>go tool pprof, flame graphs"]
        A3["HTTP Client (บทที่ 35)<br/>connection pooling<br/>timeout, retry, transport"]
        A4["Context (บทที่ 37)<br/>deadline, cancellation<br/>request-scoped values, propagation"]
        A5["Generics (บทที่ 38)<br/>type parameters, constraints<br/>generic functions and types"]
        
        A1 --> A2
        A2 --> A3
        A3 --> A4
        A4 --> A5
    end

    subgraph CodeDesign["2. การออกแบบโค้ด (Code Design)"]
        direction TB
        B1["OOP in Go (บทที่ 39)<br/>struct + methods vs classes<br/>composition over inheritance<br/>embedding, interfaces"]
        B2["Version Management (บทที่ 40)<br/>upgrading/downgrading Go<br/>go mod edit, go fix, go tool"]
        B3["Best Practices (บทที่ 41)<br/>idioms, naming conventions<br/>error handling patterns<br/>project structure, testing"]
    end

    subgraph Reference["3. เอกสารอ้างอิง (Reference)"]
        C1["Cheatsheet (บทที่ 42)<br/>commands, syntax, patterns<br/>quick reference, code snippets"]
    end

    Performance --> CodeDesign
    CodeDesign --> Reference

    subgraph Tools["เครื่องมือที่ใช้"]
        D1["go test -bench<br/>go test -benchmem"]
        D2["pprof<br/>go tool pprof"]
        D3["benchstat<br/>benchcmp"]
        D4["go fix<br/>gofmt, go vet"]
        D5["golangci-lint<br/>staticcheck"]
    end

    Performance --> Tools
    CodeDesign --> Tools
```

---

## คำอธิบายภาษาไทย (แบบละเอียด)

ภาคที่ 5 ครอบคลุมบทที่ 34–42 โดยมีเนื้อหาแบ่งเป็น 3 ช่วงหลัก ตามแผนภาพด้านบน ซึ่งจะเปลี่ยนผู้พัฒนาจากผู้ใช้ภาษาให้เป็นผู้ที่เข้าใจกลไกภายใน สามารถวิเคราะห์และปรับปรุงระบบให้มีประสิทธิภาพสูง พร้อมเขียนโค้ดที่บำรุงรักษาง่ายและยืดหยุ่น

---

## 1. การวัดและปรับปรุงประสิทธิภาพ (Performance Optimization)

### บทที่ 34: การวัดประสิทธิภาพ (Benchmarks)

**เป้าหมาย:** เขียนและวิเคราะห์ benchmark เพื่อวัดประสิทธิภาพของโค้ด

#### พื้นฐานการเขียน Benchmark

```go
// math/benchmark_test.go
package math

import (
    "testing"
)

// 1. Basic benchmark
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(10, 20)
    }
}

// 2. Benchmark with different inputs
func BenchmarkAddLargeNumbers(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(1000000, 2000000)
    }
}

// 3. Benchmark with allocation tracking
func BenchmarkStringConcat(b *testing.B) {
    b.ReportAllocs() // report memory allocations
    
    for i := 0; i < b.N; i++ {
        s := ""
        for j := 0; j < 100; j++ {
            s += "a"
        }
    }
}

// 4. Benchmark with reset timer
func BenchmarkExpensiveSetup(b *testing.B) {
    // Setup that should not be measured
    data := make([]int, 1000000)
    for i := range data {
        data[i] = i
    }
    
    b.ResetTimer() // Reset benchmark timer
    
    for i := 0; i < b.N; i++ {
        processData(data)
    }
}

// 5. Sub-benchmarks
func BenchmarkSort(b *testing.B) {
    sizes := []int{100, 1000, 10000, 100000}
    
    for _, size := range sizes {
        b.Run(fmt.Sprintf("Size-%d", size), func(b *testing.B) {
            data := generateRandomSlice(size)
            
            b.ResetTimer()
            for i := 0; i < b.N; i++ {
                sort.Ints(data)
            }
        })
    }
}

// 6. Parallel benchmark
func BenchmarkParallel(b *testing.B) {
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            // Code to benchmark in parallel
            expensiveOperation()
        }
    })
}
```

#### การรัน Benchmark

```bash
# รัน benchmark ทั้งหมด
go test -bench=.

# รันเฉพาะ benchmark ที่ match pattern
go test -bench=Add

# ระบุจำนวนรอบ (default 1)
go test -bench=. -benchtime=10s

# วัด memory allocation
go test -bench=. -benchmem

# รัน benchmark และแสดง CPU profile
go test -bench=. -cpuprofile=cpu.prof

# รัน benchmark และแสดง memory profile
go test -bench=. -memprofile=mem.prof

# รัน benchmark และแสดง block profile
go test -bench=. -blockprofile=block.prof

# เปรียบเทียบผลลัพธ์ระหว่างสองเวอร์ชัน
go test -bench=. -count=5 > old.txt
# ... modify code ...
go test -bench=. -count=5 > new.txt
benchstat old.txt new.txt
```

#### ตัวอย่างการใช้งานจริง

**1. Benchmarking String Concatenation vs StringBuilder**

```go
package benchmark

import (
    "strings"
    "testing"
)

// Bad: String concatenation in loop
func ConcatWithPlus(items []string) string {
    result := ""
    for _, item := range items {
        result += item
    }
    return result
}

// Good: Using strings.Builder
func ConcatWithBuilder(items []string) string {
    var builder strings.Builder
    builder.Grow(1024) // Pre-allocate capacity
    
    for _, item := range items {
        builder.WriteString(item)
    }
    return builder.String()
}

// Good: Using strings.Join
func ConcatWithJoin(items []string) string {
    return strings.Join(items, "")
}

func BenchmarkConcatWithPlus(b *testing.B) {
    items := make([]string, 100)
    for i := range items {
        items[i] = "a"
    }
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        ConcatWithPlus(items)
    }
}

func BenchmarkConcatWithBuilder(b *testing.B) {
    items := make([]string, 100)
    for i := range items {
        items[i] = "a"
    }
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        ConcatWithBuilder(items)
    }
}

func BenchmarkConcatWithJoin(b *testing.B) {
    items := make([]string, 100)
    for i := range items {
        items[i] = "a"
    }
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        ConcatWithJoin(items)
    }
}
```

**ผลลัพธ์:**
```
BenchmarkConcatWithPlus-8        10000    150000 ns/op    100000 B/op    99 allocs/op
BenchmarkConcatWithBuilder-8   1000000       500 ns/op      1024 B/op     1 allocs/op
BenchmarkConcatWithJoin-8      2000000       300 ns/op      1024 B/op     1 allocs/op
```

**2. Benchmarking Map Access Patterns**

```go
package benchmark

import (
    "testing"
)

type User struct {
    ID   int
    Name string
    Age  int
}

// Using map[int]User (value)
func GetUserByIDValue(users map[int]User, id int) User {
    return users[id]
}

// Using map[int]*User (pointer)
func GetUserByIDPointer(users map[int]*User, id int) *User {
    return users[id]
}

func BenchmarkMapValue(b *testing.B) {
    users := make(map[int]User)
    for i := 0; i < 10000; i++ {
        users[i] = User{ID: i, Name: "User", Age: 30}
    }
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = GetUserByIDValue(users, i%10000)
    }
}

func BenchmarkMapPointer(b *testing.B) {
    users := make(map[int]*User)
    for i := 0; i < 10000; i++ {
        users[i] = &User{ID: i, Name: "User", Age: 30}
    }
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = GetUserByIDPointer(users, i%10000)
    }
}
```

**3. Benchmarking JSON Serialization**

```go
package benchmark

import (
    "encoding/json"
    "testing"
)

type Person struct {
    Name    string   `json:"name"`
    Age     int      `json:"age"`
    Email   string   `json:"email"`
    Address string   `json:"address"`
    Phone   string   `json:"phone"`
    Tags    []string `json:"tags"`
}

func BenchmarkJSONMarshal(b *testing.B) {
    p := Person{
        Name:    "John Doe",
        Age:     30,
        Email:   "john@example.com",
        Address: "123 Main St, Bangkok",
        Phone:   "081-234-5678",
        Tags:    []string{"golang", "developer", "backend"},
    }
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _, err := json.Marshal(p)
        if err != nil {
            b.Fatal(err)
        }
    }
}

func BenchmarkJSONUnmarshal(b *testing.B) {
    data := []byte(`{"name":"John Doe","age":30,"email":"john@example.com","address":"123 Main St, Bangkok","phone":"081-234-5678","tags":["golang","developer","backend"]}`)
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        var p Person
        err := json.Unmarshal(data, &p)
        if err != nil {
            b.Fatal(err)
        }
    }
}
```

**4. Benchmarking with benchstat**

```bash
# สร้าง benchmark results
go test -bench=Concat -count=5 > concat_old.txt

# แก้ไขโค้ดให้ดีขึ้น

# สร้าง benchmark results ใหม่
go test -bench=Concat -count=5 > concat_new.txt

# เปรียบเทียบ
benchstat concat_old.txt concat_new.txt
```

**ผลลัพธ์ benchstat:**
```
name               old time/op    new time/op    delta
ConcatWithPlus-8    150µs ± 2%     150µs ± 1%    ~     (p=0.841 n=5+5)
ConcatWithBuilder-8 500ns ± 1%     500ns ± 1%    ~     (p=0.421 n=5+5)
ConcatWithJoin-8    300ns ± 1%     300ns ± 1%    ~     (p=0.548 n=5+5)

name               old alloc/op   new alloc/op   delta
ConcatWithPlus-8    100kB ± 0%     100kB ± 0%    ~     (all equal)
ConcatWithBuilder-8 1.02kB ± 0%    1.02kB ± 0%   ~     (all equal)
ConcatWithJoin-8    1.02kB ± 0%    1.02kB ± 0%   ~     (all equal)

name               old allocs/op  new allocs/op  delta
ConcatWithPlus-8      99.0 ± 0%      99.0 ± 0%    ~     (all equal)
ConcatWithBuilder-8   1.00 ± 0%      1.00 ± 0%   ~     (all equal)
ConcatWithJoin-8      1.00 ± 0%      1.00 ± 0%   ~     (all equal)
```

---

### บทที่ 36: การวิเคราะห์โปรไฟล์ (Program Profiling)

**เป้าหมาย:** ใช้ pprof เพื่อวิเคราะห์ CPU, Memory, Goroutine, และการบล็อก

#### การเปิดใช้งาน pprof

```go
package main

import (
    "log"
    "net/http"
    _ "net/http/pprof" // Import for side effects
    "runtime"
)

func main() {
    // 1. Enable profiling endpoints
    go func() {
        log.Println(http.ListenAndServe("localhost:6060", nil))
    }()
    
    // 2. Start application
    // ... your application code
}
```

#### การเก็บ CPU Profile

```go
package main

import (
    "os"
    "runtime/pprof"
    "time"
)

func main() {
    // Create CPU profile file
    f, err := os.Create("cpu.prof")
    if err != nil {
        panic(err)
    }
    defer f.Close()
    
    // Start CPU profiling
    pprof.StartCPUProfile(f)
    defer pprof.StopCPUProfile()
    
    // Run code to profile
    heavyComputation()
}

func heavyComputation() {
    // Simulate heavy computation
    for i := 0; i < 10000000; i++ {
        _ = i * i
    }
}
```

#### การเก็บ Memory Profile

```go
package main

import (
    "os"
    "runtime/pprof"
)

func main() {
    // Create memory profile file
    f, err := os.Create("mem.prof")
    if err != nil {
        panic(err)
    }
    defer f.Close()
    
    // Run code to profile
    allocateMemory()
    
    // Write heap profile
    if err := pprof.WriteHeapProfile(f); err != nil {
        panic(err)
    }
}

func allocateMemory() {
    // Simulate memory allocation
    data := make([]byte, 1024*1024*100) // 100 MB
    for i := range data {
        data[i] = byte(i % 256)
    }
    _ = data
}
```

#### การวิเคราะห์ Profile ด้วย go tool pprof

```bash
# 1. วิเคราะห์ CPU profile
go tool pprof cpu.prof
# หรือ interactive mode
go tool pprof -http=:8080 cpu.prof

# 2. วิเคราะห์ Memory profile
go tool pprof mem.prof
go tool pprof -http=:8080 mem.prof

# 3. วิเคราะห์จาก running server
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30
go tool pprof http://localhost:6060/debug/pprof/heap

# 4. วิเคราะห์ goroutine
go tool pprof http://localhost:6060/debug/pprof/goroutine

# 5. วิเคราะห์ block
go tool pprof http://localhost:6060/debug/pprof/block

# 6. วิเคราะห์ mutex
go tool pprof http://localhost:6060/debug/pprof/mutex
```

#### คำสั่งใน pprof interactive mode

```
(pprof) top           # แสดง top functions
(pprof) top10         # แสดง top 10 functions
(pprof) list function # แสดง source code ของ function
(pprof) web           # แสดง flame graph (ต้องมี graphviz)
(pprof) peek          # แสดง call tree
(pprof) traces        # แสดง trace
(pprof) help          # แสดง help
```

#### ตัวอย่างการใช้งานจริง

**1. CPU Profiling for API Handler**

```go
package main

import (
    "fmt"
    "log"
    "net/http"
    _ "net/http/pprof"
    "runtime"
    "time"
)

func main() {
    // Enable pprof endpoints
    go func() {
        log.Println(http.ListenAndServe("localhost:6060", nil))
    }()
    
    // Start application
    http.HandleFunc("/api/process", processHandler)
    http.HandleFunc("/api/compute", computeHandler)
    
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func processHandler(w http.ResponseWriter, r *http.Request) {
    // CPU intensive operation
    result := 0
    for i := 0; i < 10000000; i++ {
        result += i * i
    }
    fmt.Fprintf(w, "Result: %d", result)
}

func computeHandler(w http.ResponseWriter, r *http.Request) {
    // Memory intensive operation
    data := make([]int, 1000000)
    for i := range data {
        data[i] = i
    }
    
    // Sort
    for i := 0; i < len(data); i++ {
        for j := i + 1; j < len(data); j++ {
            if data[i] > data[j] {
                data[i], data[j] = data[j], data[i]
            }
        }
    }
    
    fmt.Fprintf(w, "Computed %d items", len(data))
}
```

**2. Generating Flame Graph**

```bash
# เก็บ CPU profile 30 วินาที
go tool pprof -http=:8080 http://localhost:6060/debug/pprof/profile?seconds=30

# เก็บ memory profile และแสดง flame graph
go tool pprof -http=:8080 http://localhost:6060/debug/pprof/heap
```

**3. Analyzing Goroutine Leaks**

```go
package main

import (
    "fmt"
    "net/http"
    _ "net/http/pprof"
    "time"
)

func main() {
    go func() {
        http.ListenAndServe("localhost:6060", nil)
    }()
    
    // Goroutine leak example
    for i := 0; i < 100; i++ {
        go leakyGoroutine()
    }
    
    time.Sleep(time.Hour)
}

func leakyGoroutine() {
    ch := make(chan int)
    go func() {
        // This goroutine never exits
        <-ch
    }()
    // ch never gets a value, goroutine leaks
}
```

**ตรวจสอบ goroutine leak:**
```bash
# ดู goroutine profile
go tool pprof http://localhost:6060/debug/pprof/goroutine

# ใน pprof interactive mode
(pprof) top
(pprof) list leakyGoroutine
```

**4. Block Profile Analysis**

```go
package main

import (
    "net/http"
    _ "net/http/pprof"
    "runtime"
    "time"
)

func main() {
    // Enable block profiling
    runtime.SetBlockProfileRate(1)
    
    go func() {
        http.ListenAndServe("localhost:6060", nil)
    }()
    
    // Block example
    ch := make(chan int)
    go func() {
        time.Sleep(5 * time.Second)
        ch <- 42
    }()
    
    // This will block for 5 seconds
    <-ch
    
    select {}
}
```

**5. Mutex Profile Analysis**

```go
package main

import (
    "net/http"
    _ "net/http/pprof"
    "runtime"
    "sync"
)

var mu sync.Mutex
var counter int

func main() {
    // Enable mutex profiling
    runtime.SetMutexProfileFraction(1)
    
    go func() {
        http.ListenAndServe("localhost:6060", nil)
    }()
    
    // Create mutex contention
    for i := 0; i < 10; i++ {
        go worker()
    }
    
    select {}
}

func worker() {
    for {
        mu.Lock()
        counter++
        mu.Unlock()
    }
}
```

---

### บทที่ 35: สร้าง HTTP Client

**เป้าหมาย:** สร้าง HTTP client ที่มีประสิทธิภาพและเชื่อถือได้

#### Basic HTTP Client

```go
package main

import (
    "bytes"
    "context"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "time"
)

func main() {
    // 1. Basic GET request
    resp, err := http.Get("https://api.example.com/users")
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    
    body, _ := io.ReadAll(resp.Body)
    fmt.Println(string(body))
    
    // 2. POST request with JSON
    user := map[string]interface{}{
        "name":  "John Doe",
        "email": "john@example.com",
    }
    
    jsonData, _ := json.Marshal(user)
    
    resp, err = http.Post(
        "https://api.example.com/users",
        "application/json",
        bytes.NewBuffer(jsonData),
    )
    defer resp.Body.Close()
    
    // 3. Custom HTTP client with timeouts
    client := &http.Client{
        Timeout: 10 * time.Second,
        Transport: &http.Transport{
            MaxIdleConns:        100,
            MaxIdleConnsPerHost: 10,
            IdleConnTimeout:     90 * time.Second,
            TLSHandshakeTimeout: 10 * time.Second,
        },
    }
    
    req, _ := http.NewRequest("GET", "https://api.example.com/users", nil)
    req.Header.Set("Authorization", "Bearer token123")
    req.Header.Set("User-Agent", "MyApp/1.0")
    
    resp, err = client.Do(req)
    defer resp.Body.Close()
}
```

#### Advanced HTTP Client with Retry

```go
package client

import (
    "context"
    "fmt"
    "io"
    "net/http"
    "time"
)

type HTTPClient struct {
    client      *http.Client
    maxRetries  int
    retryDelay  time.Duration
    backoffFunc func(int) time.Duration
}

type Config struct {
    Timeout         time.Duration
    MaxRetries      int
    RetryDelay      time.Duration
    MaxIdleConns    int
    IdleConnTimeout time.Duration
}

func NewHTTPClient(cfg Config) *HTTPClient {
    transport := &http.Transport{
        MaxIdleConns:        cfg.MaxIdleConns,
        MaxIdleConnsPerHost: cfg.MaxIdleConns / 2,
        IdleConnTimeout:     cfg.IdleConnTimeout,
        DisableCompression:  false,
    }
    
    client := &http.Client{
        Timeout:   cfg.Timeout,
        Transport: transport,
    }
    
    return &HTTPClient{
        client:     client,
        maxRetries: cfg.MaxRetries,
        retryDelay: cfg.RetryDelay,
        backoffFunc: exponentialBackoff,
    }
}

func exponentialBackoff(attempt int) time.Duration {
    return time.Duration(1<<uint(attempt)) * time.Second
}

func (c *HTTPClient) Do(ctx context.Context, req *http.Request) (*http.Response, error) {
    var lastErr error
    
    for attempt := 0; attempt <= c.maxRetries; attempt++ {
        // Clone request for each attempt
        reqClone := req.Clone(ctx)
        
        // Execute request
        resp, err := c.client.Do(reqClone)
        
        // Success
        if err == nil && resp.StatusCode < 500 {
            return resp, nil
        }
        
        // Close body if exists
        if resp != nil {
            resp.Body.Close()
        }
        
        lastErr = err
        
        // Check if we should retry
        if !c.shouldRetry(err, resp) {
            break
        }
        
        // Wait before retry
        if attempt < c.maxRetries {
            delay := c.backoffFunc(attempt)
            select {
            case <-ctx.Done():
                return nil, ctx.Err()
            case <-time.After(delay):
            }
        }
    }
    
    return nil, fmt.Errorf("max retries exceeded: %w", lastErr)
}

func (c *HTTPClient) shouldRetry(err error, resp *http.Response) bool {
    if err != nil {
        // Network errors are retryable
        return true
    }
    
    if resp != nil {
        // Retry on server errors (5xx)
        if resp.StatusCode >= 500 && resp.StatusCode < 600 {
            return true
        }
        
        // Retry on rate limiting
        if resp.StatusCode == 429 {
            return true
        }
    }
    
    return false
}

// ใช้งาน
func main() {
    cfg := Config{
        Timeout:         30 * time.Second,
        MaxRetries:      3,
        RetryDelay:      1 * time.Second,
        MaxIdleConns:    100,
        IdleConnTimeout: 90 * time.Second,
    }
    
    client := NewHTTPClient(cfg)
    ctx := context.Background()
    
    req, _ := http.NewRequestWithContext(ctx, "GET", "https://api.example.com/data", nil)
    req.Header.Set("Authorization", "Bearer token")
    
    resp, err := client.Do(ctx, req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    
    body, _ := io.ReadAll(resp.Body)
    fmt.Println(string(body))
}
```

#### Circuit Breaker Pattern

```go
package circuitbreaker

import (
    "context"
    "errors"
    "sync"
    "time"
)

type State int

const (
    StateClosed State = iota
    StateOpen
    StateHalfOpen
)

type CircuitBreaker struct {
    mu              sync.RWMutex
    state           State
    failureCount    int
    lastFailureTime time.Time
    
    maxFailures     int
    timeout         time.Duration
    halfOpenMaxRequests int
    halfOpenRequests    int
}

func NewCircuitBreaker(maxFailures int, timeout time.Duration) *CircuitBreaker {
    return &CircuitBreaker{
        state:           StateClosed,
        maxFailures:     maxFailures,
        timeout:         timeout,
        halfOpenMaxRequests: 1,
    }
}

func (cb *CircuitBreaker) Call(fn func() error) error {
    if !cb.allow() {
        return errors.New("circuit breaker is open")
    }
    
    err := fn()
    
    cb.recordResult(err == nil)
    
    return err
}

func (cb *CircuitBreaker) allow() bool {
    cb.mu.Lock()
    defer cb.mu.Unlock()
    
    switch cb.state {
    case StateClosed:
        return true
        
    case StateOpen:
        if time.Since(cb.lastFailureTime) > cb.timeout {
            cb.state = StateHalfOpen
            cb.halfOpenRequests = 0
            return true
        }
        return false
        
    case StateHalfOpen:
        if cb.halfOpenRequests < cb.halfOpenMaxRequests {
            cb.halfOpenRequests++
            return true
        }
        return false
        
    default:
        return false
    }
}

func (cb *CircuitBreaker) recordResult(success bool) {
    cb.mu.Lock()
    defer cb.mu.Unlock()
    
    if success {
        switch cb.state {
        case StateHalfOpen:
            cb.state = StateClosed
            cb.failureCount = 0
        case StateOpen:
            // Should not happen
        default:
            cb.failureCount = 0
        }
        return
    }
    
    // Failure
    cb.failureCount++
    cb.lastFailureTime = time.Now()
    
    if cb.state == StateHalfOpen {
        cb.state = StateOpen
    } else if cb.state == StateClosed && cb.failureCount >= cb.maxFailures {
        cb.state = StateOpen
    }
}

// ใช้งาน
func main() {
    cb := NewCircuitBreaker(5, 30*time.Second)
    
    for i := 0; i < 20; i++ {
        err := cb.Call(func() error {
            // Make HTTP request
            resp, err := http.Get("https://api.example.com/data")
            if err != nil {
                return err
            }
            defer resp.Body.Close()
            
            if resp.StatusCode >= 500 {
                return errors.New("server error")
            }
            
            return nil
        })
        
        if err != nil {
            fmt.Printf("Request %d failed: %v\n", i, err)
        } else {
            fmt.Printf("Request %d succeeded\n", i)
        }
        
        time.Sleep(1 * time.Second)
    }
}
```

#### Connection Pool Management

```go
package client

import (
    "crypto/tls"
    "net"
    "net/http"
    "time"
)

func CreateOptimizedClient() *http.Client {
    // Custom transport for optimal performance
    transport := &http.Transport{
        // Connection pooling
        MaxIdleConns:        100,
        MaxIdleConnsPerHost: 10,
        MaxConnsPerHost:     20,
        IdleConnTimeout:     90 * time.Second,
        
        // Connection establishment
        DialContext: (&net.Dialer{
            Timeout:   30 * time.Second,
            KeepAlive: 30 * time.Second,
        }).DialContext,
        
        // TLS configuration
        TLSHandshakeTimeout: 10 * time.Second,
        TLSClientConfig: &tls.Config{
            MinVersion: tls.VersionTLS12,
        },
        
        // Response handling
        DisableCompression: false,
        DisableKeepAlives:  false,
        
        // Expect header handling
        ExpectContinueTimeout: 1 * time.Second,
        
        // Response header timeout
        ResponseHeaderTimeout: 10 * time.Second,
    }
    
    return &http.Client{
        Transport: transport,
        Timeout:   30 * time.Second,
        CheckRedirect: func(req *http.Request, via []*http.Request) error {
            if len(via) >= 10 {
                return http.ErrUseLastResponse
            }
            return nil
        },
    }
}
```

---

### บทที่ 37: การจัดการ Context

**เป้าหมาย:** ใช้ context สำหรับ timeout, cancellation, และ request-scoped values

#### Context Basics

```go
package main

import (
    "context"
    "fmt"
    "time"
)

func main() {
    // 1. Background context (root context)
    ctx := context.Background()
    
    // 2. With cancel
    ctx, cancel := context.WithCancel(ctx)
    defer cancel()
    
    // 3. With timeout
    ctx, cancel = context.WithTimeout(ctx, 5*time.Second)
    defer cancel()
    
    // 4. With deadline
    ctx, cancel = context.WithDeadline(ctx, time.Now().Add(5*time.Second))
    defer cancel()
    
    // 5. With value
    ctx = context.WithValue(ctx, "user_id", 123)
    ctx = context.WithValue(ctx, "request_id", "abc-123")
    
    // 6. Check cancellation
    select {
    case <-ctx.Done():
        fmt.Println("Context cancelled:", ctx.Err())
    default:
        fmt.Println("Context still active")
    }
    
    // 7. Pass context to function
    processRequest(ctx)
}

func processRequest(ctx context.Context) {
    // Extract values
    userID := ctx.Value("user_id")
    requestID := ctx.Value("request_id")
    
    fmt.Printf("Processing request %v for user %v\n", requestID, userID)
    
    // Check for cancellation
    select {
    case <-ctx.Done():
        fmt.Println("Request cancelled:", ctx.Err())
        return
    default:
    }
    
    // Do work with timeout
    done := make(chan bool)
    go func() {
        // Simulate work
        time.Sleep(2 * time.Second)
        done <- true
    }()
    
    select {
    case <-done:
        fmt.Println("Work completed")
    case <-ctx.Done():
        fmt.Println("Work cancelled:", ctx.Err())
    }
}
```

#### Context in HTTP Server

```go
package main

import (
    "context"
    "fmt"
    "net/http"
    "time"
)

type key int

const (
    userKey key = iota
    requestIDKey
)

// Middleware to add request ID
func requestIDMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        requestID := r.Header.Get("X-Request-ID")
        if requestID == "" {
            requestID = generateRequestID()
        }
        
        ctx := context.WithValue(r.Context(), requestIDKey, requestID)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

// Middleware to add timeout
func timeoutMiddleware(timeout time.Duration) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            ctx, cancel := context.WithTimeout(r.Context(), timeout)
            defer cancel()
            
            // Create channel for handler completion
            done := make(chan bool)
            
            go func() {
                next.ServeHTTP(w, r.WithContext(ctx))
                done <- true
            }()
            
            select {
            case <-done:
                // Handler completed
            case <-ctx.Done():
                // Timeout occurred
                http.Error(w, "Request timeout", http.StatusGatewayTimeout)
                return
            }
        })
    }
}

// Middleware to add user info
func authMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        
        // Validate token and get user
        user, err := validateToken(token)
        if err != nil {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        
        ctx := context.WithValue(r.Context(), userKey, user)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

// Handler with context
func userHandler(w http.ResponseWriter, r *http.Request) {
    // Extract values from context
    user := r.Context().Value(userKey).(*User)
    requestID := r.Context().Value(requestIDKey).(string)
    
    // Use context for database operations
    ctx := r.Context()
    
    // Query database with timeout
    data, err := fetchUserData(ctx, user.ID)
    if err != nil {
        if ctx.Err() == context.DeadlineExceeded {
            http.Error(w, "Database timeout", http.StatusGatewayTimeout)
        } else {
            http.Error(w, "Internal error", http.StatusInternalServerError)
        }
        return
    }
    
    fmt.Fprintf(w, "User: %+v, Request ID: %s\n", data, requestID)
}

func fetchUserData(ctx context.Context, userID int) (*UserData, error) {
    // Use context in database query
    rows, err := db.QueryContext(ctx, "SELECT * FROM users WHERE id = $1", userID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    // Process rows...
    return &UserData{}, nil
}
```

#### Context in Database Operations

```go
package repository

import (
    "context"
    "database/sql"
    "fmt"
    "time"
)

type UserRepository struct {
    db *sql.DB
}

func (r *UserRepository) GetUser(ctx context.Context, id int) (*User, error) {
    // Create context with timeout for this query
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()
    
    var user User
    
    // Execute query with context
    row := r.db.QueryRowContext(ctx, "SELECT id, name, email FROM users WHERE id = $1", id)
    
    err := row.Scan(&user.ID, &user.Name, &user.Email)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, ErrNotFound
        }
        return nil, err
    }
    
    return &user, nil
}

func (r *UserRepository) GetUsers(ctx context.Context, ids []int) ([]*User, error) {
    // Create context with timeout
    ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()
    
    // Use context with query
    query := "SELECT id, name, email FROM users WHERE id = ANY($1)"
    
    rows, err := r.db.QueryContext(ctx, query, ids)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var users []*User
    for rows.Next() {
        var user User
        if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
            return nil, err
        }
        users = append(users, &user)
    }
    
    return users, rows.Err()
}

// Transaction with context
func (r *UserRepository) TransferMoney(ctx context.Context, fromID, toID int, amount float64) error {
    // Start transaction
    tx, err := r.db.BeginTx(ctx, &sql.TxOptions{
        Isolation: sql.LevelSerializable,
    })
    if err != nil {
        return err
    }
    defer tx.Rollback()
    
    // Withdraw from source
    _, err = tx.ExecContext(ctx, 
        "UPDATE accounts SET balance = balance - $1 WHERE id = $2", 
        amount, fromID)
    if err != nil {
        return err
    }
    
    // Deposit to destination
    _, err = tx.ExecContext(ctx, 
        "UPDATE accounts SET balance = balance + $1 WHERE id = $2", 
        amount, toID)
    if err != nil {
        return err
    }
    
    return tx.Commit()
}
```

#### Context with External APIs

```go
package client

import (
    "context"
    "encoding/json"
    "fmt"
    "net/http"
    "time"
)

type APIClient struct {
    client *http.Client
    baseURL string
}

func NewAPIClient(baseURL string) *APIClient {
    return &APIClient{
        client: &http.Client{
            Timeout: 30 * time.Second,
        },
        baseURL: baseURL,
    }
}

func (c *APIClient) GetUser(ctx context.Context, id int) (*User, error) {
    // Create request with context
    url := fmt.Sprintf("%s/users/%d", c.baseURL, id)
    req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
    if err != nil {
        return nil, err
    }
    
    // Add headers
    req.Header.Set("Accept", "application/json")
    
    // Execute request
    resp, err := c.client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    
    // Check status
    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
    }
    
    // Decode response
    var user User
    if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
        return nil, err
    }
    
    return &user, nil
}

// Parallel requests with context
func (c *APIClient) GetMultipleUsers(ctx context.Context, ids []int) ([]*User, error) {
    results := make([]*User, len(ids))
    errors := make([]error, len(ids))
    
    // Create context with timeout
    ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()
    
    // Channel for results
    type result struct {
        index int
        user  *User
        err   error
    }
    
    resultCh := make(chan result, len(ids))
    
    // Launch goroutines for each request
    for i, id := range ids {
        go func(idx, userID int) {
            user, err := c.GetUser(ctx, userID)
            resultCh <- result{index: idx, user: user, err: err}
        }(i, id)
    }
    
    // Collect results
    for i := 0; i < len(ids); i++ {
        select {
        case res := <-resultCh:
            results[res.index] = res.user
            errors[res.index] = res.err
        case <-ctx.Done():
            return nil, ctx.Err()
        }
    }
    
    // Check for errors
    for _, err := range errors {
        if err != nil {
            return nil, err
        }
    }
    
    return results, nil
}
```

---

### บทที่ 38: Generics – การเขียนโค้ดแบบยืดหยุ่น

**เป้าหมาย:** ใช้ type parameters เพื่อสร้างโค้ดที่ reuse ได้และ type-safe

#### Basic Generics

```go
package main

import (
    "fmt"
    "constraints"
)

// 1. Generic function
func Identity[T any](x T) T {
    return x
}

// 2. Generic function with constraints
func Min[T constraints.Ordered](a, b T) T {
    if a < b {
        return a
    }
    return b
}

// 3. Generic type
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
    
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item, true
}

func (s *Stack[T]) Peek() (T, bool) {
    if len(s.items) == 0 {
        var zero T
        return zero, false
    }
    return s.items[len(s.items)-1], true
}

// 4. Generic with multiple type parameters
type Pair[K, V any] struct {
    Key   K
    Value V
}

func (p Pair[K, V]) String() string {
    return fmt.Sprintf("(%v, %v)", p.Key, p.Value)
}

func main() {
    // Using generic functions
    fmt.Println(Identity(42))
    fmt.Println(Identity("hello"))
    
    fmt.Println(Min(10, 20))
    fmt.Println(Min(3.14, 2.71))
    
    // Using generic types
    intStack := Stack[int]{}
    intStack.Push(1)
    intStack.Push(2)
    val, _ := intStack.Pop()
    fmt.Println(val) // 2
    
    stringStack := Stack[string]{}
    stringStack.Push("hello")
    stringStack.Push("world")
    
    // Using Pair
    p1 := Pair[int, string]{Key: 1, Value: "one"}
    p2 := Pair[string, bool]{Key: "active", Value: true}
    
    fmt.Println(p1)
    fmt.Println(p2)
}
```

#### Advanced Generic Constraints

```go
package main

import (
    "fmt"
    "math"
)

// 1. Custom constraint
type Number interface {
    ~int | ~int64 | ~float64 | ~float32
}

func Sum[T Number](values []T) T {
    var sum T
    for _, v := range values {
        sum += v
    }
    return sum
}

// 2. Interface constraint
type Stringer interface {
    String() string
}

func Print[T Stringer](t T) {
    fmt.Println(t.String())
}

// 3. Comparable constraint
func Contains[T comparable](slice []T, item T) bool {
    for _, v := range slice {
        if v == item {
            return true
        }
    }
    return false
}

// 4. Type constraints with methods
type Container[T any] interface {
    Add(item T)
    Get(index int) T
    Size() int
}

type SliceContainer[T any] struct {
    items []T
}

func (c *SliceContainer[T]) Add(item T) {
    c.items = append(c.items, item)
}

func (c *SliceContainer[T]) Get(index int) T {
    return c.items[index]
}

func (c *SliceContainer[T]) Size() int {
    return len(c.items)
}

// 5. Generic with type assertion
type JSONValue interface {
    ~string | ~int | ~float64 | ~bool | []any | map[string]any
}

func ToJSON[T JSONValue](v T) string {
    // Implementation would use type switch
    return fmt.Sprintf("%v", v)
}

// 6. Generic function with type parameters in methods
type Transformable[T any] interface {
    Transform(func(T) T) T
}

type NumberWrapper[T Number] struct {
    Value T
}

func (n NumberWrapper[T]) Transform(f func(T) T) T {
    return f(n.Value)
}
```

#### ตัวอย่างการใช้งานจริง

**1. Generic Repository**

```go
package repository

import (
    "context"
    "database/sql"
    "fmt"
)

type Repository[T any] interface {
    Get(ctx context.Context, id int) (*T, error)
    Create(ctx context.Context, entity *T) error
    Update(ctx context.Context, entity *T) error
    Delete(ctx context.Context, id int) error
    List(ctx context.Context, limit, offset int) ([]*T, error)
}

type GORMRepository[T any] struct {
    db *gorm.DB
}

func NewGORMRepository[T any](db *gorm.DB) *GORMRepository[T] {
    return &GORMRepository[T]{db: db}
}

func (r *GORMRepository[T]) Get(ctx context.Context, id int) (*T, error) {
    var entity T
    result := r.db.WithContext(ctx).First(&entity, id)
    if result.Error != nil {
        return nil, result.Error
    }
    return &entity, nil
}

func (r *GORMRepository[T]) Create(ctx context.Context, entity *T) error {
    return r.db.WithContext(ctx).Create(entity).Error
}

func (r *GORMRepository[T]) Update(ctx context.Context, entity *T) error {
    return r.db.WithContext(ctx).Save(entity).Error
}

func (r *GORMRepository[T]) Delete(ctx context.Context, id int) error {
    var entity T
    return r.db.WithContext(ctx).Delete(&entity, id).Error
}

func (r *GORMRepository[T]) List(ctx context.Context, limit, offset int) ([]*T, error) {
    var entities []*T
    result := r.db.WithContext(ctx).
        Limit(limit).
        Offset(offset).
        Find(&entities)
    return entities, result.Error
}

// ใช้งาน
type User struct {
    ID    int    `gorm:"primaryKey"`
    Name  string `gorm:"size:100"`
    Email string `gorm:"uniqueIndex;size:100"`
}

type Product struct {
    ID    int     `gorm:"primaryKey"`
    Name  string  `gorm:"size:100"`
    Price float64 `gorm:"type:decimal(10,2)"`
}

func main() {
    db := connectDB()
    
    userRepo := NewGORMRepository[User](db)
    productRepo := NewGORMRepository[Product](db)
    
    ctx := context.Background()
    
    // Use user repository
    user, _ := userRepo.Get(ctx, 1)
    
    // Use product repository
    products, _ := productRepo.List(ctx, 10, 0)
}
```

**2. Generic Cache**

```go
package cache

import (
    "context"
    "encoding/json"
    "time"
    
    "github.com/redis/go-redis/v9"
)

type Cache[T any] struct {
    client *redis.Client
    prefix string
    ttl    time.Duration
}

func NewCache[T any](client *redis.Client, prefix string, ttl time.Duration) *Cache[T] {
    return &Cache[T]{
        client: client,
        prefix: prefix,
        ttl:    ttl,
    }
}

func (c *Cache[T]) key(id string) string {
    return fmt.Sprintf("%s:%s", c.prefix, id)
}

func (c *Cache[T]) Get(ctx context.Context, id string) (*T, error) {
    data, err := c.client.Get(ctx, c.key(id)).Bytes()
    if err != nil {
        return nil, err
    }
    
    var result T
    if err := json.Unmarshal(data, &result); err != nil {
        return nil, err
    }
    
    return &result, nil
}

func (c *Cache[T]) Set(ctx context.Context, id string, value *T) error {
    data, err := json.Marshal(value)
    if err != nil {
        return err
    }
    
    return c.client.Set(ctx, c.key(id), data, c.ttl).Err()
}

func (c *Cache[T]) Delete(ctx context.Context, id string) error {
    return c.client.Del(ctx, c.key(id)).Err()
}

// ใช้งาน
type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
    })
    
    userCache := NewCache[User](rdb, "user", 10*time.Minute)
    
    ctx := context.Background()
    
    // Get user from cache
    user, err := userCache.Get(ctx, "123")
    if err == nil {
        fmt.Printf("User from cache: %+v\n", user)
    } else {
        // Get from database
        user = getUserFromDB(123)
        userCache.Set(ctx, "123", user)
    }
}
```

**3. Generic Worker Pool**

```go
package worker

import (
    "context"
    "sync"
)

type Job[T any] struct {
    ID     string
    Data   T
    Result chan Result[T]
}

type Result[T any] struct {
    JobID string
    Value T
    Error error
}

type WorkerPool[T any, R any] struct {
    numWorkers int
    jobQueue   chan Job[T]
    resultQueue chan Result[R]
    workerFunc func(Job[T]) Result[R]
    wg         sync.WaitGroup
    ctx        context.Context
    cancel     context.CancelFunc
}

func NewWorkerPool[T any, R any](numWorkers int, queueSize int, workerFunc func(Job[T]) Result[R]) *WorkerPool[T, R] {
    ctx, cancel := context.WithCancel(context.Background())
    
    return &WorkerPool[T, R]{
        numWorkers:  numWorkers,
        jobQueue:    make(chan Job[T], queueSize),
        resultQueue: make(chan Result[R], queueSize),
        workerFunc:  workerFunc,
        ctx:         ctx,
        cancel:      cancel,
    }
}

func (wp *WorkerPool[T, R]) Start() {
    for i := 0; i < wp.numWorkers; i++ {
        wp.wg.Add(1)
        go wp.worker()
    }
    
    go func() {
        wp.wg.Wait()
        close(wp.resultQueue)
    }()
}

func (wp *WorkerPool[T, R]) worker() {
    defer wp.wg.Done()
    
    for {
        select {
        case <-wp.ctx.Done():
            return
        case job, ok := <-wp.jobQueue:
            if !ok {
                return
            }
            
            result := wp.workerFunc(job)
            
            select {
            case wp.resultQueue <- result:
            case <-wp.ctx.Done():
                return
            }
        }
    }
}

func (wp *WorkerPool[T, R]) Submit(job Job[T]) {
    select {
    case wp.jobQueue <- job:
    case <-wp.ctx.Done():
    }
}

func (wp *WorkerPool[T, R]) Results() <-chan Result[R] {
    return wp.resultQueue
}

func (wp *WorkerPool[T, R]) Stop() {
    wp.cancel()
    close(wp.jobQueue)
    wp.wg.Wait()
}

// ใช้งาน
func main() {
    // Define worker function
    workerFunc := func(job Job[string]) Result[int] {
        // Process job
        length := len(job.Data)
        return Result[int]{
            JobID: job.ID,
            Value: length,
        }
    }
    
    // Create worker pool
    pool := NewWorkerPool[string, int](5, 100, workerFunc)
    pool.Start()
    
    // Submit jobs
    for i := 0; i < 100; i++ {
        job := Job[string]{
            ID:   fmt.Sprintf("job-%d", i),
            Data: fmt.Sprintf("data-%d", i),
        }
        pool.Submit(job)
    }
    
    // Collect results
    for result := range pool.Results() {
        fmt.Printf("Job %s result: %d\n", result.JobID, result.Value)
    }
    
    pool.Stop()
}
```

---

## 2. การออกแบบโค้ด (Code Design)

### บทที่ 39: Go กับกระบวนทัศน์ OOP?

**เป้าหมาย:** เข้าใจว่า Go จัดการกับ OOP อย่างไร (struct, method, interface, composition)

#### Struct และ Method (แทน Class)

```go
package main

import "fmt"

// 1. Struct (แทน class)
type Rectangle struct {
    Width  float64
    Height float64
}

// 2. Constructor pattern
func NewRectangle(width, height float64) *Rectangle {
    return &Rectangle{
        Width:  width,
        Height: height,
    }
}

// 3. Methods (value receiver)
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// 4. Methods (pointer receiver)
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}

// 5. Inheritance via composition (ไม่ใช่ inheritance แบบ OOP)
type ColoredRectangle struct {
    Rectangle  // Embedded struct (composition)
    Color      string
}

func (c ColoredRectangle) Describe() string {
    return fmt.Sprintf("Rectangle with color %s, area %.2f", 
        c.Color, c.Area())
}

// 6. Polymorphism via interfaces
type Shape interface {
    Area() float64
    Perimeter() float64
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * math.Pi * c.Radius
}

// 7. Interface composition
type Drawable interface {
    Draw() string
}

type ShapeWithDraw interface {
    Shape
    Drawable
}

func main() {
    // Create objects
    rect := NewRectangle(10, 5)
    fmt.Printf("Area: %.2f\n", rect.Area())
    
    rect.Scale(2)
    fmt.Printf("After scale: %.2f x %.2f\n", rect.Width, rect.Height)
    
    // Composition
    colored := ColoredRectangle{
        Rectangle: Rectangle{Width: 10, Height: 5},
        Color:     "red",
    }
    fmt.Println(colored.Describe())
    
    // Polymorphism
    shapes := []Shape{
        Rectangle{Width: 10, Height: 5},
        Circle{Radius: 7},
    }
    
    for _, shape := range shapes {
        fmt.Printf("Area: %.2f, Perimeter: %.2f\n", 
            shape.Area(), shape.Perimeter())
    }
}
```

#### Composition over Inheritance

```go
package main

import "fmt"

// 1. Base struct (ไม่ใช่ parent class)
type Animal struct {
    Name string
    Age  int
}

func (a Animal) Speak() string {
    return "Some sound"
}

func (a Animal) Info() string {
    return fmt.Sprintf("Name: %s, Age: %d", a.Name, a.Age)
}

// 2. Composition via embedding
type Dog struct {
    Animal    // Embedded (has-a relationship)
    Breed string
}

// Override method (not override, just new method)
func (d Dog) Speak() string {
    return "Woof!"
}

// 3. Composition with interface
type Flyer interface {
    Fly() string
}

type Swimmer interface {
    Swim() string
}

type Duck struct {
    Name string
}

func (d Duck) Fly() string {
    return "Flying"
}

func (d Duck) Swim() string {
    return "Swimming"
}

// 4. Composing multiple behaviors
type SuperDuck struct {
    Duck
    Flyer
    Swimmer
}

// 5. Dependency injection (better than inheritance)
type Database interface {
    Save(data string) error
}

type Logger interface {
    Log(message string)
}

type UserService struct {
    db     Database
    logger Logger
}

func NewUserService(db Database, logger Logger) *UserService {
    return &UserService{
        db:     db,
        logger: logger,
    }
}

func (s *UserService) CreateUser(name string) error {
    s.logger.Log(fmt.Sprintf("Creating user: %s", name))
    return s.db.Save(name)
}

// 6. Functional options pattern (OOP alternative)
type Server struct {
    host    string
    port    int
    timeout time.Duration
    maxConn int
}

type Option func(*Server)

func WithHost(host string) Option {
    return func(s *Server) {
        s.host = host
    }
}

func WithPort(port int) Option {
    return func(s *Server) {
        s.port = port
    }
}

func WithTimeout(timeout time.Duration) Option {
    return func(s *Server) {
        s.timeout = timeout
    }
}

func NewServer(opts ...Option) *Server {
    // Default values
    s := &Server{
        host:    "localhost",
        port:    8080,
        timeout: 30 * time.Second,
        maxConn: 100,
    }
    
    // Apply options
    for _, opt := range opts {
        opt(s)
    }
    
    return s
}
```

#### Interface Design Principles

```go
package main

import "fmt"

// 1. Small interfaces (Interface Segregation Principle)
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

type Closer interface {
    Close() error
}

// 2. Composed interfaces
type ReadWriter interface {
    Reader
    Writer
}

type ReadWriteCloser interface {
    Reader
    Writer
    Closer
}

// 3. Accept interfaces, return structs
type User struct {
    ID   int
    Name string
}

type UserRepository interface {
    GetByID(id int) (*User, error)
    Save(user *User) error
}

// Implementation returns struct, not interface
type PostgresUserRepository struct {
    db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
    return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) GetByID(id int) (*User, error) {
    // implementation
    return &User{}, nil
}

func (r *PostgresUserRepository) Save(user *User) error {
    // implementation
    return nil
}

// 4. Interface for testing
type MockUserRepository struct {
    users map[int]*User
}

func (m *MockUserRepository) GetByID(id int) (*User, error) {
    if user, ok := m.users[id]; ok {
        return user, nil
    }
    return nil, ErrNotFound
}
```

---

### บทที่ 40: การอัปเกรดหรือดาวน์เกรดเวอร์ชัน Go

**เป้าหมาย:** จัดการเวอร์ชัน Go ในโปรเจกต์อย่างถูกต้อง

#### การอัปเกรด Go เวอร์ชัน

```bash
# 1. ตรวจสอบเวอร์ชันปัจจุบัน
go version

# 2. ดาวน์โหลดเวอร์ชันใหม่
# ไปที่ https://go.dev/dl/
# หรือใช้ brew (macOS)
brew upgrade go

# 3. ตรวจสอบเวอร์ชันใหม่
go version

# 4. อัปเกรดโมดูล
go mod edit -go=1.22

# 5. อัปเกรด dependencies
go get -u ./...

# 6. ใช้ go fix สำหรับการเปลี่ยนแปลง API
go fix ./...

# 7. ตรวจสอบโค้ด
go vet ./...

# 8. รันทดสอบ
go test ./...
```

#### การจัดการหลายเวอร์ชันด้วย gvm

```bash
# 1. ติดตั้ง gvm
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)

# 2. รายการเวอร์ชันที่สามารถติดตั้งได้
gvm listall

# 3. ติดตั้งเวอร์ชันต่างๆ
gvm install go1.20
gvm install go1.21
gvm install go1.22

# 4. ใช้เวอร์ชันเฉพาะ
gvm use go1.22 --default

# 5. สร้าง environment สำหรับโปรเจกต์
gvm pkgset create myproject
gvm pkgset use myproject
```

#### การอัปเกรดโค้ดสำหรับ Go เวอร์ชันใหม่

```go
// Go 1.18+ (Generics)
// Old code (before 1.18)
func MinInt(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func MinFloat(a, b float64) float64 {
    if a < b {
        return a
    }
    return b
}

// New code with generics
func Min[T constraints.Ordered](a, b T) T {
    if a < b {
        return a
    }
    return b
}

// Go 1.21+ (slog)
// Old code (using log)
log.Printf("User %d logged in", userID)

// New code with slog
slog.Info("User logged in", "user_id", userID)

// Go 1.22+ (range over integers)
// Old code
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// New code (Go 1.22+)
for i := range 10 {
    fmt.Println(i)
}
```

#### Compatibility Check

```go
// ใช้ build tags สำหรับเวอร์ชันเฉพาะ
// +build go1.22

package main

import "fmt"

func main() {
    // โค้ดที่ใช้เฉพาะ Go 1.22+
    for i := range 10 {
        fmt.Println(i)
    }
}
```

```go
// ใช้ runtime.Version() ตรวจสอบ
package main

import (
    "fmt"
    "runtime"
    "strconv"
    "strings"
)

func main() {
    version := runtime.Version()
    fmt.Printf("Go version: %s\n", version)
    
    // ตรวจสอบ major version
    if strings.HasPrefix(version, "go1.22") {
        fmt.Println("Using Go 1.22+ features")
    }
}

// ใช้ build constraints
//go:build go1.22
package main
```

---

### บทที่ 41: คำแนะนำในการออกแบบโค้ดที่ดี

**เป้าหมาย:** เรียนรู้ best practices และ idioms ของ Go

#### Naming Conventions

```go
// 1. Package names (lowercase, short)
package user   // good
package userService // bad (too long)

// 2. Exported names (PascalCase)
type UserService struct {}  // exported
type userRepository struct {} // unexported

// 3. Un-exported names (camelCase)
var dbConnection *sql.DB
func getUserByID(id int) {}

// 4. Acronyms (keep case)
type HTTPClient struct {}  // good
type HttpClient struct {}  // bad

// 5. Interface naming
type Reader interface {}    // -er suffix
type Writer interface {}
type Stringer interface {}

// 6. Getters (no Get prefix)
type User struct {
    name string
}

func (u *User) Name() string {  // good
    return u.name
}

func (u *User) GetName() string {  // bad (in Go)
    return u.name
}
```

#### Error Handling Patterns

```go
// 1. Check errors immediately
func readFile(filename string) ([]byte, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        return nil, fmt.Errorf("read file %s: %w", filename, err)
    }
    return data, nil
}

// 2. Wrap errors with context
func processRequest(req *http.Request) error {
    user, err := getUser(req.Context(), req.Header.Get("User-ID"))
    if err != nil {
        return fmt.Errorf("get user: %w", err)
    }
    
    if err := validateUser(user); err != nil {
        return fmt.Errorf("validate user %d: %w", user.ID, err)
    }
    
    return nil
}

// 3. Custom error types
type ValidationError struct {
    Field   string
    Value   interface{}
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation failed for %s: %s", e.Field, e.Message)
}

// 4. Sentinel errors
var (
    ErrNotFound     = errors.New("not found")
    ErrUnauthorized = errors.New("unauthorized")
    ErrInvalidInput = errors.New("invalid input")
)

// 5. Error handling in defer
func doSomething() (err error) {
    f, err := os.Open("file.txt")
    if err != nil {
        return err
    }
    defer func() {
        if closeErr := f.Close(); closeErr != nil && err == nil {
            err = closeErr
        }
    }()
    
    // process file
    return nil
}
```

#### Code Organization

```go
// 1. Project structure
/*
myapp/
├── cmd/
│   ├── api/
│   │   └── main.go
│   └── worker/
│       └── main.go
├── internal/
│   ├── config/
│   ├── domain/
│   │   ├── user/
│   │   │   ├── entity.go
│   │   │   ├── repository.go
│   │   │   └── service.go
│   │   └── order/
│   ├── infrastructure/
│   │   ├── db/
│   │   ├── cache/
│   │   └── queue/
│   └── interfaces/
│       ├── http/
│       └── grpc/
├── pkg/
│   ├── logger/
│   └── utils/
├── test/
├── go.mod
└── go.sum
*/

// 2. Function ordering in files
package user

import ()

// Constants
const ()

// Variables
var ()

// Types
type Service struct {}

// Constructor
func NewService() *Service {
    return &Service{}
}

// Public methods
func (s *Service) GetUser(id int) (*User, error) {}

// Private methods
func (s *Service) validateUser(user *User) error {}

// Helper functions
func generateID() string {}
```

#### Concurrency Patterns

```go
// 1. Use channels for communication
func worker(jobs <-chan int, results chan<- int) {
    for job := range jobs {
        results <- process(job)
    }
}

// 2. Use sync.WaitGroup for coordination
func processBatch(items []Item) {
    var wg sync.WaitGroup
    for _, item := range items {
        wg.Add(1)
        go func(item Item) {
            defer wg.Done()
            process(item)
        }(item)
    }
    wg.Wait()
}

// 3. Use context for cancellation
func queryWithTimeout(ctx context.Context, query string) error {
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()
    
    return db.QueryContext(ctx, query)
}

// 4. Don't start goroutines without knowing when to stop
type Worker struct {
    stop chan struct{}
    wg   sync.WaitGroup
}

func (w *Worker) Start() {
    w.wg.Add(1)
    go func() {
        defer w.wg.Done()
        for {
            select {
            case <-w.stop:
                return
            default:
                // do work
            }
        }
    }()
}

func (w *Worker) Stop() {
    close(w.stop)
    w.wg.Wait()
}
```

#### Performance Tips

```go
// 1. Preallocate slices
// Bad
var result []int
for i := 0; i < 1000000; i++ {
    result = append(result, i)
}

// Good
result := make([]int, 0, 1000000)
for i := 0; i < 1000000; i++ {
    result = append(result, i)
}

// 2. Use strings.Builder for concatenation
// Bad
s := ""
for i := 0; i < 1000; i++ {
    s += "a"
}

// Good
var builder strings.Builder
builder.Grow(1000)
for i := 0; i < 1000; i++ {
    builder.WriteString("a")
}
s := builder.String()

// 3. Use sync.Pool for temporary objects
var bufferPool = sync.Pool{
    New: func() interface{} {
        return new(bytes.Buffer)
    },
}

func process() {
    buf := bufferPool.Get().(*bytes.Buffer)
    defer bufferPool.Put(buf)
    buf.Reset()
    
    // use buf
}

// 4. Avoid unnecessary allocations
// Bad
func getID() string {
    return fmt.Sprintf("%d", id)
}

// Good
var idBuffer [20]byte
func getID() string {
    return strconv.Itoa(id)
}
```

---

## 3. เอกสารอ้างอิง (Reference)

### บทที่ 42: ชีทสรุป (Cheatsheet)

#### Go Commands

```bash
# Development
go mod init <module>     # Initialize module
go mod tidy              # Add missing dependencies
go mod vendor            # Create vendor directory
go build                 # Build binary
go run <file>            # Run program
go test ./...            # Run tests
go test -bench .         # Run benchmarks
go test -cover           # Test coverage
go fmt ./...             # Format code
go vet ./...             # Static analysis
go generate ./...        # Generate code

# Tools
go get <package>         # Download package
go install <package>     # Install binary
go list -m all           # List dependencies
go mod graph             # Show dependency graph
go mod why <package>     # Why is package needed

# Profiling
go test -cpuprofile cpu.prof
go test -memprofile mem.prof
go test -blockprofile block.prof
go tool pprof cpu.prof
```

#### Common Patterns

```go
// Type declarations
type MyInt int
type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}
type Reader interface {
    Read(p []byte) (n int, err error)
}

// Error handling
if err != nil {
    return fmt.Errorf("operation failed: %w", err)
}

// Defer
defer func() {
    if r := recover(); r != nil {
        log.Printf("recovered: %v", r)
    }
}()

// Channels
ch := make(chan int, 10)
go func() { ch <- 42 }()
value := <-ch

// Select
select {
case <-ch1:
    // handle
case <-ch2:
    // handle
case <-time.After(5 * time.Second):
    // timeout
default:
    // non-blocking
}

// Context
ctx := context.Background()
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel()
```

#### Useful Functions

```go
// Strings
strings.Contains(s, substr)
strings.HasPrefix(s, prefix)
strings.HasSuffix(s, suffix)
strings.Join(slice, sep)
strings.Split(s, sep)
strings.ReplaceAll(s, old, new)
strings.ToLower(s)
strings.ToUpper(s)
strings.TrimSpace(s)

// Slices
len(slice)
cap(slice)
append(slice, elem...)
copy(dst, src)
make([]T, len, cap)

// Maps
len(map)
make(map[K]V)
delete(map, key)
val, ok := map[key]

// Time
time.Now()
time.Parse(layout, value)
t.Format(layout)
t.Add(duration)
t.Sub(t2)
time.Sleep(duration)
time.After(duration)

// JSON
json.Marshal(v)
json.Unmarshal(data, &v)
json.NewEncoder(w).Encode(v)
json.NewDecoder(r).Decode(&v)
```

#### HTTP Server Patterns

```go
// Basic server
srv := &http.Server{
    Addr:         ":8080",
    Handler:      handler,
    ReadTimeout:  15 * time.Second,
    WriteTimeout: 15 * time.Second,
}
srv.ListenAndServe()

// Graceful shutdown
quit := make(chan os.Signal, 1)
signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
<-quit

ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()
srv.Shutdown(ctx)

// Middleware
func Logger(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s", r.Method, r.URL.Path)
        next.ServeHTTP(w, r)
    })
}
```

---

## สรุปภาคที่ 5: เส้นทางสู่มืออาชีพ

```mermaid
flowchart LR
    subgraph Stage1["ขั้นที่ 1: วัดผล"]
        A1[Benchmark]
        A2[Profiling]
    end
    
    subgraph Stage2["ขั้นที่ 2: ปรับปรุง"]
        B1[Optimize HTTP Client]
        B2[Use Context]
        B3[Apply Generics]
    end
    
    subgraph Stage3["ขั้นที่ 3: ออกแบบ"]
        C1[Apply Go Idioms]
        C2[Composition over Inheritance]
        C3[Small Interfaces]
    end
    
    subgraph Stage4["ขั้นที่ 4: เอกสาร"]
        D1[Cheatsheet]
        D2[Code Review Checklist]
    end
    
    Stage1 --> Stage2
    Stage2 --> Stage3
    Stage3 --> Stage4
```

**สิ่งที่ได้เรียนรู้ในภาคที่ 5:**

1. **การวัดประสิทธิภาพ (บทที่ 34, 36)**: 
   - เขียน benchmark และวิเคราะห์ผล
   - ใช้ pprof สำหรับ CPU, Memory, Goroutine profiling
   - เปรียบเทียบประสิทธิภาพระหว่างเวอร์ชัน

2. **การปรับแต่ง (บทที่ 35, 37, 38)**:
   - สร้าง HTTP client ที่มีประสิทธิภาพ
   - ใช้ context สำหรับ timeout และ cancellation
   - ใช้ generics สำหรับโค้ดที่ reuse ได้

3. **การออกแบบ (บทที่ 39, 40, 41)**:
   - เข้าใจ OOP ใน Go (struct, method, interface, composition)
   - จัดการเวอร์ชัน Go
   - ใช้ best practices และ idioms

4. **เอกสารอ้างอิง (บทที่ 42)**:
   - Cheatsheet สำหรับการพัฒนารวดเร็ว

พื้นฐานเหล่านี้จะนำไปใช้ใน **ภาคที่ 6** เพื่อเรียนรู้เครื่องมือและไลบรารียอดนิยม เช่น chi, viper, cobra, zap, GORM

## ภาคที่ 6: เครื่องมือและไลบรารียอดนิยม (บทที่ 43–45)

## แผนภาพ Data Flow (Flowchart TB)

```mermaid
flowchart TB
    subgraph Entry["จุดเริ่มต้น (Entry Points)"]
        A1["CLI Application<br/>cobra"]
        A2["HTTP Server<br/>chi"]
        A3["Configuration<br/>viper"]
    end

    subgraph Core["แกนหลัก (Core Layer)"]
        B1["Routing & Middleware<br/>chi"]
        B2["Business Logic<br/>Custom Handlers"]
        B3["Validation<br/>go-playground/validator"]
        B4["Authentication<br/>golang-jwt/jwt"]
    end

    subgraph DataLayer["ชั้นข้อมูล (Data Layer)"]
        C1["ORM<br/>GORM"]
        C2["Database<br/>PostgreSQL/MySQL"]
        C3["Migration<br/>AutoMigrate / golang-migrate"]
        C4["Query Builder<br/>GORM"]
    end

    subgraph Output["ชั้นส่งออก (Output Layer)"]
        D1["Logging<br/>zap"]
        D2["Email<br/>gomail + hermes"]
        D3["Response<br/>JSON/XML"]
    end

    subgraph External["ระบบภายนอก (External Systems)"]
        E1["SMTP Server<br/>Gmail / SendGrid"]
        E2["Database Server"]
        E3["File System"]
    end

    A1 --> B1
    A2 --> B1
    A3 --> B1
    B1 --> B2
    B2 --> B3
    B3 --> B4
    B4 --> C1
    C1 --> C2
    C1 --> C3
    C1 --> C4
    C2 --> E2
    C3 --> E2
    B2 --> D1
    B2 --> D2
    D2 --> E1
    B2 --> D3
    D3 --> A2

    subgraph Tools["เครื่องมือพัฒนา (Development Tools)"]
        F1["air<br/>(hot reload)"]
        F2["golangci-lint<br/>(linter)"]
        F3["dlv<br/>(debugger)"]
        F4["go mod<br/>(dependency)"]
    end

    A1 -.-> F1
    A2 -.-> F1
    B1 -.-> F2
    B2 -.-> F2
    C1 -.-> F3
    A1 -.-> F4
    A2 -.-> F4
```

---

## คำอธิบายภาษาไทย (แบบละเอียด)

ภาคที่ 6 ครอบคลุมบทที่ 43–45 โดยนำเสนอเครื่องมือและไลบรารียอดนิยมที่ใช้ในงานพัฒนา Go ระดับ Production แต่ละตัวมีหน้าที่และตำแหน่งในสถาปัตยกรรมของแอปพลิเคชัน ดังแสดงในแผนภาพ Data Flow ด้านบน

### 1. chi – เราเตอร์และมิดเดิลแวร์ (บทที่ 43)

**บทบาท:** รับ HTTP request, ส่งต่อผ่าน middleware chain, จับคู่ route, และส่งไปยัง handler ที่เหมาะสม

**โครงสร้างข้อมูล:** chi ใช้ `chi.Router` ซึ่งเป็น interface ที่สืบทอดจาก `http.Handler` มีความสามารถในการเพิ่ม middleware, สร้าง route group, และจับคู่ pattern เช่น `/users/{id}`

**ตัวอย่างการใช้งานจริง:**

```go
package main

import (
    "context"
    "encoding/json"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "github.com/go-chi/cors"
    "github.com/go-chi/httprate"
)

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

var users = map[int]User{
    1: {ID: 1, Name: "John"},
}

func main() {
    r := chi.NewRouter()

    // Global middleware
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)
    r.Use(middleware.RealIP)
    r.Use(middleware.Timeout(60 * time.Second))

    // CORS
    r.Use(cors.Handler(cors.Options{
        AllowedOrigins:   []string{"*"},
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
        AllowCredentials: false,
        MaxAge:           300,
    }))

    // Rate limiting
    r.Use(httprate.LimitByIP(100, 1*time.Minute))

    // Public routes
    r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("OK"))
    })

    // API v1
    r.Route("/api/v1", func(r chi.Router) {
        r.Get("/users", listUsers)
        r.Get("/users/{id}", getUser)
        r.Post("/users", createUser)
    })

    // Start server with graceful shutdown
    srv := &http.Server{
        Addr:         ":8080",
        Handler:      r,
        ReadTimeout:  15 * time.Second,
        WriteTimeout: 15 * time.Second,
    }

    go func() {
        log.Println("Server starting on :8080")
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("Server failed: %v", err)
        }
    }()

    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    log.Println("Shutting down...")

    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    srv.Shutdown(ctx)
}

func listUsers(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    // ... implementation
}

func createUser(w http.ResponseWriter, r *http.Request) {
    // ... implementation
}
```

---

### 2. viper – การจัดการ Configuration (บทที่ 43)

**บทบาท:** โหลดค่าตั้งค่าจากหลายแหล่ง (ไฟล์, environment variables, flags) และให้ API สำหรับอ่านค่า

**โครงสร้างข้อมูล:** viper ใช้ `map[string]interface{}` ภายใน เก็บค่าจากทุกแหล่ง โดยมีลำดับความสำคัญ: flags > env > config file > defaults

**ตัวอย่างการใช้งานจริง:**

```go
package config

import (
    "fmt"
    "strings"
    "time"

    "github.com/spf13/viper"
)

type Config struct {
    Server   ServerConfig
    Database DatabaseConfig
    JWT      JWTConfig
}

type ServerConfig struct {
    Port int    `mapstructure:"port"`
    Mode string `mapstructure:"mode"`
}

type DatabaseConfig struct {
    Host     string `mapstructure:"host"`
    Port     int    `mapstructure:"port"`
    User     string `mapstructure:"user"`
    Password string `mapstructure:"password"`
    Name     string `mapstructure:"name"`
}

type JWTConfig struct {
    Secret string        `mapstructure:"secret"`
    Expiry time.Duration `mapstructure:"expiry"`
}

func Load() (*Config, error) {
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(".")
    viper.AddConfigPath("./config")

    viper.AutomaticEnv()
    viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

    // Defaults
    viper.SetDefault("server.port", 8080)
    viper.SetDefault("server.mode", "debug")
    viper.SetDefault("jwt.expiry", "15m")

    if err := viper.ReadInConfig(); err != nil {
        if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
            return nil, fmt.Errorf("read config: %w", err)
        }
    }

    var cfg Config
    if err := viper.Unmarshal(&cfg); err != nil {
        return nil, fmt.Errorf("unmarshal: %w", err)
    }
    return &cfg, nil
}
```

**ไฟล์ config.yaml:**
```yaml
server:
  port: 9090
database:
  host: localhost
  port: 5432
  user: postgres
  password: ${DB_PASSWORD}   # จาก environment
  name: myapp
jwt:
  secret: ${JWT_SECRET}
```

---

### 3. cobra – การสร้าง CLI (บทที่ 43)

**บทบาท:** จัดโครงสร้างคำสั่ง (commands), flags, arguments, และ generate help/documentation

**โครงสร้างข้อมูล:** แต่ละคำสั่งคือ `cobra.Command` ที่มี fields เช่น Use, Short, Run, Flags

**ตัวอย่างการใช้งานจริง:**

```go
package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "myapp",
    Short: "MyApp CLI",
    Long:  `A sample CLI application built with cobra.`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Hello from MyApp")
    },
}

var serveCmd = &cobra.Command{
    Use:   "serve",
    Short: "Start the HTTP server",
    RunE: func(cmd *cobra.Command, args []string) error {
        port, _ := cmd.Flags().GetInt("port")
        fmt.Printf("Starting server on :%d\n", port)
        // start server...
        return nil
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

```go
// main.go
package main

import "myapp/cmd"

func main() {
    cmd.Execute()
}
```

---

### 4. zap – Structured Logging (บทที่ 43)

**บทบาท:** บันทึก log แบบมีโครงสร้าง (JSON) ด้วยประสิทธิภาพสูง รองรับระดับความรุนแรง และการแยก output

**โครงสร้างข้อมูล:** zap ใช้ `zap.Logger` ซึ่งมี methods เช่น `Info`, `Error` พร้อม fields ที่กำหนดด้วย `zap.String`, `zap.Int` เป็นต้น

**ตัวอย่างการใช้งานจริง:**

```go
package logger

import (
    "os"

    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

var Log *zap.Logger
var Sugar *zap.SugaredLogger

func Init(level string, output string) error {
    var lvl zapcore.Level
    switch level {
    case "debug":
        lvl = zapcore.DebugLevel
    case "info":
        lvl = zapcore.InfoLevel
    case "warn":
        lvl = zapcore.WarnLevel
    case "error":
        lvl = zapcore.ErrorLevel
    default:
        lvl = zapcore.InfoLevel
    }

    encoderConfig := zapcore.EncoderConfig{
        TimeKey:        "time",
        LevelKey:       "level",
        NameKey:        "logger",
        CallerKey:      "caller",
        MessageKey:     "msg",
        StacktraceKey:  "stacktrace",
        LineEnding:     zapcore.DefaultLineEnding,
        EncodeLevel:    zapcore.LowercaseLevelEncoder,
        EncodeTime:     zapcore.ISO8601TimeEncoder,
        EncodeDuration: zapcore.SecondsDurationEncoder,
        EncodeCaller:   zapcore.ShortCallerEncoder,
    }

    var ws zapcore.WriteSyncer
    if output == "stdout" {
        ws = zapcore.AddSync(os.Stdout)
    } else {
        // file output
        ws = zapcore.AddSync(&lumberjack.Logger{
            Filename:   output,
            MaxSize:    100,
            MaxBackups: 3,
            MaxAge:     28,
            Compress:   true,
        })
    }

    core := zapcore.NewCore(
        zapcore.NewJSONEncoder(encoderConfig),
        ws,
        lvl,
    )
    Log = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
    Sugar = Log.Sugar()
    return nil
}
```

**การใช้งานใน HTTP middleware:**

```go
func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        rw := &responseWriter{ResponseWriter: w, status: http.StatusOK}
        next.ServeHTTP(rw, r)
        duration := time.Since(start)

        Log.Info("HTTP request",
            zap.String("method", r.Method),
            zap.String("path", r.URL.Path),
            zap.Int("status", rw.status),
            zap.Duration("duration", duration),
            zap.String("ip", r.RemoteAddr),
        )
    })
}
```

---

### 5. GORM – ORM ทรงพลัง (บทที่ 44)

**บทบาท:** แมป struct กับตารางฐานข้อมูล, จัดการ CRUD, transactions, associations, และ migrations

**โครงสร้างข้อมูล:** ใช้ `gorm.DB` เป็น session สำหรับ query, `gorm.Model` สำหรับ fields พื้นฐาน (ID, CreatedAt, UpdatedAt, DeletedAt)

**ตัวอย่างการใช้งานจริง (model และ repository):**

```go
// models/user.go
package models

import (
    "time"

    "gorm.io/gorm"
)

type User struct {
    ID        uint           `gorm:"primaryKey" json:"id"`
    Username  string         `gorm:"size:50;uniqueIndex;not null" json:"username"`
    Email     string         `gorm:"size:100;uniqueIndex;not null" json:"email"`
    Password  string         `gorm:"size:255;not null" json:"-"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// repositories/user_repo.go
package repositories

import (
    "context"
    "errors"

    "gorm.io/gorm"
    "myapp/models"
)

type UserRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
    return r.db.WithContext(ctx).Create(user).Error
}

func (r *UserRepository) GetByID(ctx context.Context, id uint) (*models.User, error) {
    var user models.User
    err := r.db.WithContext(ctx).First(&user, id).Error
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, nil
    }
    return &user, err
}

func (r *UserRepository) Update(ctx context.Context, user *models.User) error {
    return r.db.WithContext(ctx).Save(user).Error
}
```

**การใช้งาน transaction พร้อม locking:**

```go
func (r *OrderRepository) CreateOrder(ctx context.Context, order *models.Order) error {
    return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
        // Lock product rows
        for _, item := range order.Items {
            var product models.Product
            if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&product, item.ProductID).Error; err != nil {
                return err
            }
            if product.Stock < item.Quantity {
                return errors.New("insufficient stock")
            }
            product.Stock -= item.Quantity
            if err := tx.Save(&product).Error; err != nil {
                return err
            }
        }
        if err := tx.Create(order).Error; err != nil {
            return err
        }
        return nil
    })
}
```

---

### 6. gomail + hermes – การส่งอีเมล (บทที่ 45)

**บทบาท:** สร้างและส่งอีเมลแบบ HTML ที่สวยงาม พร้อม fallback text

**โครงสร้างข้อมูล:** 
- `gomail.Dialer` สำหรับเชื่อมต่อ SMTP
- `gomail.Message` สำหรับสร้างอีเมล (From, To, Subject, Body, Attachments)
- `hermes.Hermes` สำหรับสร้าง HTML email จาก template

**ตัวอย่างการใช้งานจริง:**

```go
package email

import (
    "bytes"
    "fmt"
    "html/template"
    "time"

    "github.com/matcornic/hermes/v2"
    "gopkg.in/gomail.v2"
)

type Config struct {
    SMTPHost     string
    SMTPPort     int
    SMTPUser     string
    SMTPPassword string
    FromEmail    string
    FromName     string
    AppURL       string
    AppName      string
    AppLogo      string
}

type EmailService struct {
    dialer  *gomail.Dialer
    hermes  hermes.Hermes
    from    string
}

func New(cfg Config) *EmailService {
    dialer := gomail.NewDialer(cfg.SMTPHost, cfg.SMTPPort, cfg.SMTPUser, cfg.SMTPPassword)

    h := hermes.Hermes{
        Product: hermes.Product{
            Name: cfg.AppName,
            Link: cfg.AppURL,
            Logo: cfg.AppLogo,
        },
        Theme: new(hermes.Flat),
    }

    return &EmailService{
        dialer: dialer,
        hermes: h,
        from:   fmt.Sprintf("%s <%s>", cfg.FromName, cfg.FromEmail),
    }
}

func (s *EmailService) SendWelcome(to, name, verifyURL string) error {
    email := hermes.Email{
        Body: hermes.Body{
            Name: name,
            Intros: []string{
                fmt.Sprintf("Welcome to %s!", s.hermes.Product.Name),
                "We're excited to have you on board.",
            },
            Actions: []hermes.Action{
                {
                    Instructions: "Please verify your email address:",
                    Button: hermes.Button{
                        Color: "#22BC66",
                        Text:  "Verify Email",
                        Link:  verifyURL,
                    },
                },
            },
            Outros: []string{
                "If you didn't create an account, ignore this email.",
            },
        },
    }

    html, err := s.hermes.GenerateHTML(email)
    if err != nil {
        return err
    }
    text := fmt.Sprintf("Welcome %s! Please verify: %s", name, verifyURL)

    return s.send(to, "Welcome to "+s.hermes.Product.Name, html, text)
}

func (s *EmailService) send(to, subject, htmlBody, textBody string) error {
    m := gomail.NewMessage()
    m.SetHeader("From", s.from)
    m.SetHeader("To", to)
    m.SetHeader("Subject", subject)
    m.SetBody("text/plain", textBody)
    m.AddAlternative("text/html", htmlBody)

    return s.dialer.DialAndSend(m)
}
```

**การใช้งานใน handler:**

```go
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
    // ... create user ...
    verifyURL := fmt.Sprintf("%s/verify?token=%s", h.config.AppURL, token)
    if err := h.emailService.SendWelcome(user.Email, user.Name, verifyURL); err != nil {
        log.Printf("Failed to send welcome email: %v", err)
        // ไม่ fail การ register
    }
    // ...
}
```

---

## เทมเพลตและตัวอย่างโค้ดเพิ่มเติม

### เทมเพลตโครงสร้างโปรเจกต์สำหรับภาคนี้

```
myapp/
├── cmd/
│   ├── api/
│   │   └── main.go          # ใช้ chi
│   └── cli/
│       └── main.go          # ใช้ cobra
├── internal/
│   ├── config/
│   │   └── config.go        # viper
│   ├── handlers/
│   │   └── user.go
│   ├── models/
│   │   └── user.go
│   ├── repository/
│   │   └── user.go          # GORM
│   └── logger/
│       └── logger.go        # zap
├── pkg/
│   └── email/
│       └── email.go         # gomail + hermes
├── config.yaml
├── go.mod
└── go.sum
```

### Checklist การเลือกใช้เครื่องมือ

| เครื่องมือ | เมื่อควรใช้ | ทางเลือก |
|-----------|-----------|---------|
| chi | ต้องการ router น้ำหนักเบา เร็ว รองรับ middleware มาตรฐาน | gin, echo |
| viper | ต้องการ config ที่ยืดหยุ่น (file, env, flags) | envconfig, koanf |
| cobra | สร้าง CLI ที่มีหลายคำสั่ง | urfave/cli |
| zap | ต้องการ logging ที่เร็วมาก ใช้ structured | slog, logrus |
| GORM | ต้องการ ORM ที่ครบฟีเจอร์, associations | sqlx, sqlc |
| gomail + hermes | ต้องการส่งอีเมล HTML สวยงาม | sendgrid, mailgun |

---

## สรุป

ภาคที่ 6 นำเสนอชุดเครื่องมือที่จำเป็นสำหรับการพัฒนาแอปพลิเคชัน Go ในระดับ Production:

- **chi** จัดการ HTTP routing และ middleware อย่างมีประสิทธิภาพ
- **viper** จัดการ configuration จากหลายแหล่ง
- **cobra** สร้าง CLI ที่มีโครงสร้างชัดเจน
- **zap** บันทึก structured logs ด้วยความเร็วสูง
- **GORM** จัดการฐานข้อมูลผ่าน ORM ที่ทรงพลัง
- **gomail + hermes** สร้างและส่งอีเมล HTML ที่สวยงาม

เครื่องมือเหล่านี้ทำงานสอดคล้องกันตามแผนภาพ Data Flow ที่แสดงไว้ ช่วยให้การพัฒนาเป็นระบบ มีประสิทธิภาพ และบำรุงรักษาง่าย

## ภาคที่ 7: การออกแบบสถาปัตยกรรมและ Workflow (บทที่ 46–48)

## แผนภาพ Data Flow (Flowchart TB)

```mermaid
flowchart TB
    subgraph Delivery["Delivery Layer (Interfaces)"]
        D1["HTTP Handlers<br/>chi/gin handlers"]
        D2["Middleware<br/>Auth, Logger, CORS"]
        D3["DTOs & Validation<br/>request/response structs"]
    end

    subgraph Usecase["Usecase Layer (Business Logic)"]
        U1["Use Case Implementations<br/>RegisterUser, CreateOrder"]
        U2["Domain Services<br/>TransferService"]
        U3["Repository Interfaces<br/>UserRepository, OrderRepository"]
    end

    subgraph Repository["Repository Layer (Data Access)"]
        R1["Repository Implementations<br/>PostgresUserRepo, RedisCache"]
        R2["External API Clients<br/>PaymentGateway, SMS"]
        R3["Cache Layer<br/>Redis"]
    end

    subgraph Infrastructure["Infrastructure Layer"]
        I1["Database<br/>PostgreSQL, MySQL"]
        I2["Message Broker<br/>RabbitMQ, Kafka"]
        I3["External Systems<br/>SMS, Email"]
    end

    subgraph Models["Models (Domain Entities)"]
        M1["Entities<br/>User, Order, Product"]
        M2["Value Objects<br/>Email, Money, Address"]
        M3["Domain Events<br/>UserRegistered, OrderPlaced"]
    end

    Delivery --> Usecase
    Usecase --> Repository
    Repository --> Infrastructure
    Models -.-> Delivery
    Models -.-> Usecase
    Models -.-> Repository

    subgraph Blueprint["Production Blueprint"]
        B1["cmd/<br/>api/, worker/"]
        B2["internal/<br/>core/, application/, infrastructure/"]
        B3["pkg/<br/>shared utils"]
        B4["configs/, migrations/, deploy/"]
    end

    subgraph Workflow["Development Workflow"]
        W1["Task List Template"]
        W2["Checklist Template"]
        W3["Feature Development Flow<br/>Analyze → Design → Implement → Test → Deploy"]
    end

    Blueprint --> Delivery
    Blueprint --> Usecase
    Blueprint --> Repository
    Workflow --> Blueprint
```

---

## คำอธิบายภาษาไทย (แบบละเอียด)

ภาคที่ 7 ประกอบด้วย 3 บทหลัก:
- **บทที่ 46: Clean Architecture และโครงสร้างโปรเจกต์**
- **บทที่ 47: Blueprint สำหรับโปรเจกต์ Go ระดับ Production**
- **บทที่ 48: การออกแบบ Workflow และ Task Management**

ทั้งสามบทเชื่อมโยงกันเพื่อให้ผู้อ่านสามารถออกแบบและพัฒนาแอปพลิเคชันที่มีสถาปัตยกรรมแข็งแรง พร้อมขยาย และมีกระบวนการพัฒนาที่เป็นระบบ

---

## บทที่ 46: Clean Architecture และโครงสร้างโปรเจกต์

### หลักการ Clean Architecture

Clean Architecture แบ่งระบบออกเป็นชั้น ๆ ที่มีความรับผิดชอบชัดเจน โดยมีกฎสำคัญคือ **Dependency Rule**: dependencies ต้องชี้เข้าหาศูนย์กลาง (จากภายนอกเข้าสู่ภายใน) เท่านั้น ชั้นในสุดคือ **Entities** (Domain Models) ซึ่งไม่ควรขึ้นกับอะไรภายนอก

**ชั้นต่าง ๆ (จากนอกสุดเข้าสู่ในสุด):**

1. **Delivery (Interface Adapters)**
   - รับคำขอจากภายนอก (HTTP, gRPC, CLI)
   - แปลงข้อมูลเป็น DTO
   - เรียก Usecase
   - ส่ง Response กลับ

2. **Usecase (Application Business Rules)**
   - จัดลำดับการทำงาน (orchestration)
   - ใช้ Repository Interfaces
   - ไม่ขึ้นกับรายละเอียดของฐานข้อมูลหรือเฟรมเวิร์ก

3. **Repository (Interface)**
   - กำหนด contract สำหรับการเข้าถึงข้อมูล
   - ถูก Implement ในชั้น Infrastructure

4. **Entities (Domain Models)**
   - โมเดลหลักทางธุรกิจ (User, Order)
   - Value Objects (Email, Money)
   - Domain Events

5. **Infrastructure**
   - การ Implement จริงของ Repository (GORM, SQL)
   - การเชื่อมต่อกับ External Services
   - Cache, Queue, etc.

### ตัวอย่างโครงสร้างโปรเจกต์ตาม Clean Architecture

```
myapp/
├── cmd/
│   └── api/
│       └── main.go                 # entry point
├── internal/
│   ├── config/                     # โหลด configuration
│   ├── domain/                     # Entities, Value Objects, Events
│   │   ├── user/
│   │   │   ├── entity.go
│   │   │   ├── value_objects.go
│   │   │   ├── events.go
│   │   │   └── repository.go       # interface
│   │   └── order/
│   │       └── ...
│   ├── application/                # Usecases
│   │   ├── user/
│   │   │   ├── register.go
│   │   │   ├── login.go
│   │   │   └── dto.go
│   │   └── order/
│   │       └── ...
│   ├── infrastructure/             # Implementations
│   │   ├── persistence/
│   │   │   ├── gorm/
│   │   │   │   ├── user_repo.go
│   │   │   │   └── models.go
│   │   │   └── redis/
│   │   ├── queue/
│   │   └── email/
│   └── interfaces/                 # Delivery
│       ├── http/
│       │   ├── handlers/
│       │   ├── middleware/
│       │   └── routes.go
│       └── cli/
├── pkg/                            # reusable packages
├── configs/
├── migrations/
└── go.mod
```

### ตัวอย่างโค้ดของแต่ละชั้น

#### 1. Domain Layer (Entity)

```go
// internal/domain/user/entity.go
package user

import (
    "errors"
    "time"
)

type User struct {
    id        uint
    email     Email
    password  Password
    name      string
    isActive  bool
    createdAt time.Time
    updatedAt time.Time
    events    []DomainEvent
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
    return &User{
        id:        generateID(),
        email:     *emailVO,
        password:  *passwordVO,
        name:      name,
        isActive:  true,
        createdAt: time.Now(),
        updatedAt: time.Now(),
        events:    []DomainEvent{},
    }, nil
}

// Getters
func (u *User) ID() uint         { return u.id }
func (u *User) Email() string    { return u.email.String() }
func (u *User) Name() string     { return u.name }
func (u *User) IsActive() bool   { return u.isActive }

// Business logic
func (u *User) Deactivate() error {
    if !u.isActive {
        return errors.New("user already inactive")
    }
    u.isActive = false
    u.updatedAt = time.Now()
    u.addDomainEvent(NewUserDeactivatedEvent(u.id))
    return nil
}

func (u *User) addDomainEvent(event DomainEvent) {
    u.events = append(u.events, event)
}

func (u *User) Events() []DomainEvent { return u.events }
func (u *User) ClearEvents()          { u.events = nil }
```

#### 2. Repository Interface

```go
// internal/domain/user/repository.go
package user

import "context"

type Repository interface {
    Save(ctx context.Context, user *User) error
    FindByID(ctx context.Context, id uint) (*User, error)
    FindByEmail(ctx context.Context, email string) (*User, error)
    Update(ctx context.Context, user *User) error
    Delete(ctx context.Context, id uint) error
}
```

#### 3. Application Layer (Usecase)

```go
// internal/application/user/register.go
package user

import (
    "context"
    "fmt"

    "myapp/internal/domain/user"
)

type RegisterUseCase struct {
    userRepo user.Repository
    eventBus EventBus
}

type RegisterInput struct {
    Email    string `json:"email"`
    Password string `json:"password"`
    Name     string `json:"name"`
}

type RegisterOutput struct {
    ID    uint   `json:"id"`
    Email string `json:"email"`
    Name  string `json:"name"`
}

func NewRegisterUseCase(repo user.Repository, bus EventBus) *RegisterUseCase {
    return &RegisterUseCase{
        userRepo: repo,
        eventBus: bus,
    }
}

func (uc *RegisterUseCase) Execute(ctx context.Context, input RegisterInput) (*RegisterOutput, error) {
    // ตรวจสอบ email ซ้ำ
    existing, _ := uc.userRepo.FindByEmail(ctx, input.Email)
    if existing != nil {
        return nil, fmt.Errorf("email already registered")
    }

    // สร้าง User entity
    u, err := user.NewUser(input.Email, input.Password, input.Name)
    if err != nil {
        return nil, err
    }

    // บันทึก
    if err := uc.userRepo.Save(ctx, u); err != nil {
        return nil, err
    }

    // Dispatch domain events
    for _, event := range u.Events() {
        uc.eventBus.Publish(event)
    }

    return &RegisterOutput{
        ID:    u.ID(),
        Email: u.Email(),
        Name:  u.Name(),
    }, nil
}
```

#### 4. Infrastructure Layer (Repository Implementation with GORM)

```go
// internal/infrastructure/persistence/gorm/user_repo.go
package gormrepo

import (
    "context"

    "gorm.io/gorm"
    "myapp/internal/domain/user"
)

type userRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.Repository {
    return &userRepository{db: db}
}

func (r *userRepository) Save(ctx context.Context, u *user.User) error {
    model := &UserModel{
        ID:        u.ID(),
        Email:     u.Email(),
        Name:      u.Name(),
        IsActive:  u.IsActive(),
        CreatedAt: u.CreatedAt,
        UpdatedAt: u.UpdatedAt,
    }
    return r.db.WithContext(ctx).Create(model).Error
}

// ... other methods
```

#### 5. Delivery Layer (HTTP Handler)

```go
// internal/interfaces/http/handlers/user.go
package handlers

import (
    "encoding/json"
    "net/http"

    "myapp/internal/application/user"
)

type UserHandler struct {
    registerUC *user.RegisterUseCase
}

func NewUserHandler(registerUC *user.RegisterUseCase) *UserHandler {
    return &UserHandler{registerUC: registerUC}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
    var input user.RegisterInput
    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    output, err := h.registerUC.Execute(r.Context(), input)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(output)
}
```

---

## บทที่ 47: Blueprint สำหรับโปรเจกต์ Go ระดับ Production

### โครงสร้างโฟลเดอร์หลัก (Blueprint)

```
project-root/
├── cmd/                         # Executable entry points
│   ├── api/
│   │   └── main.go              # REST API server
│   └── worker/
│       └── main.go              # Background worker
├── internal/                    # Private code
│   ├── config/                  # Configuration loading (viper)
│   ├── domain/                  # Domain models (DDD)
│   │   ├── user/
│   │   │   ├── entity.go
│   │   │   ├── value_objects.go
│   │   │   ├── repository.go    # interface
│   │   │   └── events.go
│   │   └── order/
│   ├── application/             # Usecases
│   │   ├── user/
│   │   │   ├── register.go
│   │   │   ├── login.go
│   │   │   └── dto.go
│   │   └── order/
│   ├── infrastructure/          # Implementations
│   │   ├── persistence/
│   │   │   ├── gorm/
│   │   │   │   ├── user_repo.go
│   │   │   │   └── models.go
│   │   │   └── redis/
│   │   ├── queue/               # RabbitMQ, Kafka
│   │   ├── cache/
│   │   └── email/
│   └── interfaces/              # Delivery
│       ├── http/
│       │   ├── handlers/
│       │   ├── middleware/
│       │   └── routes.go
│       └── cli/
├── pkg/                         # Reusable packages
│   ├── logger/                  # zap wrapper
│   ├── jwt/
│   ├── validator/
│   └── errors/
├── api/                         # API definitions (OpenAPI, Protobuf)
│   └── openapi.yaml
├── configs/                     # Configuration files
│   ├── config.yaml
│   └── config.example.yaml
├── migrations/                  # Database migrations
│   └── 001_create_users_table.sql
├── scripts/                     # Build, deployment scripts
├── test/                        # Integration tests
│   └── integration/
├── deploy/                      # Docker, k8s manifests
│   ├── docker-compose.yaml
│   └── k8s/
├── go.mod
└── go.sum
```

### คำอธิบายเพิ่มเติม

- **cmd/**: มีไฟล์ `main.go` สำหรับแต่ละ executable (api, worker, cli) โดยปกติจะมีแค่ 2-3 ตัว
- **internal/**: โค้ดที่ไม่อยากให้ package อื่น import ได้ (Go compiler จะบังคับ)
  - `domain/`: แบ่งตาม bounded context (user, order) ภายในประกอบด้วย entity, value objects, repository interface, domain events
  - `application/`: use cases แต่ละตัว (RegisterUser, CreateOrder) และ DTOs
  - `infrastructure/`: การ implement repository (GORM, sqlx), cache (Redis), queue (RabbitMQ), email, etc.
  - `interfaces/`: HTTP handlers, middleware, routes; CLI commands
- **pkg/**: โค้ดที่ reusable ระหว่างโปรเจกต์ต่าง ๆ (เช่น logger, jwt, validator)
- **api/**: OpenAPI/Swagger specs หรือ protobuf
- **configs/**: ไฟล์คอนฟิก (YAML) และตัวอย่าง
- **migrations/**: SQL scripts สำหรับ migration (อาจใช้ `golang-migrate` หรือ `atlas`)
- **test/**: integration tests ที่ใช้ testcontainers หรือ environment จริง
- **deploy/**: docker-compose และ Kubernetes YAMLs

---

## บทที่ 48: การออกแบบ Workflow และ Task Management

### 1. Workflow การพัฒนา Feature ใหม่

```mermaid
flowchart LR
    A[Analyze<br/>Requirement] --> B[Design<br/>Domain Models]
    B --> C[Implement Domain<br/>Entities, Interfaces]
    C --> D[Implement Infrastructure<br/>Repository, Cache]
    D --> E[Implement Usecase<br/>Business Logic]
    E --> F[Implement Handler<br/>HTTP/CLI]
    F --> G[Write Tests<br/>Unit, Integration]
    G --> H[Document & Review]
    H --> I[Deploy]
```

### 2. Task List Template

**Phase 1: Analysis & Design**
- [ ] เข้าใจ requirement, user stories
- [ ] ระบุ bounded context และ aggregate
- [ ] ออกแบบ entity, value objects
- [ ] กำหนด repository interface methods
- [ ] ระบุ domain events
- [ ] ออกแบบ DTOs (request/response)

**Phase 2: Implementation**
- [ ] สร้าง entity struct และ business methods
- [ ] สร้าง repository interface
- [ ] สร้าง usecase struct และ method
- [ ] สร้าง HTTP handler
- [ ] ลงทะเบียน routes
- [ ] เขียน unit tests (domain, usecase with mocks)
- [ ] เขียน integration tests (repository with testcontainers)

**Phase 3: Infrastructure**
- [ ] Implement repository (GORM)
- [ ] สร้าง migration scripts
- [ ] ตั้งค่า connection pool
- [ ] Implement cache (ถ้าจำเป็น)
- [ ] ทดสอบ repository

**Phase 4: Documentation & Review**
- [ ] อัปเดต API docs (Swagger)
- [ ] รัน linter (`golangci-lint run`)
- [ ] รัน test coverage (`go test -cover`), ตั้งเป้า >80%
- [ ] รัน race detector (`go test -race`)
- [ ] Code review

**Phase 5: Deployment**
- [ ] สร้าง Docker image
- [ ] อัปเดต docker-compose (ถ้ามี)
- [ ] Deploy to staging
- [ ] ทดสอบ end-to-end
- [ ] Deploy to production

### 3. Checklist Template

**Code Quality**
- [ ] All exported functions have comments
- [ ] No unused imports/variables (`go vet`)
- [ ] Code formatted with `go fmt`
- [ ] Error handling explicit (no ignored errors)
- [ ] No panic in library code (only main/init)
- [ ] Context passed as first parameter
- [ ] Interfaces are small and focused
- [ ] No global state except configuration

**Security**
- [ ] Passwords hashed with bcrypt
- [ ] JWT secret loaded from environment
- [ ] CORS properly configured
- [ ] Input validation on all endpoints
- [ ] SQL injection prevented (parameterized queries)
- [ ] No sensitive data in logs
- [ ] HTTPS enforced in production
- [ ] Rate limiting on auth endpoints

**Performance**
- [ ] Database indexes on frequently queried columns
- [ ] Use of caching (Redis) for hot data
- [ ] Connection pools configured for DB and Redis
- [ ] Avoid N+1 queries (use Preload in GORM)
- [ ] Benchmarks for critical paths

**Testing**
- [ ] Unit tests cover business logic (usecase)
- [ ] Repository tests with testcontainers
- [ ] HTTP handler tests with httptest
- [ ] Mock external dependencies
- [ ] Race condition tests with `-race`
- [ ] Coverage >80%

**Deployment**
- [ ] Configurable via environment variables
- [ ] Graceful shutdown implemented
- [ ] Health check endpoints (/health, /ready)
- [ ] Logging to stdout (for container)
- [ ] Docker image built with non-root user
- [ ] Secrets not baked into image
- [ ] Database migration runs automatically on startup
- [ ] Readiness and liveness probes configured
- [ ] Monitoring metrics exposed (Prometheus)

---

## ตัวอย่างการนำไปใช้จริง: การเพิ่มโมดูล "Product"

สมมติเราต้องเพิ่มฟีเจอร์จัดการสินค้า (Product) ในแอปพลิเคชัน e-commerce

### ขั้นตอนตาม Workflow

1. **Analyze**: Product มี fields: id, name, description, price, stock, status (active/inactive), created_at, updated_at. Use cases: create, get, list, update, delete, search.

2. **Design**: 
   - Entity `Product` (id, name, price, stock, ...) with business methods like `ReduceStock(quantity)`, `IncreaseStock(quantity)`, `IsAvailable()`.
   - Repository interface: `Save`, `FindByID`, `FindAll`, `Update`, `Delete`, `Search`.
   - Use cases: `CreateProduct`, `GetProduct`, `UpdateProduct`, `DeleteProduct`, `ListProducts`, `SearchProducts`.
   - DTOs: `CreateProductRequest`, `ProductResponse`, `ListProductsRequest`, etc.

3. **Implement Domain**:
   ```go
   // internal/domain/product/entity.go
   type Product struct {
       id        uint
       name      string
       price     Money
       stock     int
       status    ProductStatus
       createdAt time.Time
       updatedAt time.Time
   }
   // business methods...
   ```

4. **Implement Infrastructure**:
   - GORM model with tags
   - Repository implementation using GORM

5. **Implement Usecase**:
   - `CreateProductUseCase` that validates and saves product

6. **Implement Handler**:
   - HTTP handlers with validation

7. **Write Tests**:
   - Unit tests for entity business logic
   - Usecase tests with mock repository
   - Repository tests with testcontainers

8. **Document & Review**:
   - Update OpenAPI spec
   - Run linter and coverage

9. **Deploy**:
   - Create migration for `products` table
   - Update Docker image

### ผลลัพธ์: โครงสร้างไฟล์ใหม่

```
internal/
├── domain/
│   └── product/
│       ├── entity.go
│       ├── repository.go
│       └── events.go
├── application/
│   └── product/
│       ├── create.go
│       ├── get.go
│       ├── list.go
│       ├── update.go
│       ├── delete.go
│       └── dto.go
├── infrastructure/
│   └── persistence/
│       └── gorm/
│           ├── product_repo.go
│           └── product_model.go
└── interfaces/
    └── http/
        └── handlers/
            └── product.go
```

---

## สรุป

ภาคที่ 7 ได้เสนอแนวทางในการออกแบบสถาปัตยกรรมและกระบวนการพัฒนาสำหรับโปรเจกต์ Go ระดับ Production:

- **Clean Architecture** ช่วยแยกความรับผิดชอบและทำให้โค้ดทดสอบง่าย
- **Blueprint** ที่นำเสนอเป็นโครงสร้างเริ่มต้นที่ผ่านการพิสูจน์แล้วจากโปรเจกต์จริง
- **Workflow และ Task Management** ช่วยให้ทีมทำงานเป็นระบบ ลดความเสี่ยง และเพิ่มคุณภาพ

การนำหลักการเหล่านี้ไปใช้จะช่วยให้แอปพลิเคชันมีความยืดหยุ่น บำรุงรักษาง่าย และสามารถขยายขนาดได้ตามความต้องการทางธุรกิจ

## # ภาคที่ 8: Domain-Driven Design (DDD) กับ Go (บทที่ 49–51)

## แผนภาพ Data Flow (Flowchart TB)

```mermaid
flowchart TB
    subgraph Ubiquitous["Ubiquitous Language (ภาษาร่วม)"]
        U1["Domain Experts & Developers<br/>ใช้คำศัพท์เดียวกัน"]
    end

    subgraph BoundedContext["Bounded Contexts (บริบทที่จำกัด)"]
        BC1["Order Context<br/>Order, OrderItem, Payment"]
        BC2["User Context<br/>User, Profile, Authentication"]
        BC3["Inventory Context<br/>Product, Stock, Warehouse"]
    end

    subgraph DomainModel["Domain Model (โมเดลโดเมน)"]
        direction TB
        DM1["Entities<br/>- มี identity<br/>- mutable"]
        DM2["Value Objects<br/>- ไม่มี identity<br/>- immutable"]
        DM3["Aggregates<br/>- กลุ่มของ entities/value objects<br/>- Aggregate Root เป็นทางเข้า"]
        DM4["Domain Events<br/>- เหตุการณ์สำคัญในโดเมน<br/>- ใช้สื่อสารระหว่าง aggregates"]
        DM5["Repositories<br/>- abstraction สำหรับ aggregate roots<br/>- ซ่อน infrastructure"]
        DM6["Domain Services<br/>- business logic ที่ไม่เหมาะกับ entity/value object"]
    end

    subgraph Application["Application Layer (Use Cases)"]
        A1["Commands<br/>PlaceOrder, RegisterUser"]
        A2["Queries<br/>GetOrder, ListUsers"]
        A3["DTOs<br/>Request/Response objects"]
        A4["Event Handlers<br/>- สนใจ domain events<br/>- ทำงาน async"]
    end

    subgraph Infrastructure["Infrastructure Layer"]
        I1["Persistence<br/>PostgreSQL, GORM"]
        I2["Event Store<br/>PostgreSQL JSONB, EventStoreDB"]
        I3["Message Broker<br/>RabbitMQ, Kafka"]
        I4["Cache<br/>Redis"]
    end

    Ubiquitous --> DomainModel
    DomainModel --> Application
    Application --> Infrastructure

    subgraph EventStorming["Event Storming (Workshop)"]
        ES1["Domain Events<br/>(orange)"]
        ES2["Commands<br/>(blue)"]
        ES3["Aggregates<br/>(yellow)"]
        ES4["Policies / Rules<br/>(purple)"]
        ES5["Read Models<br/>(green)"]
        ES6["External Systems<br/>(pink)"]
    end

    subgraph CQRS["CQRS (Command Query Responsibility Segregation)"]
        C1["Command Model<br/>- Write side<br/>- Domain logic, consistency"]
        C2["Query Model<br/>- Read side<br/>- Optimized for queries, denormalized"]
        C3["Event Sourcing<br/>- Store events as source of truth<br/>- Rebuild state from events"]
        C4["Projections<br/>- Build read models from events"]
    end

    EventStorming --> DomainModel
    CQRS --> Application
    CQRS --> Infrastructure
```

---

## คำอธิบายภาษาไทย (แบบละเอียด)

ภาคที่ 8 ครอบคลุมบทที่ 49–51 โดยมีเนื้อหาเกี่ยวกับ **Domain-Driven Design (DDD)** และการนำไปใช้ในภาษา Go:

- **บทที่ 49:** หลักการ DDD และการนำไปใช้ใน Go
- **บทที่ 50:** Aggregates, Event Storming และ CQRS
- **บทที่ 51:** การออกแบบบริการด้วย Go-DDD

DDD เป็นแนวทางในการออกแบบซอฟต์แวร์ที่เน้นการสร้างโมเดลที่สะท้อนความรู้ความเข้าใจ (domain knowledge) อย่างแท้จริง โดยใช้ภาษากลาง (Ubiquitous Language) ร่วมกันระหว่างนักพัฒนาและผู้เชี่ยวชาญโดเมน

---

## บทที่ 49: หลักการ DDD และการนำไปใช้ใน Go

### 1. หลักการสำคัญของ DDD

- **Ubiquitous Language (ภาษาร่วม)**: สร้างภาษากลางที่ใช้ร่วมกันระหว่างนักพัฒนาและผู้เชี่ยวชาญโดเมน ใช้ศัพท์เดียวกันในโค้ด, การสนทนา, และเอกสาร
- **Bounded Context (บริบทที่จำกัด)**: แบ่งโดเมนขนาดใหญ่ออกเป็นบริทย่อยที่มีขอบเขตชัดเจน แต่ละ Bounded Context มีโมเดลของตัวเองและภาษาร่วมของตัวเอง
- **Entities**: วัตถุที่มีเอกลักษณ์ (identity) และสามารถเปลี่ยนแปลงได้ (mutable) เช่น `User`, `Order`
- **Value Objects**: วัตถุที่ไม่มีเอกลักษณ์ในตัวเอง ถูกกำหนดโดยคุณสมบัติ (immutable) เช่น `Address`, `Money`
- **Aggregates**: กลุ่มของ Entities และ Value Objects ที่ถูกจัดการเป็นหน่วยเดียวกัน มี Aggregate Root เป็นตัวควบคุมความสอดคล้องของข้อมูล
- **Domain Events**: เหตุการณ์สำคัญในโดเมนที่เกิดขึ้น เช่น `OrderPlaced`, `UserRegistered`
- **Repositories**: ให้ abstraction ในการเข้าถึง aggregate roots ซ่อนรายละเอียดของแหล่งข้อมูล
- **Domain Services**: ใช้สำหรับ logic ที่ไม่เหมาะจะอยู่ใน entity หรือ value object เช่น `TransferService` ที่โอนเงินระหว่างบัญชี

### 2. การนำ DDD ไปใช้ใน Go

Go ไม่มี class และ inheritance แต่ใช้ **struct + method** และ **interface** เพื่อสร้าง abstraction การสืบทอดทำผ่าน **composition** (embedding) แทน

#### 2.1 Entities และ Value Objects

```go
// internal/domain/user/entity.go
package user

import (
    "errors"
    "time"
)

// Entity: User (มี identity)
type User struct {
    id        uint
    email     Email      // Value Object
    password  Password   // Value Object
    name      string
    isActive  bool
    createdAt time.Time
    updatedAt time.Time
    events    []DomainEvent
}

// Constructor (factory method)
func NewUser(email, password, name string) (*User, error) {
    emailVO, err := NewEmail(email)
    if err != nil {
        return nil, err
    }
    passwordVO, err := NewPassword(password)
    if err != nil {
        return nil, err
    }
    return &User{
        id:        generateID(),
        email:     *emailVO,
        password:  *passwordVO,
        name:      name,
        isActive:  true,
        createdAt: time.Now(),
        updatedAt: time.Now(),
        events:    []DomainEvent{},
    }, nil
}

// Getters (exported)
func (u *User) ID() uint         { return u.id }
func (u *User) Email() string    { return u.email.String() }
func (u *User) Name() string     { return u.name }
func (u *User) IsActive() bool   { return u.isActive }

// Business method
func (u *User) Deactivate() error {
    if !u.isActive {
        return errors.New("user already inactive")
    }
    u.isActive = false
    u.updatedAt = time.Now()
    u.addDomainEvent(NewUserDeactivatedEvent(u.id))
    return nil
}

func (u *User) addDomainEvent(event DomainEvent) {
    u.events = append(u.events, event)
}

func (u *User) Events() []DomainEvent { return u.events }
func (u *User) ClearEvents()          { u.events = nil }
```

```go
// internal/domain/user/value_objects.go
package user

import (
    "regexp"
    "strings"

    "golang.org/x/crypto/bcrypt"
)

// Value Object: Email (immutable)
type Email struct {
    value string
}

func NewEmail(email string) (*Email, error) {
    email = strings.TrimSpace(email)
    if !isValidEmail(email) {
        return nil, errors.New("invalid email format")
    }
    return &Email{value: email}, nil
}

func (e Email) String() string { return e.value }

// Value Object: Password (immutable, hashed)
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

func (p *Password) Compare(plain string) bool {
    return bcrypt.CompareHashAndPassword([]byte(p.hash), []byte(plain)) == nil
}
```

#### 2.2 Aggregates และ Aggregate Root

```go
// internal/domain/order/aggregate.go
package order

import (
    "errors"
    "time"
)

// Aggregate Root: Order
type Order struct {
    id         uint
    customerID uint
    items      []OrderItem   // slice of value objects
    status     OrderStatus
    total      Money
    createdAt  time.Time
    updatedAt  time.Time
    events     []DomainEvent
}

func NewOrder(customerID uint) *Order {
    o := &Order{
        id:         generateID(),
        customerID: customerID,
        items:      []OrderItem{},
        status:     StatusDraft,
        total:      Money{},
        createdAt:  time.Now(),
        updatedAt:  time.Now(),
    }
    o.addDomainEvent(NewOrderCreatedEvent(o.id, customerID))
    return o
}

func (o *Order) AddItem(productID uint, name string, price Money, quantity int) error {
    if o.status != StatusDraft {
        return errors.New("cannot add item to non-draft order")
    }
    if quantity <= 0 {
        return errors.New("quantity must be positive")
    }
    // Create value object
    item := NewOrderItem(productID, name, price, quantity)
    o.items = append(o.items, item)
    o.recalculateTotal()
    o.updatedAt = time.Now()
    o.addDomainEvent(NewOrderItemAddedEvent(o.id, productID, quantity))
    return nil
}

func (o *Order) Submit() error {
    if len(o.items) == 0 {
        return errors.New("cannot submit empty order")
    }
    if o.status != StatusDraft {
        return errors.New("order already submitted")
    }
    o.status = StatusSubmitted
    o.updatedAt = time.Now()
    o.addDomainEvent(NewOrderSubmittedEvent(o.id))
    return nil
}

func (o *Order) recalculateTotal() {
    var sum int64
    for _, item := range o.items {
        sum += item.Total().Cents()
    }
    o.total = NewMoneyFromCents(sum)
}

// Repository interface for Order aggregate root
type OrderRepository interface {
    Save(ctx context.Context, order *Order) error
    FindByID(ctx context.Context, id uint) (*Order, error)
    FindByCustomer(ctx context.Context, customerID uint) ([]*Order, error)
}
```

#### 2.3 Domain Events

```go
// internal/domain/order/events.go
package order

import "time"

type DomainEvent interface {
    OccurredAt() time.Time
}

type OrderCreatedEvent struct {
    OrderID    uint      `json:"order_id"`
    CustomerID uint      `json:"customer_id"`
    Occurred   time.Time `json:"occurred_at"`
}

func NewOrderCreatedEvent(orderID, customerID uint) OrderCreatedEvent {
    return OrderCreatedEvent{
        OrderID:    orderID,
        CustomerID: customerID,
        Occurred:   time.Now(),
    }
}

func (e OrderCreatedEvent) OccurredAt() time.Time { return e.Occurred }

type OrderSubmittedEvent struct {
    OrderID  uint      `json:"order_id"`
    Occurred time.Time `json:"occurred_at"`
}

func NewOrderSubmittedEvent(orderID uint) OrderSubmittedEvent {
    return OrderSubmittedEvent{
        OrderID:  orderID,
        Occurred: time.Now(),
    }
}

func (e OrderSubmittedEvent) OccurredAt() time.Time { return e.Occurred }
```

#### 2.4 Domain Services

```go
// internal/domain/order/transfer_service.go
package order

type TransferService interface {
    Transfer(ctx context.Context, fromID, toID uint, amount Money) error
}

// Implementation in infrastructure layer
type BankTransferService struct {
    client *http.Client
    // ...
}

func (s *BankTransferService) Transfer(ctx context.Context, fromID, toID uint, amount Money) error {
    // Call external bank API
    // ...
}
```

#### 2.5 Repositories (interface ใน domain, implementation ใน infrastructure)

```go
// internal/domain/order/repository.go (interface)
package order

import "context"

type OrderRepository interface {
    Save(ctx context.Context, order *Order) error
    FindByID(ctx context.Context, id uint) (*Order, error)
}
```

```go
// internal/infrastructure/persistence/gorm/order_repo.go (implementation)
package gormrepo

import (
    "context"
    "encoding/json"

    "gorm.io/gorm"
    "myapp/internal/domain/order"
)

type orderRepository struct {
    db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) order.OrderRepository {
    return &orderRepository{db: db}
}

// GORM model for persistence
type OrderModel struct {
    ID         uint           `gorm:"primaryKey"`
    CustomerID uint           `gorm:"index"`
    Status     string         `gorm:"size:20"`
    TotalCents int64
    Items      json.RawMessage `gorm:"type:jsonb"`
    CreatedAt  time.Time
    UpdatedAt  time.Time
}

func (r *orderRepository) Save(ctx context.Context, o *order.Order) error {
    // Convert domain Order to OrderModel
    itemsJSON, _ := json.Marshal(o.Items()) // assume items are serializable
    model := &OrderModel{
        ID:         o.ID(),
        CustomerID: o.CustomerID(),
        Status:     o.Status().String(),
        TotalCents: o.Total().Cents(),
        Items:      itemsJSON,
        CreatedAt:  o.CreatedAt(),
        UpdatedAt:  o.UpdatedAt(),
    }
    return r.db.WithContext(ctx).Save(model).Error
}
```

---

## บทที่ 50: Aggregates, Event Storming และ CQRS

### 1. Aggregates (กลุ่มวัตถุที่มีความสอดคล้อง)

Aggregate คือกลุ่มของ Entities และ Value Objects ที่ถูกยึดเข้าด้วยกันโดย Aggregate Root ซึ่งเป็นตัวเดียวที่อนุญาตให้เข้าถึงหรือแก้ไขวัตถุอื่นภายในกลุ่มจากภายนอก

**หลักการ:**
- Aggregate Root มี identity และเป็นจุดเดียวที่ภายนอกเข้าถึงได้
- การเปลี่ยนแปลงใด ๆ ภายใน Aggregate ต้องทำผ่าน Root เท่านั้น
- ภายใน Aggregate เดียวกันต้องมีความสอดคล้องกันในทันที (consistency boundary)
- ระหว่าง Aggregates ควรใช้ eventual consistency ผ่าน Domain Events

### 2. Event Storming (เทคนิคค้นพบโดเมน)

Event Storming เป็นเวิร์กช็อปแบบมีส่วนร่วมที่ช่วยให้ทีม (นักพัฒนา, นักธุรกิจ, ผู้เชี่ยวชาญโดเมน) ระบุและเข้าใจโดเมนผ่านเหตุการณ์สำคัญที่เกิดขึ้นในระบบ

**สัญลักษณ์ทั่วไป:**
- **สีส้ม** – Domain Events (สิ่งที่เกิดขึ้นแล้ว) เช่น `OrderPlaced`, `PaymentReceived`
- **สีน้ำเงิน** – Commands (คำสั่งที่ทำให้เกิดเหตุการณ์) เช่น `PlaceOrder`, `RefundPayment`
- **สีเหลือง** – Aggregates (กลุ่มข้อมูลที่ถูกคำสั่งเรียก) เช่น `Order`, `Customer`
- **สีม่วง** – Policies / Rules (กฎที่ทำงานอัตโนมัติ)
- **สีเขียว** – Read Models (ข้อมูลที่แสดงผล)
- **สีชมพู** – External Systems (ระบบภายนอก)

### 3. CQRS (Command Query Responsibility Segregation)

CQRS แยกโมเดลการ **เขียน** (Command) และ **อ่าน** (Query) ออกจากกัน ทำให้สามารถปรับแต่งให้เหมาะสมกับแต่ละฝั่งได้

**ตัวอย่างโครงสร้าง CQRS ใน Go:**

```go
// Command side (write model)
type PlaceOrderCommand struct {
    CustomerID uint
    Items      []OrderItemDTO
}

type PlaceOrderHandler struct {
    orderRepo order.OrderRepository
    eventBus  EventBus
}

func (h *PlaceOrderHandler) Handle(ctx context.Context, cmd PlaceOrderCommand) error {
    // Create order aggregate
    ord := order.NewOrder(cmd.CustomerID)
    for _, item := range cmd.Items {
        ord.AddItem(item.ProductID, item.Name, item.Price, item.Quantity)
    }
    if err := ord.Submit(); err != nil {
        return err
    }
    if err := h.orderRepo.Save(ctx, ord); err != nil {
        return err
    }
    // Publish events
    for _, evt := range ord.Events() {
        h.eventBus.Publish(evt)
    }
    return nil
}

// Query side (read model)
type OrderReadModel struct {
    ID         uint
    CustomerID uint
    Status     string
    Total      float64
    Items      []OrderItemReadModel
}

type OrderQueryHandler struct {
    db *sql.DB // หรือ repository อ่านอย่างเดียว
}

func (h *OrderQueryHandler) GetOrder(ctx context.Context, id uint) (*OrderReadModel, error) {
    var order OrderReadModel
    // query จาก read database (อาจเป็นตารางที่ denormalized)
    // ...
    return &order, nil
}
```

#### Event Sourcing + Projections

Event Sourcing คือการเก็บ event ทั้งหมดเป็น source of truth แล้วสร้าง state ปัจจุบันจากการ replay events

```go
// Event store
type EventStore interface {
    Save(ctx context.Context, streamID string, events []DomainEvent, expectedVersion int) error
    Load(ctx context.Context, streamID string) ([]DomainEvent, error)
}

// Projection: สร้าง read model จาก events
type OrderProjection struct {
    db *sql.DB
}

func (p *OrderProjection) HandleEvent(event DomainEvent) {
    switch e := event.(type) {
    case OrderCreatedEvent:
        p.db.Exec("INSERT INTO order_read_models (id, customer_id, status, total, created_at) VALUES ($1, $2, $3, $4, $5)",
            e.OrderID, e.CustomerID, "Draft", 0, e.Occurred)
    case OrderItemAddedEvent:
        p.db.Exec("INSERT INTO order_item_read_models (order_id, product_id, quantity, price) VALUES ($1, $2, $3, $4)",
            e.OrderID, e.ProductID, e.Quantity, e.Price)
        p.db.Exec("UPDATE order_read_models SET total = total + $1 WHERE id = $2", e.Price*float64(e.Quantity), e.OrderID)
    case OrderSubmittedEvent:
        p.db.Exec("UPDATE order_read_models SET status = 'Submitted' WHERE id = $1", e.OrderID)
    }
}
```

---

## บทที่ 51: การออกแบบบริการด้วย Go-DDD

### โครงสร้างโปรเจกต์แบบ DDD เต็มรูปแบบ

```
myapp/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── config/               # configuration
│   ├── domain/               # domain layer
│   │   ├── user/
│   │   │   ├── entity.go
│   │   │   ├── value_objects.go
│   │   │   ├── repository.go
│   │   │   ├── events.go
│   │   │   └── service.go (domain services)
│   │   ├── order/
│   │   │   ├── aggregate.go
│   │   │   ├── repository.go
│   │   │   └── events.go
│   │   └── shared/          # shared value objects
│   ├── application/          # use cases
│   │   ├── user/
│   │   │   ├── register.go
│   │   │   ├── login.go
│   │   │   └── dto.go
│   │   └── order/
│   │       ├── place_order.go
│   │       ├── get_order.go
│   │       └── dto.go
│   ├── infrastructure/       # implementations
│   │   ├── persistence/
│   │   │   ├── gorm/
│   │   │   │   ├── user_repo.go
│   │   │   │   └── order_repo.go
│   │   │   └── event_store/
│   │   ├── bus/              # event bus
│   │   └── mail/
│   └── interfaces/           # delivery
│       ├── http/
│       │   ├── handlers/
│       │   ├── middleware/
│       │   └── routes.go
│       └── cli/
├── pkg/                      # reusable packages
│   ├── logger/
│   └── errors/
├── migrations/
└── go.mod
```

### ตัวอย่าง Use Case (Application Layer)

```go
// internal/application/order/place_order.go
package order

import (
    "context"

    "myapp/internal/domain/order"
    "myapp/internal/domain/user"
)

type PlaceOrderUseCase struct {
    orderRepo order.OrderRepository
    userRepo  user.UserRepository
    eventBus  EventBus
}

type PlaceOrderInput struct {
    UserID uint
    Items  []OrderItemInput
}

type OrderItemInput struct {
    ProductID uint
    Name      string
    Price     float64
    Quantity  int
}

type PlaceOrderOutput struct {
    OrderID uint
    Status  string
}

func NewPlaceOrderUseCase(
    orderRepo order.OrderRepository,
    userRepo user.UserRepository,
    eventBus EventBus,
) *PlaceOrderUseCase {
    return &PlaceOrderUseCase{
        orderRepo: orderRepo,
        userRepo:  userRepo,
        eventBus:  eventBus,
    }
}

func (uc *PlaceOrderUseCase) Execute(ctx context.Context, input PlaceOrderInput) (*PlaceOrderOutput, error) {
    // 1. Validate user exists and is active
    u, err := uc.userRepo.FindByID(ctx, input.UserID)
    if err != nil {
        return nil, fmt.Errorf("user not found: %w", err)
    }
    if !u.IsActive() {
        return nil, errors.New("inactive user cannot place order")
    }

    // 2. Create order aggregate
    ord := order.NewOrder(input.UserID)
    for _, item := range input.Items {
        money := order.NewMoney(item.Price)
        ord.AddItem(item.ProductID, item.Name, money, item.Quantity)
    }

    // 3. Submit order
    if err := ord.Submit(); err != nil {
        return nil, err
    }

    // 4. Persist
    if err := uc.orderRepo.Save(ctx, ord); err != nil {
        return nil, err
    }

    // 5. Dispatch events
    for _, evt := range ord.Events() {
        uc.eventBus.Publish(evt)
    }

    return &PlaceOrderOutput{
        OrderID: ord.ID(),
        Status:  ord.Status().String(),
    }, nil
}
```

### Dependency Injection (Wire)

```go
// internal/wire/wire.go
package wire

import (
    "myapp/internal/application/order"
    "myapp/internal/domain/order"
    "myapp/internal/infrastructure/persistence/gorm"
    "myapp/internal/interfaces/http/handlers"
    "myapp/pkg/eventbus"
)

func InitializeOrderAPI(db *gorm.DB, bus eventbus.EventBus) *handlers.OrderHandler {
    orderRepo := gorm.NewOrderRepository(db)
    userRepo := gorm.NewUserRepository(db)

    placeOrderUC := order.NewPlaceOrderUseCase(orderRepo, userRepo, bus)
    getOrderUC := order.NewGetOrderUseCase(orderRepo)

    return handlers.NewOrderHandler(placeOrderUC, getOrderUC)
}
```

---

## สรุป

ภาคที่ 8 ได้แนะนำหลักการ Domain-Driven Design และการนำไปใช้ใน Go อย่างเป็นรูปธรรม:

- **หลักการ DDD**: Ubiquitous Language, Bounded Context, Entities, Value Objects, Aggregates, Repositories, Domain Services, Domain Events
- **Event Storming**: เทคนิคการวิเคราะห์โดเมนผ่านเหตุการณ์
- **CQRS**: แยก command และ query models, event sourcing, projections

การนำ DDD มาปรับใช้ใน Go ช่วยให้โค้ดสะท้อนภาษาธุรกิจ ทดสอบง่าย และปรับเปลี่ยน infrastructure ได้โดยไม่กระทบโดเมน

## ภาคที่ 9: การผสานระบบภายนอก (บทที่ 52–58)

## แผนภาพ Data Flow (Flowchart TB)

```mermaid
flowchart TB
    subgraph External["ระบบภายนอก (External Systems)"]
        direction LR
        IoT["IoT Sensors"] -->|MQTT| MQTT
        User["ผู้ใช้"] -->|HTTP Request| API
    end

    subgraph Core["ระบบหลัก Go"]
        API[HTTP API / WebSocket]
        Redis[(Redis)]
        Rabbit[(RabbitMQ)]
        Influx[(InfluxDB)]
    end

    subgraph Processing["การประมวลผล"]
        Workers[Worker Pool]
        Cache[Cache Layer]
        Stream[Event Stream]
    end

    subgraph Notifications["การแจ้งเตือน"]
        SMS[SMS Gateway]
        LINE[LINE Notify]
        Discord[Discord Webhook]
    end

    subgraph Clients["ไคลเอนต์"]
        Browser[เว็บเบราว์เซอร์]
        Mobile[แอปมือถือ]
    end

    %% Data flow
    MQTT -->|Telemetry| Influx
    API -->|Read/Write| Redis
    API -->|Publish Event| Rabbit
    Rabbit -->|Consume| Workers
    Workers -->|Store Metrics| Influx
    Workers -->|Update Cache| Redis
    Influx -->|Query| API
    Redis -->|Cached Data| API

    API -->|WebSocket| Browser
    API -->|WebSocket| Mobile

    Workers -->|Alert| SMS
    Workers -->|Alert| LINE
    Workers -->|Alert| Discord
    Influx -->|Threshold Exceeded| Workers
```

---

## คำอธิบายภาษาไทย (แบบละเอียด)

ภาคที่ 9 ครอบคลุมบทที่ 52–58 โดยนำเสนอการผสานระบบภายนอกที่พบได้บ่อยในโลกแห่งความจริง:

- **Redis (บทที่ 52)**: Cache, Message Queue
- **RabbitMQ (บทที่ 53)**: Message Broker มาตรฐานองค์กร
- **MQTT (บทที่ 54)**: IoT และระบบเรียลไทม์
- **InfluxDB (บทที่ 55)**: Time‑Series Database
- **WebSocket (บทที่ 56)**: การสื่อสารแบบ real‑time
- **SMS และ LINE Notify (บทที่ 57)**: การแจ้งเตือนผ่านมือถือ
- **Discord Webhook (บทที่ 58)**: การแจ้งเตือนใน Discord

---

## บทที่ 52: Redis สำหรับ Cache และ Message Queue

### Redis คืออะไร?

Redis (Remote Dictionary Server) เป็นฐานข้อมูลแบบ in‑memory ที่รวดเร็วสูง นิยมใช้สำหรับ caching, session storage, message queue (Pub/Sub, List), และ real‑time analytics

### การติดตั้งและเชื่อมต่อ

```go
go get github.com/redis/go-redis/v9
```

```go
package cache

import (
    "context"
    "encoding/json"
    "time"

    "github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func InitRedis(addr, password string, db int) {
    rdb = redis.NewClient(&redis.Options{
        Addr:     addr,
        Password: password,
        DB:       db,
    })
}

func GetRedis() *redis.Client {
    return rdb
}
```

### 1. Cache-aside Pattern

```go
// GetUserWithCache
func GetUserWithCache(ctx context.Context, userID int) (*User, error) {
    key := fmt.Sprintf("user:%d", userID)

    // 1. Try to get from Redis
    val, err := rdb.Get(ctx, key).Result()
    if err == nil {
        var user User
        if err := json.Unmarshal([]byte(val), &user); err == nil {
            return &user, nil
        }
    }

    // 2. Cache miss: query database
    user, err := db.GetUserByID(ctx, userID)
    if err != nil {
        return nil, err
    }

    // 3. Store in cache (async)
    go func() {
        data, _ := json.Marshal(user)
        rdb.Set(context.Background(), key, data, 10*time.Minute)
    }()

    return user, nil
}
```

### 2. Redis Pub/Sub (Message Queue)

**Publisher**
```go
func PublishOrderEvent(ctx context.Context, orderID int) error {
    event := map[string]interface{}{
        "event": "order_created",
        "order_id": orderID,
    }
    data, _ := json.Marshal(event)
    return rdb.Publish(ctx, "orders", data).Err()
}
```

**Subscriber**
```go
func SubscribeOrders() {
    pubsub := rdb.Subscribe(context.Background(), "orders")
    defer pubsub.Close()

    ch := pubsub.Channel()
    for msg := range ch {
        var event map[string]interface{}
        json.Unmarshal([]byte(msg.Payload), &event)
        // process event
        log.Printf("Received event: %v", event)
    }
}
```

### 3. Redis List as Queue (Producer-Consumer)

**Producer (LPush)**
```go
func EnqueueJob(job interface{}) error {
    data, _ := json.Marshal(job)
    return rdb.LPush(context.Background(), "jobs", data).Err()
}
```

**Consumer (BRPop)**
```go
func Worker() {
    for {
        result, err := rdb.BRPop(context.Background(), 0, "jobs").Result()
        if err != nil {
            continue
        }
        // result[1] is the value
        var job Job
        json.Unmarshal([]byte(result[1]), &job)
        processJob(job)
    }
}
```

---

## บทที่ 53: RabbitMQ – Message Broker มาตรฐานองค์กร

### RabbitMQ คืออะไร?

RabbitMQ เป็น message broker ที่ใช้โปรโตคอล AMQP รองรับ routing ที่ซับซ้อน (direct, topic, fanout exchanges), message persistence, และ reliability features

### การติดตั้งและเชื่อมต่อ

```go
go get github.com/rabbitmq/amqp091-go
```

```go
package queue

import (
    "context"
    "log"

    amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
    conn *amqp.Connection
    ch   *amqp.Channel
}

func NewRabbitMQ(url string) (*RabbitMQ, error) {
    conn, err := amqp.Dial(url)
    if err != nil {
        return nil, err
    }
    ch, err := conn.Channel()
    if err != nil {
        conn.Close()
        return nil, err
    }
    return &RabbitMQ{conn: conn, ch: ch}, nil
}

func (r *RabbitMQ) Close() {
    r.ch.Close()
    r.conn.Close()
}
```

### 1. Simple Queue (Work Queue)

**Producer**
```go
func (r *RabbitMQ) Publish(queueName string, body []byte) error {
    q, err := r.ch.QueueDeclare(
        queueName, // name
        true,      // durable
        false,     // delete when unused
        false,     // exclusive
        false,     // no-wait
        nil,       // arguments
    )
    if err != nil {
        return err
    }

    return r.ch.Publish(
        "",          // exchange
        q.Name,      // routing key
        false,       // mandatory
        false,       // immediate
        amqp.Publishing{
            ContentType:  "application/json",
            Body:         body,
            DeliveryMode: amqp.Persistent, // make message persistent
        },
    )
}
```

**Consumer**
```go
func (r *RabbitMQ) Consume(queueName string, handler func([]byte) error) error {
    q, err := r.ch.QueueDeclare(queueName, true, false, false, false, nil)
    if err != nil {
        return err
    }

    // prefetch 1 (fair dispatch)
    r.ch.Qos(1, 0, false)

    msgs, err := r.ch.Consume(
        q.Name, // queue
        "",     // consumer
        false,  // auto-ack (we'll ack manually)
        false,  // exclusive
        false,  // no-local
        false,  // no-wait
        nil,    // args
    )
    if err != nil {
        return err
    }

    go func() {
        for d := range msgs {
            if err := handler(d.Body); err != nil {
                // if error, reject and requeue
                d.Nack(false, true)
            } else {
                d.Ack(false)
            }
        }
    }()
    return nil
}
```

### 2. Publish/Subscribe (Fanout Exchange)

**Producer**
```go
func (r *RabbitMQ) PublishFanout(exchangeName string, body []byte) error {
    err := r.ch.ExchangeDeclare(
        exchangeName, // name
        "fanout",     // type
        true,         // durable
        false,        // auto-deleted
        false,        // internal
        false,        // no-wait
        nil,          // arguments
    )
    if err != nil {
        return err
    }

    return r.ch.Publish(
        exchangeName, // exchange
        "",           // routing key (ignored for fanout)
        false,
        false,
        amqp.Publishing{
            ContentType: "application/json",
            Body:        body,
        },
    )
}
```

**Consumer**
```go
func (r *RabbitMQ) SubscribeFanout(exchangeName string, handler func([]byte) error) error {
    err := r.ch.ExchangeDeclare(exchangeName, "fanout", true, false, false, false, nil)
    if err != nil {
        return err
    }

    q, err := r.ch.QueueDeclare(
        "",    // name (auto-generated)
        false, // durable
        false, // delete when unused
        true,  // exclusive (temporary)
        false, // no-wait
        nil,
    )
    if err != nil {
        return err
    }

    err = r.ch.QueueBind(q.Name, "", exchangeName, false, nil)
    if err != nil {
        return err
    }

    msgs, err := r.ch.Consume(q.Name, "", true, false, false, false, nil)
    if err != nil {
        return err
    }

    go func() {
        for d := range msgs {
            handler(d.Body)
        }
    }()
    return nil
}
```

### 3. Routing (Direct Exchange) และ Topic Exchange

คล้ายกันแต่ใช้ exchange type ต่างกันและ routing key มี pattern สำหรับ topic

---

## บทที่ 54: MQTT สำหรับ IoT และระบบเรียลไทม์

### MQTT คืออะไร?

MQTT (Message Queuing Telemetry Transport) เป็นโปรโตคอล lightweight สำหรับการสื่อสารแบบ publish/subscribe เหมาะสำหรับ IoT, mobile apps, และระบบที่ต้องการการส่งข้อมูลแบบ real‑time ด้วยแบนด์วิดท์ต่ำ

### การติดตั้งและเชื่อมต่อ

```go
go get github.com/eclipse/paho.mqtt.golang
```

```go
package mqtt

import (
    "fmt"
    "log"
    "time"

    mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MQTTClient struct {
    client mqtt.Client
}

func NewMQTTClient(broker, clientID, username, password string) (*MQTTClient, error) {
    opts := mqtt.NewClientOptions()
    opts.AddBroker(broker)
    opts.SetClientID(clientID)
    if username != "" {
        opts.SetUsername(username)
        opts.SetPassword(password)
    }
    opts.SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {
        log.Printf("Received: %s from topic: %s", msg.Payload(), msg.Topic())
    })
    opts.SetAutoReconnect(true)
    opts.SetConnectTimeout(10 * time.Second)

    client := mqtt.NewClient(opts)
    if token := client.Connect(); token.Wait() && token.Error() != nil {
        return nil, token.Error()
    }
    return &MQTTClient{client: client}, nil
}

func (m *MQTTClient) Publish(topic string, qos byte, retained bool, payload []byte) error {
    token := m.client.Publish(topic, qos, retained, payload)
    token.Wait()
    return token.Error()
}

func (m *MQTTClient) Subscribe(topic string, qos byte, handler mqtt.MessageHandler) error {
    token := m.client.Subscribe(topic, qos, handler)
    token.Wait()
    return token.Error()
}

func (m *MQTTClient) Close() {
    m.client.Disconnect(250)
}
```

### ตัวอย่าง: IoT Sensor Data Collector

```go
func main() {
    mqttClient, err := NewMQTTClient("tcp://localhost:1883", "go_collector", "", "")
    if err != nil {
        panic(err)
    }
    defer mqttClient.Close()

    // Subscribe to all sensor data topics
    err = mqttClient.Subscribe("sensors/+/data", 0, func(client mqtt.Client, msg mqtt.Message) {
        var data SensorData
        if err := json.Unmarshal(msg.Payload(), &data); err != nil {
            log.Printf("Invalid JSON: %v", err)
            return
        }
        // Save to InfluxDB or process
        fmt.Printf("Received from %s: temp=%.2f, humidity=%.2f\n",
            data.DeviceID, data.Temp, data.Humidity)
    })
    if err != nil {
        panic(err)
    }

    // Keep the program running
    select {}
}
```

---

## บทที่ 55: InfluxDB – Time‑Series Database

### InfluxDB คืออะไร?

InfluxDB เป็นฐานข้อมูล time‑series ที่ออกแบบมาเพื่อจัดเก็บและสืบค้นข้อมูลที่มี timestamp เช่น ข้อมูลเซ็นเซอร์, metrics, logs มีประสิทธิภาพสูงและมี query language เฉพาะ (Flux หรือ InfluxQL)

### การติดตั้งและเชื่อมต่อ

```go
go get github.com/influxdata/influxdb-client-go/v2
```

```go
package influx

import (
    "context"
    "time"

    influxdb2 "github.com/influxdata/influxdb-client-go/v2"
    "github.com/influxdata/influxdb-client-go/v2/api"
)

type InfluxDB struct {
    client   influxdb2.Client
    writeAPI api.WriteAPIBlocking
    queryAPI api.QueryAPI
    org      string
    bucket   string
}

func NewInfluxDB(url, token, org, bucket string) *InfluxDB {
    client := influxdb2.NewClient(url, token)
    writeAPI := client.WriteAPIBlocking(org, bucket)
    queryAPI := client.QueryAPI(org)
    return &InfluxDB{
        client:   client,
        writeAPI: writeAPI,
        queryAPI: queryAPI,
        org:      org,
        bucket:   bucket,
    }
}

func (i *InfluxDB) WritePoint(measurement string, tags map[string]string, fields map[string]interface{}, t time.Time) error {
    p := influxdb2.NewPoint(measurement, tags, fields, t)
    return i.writeAPI.WritePoint(context.Background(), p)
}

func (i *InfluxDB) Query(flux string) ([]api.FluxRecord, error) {
    result, err := i.queryAPI.Query(context.Background(), flux)
    if err != nil {
        return nil, err
    }
    var records []api.FluxRecord
    for result.Next() {
        records = append(records, result.Record())
    }
    return records, result.Err()
}
```

### ตัวอย่างการใช้งาน

```go
// Write sensor data
influx.WritePoint("temperature",
    map[string]string{"device": "sensor1", "location": "room101"},
    map[string]interface{}{"value": 25.6},
    time.Now(),
)

// Query average temperature for last hour
flux := `from(bucket:"mybucket")
  |> range(start: -1h)
  |> filter(fn: (r) => r._measurement == "temperature")
  |> filter(fn: (r) => r.device == "sensor1")
  |> aggregateWindow(every: 5m, fn: mean)`
records, err := influx.Query(flux)
```

---

## บทที่ 56: WebSocket และ Socket.IO

### WebSocket พื้นฐาน

WebSocket เป็นโปรโตคอลที่ช่วยให้ client-server สื่อสารแบบ full-duplex เหนือ TCP เหมาะสำหรับ real-time applications เช่น chat, live notifications

**การติดตั้ง gorilla/websocket**
```go
go get github.com/gorilla/websocket
```

### ตัวอย่าง WebSocket Server

```go
package websocket

import (
    "log"
    "net/http"
    "sync"

    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool { return true },
}

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
    mu         sync.RWMutex
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
            h.mu.Lock()
            h.clients[client] = true
            h.mu.Unlock()
        case client := <-h.unregister:
            h.mu.Lock()
            if _, ok := h.clients[client]; ok {
                delete(h.clients, client)
                close(client.send)
            }
            h.mu.Unlock()
        case message := <-h.broadcast:
            h.mu.RLock()
            for client := range h.clients {
                select {
                case client.send <- message:
                default:
                    close(client.send)
                    delete(h.clients, client)
                }
            }
            h.mu.RUnlock()
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

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool { return true },
}

func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
        return
    }
    client := &Client{conn: conn, send: make(chan []byte, 256), hub: hub}
    hub.register <- client

    go client.writePump()
    go client.readPump()
}
```

### การใช้งานใน main

```go
hub := NewHub()
go hub.Run()
http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
    ServeWs(hub, w, r)
})
```

---

## บทที่ 57: การส่ง SMS และ LINE Notify

### 1. ส่ง SMS ผ่าน Twilio (ตัวอย่าง)

```go
package sms

import (
    "github.com/twilio/twilio-go"
    twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

type TwilioSender struct {
    client     *twilio.RestClient
    fromNumber string
}

func NewTwilioSender(accountSid, authToken, fromNumber string) *TwilioSender {
    client := twilio.NewRestClientWithParams(twilio.ClientParams{
        Username: accountSid,
        Password: authToken,
    })
    return &TwilioSender{client: client, fromNumber: fromNumber}
}

func (t *TwilioSender) Send(to, message string) error {
    params := &twilioApi.CreateMessageParams{}
    params.SetTo(to)
    params.SetFrom(t.fromNumber)
    params.SetBody(message)

    _, err := t.client.Api.CreateMessage(params)
    return err
}
```

### 2. LINE Notify – การแจ้งเตือนผ่าน LINE

**LINE Notify** ให้บริการแจ้งเตือนฟรี โดยต้องสร้าง token ที่ https://notify-bot.line.me/

```go
package linenotify

import (
    "bytes"
    "fmt"
    "io"
    "mime/multipart"
    "net/http"
    "os"
)

type LINE struct {
    token string
}

func NewLINE(token string) *LINE {
    return &LINE{token: token}
}

func (l *LINE) SendMessage(message string) error {
    url := "https://notify-api.line.me/api/notify"
    data := []byte(fmt.Sprintf("message=%s", message))
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
    if err != nil {
        return err
    }
    req.Header.Set("Authorization", "Bearer "+l.token)
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("LINE Notify returned status %d", resp.StatusCode)
    }
    return nil
}

func (l *LINE) SendImage(message, imagePath string) error {
    var b bytes.Buffer
    w := multipart.NewWriter(&b)

    w.WriteField("message", message)

    file, err := os.Open(imagePath)
    if err != nil {
        return err
    }
    defer file.Close()

    part, err := w.CreateFormFile("imageFile", imagePath)
    if err != nil {
        return err
    }
    io.Copy(part, file)
    w.Close()

    req, err := http.NewRequest("POST", "https://notify-api.line.me/api/notify", &b)
    if err != nil {
        return err
    }
    req.Header.Set("Authorization", "Bearer "+l.token)
    req.Header.Set("Content-Type", w.FormDataContentType())

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("LINE Notify returned status %d", resp.StatusCode)
    }
    return nil
}
```

---

## บทที่ 58: Discord Webhook สำหรับแจ้งเตือน

### Discord Webhook คืออะไร?

Discord Webhook เป็นวิธีส่งข้อความอัตโนมัติไปยัง channel ใน Discord โดยไม่ต้องใช้ bot token เหมาะสำหรับการแจ้งเตือนจากระบบ (เช่น deployment status, error logs)

### การส่งข้อความผ่าน Go

```go
package discord

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
)

type Webhook struct {
    url string
}

type WebhookPayload struct {
    Content   string  `json:"content,omitempty"`
    Username  string  `json:"username,omitempty"`
    AvatarURL string  `json:"avatar_url,omitempty"`
    Embeds    []Embed `json:"embeds,omitempty"`
}

type Embed struct {
    Title       string  `json:"title,omitempty"`
    Description string  `json:"description,omitempty"`
    Color       int     `json:"color,omitempty"` // decimal color
    Fields      []Field `json:"fields,omitempty"`
    Timestamp   string  `json:"timestamp,omitempty"`
}

type Field struct {
    Name   string `json:"name"`
    Value  string `json:"value"`
    Inline bool   `json:"inline,omitempty"`
}

func NewWebhook(url string) *Webhook {
    return &Webhook{url: url}
}

func (w *Webhook) SendMessage(content string) error {
    payload := WebhookPayload{Content: content}
    return w.send(payload)
}

func (w *Webhook) SendEmbed(embed Embed) error {
    payload := WebhookPayload{Embeds: []Embed{embed}}
    return w.send(payload)
}

func (w *Webhook) send(payload WebhookPayload) error {
    data, err := json.Marshal(payload)
    if err != nil {
        return err
    }
    resp, err := http.Post(w.url, "application/json", bytes.NewBuffer(data))
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
        return fmt.Errorf("discord webhook returned status %d", resp.StatusCode)
    }
    return nil
}
```

### ตัวอย่างการแจ้งเตือน Deployment

```go
webhook := discord.NewWebhook("https://discord.com/api/webhooks/...")

// Simple message
webhook.SendMessage("🚀 Deployment started: v1.2.3")

// Embed with details
embed := discord.Embed{
    Title:       "Deployment Successful",
    Description: "Version v1.2.3 has been deployed to production.",
    Color:       0x00ff00, // green
    Fields: []discord.Field{
        {Name: "Environment", Value: "Production", Inline: true},
        {Name: "Version", Value: "v1.2.3", Inline: true},
        {Name: "Deployed By", Value: "CI/CD", Inline: false},
    },
    Timestamp: time.Now().Format(time.RFC3339),
}
webhook.SendEmbed(embed)
```

---

## สรุป

ภาคที่ 9 ได้นำเสนอการผสานระบบภายนอกที่จำเป็นสำหรับแอปพลิเคชันสมัยใหม่:

| ระบบ | การใช้งานหลัก | ตัวอย่าง |
|------|-------------|---------|
| Redis | Cache, Message Queue | Caching user data, job queue |
| RabbitMQ | Message Broker | Async email sending, order processing |
| MQTT | IoT Communication | Sensor data collection |
| InfluxDB | Time‑Series Database | Metrics storage, telemetry |
| WebSocket | Real‑time Communication | Chat, live notifications |
| SMS | Mobile alerts | OTP, transaction alerts |
| LINE Notify | Instant messaging | System alerts, notifications |
| Discord | Team communication | Deployment status, error logs |

การรวมระบบเหล่านี้เข้ากับ Go application ช่วยให้สามารถสร้างแอปพลิเคชันที่มีความสามารถครบครัน รองรับการทำงานแบบ real‑time, event‑driven, และการแจ้งเตือนหลากหลายช่องทาง

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