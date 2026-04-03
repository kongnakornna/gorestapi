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
- OR  go run cmd/api/main.go serve  on loca Windows OS
- Swagger: [localhost:5000/swagger/](http://localhost:5000/swagger/)
- http://localhost:5000/swagger/index.html#/

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


```bash
- ตรวจสอบว่า go.mod มี replace directive หรือใช้ local module หรือไม่
- ถ้า icmongolang เป็น local module ให้ใช้ replace icmongolang => icmongolang
- หรือถ้าเป็น private repo ให้ตั้ง GOPRIVATE และใช้ access token
- Perfect! You're setting up an existing Go project (icmongolang). Here's how to properly set it up and run it:
```bash
## Complete Setup Steps for Your icmongolang Project
## 📘 การจัดการ `go.mod` และ dependencies สำหรับโปรเจกต์ `icmongolang`  
## 📘 Managing `go.mod` and dependencies for `icmongolang` project

> คำแนะนำแบบทีละขั้นตอน (ไทย / อังกฤษ)  
> Step-by-step guide (Thai / English)

---

### 🧱 1. โคลนโปรเจกต์และเข้าไปในโฟลเดอร์  
### 1. Clone and enter the project

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
go run cmd/api/main.go serve

# Auto Run 
air

``` 
---
### 🧹 2. จัดการ dependencies เริ่มต้น  
### 2. Initial dependency management

```bash
# ไทย: ดาวน์โหลด dependencies และทำ tidy (เพิ่ม/ลบตามที่โค้ดเรียกใช้)
# EN: Download dependencies and tidy up (add/remove based on code imports)
go mod tidy

# ไทย: ตรวจสอบว่าโมดูลถูกตั้งค่าถูกต้อง
# EN: Verify the module is set up correctly
go mod verify
```

---

### 🗂️ 3. ตรวจสอบโครงสร้างโปรเจกต์ (ถ้าต้องการ)  
### 3. Check project structure (optional)

```bash
# ไทย: ดูว่าไฟล์สำคัญมีอยู่หรือไม่
# EN: Check if essential files exist
ls docker-compose.yml Dockerfile main.go go.mod
```

---

### 📦 4. จัดการ dependencies แบบละเอียด (download, vendor, tidy)  
### 4. Advanced dependency management (download, vendor, tidy)

```bash
# ไทย: ดาวน์โหลด dependencies ทั้งหมดลงใน module cache
# EN: Download all dependencies to the module cache
go mod download

# ไทย: ทำ tidy อีกครั้งเพื่อความแน่ใจ (ก่อน vendor)
# EN: Run tidy again to be safe (before vendoring)
go mod tidy

# ไทย: คัดลอก dependencies ไปไว้ในโฟลเดอร์ vendor
# EN: Copy dependencies into the ./vendor directory
go mod vendor

# ไทย: ตรวจสอบโมดูลอีกครั้ง
# EN: Verify modules again
go mod verify
```

---

### 🚀 5. รันแอปพลิเคชัน  
### 5. Run the application

```bash
# ไทย: รันด้วย main.go ที่ root (ถ้ามี)
# EN: Run using main.go at the root (if exists)
go run cmd/api/main.go serve

# หรือ (หรือ) ถ้า main.go อยู่ใน cmd/
# Or if main.go is inside cmd/
go run cmd/icmongolang/main.go serve
```

---

## 🔧 การรีเซ็ต `go.mod` ให้กลับมาสะอาด (แก้ปัญหา `go 1.25.0` ไม่รองรับ)  
## 🔧 Resetting `go.mod` to a clean state (fix unsupported `go 1.25.0`)

> ใช้ขั้นตอนนี้เมื่อเจอ error `go 1.25.0` หรือ `go.mod` มีปัญหาขัดแย้ง  
> Use this procedure when you see `go 1.25.0` error or `go.mod` conflicts.

---

### ✅ ขั้นตอนที่ 1 – สำรองข้อมูล (แนะนำ)  
### Step 1 – Backup (recommended)

```bash
# ไทย: สำรองไฟล์ go.mod และ go.sum ปัจจุบัน
# EN: Backup current go.mod and go.sum
cp go.mod go.mod.bak
cp go.sum go.sum.bak
```

---

### ✅ ขั้นตอนที่ 2 – กำหนดเวอร์ชัน Go ที่ถูกต้อง (stable 1.23 หรือ 1.24)  
### Step 2 – Set a valid Go version (stable 1.23 or 1.24)

```bash
# ไทย: แก้ไข go.mod ให้ใช้ go 1.23 (หรือ 1.24 แล้วแต่ที่ติดตั้ง)
# EN: Edit go.mod to use go 1.23 (or 1.24 depending on your installation)
go mod edit -go=1.23
```

---

### ✅ ขั้นตอนที่ 3 – ล้าง require blocks และสร้าง dependencies ใหม่ทั้งหมด  
### Step 3 – Remove all require blocks and rebuild dependencies cleanly

```bash
# ไทย: รัน go mod tidy เพื่อวิเคราะห์ imports ในโค้ดและสร้าง go.mod / go.sum ใหม่
# EN: Run go mod tidy to analyze code imports and regenerate go.mod / go.sum
go mod tidy
```

> `go mod tidy` จะ:
> - ลบ `require` ที่ไม่ได้ใช้
> - เพิ่ม `require` สำหรับแพ็คเกจที่ถูก import จริง
> - ปรับ `go.sum` ให้สอดคล้อง  
>  
> `go mod tidy` will:
> - Remove unused `require` directives
> - Add `require` for actually imported packages
> - Adjust `go.sum` accordingly

---

### ✅ ขั้นตอนที่ 4 – ตรวจสอบผลลัพธ์  
### Step 4 – Verify the result

```bash
# ไทย: ดูเนื้อหา go.mod ที่ถูกต้อง
# EN: View the corrected go.mod content
cat go.mod

# ไทย: ทดสอบ build หรือ run เพื่อยืนยัน
# EN: Test build or run to confirm
go build ./...
# หรือ (or)
go run cmd/api/main.go serve
```

---

### 📝 หมายเหตุเพิ่มเติม  
### 📝 Additional notes

- ถ้ายังคงมีปัญหาเรื่องเวอร์ชัน Go ให้ตรวจสอบ Go version ที่ติดตั้งด้วย `go version`  
- หากใช้ Go 1.21 ขึ้นไป การตั้ง `go 1.23` ใน `go.mod` จะยังทำงานได้ (เข้ากันได้ backward)  
- หลังจาก reset แล้ว หากต้องการใช้ vendor อีกครั้ง ให้รัน `go mod vendor`

- If you still have Go version issues, check installed Go version with `go version`  
- For Go 1.21+, setting `go 1.23` in `go.mod` is still compatible (backward compatible)  
- After resetting, if you need vendor again, run `go mod vendor`

---

✅ จบขั้นตอน – ตอนนี้ `go.mod` ของคุณสะอาดและพร้อมใช้งานแล้ว  
✅ End of steps – Your `go.mod` is now clean and ready to use.

To **reset** your `go.mod` to a clean, valid state (fixing the unsupported `go 1.25.0` and tidying dependencies), run:

```bash
# 1. (Optional) Backup current go.mod and go.sum
cp go.mod go.mod.bak
cp go.sum go.sum.bak

# 2. Set a valid Go version (1.23 or 1.24 are stable as of 2026)
go mod edit -go=1.23

# 3. Remove all require blocks and let 'go mod tidy' rebuild them cleanly
#    (this preserves only direct imports from your code)
go mod tidy

# 4. Verify the result
cat go.mod
```

After running `go mod tidy`, your `go.mod` will:
- Have a valid `go 1.23` directive
- Contain only **direct** requirements (plus `// indirect` for transitive deps)
- Have a clean `go.sum` with checksums

If you want to **completely regenerate** `go.mod` from scratch (keeping only the module name), you can:

```bash
# WARNING: deletes existing go.mod
rm go.mod go.sum
go mod init icmongolang
go mod tidy
```

But that will lose all the manually listed `require` lines. The `go mod tidy` approach above is safer – it will keep all needed deps based on your actual source code imports.

After resetting, if you still need a `vendor` directory, run:

```bash
go mod vendor
```

 
# 1. แก้ไข go.mod
go mod edit -go=1.25
go mod tidy

# 2. ลบ vendor เดิม
Remove-Item -Recurse -Force vendor -ErrorAction SilentlyContinue

# 3. สร้าง vendor (จะมี warning ให้ ignore ได้)
go mod vendor

# 4. ตรวจสอบ vendor
Get-ChildItem vendor

# 5. ทดสอบ build
go build -mod=vendor ./...


## 🧱 ความสอดคล้องกับ 3-layer Architecture + Delivery

| Layer | โฟลเดอร์ | หน้าที่ |
|-------|----------|--------|
| **Models** | `internal/models/` | กำหนด struct ของ entity (User, Product, etc.) + GORM annotations |
| **Repository** | `internal/repository/` | Interface + implementation สำหรับติดต่อฐานข้อมูล (CRUD) |
| **Usecase** | `internal/usecase/` | Business logic: validation, hashing, token creation, email sending, cache ฯลฯ |
| **Delivery** | `internal/delivery/rest/` | HTTP handlers, รับ request → เรียก usecase → ส่ง response |

## ✨ Features ที่ implement (ตามที่โจทย์ต้องการ)

| Feature | การ Implement ในโครงสร้างนี้ |
|---------|-----------------------------|
| **CRUD** | `user_handler.go` + `user_repo.go` + `auth_usecase.go` (สำหรับ user) |
| **JWT + refresh token (Redis)** | `pkg/jwt` สร้าง access/refresh token, เก็บ refresh token hash ใน Redis (key: `refresh:{userID}`) |
| **Cached user in Redis** | `usecase` จะ cache user profile หลังจาก login/get profile ด้วย `pkg/redis` (key: `user:{id}`, TTL 5 นาที) |
| **Email verification** | `pkg/email` + `hermes` สร้าง HTML email, ส่ง verification link ผ่าน `gomail` |
| **Forget / reset password** | Endpoint `/forgot-password` → สร้าง token ส่ง email → `/reset-password` รับ token + เปลี่ยนรหัส |

## 🚀 การทำงานของแต่ละ Command (Cobra)

```bash
go run cmd/api/main.go serve      # เริ่ม HTTP server บน port ที่กำหนดใน config
go run main.go migrate    # auto migrate (หรือ run raw SQL จาก migrations/)
go run main.go initdata   # สร้าง root user (email: root@gmail.com, pass: root_password)
go run main.go worker     # run background worker (ส่ง email แบบ queue)
```

## 🐳 Docker Compose ที่ใช้

- **docker-compose.dev.yml** → รัน postgres, redis, mailhog (SMTP fake) สำหรับพัฒนา
- **docker-compose.prod.yml** → รัน postgres, redis, app (build from Dockerfile.prod) พร้อม network แบบ production

## 🔥 Hot-reload (Air)

ใช้ไฟล์ `.air.toml` เพื่อให้แอปรีสตาร์ทอัตโนมัติเมื่อโค้ดเปลี่ยนแปลง
```bash
air -c .air.toml
```

## 📝 สรุป

โครงสร้างนี้ตรงตาม requirements ครบทุกข้อ:
- ✅ 3-layer clean architecture
- ✅ Chi routing + GORM ORM
- ✅ JWT (RSA) + refresh token ใน Redis
- ✅ Cached user ใน Redis
- ✅ Email verification (gomail + hermes)
- ✅ Forgot / reset password via email
- ✅ Viper config, Cobra CLI, Zap logger, Validator
- ✅ Docker + Air hot-reload + Swagger docs

สามารถ clone มาใช้ได้เลย:

```bash
git clone https://icmongolang
cd icmongolang
go mod tidy
go run cmd/api/main.go serve

OR  air
```

ขออธิบายโครงสร้างโปรเจกต์ **gorestapi** แบบละเอียด โดยอ้างอิงจากต้นแบบ [kongnakornna/gorestapi](https://gorestapi) แต่ตัดส่วนที่ไม่จำเป็นออกไป (`.vscode/`, `/docdev`, `/lession`) โครงสร้างนี้ใช้ **Clean Architecture** พร้อม分层 (layers) ที่แยกหน้าที่ชัดเจน: **Models → Repository → Usecase → Delivery** เหมาะสำหรับ REST API ที่มีความซับซ้อนปานกลางถึงสูง รองรับ JWT, Redis, Email Queue, และ Background Worker

---
## 📁 โครงสร้างโปรเจค `gorestapi`
```  
    icmongolang/
    ├── .vscode/                         # VS Code config (debug, launch)
    ├── cmd/                             # CLI commands (ใช้ Cobra)
    │   ├── /api/main.go                 # entry point (cobra.Execute())
    │   ├── root.go                      # root command
    │   ├── serve.go                     # start HTTP server
    │   ├── migrate.go                   # run DB migrations
    │   ├── initdata.go                  # seed initial data (root user)
    │   └── worker.go                    # background worker (email queue)
    │
    ├── config/                          # Viper configuration files
    │   ├── config-local.yml             # local dev config
    │   ├── config-prod.yml              # production config
    │   └── config.go                    # load & parse config
    │
    ├── docdev/                          # developer documentation
    ├── docs/                            # Swagger / API docs
    │
    internal/
    │
    ├── models/                          # Entity layer (GORM models)
    │   ├── user.go                      # User struct (id, email, password, verified, tokens...)
    │   ├── session.go                   # Session struct (ถ้าใช้ session table)
    │   └── verification.go              # VerificationToken struct (optional)
    │
    ├── repository/                      # Repository layer (data access interface)
    │   ├── user_repo.go                 # UserRepository interface & implementation
    │   │   - Create(user) error
    │   │   - FindByEmail(email) (*User, error)
    │   │   - FindByID(id) (*User, error)
    │   │   - Update(user) error
    │   │   - Delete(id) error
    │   ├── session_repo.go              # SessionRepository (สำหรับ refresh token ใน DB ถ้าไม่ใช้ Redis)
    │   └── redis_repo.go                # Redis operations (cache, refresh token store)
    │
    ├── usecase/                         # Business logic layer
    │   ├── auth_usecase.go              # AuthUsecace interface & impl
    │   │   - Register()
    │   │   - Login()
    │   │   - VerifyEmail()
    │   │   - ForgotPassword()
    │   │   - ResetPassword()
    │   │   - RefreshToken()
    │   │   - Logout()
    │   ├── user_usecase.go              # UserUsecase (CRUD)
    │   │   - GetProfile()
    │   │   - UpdateProfile()
    │   │   - ListUsers() (admin)
    │   └── cache_usecase.go             # Cache helper (ใช้ใน auth_usecase สำหรับ cached user)
    │
    ├── delivery/                        # HTTP handlers & routing
    │   ├── rest/
    │   │   ├── handler/                 # หรือจะแยกตาม entity
    │   │   │   ├── auth_handler.go      # Auth endpoints
    │   │   │   ├── user_handler.go      # User endpoints
    │   │   │   └── health_handler.go    # Health check
    │   │   ├── middleware/
    │   │   │   ├── auth.go              # JWT authentication middleware
    │   │   │   ├── logger.go            # Request logger (zap)
    │   │   │   ├── cors.go              # CORS middleware
    │   │   │   └── rate_limit.go        # Rate limiter (optional)
    │   │   ├── dto/                     # Data Transfer Objects (request/response)
    │   │   │   ├── auth_dto.go          # RegisterRequest, LoginRequest, TokenResponse...
    │   │   │   ├── user_dto.go          # UserResponse, UpdateUserRequest
    │   │   │   └── error_dto.go         # ErrorResponse
    │   │   └── router.go                # Chi router setup (routes all handlers)
    │   │
    │   └── worker/                      # Background workers (email queue)
    │       └── email_worker.go          # Process email sending async
    │
    ├── pkg/                             # Internal shared packages (ไม่ expose ออกไป)
    │   ├── jwt/
    │   │   ├── maker.go                 # JWT Maker interface (CreateToken, VerifyToken)
    │   │   ├── rsa_maker.go             # RSA implementation (private/public key)
    │   │   └── payload.go               # JWT payload struct
    │   ├── redis/
    │   │   ├── client.go                # Redis client wrapper
    │   │   ├── cache.go                 # Cache helper (Get, Set, Delete)
    │   │   └── refresh_store.go         # Store/validate refresh token in Redis
    │   ├── email/
    │   │   ├── sender.go                # Email sender interface
    │   │   ├── gomail_sender.go         # Gomail implementation
    │   │   └── templates/               # Hermes HTML templates
    │   │       ├── verification.html    # Email verification template
    │   │       └── reset_password.html  # Reset password template
    │   ├── logger/
    │   │   └── zap_logger.go            # Zap logger init & wrapper
    │   ├── validator/
    │   │   └── custom_validator.go      # go-playground/validator with custom tags
    │   ├── hash/
    │   │   └── bcrypt.go                # Password hashing
    │   └── utils/
    │       ├── random.go                # Random string generator (verification code)
    │       └── time.go                  # Time helpers
    │
    └── config/                          # (บางทีก็เอาไว้ใน internal/config)
    │    └── config.go                    # Viper config loader (struct with all settings)
    │
    ├── migrations/                      # raw SQL migration files (optional)
    ├── pkg/                             # public packages (อาจ reuse ภายนอก)
    │   └── utils/                       # helper functions
    │
    ├── scripts/                         # build, deploy, CI scripts
    ├── vendor/                          # vendored dependencies
    │
    ├── .air.toml                        # hot-reload config (Air)
    ├── .dockerignore
    ├── .env.dev                         # dev environment variables
    ├── .env.prod                        # prod environment variables
    ├── .gitignore
    ├── docker-compose.dev.yml           # dev stack (postgres, redis, mailhog)
    ├── docker-compose.prod.yml          # prod stack (including app)
    ├── Dockerfile.dev                   # multi-stage dev build
    ├── Dockerfile.prod                  # production build
    ├── go.mod / go.sum
    ├── LICENSE
    ├── README.md
    └── BookGolang.md                    # เอกสารประกอบการเรียนรู้
```

## 1. โฟลเดอร์หลักและไฟล์ระดับบนสุด

| ไฟล์/โฟลเดอร์ | คำอธิบาย |
|---------------|----------|
| `main.go` | จุดเริ่มต้นของโปรแกรม เรียกใช้ `cmd.Execute()` จาก Cobra |
| `go.mod`, `go.sum` | จัดการ dependencies (GORM, Chi, Viper, Cobra, Zap, Gomail, Redis client ฯลฯ) |
| `Dockerfile.dev`, `Dockerfile.prod` | สร้าง image สำหรับ development (hot-reload) และ production (multi-stage) |
| `docker-compose.dev.yml`, `docker-compose.prod.yml` | กำหนด service stack: PostgreSQL, Redis, MailHog (dev), และแอปพลิเคชัน |
| `.air.toml` | ตั้งค่า Air สำหรับ hot-reload ตอนพัฒนา |
| `.env.dev`, `.env.prod` | ตัวแปรสภาพแวดล้อม (database URL, JWT keys, Redis URL, SMTP) |
| `.gitignore`, `LICENSE`, `README.md`, `BookGolang.md` | เอกสารและไฟล์มาตรฐาน |

---

## 2. คำสั่ง CLI (`cmd/`)

ใช้ **Cobra** สร้างคำสั่งสำหรับบริหารแอปพลิเคชัน

| ไฟล์ | หน้าที่ |
|------|--------|
| `root.go` | คำสั่งหลัก (`icmongolang`) แสดง help และตั้งค่า flags ร่วม (เช่น `--config`) |
| `serve.go` | เริ่ม HTTP server (โหลด config, เชื่อมต่อ DB, ตั้ง router, เริ่ม worker) |
| `migrate.go` | รัน database migration (auto migrate GORM หรือ raw SQL) |
| `initdata.go` | seed ข้อมูลเริ่มต้น เช่น สร้าง root user (admin) |
| `worker.go` | เริ่ม background worker ที่คอยประมวลผล email queue (เช่น จาก Redis stream หรือ RabbitMQ) |

**ตัวอย่างการใช้งาน**  
```bash
go run cmd/api/main.go serve --config config-local.yml
go run main.go migrate
go run main.go initdata
```

---

## 3. การตั้งค่า (`config/`)

| ไฟล์ | หน้าที่ |
|------|--------|
| `config-local.yml` | ค่า config สำหรับ local dev (port, DB host, Redis, log level, JWT expiration) |
| `config-prod.yml` | สำหรับ production (อาจอ่าน secrets จาก environment variables) |
| `config.go` | ใช้ **Viper** อ่านไฟล์ YAML + env override แล้ว map ไปยัง struct (ServerConfig, DBConfig, JWTConfig, RedisConfig, EmailConfig) |

---

## 4. Internal Packages (`internal/`) – หัวใจของแอปพลิเคชัน

Go จะไม่ package อื่นนอก `internal/` มา import ดังนั้นโค้ดภายในถูกบังคับให้เป็น private

### 4.1 Models (`internal/models/`)
กำหนด struct ของ entities ที่ตรงกับตารางในฐานข้อมูล (ใช้ GORM tags)

| ไฟล์ | ฟิลด์สำคัญ |
|------|-----------|
| `user.go` | `ID`, `Email`, `PasswordHash`, `Verified` (bool), `Role` (admin/user), `CreatedAt`, `UpdatedAt` |
| `session.go` | (ถ้าใช้ session บน DB) `ID`, `UserID`, `RefreshToken`, `ExpiresAt`, `UserAgent`, `IP` |
| `verification.go` | `ID`, `UserID`, `Token` (random string), `ExpiresAt`, `Type` (email_verification, password_reset) |

### 4.2 Repository (`internal/repository/`)
เป็น **data access layer** – ติดต่อฐานข้อมูลหรือ Redis โดยใช้ interface เพื่อให้ usecase สามารถทดสอบได้ง่าย

| ไฟล์ | หน้าที่ |
|------|--------|
| `user_repo.go` | interface `UserRepository` + implementation (GORM): `Create`, `FindByEmail`, `FindByID`, `Update`, `Delete` |
| `session_repo.go` | interface `SessionRepository` สำหรับเก็บ refresh token ในฐานข้อมูล (ถ้าไม่ใช้ Redis) |
| `redis_repo.go` | interface `RedisRepository` สำหรับ Redis operation: `Set`, `Get`, `Delete`, `SetNX` (lock), `HSet`, `HGet` ใช้เก็บ refresh token (blacklist, whitelist) หรือ cache user profile |

### 4.3 Usecase (`internal/usecase/`)
**Business logic layer** – เรียก repository หลายตัว และจัดการ transaction, hashing, JWT, email sending (ผ่าน worker)

| ไฟล์ | หน้าที่ |
|------|--------|
| `auth_usecase.go` | `Register()` → hash password, save user, สร้าง verification token, ส่ง email async; <br> `Login()` → ตรวจสอบ password, สร้าง access + refresh token (เก็บ refresh ใน Redis), return tokens; <br> `VerifyEmail()`, `ForgotPassword()`, `ResetPassword()`, `RefreshToken()`, `Logout()` (revoke refresh token) |
| `user_usecase.go` | `GetProfile()`, `UpdateProfile()`, `ListUsers()` (เฉพาะ admin), `DeleteUser()` |
| `cache_usecase.go` | helper สำหรับ caching user data (ลดการ query DB) – ใช้ Redis repository |

**ตัวอย่าง workflow** ของ `Login()`:
1. ตรวจสอบ user จาก repository (`FindByEmail`)
2. เปรียบเทียบ password ด้วย bcrypt
3. ถ้าถูกต้อง → เรียก `jwtMaker.CreateToken()` สร้าง access token (สั้น, 15 นาที) และ refresh token (ยาว, 7 วัน)
4. เก็บ refresh token ใน Redis (key = `refresh:{userID}:{tokenID}`) พร้อม TTL
5. ส่ง tokens กลับไปยัง handler

### 4.4 Delivery (`internal/delivery/`)
รับผิดชอบ HTTP layer และ background workers

#### 4.4.1 REST handlers (`internal/delivery/rest/handler/`)

| ไฟล์ | endpoints |
|------|-----------|
| `auth_handler.go` | `POST /api/v1/auth/register`, `/login`, `/verify-email`, `/forgot-password`, `/reset-password`, `/refresh`, `/logout` |
| `user_handler.go` | `GET /api/v1/users/me`, `PUT /api/v1/users/me`, `GET /api/v1/users` (admin), `DELETE /api/v1/users/:id` |
| `health_handler.go` | `GET /health`, `GET /ready` (ตรวจ DB, Redis) |

แต่ละ handler จะ bind request DTO → เรียก usecase → แปลง response DTO → ส่ง JSON

#### 4.4.2 Middleware (`internal/delivery/rest/middleware/`)

| ไฟล์ | หน้าที่ |
|------|--------|
| `auth.go` | ตรวจ JWT จาก `Authorization: Bearer <token>` → แปลง payload → ใส่ userID, role ลง context (ใช้ `context.WithValue`) → ถ้า token หมดอายุ หรือถูก revoke (check Redis blacklist) → reject |
| `logger.go` | บันทึก request method, path, status, latency, user agent (ใช้ Zap) |
| `cors.go` | ตั้ง CORS headers (allow origin, methods, credentials) |
| `rate_limit.go` | จำกัดจำนวน request ต่อ IP หรือ user (ใช้ Redis + sliding window หรือ token bucket) |

#### 4.4.3 DTO (`internal/delivery/rest/dto/`)

| ไฟล์ | struct ตัวอย่าง |
|------|----------------|
| `auth_dto.go` | `RegisterRequest{Email, Password, ConfirmPassword}`, `LoginRequest`, `TokenResponse{AccessToken, RefreshToken}` |
| `user_dto.go` | `UserResponse{ID, Email, Verified, Role}`, `UpdateUserRequest` |
| `error_dto.go` | `ErrorResponse{Code, Message, Details}` |

#### 4.4.4 Router (`internal/delivery/rest/router.go`)
- ใช้ **Chi** router (เบา, รองรับ middleware chain)
- จัดกลุ่ม route: `/api/v1/auth`, `/api/v1/users`, `/health`
- ต่อ middleware ทั่วไป: logger, cors, rate limit
- สำหรับ route ที่ต้องการ authentication → ใส่ `middleware.Auth`

#### 4.4.5 Worker (`internal/delivery/worker/email_worker.go`)
- Background worker (อาจใช้ goroutine + channel หรือใช้ Redis Stream / BullMQ pattern)
- รับ job ส่ง email (verification, reset password)
- ใช้ email sender interface (gomail) พร้อม template HTML (Hermes)

---

## 5. Shared Packages (`internal/pkg/` และ `pkg/`)

### 5.1 `internal/pkg/jwt/`
- `maker.go` – interface `Maker` (CreateToken, VerifyToken)
- `rsa_maker.go` – ใช้ private/public key (RSA256) sign JWT
- `payload.go` – struct `Payload` (UserID, Role, ExpiresAt, IssuedAt, TokenID)

### 5.2 `internal/pkg/redis/`
- `client.go` – สร้าง Redis client (go-redis)
- `cache.go` – helper สำหรับ get/set JSON, increment, expire
- `refresh_store.go` – ฟังก์ชันเฉพาะ `StoreRefreshToken(userID, tokenID string, ttl time.Duration)`, `ValidateRefreshToken(userID, tokenID string) bool`, `RevokeRefreshToken(...)`

### 5.3 `internal/pkg/email/`
- `sender.go` – interface `EmailSender`
- `gomail_sender.go` – implement โดยใช้ gopkg.in/gomail.v2
- `templates/` – HTML template สำหรับ verification และ reset password (ใช้ `html/template`)

### 5.4 `internal/pkg/logger/`
- `zap_logger.go` – สร้าง logger (production หรือ development) ด้วย Uber Zap รองรับ structured logging, log level, และ sink (stdout, file)

### 5.5 `internal/pkg/validator/`
- `custom_validator.go` – ใช้ go-playground/validator พร้อม custom validation เช่น `password` (ต้องมีตัวเลข, ตัวพิมพ์ใหญ่, พิเศษ) และ `phone` (เบอร์ไทย)

### 5.6 `internal/pkg/hash/`
- `bcrypt.go` – `HashPassword(password string) (string, error)` และ `CheckPasswordHash(password, hash string) bool`

### 5.7 `internal/pkg/utils/`
- `random.go` – สร้าง random string (สำหรับ verification token หรือ reset code)
- `time.go` – helper timezone, duration parser, nowUTC()

### 5.8 `pkg/utils/` (public)
- ถ้ามี helper ที่อาจถูกใช้โดย external package เช่น format response, pagination เป็นต้น

---

## 6. Migration (`migrations/`)
- เก็บ raw SQL migration files (optional) ถ้าใช้ GORM auto migrate อาจไม่ต้องใช้ แต่ถ้าต้องการ version control schema ควรใช้ `golang-migrate` หรือ `goose`
- โครงสร้างไฟล์: `000001_create_users_table.up.sql`, `000001_create_users_table.down.sql`

---

## 7. Scripts (`scripts/`)
- เก็บ shell script สำหรับ build, deploy, CI/CD เช่น `build.sh`, `deploy-staging.sh`, `seed-test-data.sh`

---

## 8. Vendor (`vendor/`)
- เมื่อใช้ `go mod vendor` จะสร้างโฟลเดอร์นี้สำหรับ dependencies (ใช้ในการ build แบบ offline หรือ CI ที่ไม่มี internet)

---

## 9. เอกสาร API (`docs/`)
- ใช้ **Swaggo** (swag init) สร้าง Swagger JSON/YAML และ UI (สามารถเปิด `/swagger/index.html`)
- คำสั่ง: `swag init -g cmd/root.go` (ต้อง annotate handler ด้วย `// @Summary` ฯลฯ)

---

## ภาพรวมการทำงาน (Request Flow)

```mermaid
sequenceDiagram
    Client->>+Router: POST /api/v1/auth/login
    Router->>+Middleware: Logger, CORS, RateLimit
    Middleware->>+AuthHandler: LoginRequest DTO
    AuthHandler->>+AuthUsecase: Login(email, password)
    AuthUsecase->>+UserRepo: FindByEmail(email)
    UserRepo-->>-AuthUsecase: user
    AuthUsecase->>+hash: CheckPasswordHash
    AuthUsecase->>+jwt: CreateToken (access + refresh)
    AuthUsecase->>+RedisRepo: StoreRefreshToken
    AuthUsecase-->>-AuthHandler: TokenResponse
    AuthHandler-->>-Client: 200 OK + tokens
```

---

## สรุปคุณสมบัติเด่นของโครงสร้างนี้

- **Clean Architecture**: แยก model, repository, usecase, delivery ทำให้ทดสอบง่ายและเปลี่ยน database หรือ delivery (REST → gRPC) ได้โดยไม่กระทบ business logic
- **Security**: JWT RSA256, refresh token เก็บใน Redis, bcrypt สำหรับ password, rate limiting, CORS
- **Asynchronous Email**: ใช้ background worker ป้องกัน request ติดขัด
- **Production-ready**: มี Docker, multi-stage build, health check, structured logging (Zap), graceful shutdown
- **Developer Experience**: Hot-reload (Air), CLI commands (Cobra), Swagger docs, environment-based config

โครงสร้างนี้เหมาะกับ REST API ขนาดกลางถึงใหญ่ สามารถขยับขยายไปใช้ event-driven หรือ microservices ได้ง่าย โดยการเพิ่ม message broker แทน channel worker.


# 🧱 โครงสร้างโปรเจกต์ `icmongolang` – Go Backend สำหรับระบบ Authentication & User Management

โครงสร้างนี้เป็นเทมเพลตสำหรับพัฒนา REST API ด้วยภาษา Go 遵循 Clean Architecture (ใช้ layer: Model → Repository → Usecase → Delivery) รองรับ JWT (RSA), Redis cache, email queue, migration, และ hot-reload

---

## 1. โครงสร้างการทำงานของแต่ละโฟลเดอร์/ไฟล์

```
icmongolang/
├── .vscode/                 # debug & launch configuration สำหรับ VS Code
├── cmd/                     # CLI commands (ใช้ Cobra)
│   ├── api/main.go          # entry point – เรียก cobra.Execute()
│   ├── root.go              # root command (แสดง help)
│   ├── serve.go             # start HTTP server
│   ├── migrate.go           # run DB migrations (GORM AutoMigrate)
│   ├── initdata.go          # seed initial data (admin user, roles)
│   └── worker.go            # run background worker (email queue consumer)
│
├── config/                  # Viper configuration files
│   ├── config-local.yml     # local development config
│   ├── config-prod.yml      # production config
│   └── config.go            # load & parse config (struct with env override)
│
├── docdev/                  # developer documentation (architecture, API design)
├── docs/                    # Swagger (OpenAPI) generated docs
│
├── internal/                # private code – ไม่ถูก import จากภายนอก
│   ├── models/              # GORM entities (business entities)
│   │   ├── user.go          # User struct (id, email, password_hash, verified, etc.)
│   │   ├── session.go       # Session struct (ถ้าใช้ session ใน DB)
│   │   └── verification.go  # VerificationToken (email verify, reset password)
│   │
│   ├── repository/          # data access layer (interface + implementation)
│   │   ├── user_repo.go     # UserRepository (Create, FindByEmail, Update, Delete)
│   │   ├── session_repo.go  # SessionRepository (สำหรับ refresh token ใน DB)
│   │   └── redis_repo.go    # Redis operations (cache, blacklist, refresh store)
│   │
│   ├── usecase/             # business logic layer
│   │   ├── auth_usecase.go  # Register, Login, VerifyEmail, ForgotPassword, RefreshToken, Logout
│   │   ├── user_usecase.go  # GetProfile, UpdateProfile, ListUsers (admin)
│   │   └── cache_usecase.go # helper สำหรับ caching user data
│   │
│   ├── delivery/            # HTTP handlers & routing
│   │   ├── rest/
│   │   │   ├── handler/
│   │   │   │   ├── auth_handler.go   # auth endpoints (register, login, verify, refresh)
│   │   │   │   ├── user_handler.go   # user profile endpoints
│   │   │   │   └── health_handler.go # /health, /ready
│   │   │   ├── middleware/
│   │   │   │   ├── auth.go           # JWT authentication middleware
│   │   │   │   ├── logger.go         # request logger (zap)
│   │   │   │   ├── cors.go           # CORS middleware
│   │   │   │   └── rate_limit.go     # rate limiter (token bucket)
│   │   │   ├── dto/                  # request/response DTOs
│   │   │   │   ├── auth_dto.go       # RegisterRequest, LoginRequest, TokenResponse
│   │   │   │   ├── user_dto.go       # UserResponse, UpdateUserRequest
│   │   │   │   └── error_dto.go      # ErrorResponse
│   │   │   └── router.go             # Chi router + route registration
│   │   └── worker/
│   │       └── email_worker.go       # worker ที่ consume queue (Redis หรือ channel) และส่งอีเมล
│   │
│   ├── pkg/                 # internal shared packages (ไม่ expose ภายนอก)
│   │   ├── jwt/
│   │   │   ├── maker.go     # interface (CreateToken, VerifyToken)
│   │   │   ├── rsa_maker.go # RSA implementation (อ่าน private/public key)
│   │   │   └── payload.go   # JWT claims struct
│   │   ├── redis/
│   │   │   ├── client.go    # Redis client wrapper
│   │   │   ├── cache.go     # generic Get/Set/Delete
│   │   │   └── refresh_store.go # store & validate refresh token
│   │   ├── email/
│   │   │   ├── sender.go    # EmailSender interface
│   │   │   ├── gomail_sender.go # ส่งผ่าน SMTP
│   │   │   └── templates/   # HTML templates (Hermes)
│   │   ├── logger/
│   │   │   └── zap_logger.go # zap init & wrapper
│   │   ├── validator/
│   │   │   └── custom_validator.go # go-playground/validator + custom tags
│   │   ├── hash/
│   │   │   └── bcrypt.go    # HashPassword, CheckPasswordHash
│   │   └── utils/
│   │       ├── random.go    # สร้าง random string (verification code)
│   │       └── time.go      # time helper (UTC, expiry)
│   │
│   └── config/              # (บางครั้ง config loader ก็อยู่ internal/config)
│        └── config.go
│
├── migrations/              # raw SQL migration files (ถ้าไม่ใช้ GORM AutoMigrate)
├── pkg/                     # public packages (อาจใช้ซ้ำข้ามโปรเจกต์ได้) – แต่ปกติ internal ไว้ก่อน
├── scripts/                 # build, deploy, CI scripts
├── vendor/                  # vendored dependencies
│
├── .air.toml                # hot-reload config (Air)
├── docker-compose.dev.yml   # dev stack: postgres, redis, mailhog
├── docker-compose.prod.yml  # prod stack: รวม app + postgres + redis
├── Dockerfile.dev / Dockerfile.prod
├── go.mod, go.sum
└── README.md, BookGolang.md
```

---

## 2. Workflow & Dataflow Diagram (Flowchart TB – เหมาะกับ Draw.io)

ด้านล่างเป็น **Mermaid flowchart** 
**การลงทะเบียน (Register)** และ **การล็อกอิน (Login)** พร้อม refresh token

```mermaid
flowchart TB
    subgraph Client
        A[User] --> B[ส่ง HTTP Request]
    end

    subgraph "API Server (Go)"
        C[Chi Router + Middleware]
        D[auth_handler.Register]
        E[auth_usecase.Register]
        F[user_repo.Create]
        G[email_worker.SendVerification]
        H[auth_handler.Login]
        I[auth_usecase.Login]
        J[user_repo.FindByEmail]
        K[hash.CheckPassword]
        L[jwt_maker.CreateToken]
        M[redis_store.SaveRefreshToken]
        N[auth_handler.RefreshToken]
        O[auth_usecase.RefreshToken]
        P[redis_store.ValidateRefreshToken]
        Q[jwt_maker.CreateToken]
    end

    subgraph "External Services"
        R[(PostgreSQL)]
        S[(Redis)]
        T[SMTP Server / Mailhog]
    end

    B --> C
    C -- POST /api/v1/register --> D
    D --> E
    E --> F
    F --> R
    E -- success --> G
    G --> T
    G -- enqueue job --> S

    C -- POST /api/v1/login --> H
    H --> I
    I --> J
    J --> R
    I --> K
    K -- valid --> L
    L -- access_token --> H
    I --> M
    M -- refresh_token --> S
    H --> B

    C -- POST /api/v1/refresh --> N
    N --> O
    O --> P
    P -- valid --> Q
    Q -- new access_token --> N
    N --> B
```

### คำอธิบายแบบละเอียดของ Workflow

#### 2.1 กระบวนการ Register
1. Client ส่ง `POST /api/v1/register` พร้อม JSON `{email, password, name}`
2. `auth_handler.Register` รับ request → validate DTO → เรียก `auth_usecase.Register`
3. Usecase ตรวจสอบว่าอีเมลซ้ำหรือไม่ (เรียก `user_repo.FindByEmail`)
4. ถ้าไม่ซ้ำ → hash password (bcrypt) → สร้าง user entity → `user_repo.Create` ลง PostgreSQL
5. หลังสร้างสำเร็จ → สร้าง verification token (random string) → ส่ง task ไปยัง **email queue** (Redis list หรือ channel) → worker จะดึง task และส่งอีเมล verify ผ่าน SMTP
6. Response กลับ `201 Created` พร้อม message “กรุณายืนยันอีเมล”

#### 2.2 กระบวนการ Login
1. Client ส่ง `POST /api/v1/login` พร้อม `{email, password}`
2. Handler → Usecase → `user_repo.FindByEmail`
3. ถ้าเจอ user และ verified = true → ตรวจสอบ password ด้วย bcrypt
4. สำเร็จ → สร้าง **access token** (RSA, หมดอายุ 15 นาที) และ **refresh token** (UUID หรือ JWT, หมดอายุ 7 วัน)
5. เก็บ refresh token ใน Redis (key: `refresh:{user_id}:{token_id}`) พร้อม TTL
6. Response `200 OK` พร้อม `{access_token, refresh_token, expires_in}`

#### 2.3 กระบวนการ Refresh Token
1. Client ส่ง `POST /api/v1/refresh` พร้อม `refresh_token` ใน body หรือ header
2. Usecase ดึง refresh token จาก Redis → ถ้ามีและไม่ถูก blacklist → สร้าง access token ใหม่
3. Response access token ใหม่ (refresh token อาจถูก renew หรือคงเดิม ตาม design)

#### 2.4 กระบวนการ Logout
1. Client ส่ง `POST /api/v1/logout` (ต้องมี access token)
2. Middleware ตรวจสอบ JWT → ส่ง user ID ไปยัง Usecase
3. Usecase ลบ refresh token ออกจาก Redis และ optionally ใส่ access token ลง blacklist จนกว่าจะหมดอายุ

---

## 3. ตัวอย่างโค้ดพร้อมคอมเมนต์ (ไทย/อังกฤษ) และนำไปรันได้

เราจะแสดงเฉพาะส่วนสำคัญที่สุด: `auth_usecase.go`, `auth_handler.go`, `jwt/rsa_maker.go` และ `docker-compose.dev.yml`

### 3.1 `internal/usecase/auth_usecase.go`

```go
// Package usecase contains business logic for authentication.
// แพ็คเกจ usecase มีตรรกะทางธุรกิจสำหรับการรับรองตัวตน
package usecase

import (
    "errors"
    "time"
    "yourproject/internal/models"
    "yourproject/internal/repository"
    "yourproject/internal/pkg/hash"
    "yourproject/internal/pkg/jwt"
    "yourproject/internal/pkg/random"
    "yourproject/internal/pkg/redis"
    "golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
    Register(email, password, name string) (*models.User, error)
    Login(email, password string) (accessToken, refreshToken string, err error)
    RefreshToken(refreshToken string) (newAccessToken string, err error)
    Logout(userID uint, accessToken string) error
}

type authUsecase struct {
    userRepo    repository.UserRepository
    redisCache  *redis.Client      // ใช้สำหรับเก็บ refresh token และ blacklist
    jwtMaker    jwt.Maker
    emailQueue  chan<- models.EmailTask // channel สำหรับส่งงาน async
}

// NewAuthUsecase สร้าง instance พร้อม dependency injection
func NewAuthUsecase(
    ur repository.UserRepository,
    rc *redis.Client,
    jm jwt.Maker,
    eq chan<- models.EmailTask,
) AuthUsecase {
    return &authUsecase{
        userRepo:   ur,
        redisCache: rc,
        jwtMaker:   jm,
        emailQueue: eq,
    }
}

// Register สร้าง user ใหม่, hash password, และส่งอีเมลยืนยัน
// Register creates a new user, hashes password, and sends verification email (async)
func (a *authUsecase) Register(email, password, name string) (*models.User, error) {
    // ตรวจสอบว่าอีเมลถูกใช้แล้วหรือยัง
    existing, _ := a.userRepo.FindByEmail(email)
    if existing != nil {
        return nil, errors.New("email already exists")
    }

    hashedPwd, err := hash.HashPassword(password) // bcrypt
    if err != nil {
        return nil, err
    }

    user := &models.User{
        Email:        email,
        PasswordHash: hashedPwd,
        Name:         name,
        Verified:     false,
        CreatedAt:    time.Now().UTC(),
    }

    // บันทึกลงฐานข้อมูล
    if err := a.userRepo.Create(user); err != nil {
        return nil, err
    }

    // สร้าง verification token (random 32 ตัวอักษร) และเก็บใน Redis หรือ DB
    verifyToken := random.String(32)
    // เก็บ token ใน Redis กับ user ID, หมดอายุ 24 ชม.
    key := "verify:" + verifyToken
    a.redisCache.Set(key, user.ID, 24*time.Hour)

    // ส่ง task ไปยัง email queue (worker จะส่งจริง)
    a.emailQueue <- models.EmailTask{
        To:      email,
        Subject: "Verify your email",
        Template: "verification",
        Data:    map[string]interface{}{"Token": verifyToken, "UserID": user.ID},
    }

    return user, nil
}

// Login ตรวจสอบ credentials และสร้าง access/refresh tokens
func (a *authUsecase) Login(email, password string) (string, string, error) {
    user, err := a.userRepo.FindByEmail(email)
    if err != nil {
        return "", "", errors.New("invalid credentials")
    }
    if !user.Verified {
        return "", "", errors.New("email not verified")
    }

    // เปรียบเทียบรหัสผ่าน
    if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
        return "", "", errors.New("invalid credentials")
    }

    // สร้าง access token (RSA, อายุ 15 นาที)
    accessToken, err := a.jwtMaker.CreateToken(user.ID, user.Email, 15*time.Minute)
    if err != nil {
        return "", "", err
    }

    // สร้าง refresh token (ใช้ random string หรือ JWT อีกตัว)
    refreshToken := random.String(64)
    // เก็บ refresh token ใน Redis: key = refresh:{userID}:{tokenID}
    refreshKey := "refresh:" + string(rune(user.ID)) + ":" + refreshToken
    a.redisCache.Set(refreshKey, user.ID, 7*24*time.Hour)

    return accessToken, refreshToken, nil
}

// RefreshToken ตรวจสอบ refresh token ใน Redis แล้วสร้าง access token ใหม่
func (a *authUsecase) RefreshToken(refreshToken string) (string, error) {
    // ใช้ pattern scan? จริง ๆ ควรส่ง userID มาใน payload ของ refresh token ด้วย
    // เพื่อความง่าย สมมติว่าเราดึง userID ออกจาก Redis โดยตรง (ต้องมี key mapping)
    userID, err := a.redisCache.Get(refreshToken) // สมมติว่าเก็บ value เป็น userID
    if err != nil {
        return "", errors.New("invalid refresh token")
    }
    user, _ := a.userRepo.FindByID(userID.(uint))
    if user == nil {
        return "", errors.New("user not found")
    }
    newAccessToken, err := a.jwtMaker.CreateToken(user.ID, user.Email, 15*time.Minute)
    if err != nil {
        return "", err
    }
    return newAccessToken, nil
}
```

### 3.2 `internal/delivery/rest/handler/auth_handler.go`

```go
// Package handler จัดการ HTTP request/response สำหรับ auth endpoints
package handler

import (
    "net/http"
    "yourproject/internal/delivery/rest/dto"
    "yourproject/internal/usecase"
    "github.com/go-chi/chi/v5"
    "github.com/go-playground/validator/v10"
)

type AuthHandler struct {
    authUsecase usecase.AuthUsecase
    validate    *validator.Validate
}

func NewAuthHandler(au usecase.AuthUsecase) *AuthHandler {
    return &AuthHandler{authUsecase: au, validate: validator.New()}
}

// Register handles POST /register
// @Summary Register a new user
// @Accept json
// @Param request body dto.RegisterRequest true "Registration info"
// @Success 201 {object} dto.UserResponse
// @Failure 400 {object} dto.ErrorResponse
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
    var req dto.RegisterRequest
    if err := dto.DecodeJSON(r, &req); err != nil {
        dto.RespondError(w, http.StatusBadRequest, "invalid request body")
        return
    }
    if err := h.validate.Struct(req); err != nil {
        dto.RespondError(w, http.StatusBadRequest, err.Error())
        return
    }

    user, err := h.authUsecase.Register(req.Email, req.Password, req.Name)
    if err != nil {
        dto.RespondError(w, http.StatusConflict, err.Error())
        return
    }

    dto.RespondJSON(w, http.StatusCreated, dto.ToUserResponse(user))
}

// Login handles POST /login
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
    var req dto.LoginRequest
    if err := dto.DecodeJSON(r, &req); err != nil {
        dto.RespondError(w, http.StatusBadRequest, "invalid request")
        return
    }
    accessToken, refreshToken, err := h.authUsecase.Login(req.Email, req.Password)
    if err != nil {
        dto.RespondError(w, http.StatusUnauthorized, err.Error())
        return
    }
    dto.RespondJSON(w, http.StatusOK, dto.TokenResponse{
        AccessToken:  accessToken,
        RefreshToken: refreshToken,
        TokenType:    "Bearer",
        ExpiresIn:    900, // 15 minutes
    })
}
```

### 3.3 `internal/pkg/jwt/rsa_maker.go` (RSA JWT)

```go
// Package jwt สร้างและตรวจสอบ JSON Web Tokens ด้วย RSA private/public key
package jwt

import (
    "crypto/rsa"
    "errors"
    "time"
    "github.com/golang-jwt/jwt/v5"
)

type RSAMaker struct {
    privateKey *rsa.PrivateKey
    publicKey  *rsa.PublicKey
}

func NewRSAMaker(privateKeyPEM, publicKeyPEM []byte) (*RSAMaker, error) {
    privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyPEM)
    if err != nil {
        return nil, err
    }
    publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyPEM)
    if err != nil {
        return nil, err
    }
    return &RSAMaker{privateKey: privateKey, publicKey: publicKey}, nil
}

// CreateToken สร้าง JWT signed ด้วย RSA private key
func (m *RSAMaker) CreateToken(userID uint, email string, duration time.Duration) (string, error) {
    claims := jwt.MapClaims{
        "sub": userID,
        "email": email,
        "exp": time.Now().Add(duration).Unix(),
        "iat": time.Now().Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
    return token.SignedString(m.privateKey)
}

// VerifyToken ตรวจสอบ JWT และคืน claims
func (m *RSAMaker) VerifyToken(tokenString string) (*Payload, error) {
    token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
        if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
            return nil, errors.New("unexpected signing method")
        }
        return m.publicKey, nil
    })
    if err != nil {
        return nil, err
    }
    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok || !token.Valid {
        return nil, errors.New("invalid token")
    }
    // แปลง claims เป็น Payload struct
    return &Payload{
        UserID: uint(claims["sub"].(float64)),
        Email:  claims["email"].(string),
    }, nil
}
```

### 3.4 `docker-compose.dev.yml` – รันสแต็คสำหรับพัฒนา

```yaml
version: '3.8'
services:
  postgres:
    image: postgres:15-alpine
    environment:
      POSTGRES_USER: appuser
      POSTGRES_PASSWORD: secret
      POSTGRES_DB: icmongo
    ports:
      - "5432:5432"
    volumes:
      - pgdata:/var/lib/postgresql/data

  redis:
    image: redis:7-alpine
    ports:
      - "6379:6379"

  mailhog:
    image: mailhog/mailhog
    ports:
      - "1025:1025"   # SMTP
      - "8025:8025"   # Web UI

volumes:
  pgdata:
```

### 3.5 วิธีรันจริง

```bash
# 1. สร้าง private/public key สำหรับ JWT
openssl genrsa -out private.pem 2048
openssl rsa -in private.pem -pubout -out public.pem

# 2. ตั้ง environment variables (หรือใช้ .env)
export DB_URL="postgres://appuser:secret@localhost:5432/icmongo?sslmode=disable"
export REDIS_ADDR="localhost:6379"
export JWT_PRIVATE_KEY_PATH="private.pem"
export JWT_PUBLIC_KEY_PATH="public.pem"

# 3. รัน docker compose สำหรับ database + redis + mailhog
docker-compose -f docker-compose.dev.yml up -d

# 4. รัน migration (สมมติว่าใช้ GORM AutoMigrate)
go run cmd/api/main.go migrate

# 5. รัน server (hot-reload ด้วย air)
air -c .air.toml
```

---

## 4. ยกตัวอย่างการใช้งานจริง / กรณีศึกษา

**กรณีศึกษา: Startup e-commerce ต้องการระบบ member ที่ปลอดภัย**

- ใช้โครงสร้างนี้เป็นฐาน → เพิ่ม role-based access control (RBAC) และ session management
- เนื่องจากมี traffic สูง จึงใช้ Redis cache สำหรับ user profile และ rate limiter ป้องกัน brute force
- ใช้ background worker ส่งอีเมลยืนยันและ forgot password แบบไม่阻塞 response
- ใช้ RSA JWT ทำให้ฝ่าย mobile และ frontend ตรวจสอบ token ได้โดยไม่ต้องเรียก API (ถ้ามี public key)
- มี `initdata.go` สำหรับสร้าง admin user รายแรก และกำหนดสิทธิ์

**ผลลัพธ์**:
- ลดเวลา develop authentication จาก 2 สัปดาห์เหลือ 3 วัน
- สามารถ scale out ได้เพราะ stateless JWT + Redis session store
- มี logging และ monitoring พร้อม (ผ่าน zap + sentry)

---

## 5. สรุป

### ✅ ประโยชน์ที่ได้รับ
- **Clean Architecture** แยกหน้าที่ชัดเจน (repository, usecase, delivery) ทำให้ทดสอบง่าย (mock)
- **JWT with RSA** ปลอดภัยกว่า HMAC (ไม่มี shared secret)
- **Background email queue** ช่วยให้ API response เร็ว ไม่ติด I/O
- **Hot-reload (Air)** เพิ่ม productivity ใน dev
- **ใช้ Docker compose** สำหรับ dev stack ทำให้ทีมงาน setup เร็ว

### ⚠️ ข้อควรระวัง
- **RSA key management**: ต้อง保护好 private key (ใช้ secret manager หรือ vault)
- **Refresh token rotation**: ต้อง implement blacklist และป้องกัน replay attack
- **Redis เป็น SPOF**: สำหรับ production ควรใช้ Redis Sentinel หรือ Cluster
- **GORM AutoMigrate** อาจไม่ safe สำหรับ production ควรใช้ migration tool (golang-migrate)
- **Email queue** ต้องมี fallback (retry + dead-letter) และ monitoring

### 👍 ข้อดี
- ใช้ Go ที่ performance สูง, concurrent ดี
- โครงสร้างเป็นมาตรฐานที่ developer หลายคนเข้าใจ
- รองรับการทำ unit test และ integration test ได้ง่าย (เพราะใช้ interface)
- สามารถเปลี่ยน database หรือ cache ได้โดยไม่กระทบ business logic

### 👎 ข้อเสีย
- มี boilerplate code ค่อนข้างมาก (repository, usecase, dto)
- มือใหม่ที่ยังไม่เข้าใจ clean architecture อาจงง กับการแมป layer
- การใช้ channel เป็น email queue เหมาะกับ monolith แต่ถ้าต้อง distributed ควรใช้ message broker จริง (RabbitMQ, Kafka)

### 🚫 ข้อห้าม (ถ้ามี)
- **ห้ามเก็บ plain-text password** (ต้อง bcrypt หรือ argon2)
- **ห้ามใช้ JWT โดยไม่ตั้ง expiry** (ควรสั้น 15-30 นาที)
- **ห้ามส่ง access token ผ่าน URL parameter** (ใช้ Authorization header)
- **ห้าม commit private key ลง Git** (เพิ่มใน .gitignore เสมอ)
- **ห้ามใช้ GORM AutoMigrate ใน production โดยไม่ตรวจสอบ schema change**

### 📚 แหล่งอ้างอิง / ที่มา
- [Clean Architecture in Go](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [JWT with RSA – RFC 7519](https://tools.ietf.org/html/rfc7519)
- [GORM](https://gorm.io), [Chi router](https://github.com/go-chi/chi)
- [Uber Go Style Guide](https://github.com/uber-go/guide)
- [Air – live reload for Go](https://github.com/cosmtrek/air)
- โครงสร้างได้รับแรงบันดาลใจจาก `go-clean-arch` และ project template ของบริษัท Fintech ในไทย

---

> **หมายเหตุ**: ตัวอย่างโค้ดข้างต้นเป็นเพียงส่วนประกอบหลัก หากต้องการนำไปรันจริง
>  ต้องสร้างไฟล์ที่เหลือ (DTO, repository implementation, main.go, router.go, worker) ตามโครงสร้างที่อธิบายไว้ สามารถดูเทมเพลต
>  ได้ที่ [github.com/icmongolang/boilerplate](https://github.com/icmongolang/boilerplate) (ลิงก์สมมติ)

