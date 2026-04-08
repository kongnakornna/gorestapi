
### โฟลเดอร์หลัก  icmongolang
ขออธิบายโครงสร้างโปรเจกต์ **icmongolang** ซึ่งเป็น Go Backend Project ที่ใช้สถาปัตยกรรมแบบ Clean Architecture (หรือ Modular) พร้อมคอมเมนต์ทั้งภาษาไทยและอังกฤษครับ

## 📁 โครงสร้างโปรเจกต์ `icmongolang` – คำอธิบาย (Thai & English)

```text
icmongolang/
├── cmd/                          # Entry points ของแอปพลิเคชัน (command line)
│   ├── api/                      # เริ่มต้น REST API server
│   │   └── main.go               # จุดเริ่มต้นหลักของโปรแกรม
│   ├── initdata.go               # คำสั่งสำหรับเติมข้อมูลเริ่มต้น (seed data)
│   ├── migrate.go                # คำสั่ง run database migration
│   ├── root.go                   # root command ของ CLI (ใช้ Cobra)
│   ├── serve.go                  # คำสั่ง start server
│   └── worker.go                 # คำสั่ง start background worker (เช่น email)
│
├── config/                       # จัดการคอนฟิกเรชัน
│   ├── config-local.yml          # คอนฟิกสำหรับ environment local
│   ├── config-prod.yml           # คอนฟิกสำหรับ production
│   └── config.go                 # โค้ดโหลดและ parse คอนฟิก
│
├── docdev/                       # เอกสารสำหรับนักพัฒนาภายใน
├── docs/                         # เอกสารทั่วไป (API docs, specs)
│
├── internal/                     # โค้ดส่วน private ของโปรเจกต์ (ใช้ภายในเท่านั้น)
│   ├── models/                   # Entity / data model (structs)
│   │   ├── base.go               # Base model (เช่น ID, timestamps)
│   │   ├── session.go            # Session model
│   │   ├── user.go               # User model
│   │   └── verification.go       # Email verification model
│   │
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
├── README.md                     # คำอธิบายโปรเจกต์
└── BookGolang.md                 # เอกสารหรือบันทึกส่วนตัวเกี่ยวกับ Go
```

---

## 🧠 สถาปัตยกรรมหลัก (Core Architecture)

โครงสร้างนี้ใช้ **Clean Architecture** (หรือ Hexagonal) แบ่งเป็น 3 ชั้นหลัก:

| Layer          | หน้าที่ (Thai)                                 | Responsibility (English)                    |
|----------------|-----------------------------------------------|---------------------------------------------|
| **Delivery**   | จัดการ request/response (HTTP, worker)         | Handles incoming requests (HTTP, workers)   |
| **Usecase**    | บรรจุ business logic ของแอป                    | Contains application business logic         |
| **Repository** | ติดต่อกับ data source (DB, Redis, external API) | Communicates with data sources (DB, Redis)  |
| **Models**     | Entity / domain object ที่ใช้ร่วมกันทุกชั้น      | Shared domain entities across all layers    |

---

## ✅ จุดเด่นของโครงสร้างนี้ (Key strengths)

### ภาษาไทย
- **แยกความรับผิดชอบชัดเจน** (Separation of concerns) – แต่ละชั้นทำหน้าที่ของตัวเอง ทำให้ทดสอบและบำรุงรักษาง่าย
- **รองรับการทำงานหลาย environment** (local, prod) ผ่านไฟล์ config และ Docker compose
- **มีระบบย่อยครบครัน**: JWT auth, rate limiting, logging (Zap), migration, email template, background worker
- **ใช้ dependency injection** ผ่าน constructor ทำให้โค้ด loose coupling
- **พร้อม deploy** ด้วย Docker และมี scripts ช่วย

### English
- **Clear separation of concerns** – each layer has a single responsibility, easy to test and maintain.
- **Multi‑environment support** (local, prod) via config files and Docker compose.
- **Complete sub‑systems**: JWT auth, rate limiting, logging (Zap), migrations, email templates, background workers.
- **Dependency injection** via constructors – results in loose coupling.
- **Ready to deploy** with Docker and helper scripts.

---

## 📌 หมายเหตุเพิ่มเติม (Additional notes)

1. มีโฟลเดอร์ `internal/pkg` และ `pkg/utils` ซ้อนกัน – อาจต้องการปรับโครงสร้างให้รวมกัน (หรือ `pkg/utils` ถูกสร้างไว้เผื่อ public reuse)
2. ไฟล์ `BookGolang.md` – น่าจะเป็นเอกสารการเรียนรู้ส่วนตัว
3. ใช้ `cmd/` แยกหลาย entry points (api, worker, migrate, initdata) – เหมาะกับ microservices หรือ CLI tools

ถ้าต้องการให้อธิบายส่วนไหนเพิ่มเติม บอกได้เลยครับ 🙌
โครงสร้างโปรเจกต์ `icmongolang` ที่คุณส่งมา เป็นตัวอย่างของ **Go Project Layout แบบ Standard (หรือที่เรียกว่า Standard Go Project Layout)** ซึ่งได้รับความนิยมในชุมชน Go สำหรับโปรเจกต์ขนาดกลางถึงใหญ่ โดยเฉพาะที่ใช้ Clean Architecture, DDD หรือ Hexagonal Architecture

 

---

## 🧱 1. โครงสร้างนี้คืออะไร?

คือ **รูปแบบการจัดเรียงโฟลเดอร์และไฟล์ในโปรเจกต์ Go** ที่แบ่งตามหน้าที่ (separation of concerns) อย่างชัดเจน โดยยึดตามแนวทางของ `golang-standards/project-layout` (ไม่ใช่ official standard แต่เป็น de facto standard ที่ใช้กันแพร่หลาย)

โฟลเดอร์หลัก ๆ มีความหมายดังนี้:

| โฟลเดอร์ | หน้าที่ |
|----------|--------|
| `cmd/` | เก็บ executable entry points (main packages) แต่ละตัวแยก subfolder เช่น `api/`, `worker/`, `migrate` |
| `internal/` | โค้ดส่วนตัวที่ **ไม่ถูก import** จากโปรเจกต์อื่น (private) ใช้สำหรับ business logic, repository, delivery |
| `pkg/` | โค้ดที่สามารถ **ถูก import** จากโปรเจกต์อื่นได้ (public library) เช่น utilities, email sender, JWT maker |
| `config/` | ไฟล์ configuration (yml, env) และโค้ดสำหรับโหลด config |
| `migrations/` | SQL migration files (up/down) สำหรับ database schema |
| `docs/` | เอกสาร API (เช่น Swagger/OpenAPI) หรือเอกสารทั่วไป |
| `scripts/` | build, deploy, หรือ utility scripts |
| `vendor/` | dependencies ที่ vendor ไว้ (optional) |
| `docker-compose*.yml`, `Dockerfile*` | container configuration |

---

## 📚 2. มีกี่แบบ? (รูปแบบโครงสร้างโปรเจกต์ Go)

โครงสร้าง Go มีหลายแบบ ขึ้นอยู่กับขนาดทีมและความซับซ้อน ที่พบบ่อย:

| รูปแบบ | คำอธิบาย | เหมาะกับ |
|--------|----------|----------|
| **Flat / Simple** | วาง `main.go`, `handlers.go`, `models.go` ไว้ที่ root | โปรเจกต์เล็ก, ไลบรารี, ตัวอย่าง |
| **Package-oriented** | จัดตาม package หลัก เช่น `user/`, `order/`, แต่ละ package มีทุกอย่าง (handler, repo, model) | โมดูลที่แยกอิสระ |
| **Standard layout** (แบบนี้) | มี `cmd`, `internal`, `pkg` ชัดเจน ตามแนวทาง community | โปรเจกต์ขนาดกลาง-ใหญ่, หลาย executable, ต้องการ clean architecture |
| **Modular monolith** | เหมือน standard layout แต่เพิ่ม modules ชัดเจน (เช่น `internal/users`, `internal/products`) | monolith ที่พร้อมแยกเป็น microservices ในอนาคต |
| **Hexagonal / Clean** | แยกตาม layer: `domain`, `application`, `infrastructure`, `interfaces` | เน้น decoupling สูง, เปลี่ยน framework ได้ง่าย |

**แบบที่คุณส่งมา** คือ **Standard layout** ผสมกับ **Clean architecture** (มี `delivery`, `usecase`, `repository`, `models`)

---

## 🛠️ 3. ใช้อย่างไร?

### 3.1 การวางโค้ดในแต่ละโฟลเดอร์ (ตามตัวอย่าง)

- **`cmd/api/main.go`** → เริ่มต้น HTTP server, inject dependencies, เรียก `router.go`
- **`internal/delivery/rest/router.go`** → ตั้งค่า Chi router, เรียก `MapUserRoute` และอื่น ๆ
- **`internal/delivery/rest/handler/`** → รับ request, validate, call usecase, return response
- **`internal/usecase/`** → business logic (ไม่รู้เรื่อง HTTP หรือ DB)
- **`internal/repository/`** → อ่าน/เขียน database, cache (implement interface ที่ usecase กำหนด)
- **`internal/models/`** → structs สำหรับ DB และ domain
- **`internal/pkg/`** → utilities เฉพาะภายในโปรเจกต์ (เช่น hash, jwt, logger, redis client) – *แต่จริง ๆ แล้วควรย้ายไป `pkg/` ถ้าสามารถ reuse ข้ามโปรเจกต์ได้*
- **`pkg/`** → ของที่ reusable จริง ๆ (เช่น `utils/random.go` แต่ในตัวอย่างมีแค่ `pkg/utils/` ว่าง)

### 3.2 ขั้นตอนการ build และ run

```bash
# รัน API
go run cmd/api/main.go

# รัน migration
go run cmd/migrate.go

# รัน worker (เช่น email worker)
go run cmd/worker.go
```

---

## 🎯 4. นำไปใช้ในกรณีไหน?

| กรณี | เหมาะสม? |
|------|----------|
| โปรเจกต์ Go ขนาดกลาง-ใหญ่ (หลาย thousand lines) | ✅ มาก |
| มีหลาย executables (API, worker, CLI, migrate) | ✅ มาก |
| มีทีมพัฒนา 2+ คน | ✅ มาก |
| ต้องการ clean architecture / testable business logic | ✅ มาก |
| ต้องการแยก public library (pkg) ออกจาก internal code | ✅ มาก |
| โปรเจกต์เล็ก (1-2 packages, < 1000 บรรทัด) | ❌ overkill |
| ต้องการ prototype หรือ hackathon | ❌ ไม่จำเป็น |

---

## ❓ 5. ทำไมต้องใช้? (ประโยชน์ที่ได้รับ)

### ✅ ประโยชน์ที่ได้รับ

1. **Separation of concerns** – แต่ละ layer มีหน้าที่ชัดเจน เปลี่ยน HTTP framework ได้โดยไม่กระทบ usecase
2. **Testability** – usecase และ repository สามารถ mock ได้ง่าย
3. **Reusability** – โค้ดใน `pkg/` สามารถใช้ข้ามโปรเจกต์ได้ (โดย import path)
4. **Maintainability** – โครงสร้างเป็นระเบียบ หาไฟล์ง่าย แก้ไขได้ตรงจุด
5. **รองรับการเติบโต** – เพิ่ม executable ใหม่ได้โดยไม่รก root
6. **Standard community practice** – developer ใหม่ที่เคยเห็น standard layout จะทำความเข้าใจได้เร็ว

### ⚠️ ข้อควรระวัง

- **อย่าวางทุกอย่างไว้ใน `internal/` แล้ว import จากนอกโปรเจกต์ไม่ได้** – ถ้าต้องการให้ package อื่นใช้ได้ ต้องย้ายไป `pkg/`
- **อย่าให้เกิด circular dependency** – โดยเฉพาะระหว่าง delivery ↔ usecase ↔ repository ควรใช้ interface
- **อย่าสร้างโฟลเดอร์ที่ไม่จำเป็น** – เช่น `internal/pkg/` ซ้อนกัน (ควรเป็น `internal/...` หรือ `pkg/` อย่างใดอย่างหนึ่ง)
- **ระวังเรื่อง import path** – ถ้าเปลี่ยนชื่อ module ใน `go.mod` จะต้องเปลี่ยน import ทุกที่

### 👍 ข้อดี

- ชุมชน Go รับรู้และมีตัวอย่างเยอะ
- ใช้กับหลายเครื่องมือ (Docker, CI/CD, codegen) ได้ง่าย
- แยก public/private code ได้ตาม Go idiom (internal, pkg)

### 👎 ข้อเสีย

- **ซับซ้อนเกินไปสำหรับโปรเจกต์เล็ก** – ต้องสร้างหลายโฟลเดอร์ boilerplate
- **ไม่มี official standard** – แต่ละองค์กรอาจปรับเปลี่ยน ทำให้新人สับสนได้
- **over-engineering** – ถ้าใช้ clean architecture 100% อาจยุ่งยากสำหรับ CRUD ธรรมดา

### 🚫 ข้อห้าม (ถ้ามี)

| ข้อห้าม | เหตุผล |
|---------|--------|
| ห้าม import โค้ดจาก `internal/` ของโปรเจกต์อื่น | Go compiler จะไม่อนุญาต (เป็น private) |
| ห้ามวาง `main.go` ไว้ที่ root ถ้ามีหลาย executables | จะสับสนว่า entry point ไหนคืออะไร |
| ห้ามใช้ `pkg/` เป็นที่ทิ้งขยะของ utils ที่ไม่เป็นระเบียบ | ควรจัดหมวดหมู่ภายใน `pkg/` เช่น `pkg/stringutil`, `pkg/timeutil` |
| ห้ามสร้าง circular dependencies ระหว่าง delivery, usecase, repository | ทำให้ compile fail และทดสอบยาก |
| ห้ามใช้ package name ซ้ำกับ standard library (เช่น `utils`, `errors`) | อาจชนะและสับสน |

---

## 📌 สรุปภาพรวมของโครงสร้างนี้

`icmongolang` เป็นโปรเจกต์ Go ที่ใช้ **Standard Go Project Layout + Clean Architecture** มีการแยก:
- **Entry points** (`cmd/`)
- **Private code** (`internal/` – delivery, usecase, repository, models, pkg เฉพาะโปรเจกต์)
- **Public code** (`pkg/` – reusable utilities)
- **Config & migrations** (`config/`, `migrations/`)
- **DevOps** (`docker-compose`, `Dockerfile`, `scripts/`)

เหมาะกับโปรเจกต์ production ขนาดกลางถึงใหญ่ที่ต้องการความยั่งยืนระยะยาว แต่ไม่เหมาะกับโปรเจกต์เล็กหรือ prototype เพราะมีความซับซ้อนเกินความจำเป็น

หากคุณต้องการปรับใช้โครงสร้างนี้กับโปรเจกต์ของตัวเอง แนะนำให้เริ่มจาก `cmd/` และ `internal/` ก่อน แล้วค่อย ๆ เพิ่ม `pkg/` เมื่อมีโค้ดที่ reuse ได้จริงครับ



```bash
# ไทย: โคลน repository จาก GitHub และเปลี่ยนไปยังไดเรกทอรีโปรเจกต์
# EN: Clone the repository from GitHub and change into the project directory
git clone github.com/kongnakornna/icmongolang.git
cd icmongolang
```
```bash

go mod tidy
go mod download
go mod verify

# ล้างฐานข้อมูลเก่า (ระวังข้อมูล)
# go run cmd/api/main.go migrate:reset (ถ้ามีคำสั่ง)

# รัน migrate ใหม่
go run main.go migrate



OR ใช้ go run โดยตรง (ไม่ต้อง build exe)

go run ./main.go serve

 ```
# Auto Run 

air

```bash


- for windows 10
 
 - .air.toml
 
    root = "."
    tmp_dir = "tmp"
    env_files = [".env.dev"]   # โหลด env โดยอัตโนมัติ

    [build]
    # ใช้ array: [binary, argument1, argument2, ...]
    entrypoint = ["./tmp/main.exe", "serve"]
    cmd = "go build -o ./tmp/main.exe ./cmd/api"
    env = ["GOOS=windows", "GOARCH=amd64"]
    clean_on_exit = true

    [log]
    time = true

    [misc]
    clean_on_exit = true

 ```

# ล้าง docs เก่า
rm -rf docs/

# สร้าง docs ใหม่
swag init
 
# รันแอปพลิเคชัน
Remove-Item -Recurse -Force docs
swag init
go mod tidy
go mod vendor
go run main.go serve


# Delete the entire vendor folder
Remove-Item -Recurse -Force vendor
go clean -modcache
# Clean up go.mod and download fresh modules
go mod tidy

# Re‑create the vendor directory
go mod vendor

# Now run the server
go run main.go serve
 
 
 
 
Remove-Item -Recurse -Force vendor
go clean -modcache
go mod tidy
go mod vendor
go run main.go serve