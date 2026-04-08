เรื่อง :คู่มือการเพิ่มระบบ Monitoring สำหรับทีมพัฒนา Go  
คุณคือ :Technical lead  and Golang Developer
วัตุประสงค์ : คู่มือนี้ช่วยให้ทีม Go เพิ่ม **Monitoring และ Health Checks Logging, Metrics, Tracing, Error Tracking, Database Monitoring,Network,CPU, 
RAM Memory, และ Performance Monitoring** 
ใช้ slog, Prometheus, OpenTelemetry (Jaeger), Sentry พร้อมเทมเพลตโค้ดและ Checklist สำหรับ Developer และ DevOps
 Complete Monitoring   Golang.
ทำ เป็น Rest API module ใหม่ แแยกจากของเดิม จะได้ไม่กระลของเดิม
Alert to email

ข้อกำหนด:
เอกสาร รูปแบบ Markdown ที่มีเนื้อหาทั้งภาษาไทยและภาษาอังกฤษ (สลับกันตามหัวข้อ)

## 📁 โครงสร้างโปรเจกต์ `icmongolang` เดิม

```text
icmongolang/
├── main.go
├── cmd/                          # Entry points ของแอปพลิเคชัน (command line)
│   ├── api/                      # เริ่มต้น REST API server
│   │   └── main.go               # จุดเริ่มต้นหลักของโปรแกรม
│   ├── initdata.go               # คำสั่งสำหรับเติมข้อมูลเริ่มต้น (seed data)
│   ├── migrate.go                # คำสั่ง run database migration
│   ├── root.go                   # root command ของ CLI (ใช้ Cobra)
│   ├── serve.go                  # คำสั่ง start server
│   └── worker.go                 # คำสั่ง start background worker (เช่น email)
│
├── config/                        
│   ├── swagger.json          
│   ├── swagger.yaml           
│   └── docs.go                 
│
├── docs/                         # เอกสารทั่วไป (API docs, specs)
│   ├── monitoring.go            
│   ├── rate_limit.go           
│   └── security.go    
│
├── internal/                     # โค้ดส่วน private ของโปรเจกต์ (ใช้ภายในเท่านั้น)
│   ├── auth/                    
│ 	│   ├── delivery/                 
│ 	│   │   └── http/                  
│ 	│   │       ├── handlers.go
│ 	│   │       └── routes.go  
│   │   ├── handler.go                
│   ├── distributor/                  
│   │   └──distributor.go 
│   ├── distributor/                  
│   │   └──distributor.go                          
│   ├── items/                 
│ 	│   ├── delivery/                 
│ 	│   │   └── http/                  
│ 	│   │       ├── handlers.go
│ 	│   │       └── routes.go  
│ 	│   ├── presenter/                                 
│ 	│   │    └── presenter.go  
│ 	│   ├── repository/                                 
│ 	│   │    └── pg_repository.go  
│ 	│   ├── usecase/                                 
│ 	│   │    └── usecase.go 
│ 	│   ├──handler.go   
│ 	│   ├──pg_repository.go   
│ 	│   ├──usecase.go   
│   ├── server/                  
│   │   ├── handlers.go          
│   │   └── server.go  
│   ├── usecase/                   
│   │   └── usecase.go  
│   ├── users/                 
│ 	│   ├── delivery/                 
│ 	│   │   └── http/                  
│ 	│   │       ├── handlers.go
│ 	│   │       └── routes.go  
│ 	│   ├── distributor/                 
│ 	│   │   └──distributor.go  
│ 	│   ├── presenter/                                 
│ 	│   │    └── presenter.go  
│ 	│   ├── processor/                                 
│ 	│   │    └── processor.go  
│ 	│   ├── repository/          
│ 	│   │    ├── pg_repository.go                            
│ 	│   │    └── redis_repository.go  
│ 	│   ├── usecase/                                 
│ 	│   │    └── usecase.go 
│   ├── worker/                   
│   │   └── worker.go  
│ 	│   ├──pg_repository.go   
│ 	│   ├──redis_repository.go   
│ 	│   └──usecase.go 
│   ├── models/                   # Entity / data model (structs)
│   │   ├── base.go               # Base model (เช่น ID, timestamps)
│   │   ├── session.go            # Session model
│   │   ├── models.go             # models model
│   │   ├── user.go               # User model
│   │   └── verification.go       # Email verification model
│   ├── middleware/                   # Entity / data model (structs)
│   │   ├── cors.go                
│   │   ├── jwtauth.go            
│   │   ├── logging.go          
│   │   ├── middleware.go           
│   │   ├── monitoring.go            
│   │   ├── rate_limit.go           
│   │   └── security.go        
│   ├── repository/               # Data access layer (database, cache)
│   │   ├── pg_repository.go      # PostgreSQL interface/repo
│   │   ├── redis_repo.go         # Redis interface/repo
│   │   ├── session_repo.go       # Session repository methods
│   │   └── user_repo.go          # User repository methods
│   │
│   ├── usecase/                  # Business logic layer
│   │   ├── auth_usecase.go       # Auth logic (login, register, logout)
│   │   ├── cache_usecase.go      # Cache logic
│   │   └── user_usecase.go       # User business logic
│   │
│   ├── delivery/                 # Delivery mechanism (HTTP, workers)
│   │   ├── rest/                 # REST API delivery
│   │   │   ├── handler/          # HTTP handlers
│   │   │   │   ├── auth_handler.go
│   │   │   │   ├── health_handler.go
│   │   │   │   └── user_handler.go
│   │   │   ├── middleware/       # HTTP middlewares
│   │   │   │   ├── auth.go       # JWT auth
│   │   │   │   ├── cors.go       # CORS
│   │   │   │   ├── logger.go     # Logging
│   │   │   │   ├── monitoring.go # Metrics
│   │   │   │   ├── rate_limit.go # Rate limiting
│   │   │   │   └── security.go   # Security headers
│   │   │   ├── dto/              # Data Transfer Objects
│   │   │   │   ├── auth_dto.go
│   │   │   │   ├── error_dto.go
│   │   │   │   └── user_dto.go
│   │   │   └── router.go         # Route registration
│   │   └── worker/               # Background worker handlers
│   │       └── email_worker.go   # จัดการส่งอีเมล async
│   │
│   └── pkg/                      # Internal shared packages
│       ├── email/                # Email sending logic
│       │   ├── gomail_sender.go
│       │   ├── sender.go
│       │   └── templates/        # HTML email templates
│       │       ├── reset_password.html
│       │       └── verification.html
│       ├── hash/                 # Password hashing (bcrypt)
│       │   └── bcrypt.go
│       ├── jwt/                  # JWT generation & verification
│       │   ├── maker.go
│       │   ├── payload.go
│       │   └── rsa_maker.go      # RSA-based JWT
│       ├── logger/               # Structured logging (Zap)
│       │   └── zap_logger.go
│       ├── redis/                # Redis client & cache helpers
│       │   ├── cache.go
│       │   ├── client.go
│       │   └── refresh_store.go  # Store refresh tokens
│       ├── utils/                # Utility functions
│       │   ├── random.go         # สุ่ม string, numbers
│       │   └── time.go           # time helpers
│       └── validator/            # Custom validation
│           └── custom_validator.go
│
├── migrations/                   # SQL migration files (up/down)
│   ├── 000001_create_users_table.down.sql
│   └── 000001_create_users_table.up.sql
│
├── pkg/                          # โค้ดที่สามารถนำไปใช้ภายนอกได้ (public)
│   └── utils/                    # (ซ้ำกับ internal/pkg? อาจจะแยก)
│
├── scripts/                      # Scripts สำหรับ build, deploy
│   ├── build.sh
│   └── deploy.sh
│
├── vendor/                       # Dependency ที่ vendor ไว้ (ถ้าใช้ vendor mode)
│
├── .air.toml                     # คอนฟิกสำหรับ live reload (Air)
├── .dockerignore                 # ไฟล์ที่ docker ignore
├── .env.dev                      # environment variables สำหรับ dev
├── .env.prod                     # environment variables สำหรับ prod
├── .gitignore
├── docker-compose.dev.yml        # Docker compose สำหรับ development
├── docker-compose.prod.yml       # Docker compose สำหรับ production
├── Dockerfile.dev                # Docker image สำหรับ dev
├── Dockerfile.prod               # Docker image สำหรับ prod
├── go.mod                        # Go module definition
├── go.sum                        # Checksums ของ dependencies
├── LICENSE                       # สัญญาอนุญาต
└── README.md                     # คำอธิบายโปรเจกต์
```

ออกแบบ ใช้ใชได้กับ ของเดิม

**คำอธิบายแต่ละโฟลเดอร์/ไฟล์** (ไทย/อังกฤษ):
   - โครงสร้างการทำงาน
   - วัตุประสงค์
   - กลุ่มเป้าหมาย
   - ความรู้พื้นฐาน
   - เนื้อหา โดยย่อ กระชับ เน้น วัตถุประสงค์  ประโยชน์ของการใช้
   - บทนำ
   - บทนิยาม
   - ออกแบบ workflow
     - วาดรูป dataflow  2D
     - วาดรูป dataflow สร้าง รูปแบบ เหมือนจริง ลักษณะ flowchart คลาดกับ draw.io Diagrams เพื่ออธิบายกระบวนการ ทำความเข้าใจ ระวัง อักขระพิเศษ ป้องการการ เปิดอ่านแล้ว error
     - พร้อมอธิบาย แบบ ละเอียด 
     - คอมเม้น code ภาษาไทย และ ภาษาอังถถษ อธิบาย การทำงาน แต่ละจุด
     - ยกตัวอย่างการใช้งานจริง หรือ กรณีศึกษา แนวทางแก้ไขปัญหา ที่อาจจะเกิดขึ้น
     - เทมเพลต และ ตัวอย่างโค้ด พร้อมนำไป run ได้ทันที  มีคำอธิบายการใช้งานแต่ละจุด การคอมเม้น  
	-  พร้อมสร้างเอกสาร
	-  หลักการ (Concept)
	-  คืออะไร?
	-  มีกี่แบบ?  
	-  ใช้อย่างไร 
	-  นำไปใช้กรณีไหน
	-  ข้อห้ามสำคัญ 
	-  ประโยชน์ที่ได้รับ
	-  ข้อดี
	-  ข้อเสีย
	-  ข้อห้าม 
	-  ข้อควรระวัง
	-  คู่มือการทดสอบ
	-  CHECK List  การทดสอบ
	-  คู่มือการการใช้งาน
	-  CHECK List การการใช้งาน 
	-  คู่มือการบำรุงรักษา
	-  CHECK List การบำรุงรักษา   
	-  การตรวจสอบคสามปลดภัย และความเสี่ยง
	-  คู่มือการขยาย หรือแก้ไข หรือ เพิมเติม ในอนาคต
	-  CHECK List Test Module
   - สรุป
      -ประโยชน์ที่ได้รับ
      -ข้อควรระวัง
      -ข้อดี
      -ข้อเสีย
   -ข้อห้าม ถ้ามี
   -  การออกแบบ Workflow และ Dataflow ระวัง อักขระ พิเศษ จำทำให้รูปแสดงไม่ได้ ระวังให้ดี
   -  สร้าง CODE folder Proje / Modules โครงสร้างโฟลเดอร์/ไฟล์ (Folder/File Structure)
   -   สร้างโค้ดที่รันได้จริง ตามโครงสร้งนี้ ทุกส่วน
   -  คอมเม้น CODE ไทย อังกถษ คนละบรรทัด
   -  ออกแบบ Git Flow แตก Branch  
   -  ออกแบบ การ push pull mearh code
	-  การออกแบบ Git Flow (Branching Strategy)
		- ใช้แนวทางมาตรฐาน (Git Flow) ที่แบ่งสาขาชัดเจนเพื่อแยกงาน  
		- main (หรือ master): เก็บโค้ดที่เสถียรพร้อมใช้งานจริง (Stable & Production)
		- develop: สาขาหลักที่ใช้รวมฟีเจอร์ต่างๆ เพื่อทดสอบ
		- feature/...: แตกจาก develop เพื่อทำฟีเจอร์ใหม่ เมื่อเสร็จแล้วให้ทำ Pull Request กลับมาที่ develop
		- release/...: แตกจาก develop เพื่อเตรียมปล่อยเวอร์ชันใหม่ ทำ QA และแก้บั๊กสุดท้าย ก่อนจะ Merge เข้า main และ develop
		- hotfix/...: แตกจาก main เพื่อแก้บั๊กเร่งด่วนบน Production แล้ว Merge เข้าทั้ง main และ develop 
	-  ขั้นตอนการทำงาน (Workflow):
		- Start: แตก feature/login จาก develop
		- Dev: พัฒนาและ commit ใน feature/login
		- Review: สร้าง Pull Request (PR) จาก feature/login -> develop
		- Test: เมื่อได้รอบการปล่อย ให้ทำ release/1.0.0 จาก develop เพื่อ QA
		- Deploy: Merge release/1.0.0 -> main และทำ Tag เวอร์ชัน
		- CHECK List
	
	-  การกำหนดค่า environment
	- environment server localhost 
	- environment server dev 
	- environment server uat 
	- environment server production 
		- Localhost (Local Development): สภาพแวดล้อมบนเครื่องคอมพิวเตอร์ของนักพัฒนา (Local Machine) 
		ใช้สำหรับการเขียนโค้ดและทดสอบเบื้องต้น โดยจะเชื่อมต่อกับฐานข้อมูลหรือบริการภายในเครื่องตัวเอง
		- Dev Server (Development): เซิร์ฟเวอร์ที่รวมโค้ดจากนักพัฒนาหลายๆ คนเพื่อทดสอบร่วมกัน (Integration) ซึ่งอาจจะยังไม่เสถียรเท่าที่ควร
		- UAT Server (User Acceptance Testing): สภาพแวดล้อมจำลองที่เหมือน Production ที่สุด 
		เพื่อให้ผู้ใช้งานจริง (User/Client) ทดสอบฟังก์ชันการทำงานก่อนใช้งานจริง
		- Production Server (PROD): สภาพแวดล้อมจริงที่แอปพลิเคชันเปิดใช้งานสำหรับผู้ใช้งานทั่วไป 
		เป็นจุดที่ต้องการความเสถียรและความปลอดภัยสูงสุด 
		วิธีการตั้งค่าตัวแปร (Environment Variables):
		การกำหนดค่าสำหรับแต่ละสภาพแวดล้อมมักจะทำผ่านไฟล์ .env ที่แตกต่างกัน (เช่น .env.dev, .env.prod) 
		โดยจะ ไม่ เก็บไฟล์เหล่านี้ไว้ในระบบ Git เพื่อความปลอดภัยของข้อมูล 
 	-  Global Protect
	-  สรุป ภาษาไทย ภาษาอังถถษ
 
 
 
เรื่อง :คู่มือการเพิ่มระบบ Monitoring สำหรับทีมพัฒนา Go  
คุณคือ :Technical lead  and Golang Developer
วัตุประสงค์ : คู่มือนี้ช่วยให้ทีม Go เพิ่ม **Logging, Metrics, Tracing, Error Tracking, Database Monitoring,Network,CPU, RAM Memory, และ Performance Monitoring** 
ใช้ slog, Prometheus, OpenTelemetry (Jaeger), Sentry พร้อมเทมเพลตโค้ดและ Checklist สำหรับ Developer และ DevOps
Prometheus & Grafana. Complete Monitoring   Golang.
Alert to email

ข้อกำหนด:
เอกสาร รูปแบบ Markdown ที่มีเนื้อหาทั้งภาษาไทยและภาษาอังกฤษ (สลับกันตามหัวข้อ)
**คำอธิบายแต่ละโฟลเดอร์/ไฟล์** (ไทย/อังกฤษ):
   - โครงสร้างการทำงาน
   - วัตุประสงค์
   - กลุ่มเป้าหมาย
   - ความรู้พื้นฐาน
   - เนื้อหา โดยย่อ กระชับ เน้น วัตถุประสงค์  ประโยชน์ของการใช้
   - บทนำ
   - บทนิยาม
   - ออกแบบ workflow
     - วาดรูป dataflow  2D
     - วาดรูป dataflow สร้าง รูปแบบ เหมือนจริง ลักษณะ flowchart คลาดกับ draw.io Diagrams เพื่ออธิบายกระบวนการ ทำความเข้าใจ ระวัง อักขระพิเศษ ป้องการการ เปิดอ่านแล้ว error
     - พร้อมอธิบาย แบบ ละเอียด 
     - คอมเม้น code ภาษาไทย และ ภาษาอังถถษ อธิบาย การทำงาน แต่ละจุด
     - ยกตัวอย่างการใช้งานจริง หรือ กรณีศึกษา แนวทางแก้ไขปัญหา ที่อาจจะเกิดขึ้น
     - เทมเพลต และ ตัวอย่างโค้ด พร้อมนำไป run ได้ทันที  มีคำอธิบายการใช้งานแต่ละจุด การคอมเม้น  
	-  พร้อมสร้างเอกสาร
	-  หลักการ (Concept)
	-  คืออะไร?
	-  มีกี่แบบ?  
	-  ใช้อย่างไร 
	-  นำไปใช้กรณีไหน
	-  ข้อห้ามสำคัญ 
	-  ประโยชน์ที่ได้รับ
	-  ข้อดี
	-  ข้อเสีย
	-  ข้อห้าม 
	-  ข้อควรระวัง
	-  คู่มือการทดสอบ
	-  CHECK List  การทดสอบ
	-  คู่มือการการใช้งาน
	-  CHECK List การการใช้งาน 
	-  คู่มือการบำรุงรักษา
	-  CHECK List การบำรุงรักษา   
	-  การตรวจสอบคสามปลดภัย และความเสี่ยง
	-  คู่มือการขยาย หรือแก้ไข หรือ เพิมเติม ในอนาคต
	-  CHECK List Test Module
   - สรุป
      -ประโยชน์ที่ได้รับ
      -ข้อควรระวัง
      -ข้อดี
      -ข้อเสีย
   -ข้อห้าม ถ้ามี
   -  การออกแบบ Workflow และ Dataflow ระวัง อักขระ พิเศษ จำทำให้รูปแสดงไม่ได้ ระวังให้ดี
   -  สร้าง CODE folder Proje / Modules โครงสร้างโฟลเดอร์/ไฟล์ (Folder/File Structure)
   -   สร้างโค้ดที่รันได้จริง ตามโครงสร้งนี้ ทุกส่วน
   -  คอมเม้น CODE ไทย อังกถษ คนละบรรทัด
   -  ออกแบบ Git Flow แตก Branch  
   -  ออกแบบ การ push pull mearh code
	-  การออกแบบ Git Flow (Branching Strategy)
		- ใช้แนวทางมาตรฐาน (Git Flow) ที่แบ่งสาขาชัดเจนเพื่อแยกงาน  
		- main (หรือ master): เก็บโค้ดที่เสถียรพร้อมใช้งานจริง (Stable & Production)
		- develop: สาขาหลักที่ใช้รวมฟีเจอร์ต่างๆ เพื่อทดสอบ
		- feature/...: แตกจาก develop เพื่อทำฟีเจอร์ใหม่ เมื่อเสร็จแล้วให้ทำ Pull Request กลับมาที่ develop
		- release/...: แตกจาก develop เพื่อเตรียมปล่อยเวอร์ชันใหม่ ทำ QA และแก้บั๊กสุดท้าย ก่อนจะ Merge เข้า main และ develop
		- hotfix/...: แตกจาก main เพื่อแก้บั๊กเร่งด่วนบน Production แล้ว Merge เข้าทั้ง main และ develop 
	-  ขั้นตอนการทำงาน (Workflow):
		- Start: แตก feature/login จาก develop
		- Dev: พัฒนาและ commit ใน feature/login
		- Review: สร้าง Pull Request (PR) จาก feature/login -> develop
		- Test: เมื่อได้รอบการปล่อย ให้ทำ release/1.0.0 จาก develop เพื่อ QA
		- Deploy: Merge release/1.0.0 -> main และทำ Tag เวอร์ชัน
		- CHECK List
	
	-  การกำหนดค่า environment
	- environment server localhost 
	- environment server dev 
	- environment server uat 
	- environment server production 
		- Localhost (Local Development): สภาพแวดล้อมบนเครื่องคอมพิวเตอร์ของนักพัฒนา (Local Machine) 
		ใช้สำหรับการเขียนโค้ดและทดสอบเบื้องต้น โดยจะเชื่อมต่อกับฐานข้อมูลหรือบริการภายในเครื่องตัวเอง
		- Dev Server (Development): เซิร์ฟเวอร์ที่รวมโค้ดจากนักพัฒนาหลายๆ คนเพื่อทดสอบร่วมกัน (Integration) ซึ่งอาจจะยังไม่เสถียรเท่าที่ควร
		- UAT Server (User Acceptance Testing): สภาพแวดล้อมจำลองที่เหมือน Production ที่สุด 
		เพื่อให้ผู้ใช้งานจริง (User/Client) ทดสอบฟังก์ชันการทำงานก่อนใช้งานจริง
		- Production Server (PROD): สภาพแวดล้อมจริงที่แอปพลิเคชันเปิดใช้งานสำหรับผู้ใช้งานทั่วไป 
		เป็นจุดที่ต้องการความเสถียรและความปลอดภัยสูงสุด 
		วิธีการตั้งค่าตัวแปร (Environment Variables):
		การกำหนดค่าสำหรับแต่ละสภาพแวดล้อมมักจะทำผ่านไฟล์ .env ที่แตกต่างกัน (เช่น .env.dev, .env.prod) 
		โดยจะ ไม่ เก็บไฟล์เหล่านี้ไว้ในระบบ Git เพื่อความปลอดภัยของข้อมูล 
 	-  Global Protect
	-  สรุป ภาษาไทย ภาษาอังถถษ
 
 
 