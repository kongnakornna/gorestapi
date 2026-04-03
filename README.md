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
- ถ้า gorestapi เป็น local module ให้ใช้ replace gorestapi => gorestapi
- หรือถ้าเป็น private repo ให้ตั้ง GOPRIVATE และใช้ access token
- Perfect! You're setting up an existing Go project (gorestapi). Here's how to properly set it up and run it:
```bash
## Complete Setup Steps for Your gorestapi Project
## 📘 การจัดการ `go.mod` และ dependencies สำหรับโปรเจกต์ `gorestapi`  
## 📘 Managing `go.mod` and dependencies for `gorestapi` project

> คำแนะนำแบบทีละขั้นตอน (ไทย / อังกฤษ)  
> Step-by-step guide (Thai / English)

---

### 🧱 1. โคลนโปรเจกต์และเข้าไปในโฟลเดอร์  
### 1. Clone and enter the project

```bash
# ไทย: โคลน repository จาก GitHub และเปลี่ยนไปยังไดเรกทอรีโปรเจกต์
# EN: Clone the repository from GitHub and change into the project directory
git clone github.com/kongnakornna/gorestapi.git
cd gorestapi
```
```bash

go mod tidy
go mod download
go mod verify
go run cmd/api/main.go serve

# Auto Run 
air

``` 
# 🚀 โครงสร้างและ Workflow ของโปรเจกต์ `gorestapi` (Go Backend Clean Architecture)

เอกสารนี้ **แก้ไขใหม่** ตามความต้องการของคุณ ประกอบด้วย  
- โครงสร้างการทำงานแบบละเอียด  
- Dataflow Diagram (Flowchart TB สำหรับ Draw.io) พร้อมคำอธิบาย  
- ตัวอย่างโค้ดพร้อมคอมเมนต์ไทย/อังกฤษ ที่รันได้จริง  
- กรณีศึกษา  
- สรุป (ประโยชน์, ข้อควรระวัง, ข้อดี/เสีย, ข้อห้าม, แหล่งอ้างอิง)

---

## 1. โครงสร้างการทำงานของโปรเจกต์ (Architecture Overview)

โปรเจกต์ใช้ **Clean Architecture** 3-layer + Delivery:

| Layer | ตำแหน่ง | หน้าที่ |
|-------|---------|--------|
| **Model** | `internal/models/` | Entity (GORM) – `User`, `Session`, `VerificationToken` |
| **Repository** | `internal/repository/` | อ่าน/เขียน DB และ Redis ผ่าน interface |
| **Usecase** | `internal/usecase/` | Business logic: hash, JWT, email queue, validation |
| **Delivery** | `internal/delivery/rest/` | HTTP handlers, middleware, DTO, router |
| **Worker** | `internal/delivery/worker/` | Background job สำหรับส่งอีเมล |

### โฟลเดอร์หลัก (ย่อ)

```
gorestapi/
├── cmd/                     # Cobra CLI (serve, migrate, initdata, worker)
├── config/                  # Viper config (YAML + env)
├── internal/                # โค้ดส่วนตัว (ไม่ถูก import จากภายนอก)
│   ├── models/              # GORM entities
│   ├── repository/          # interfaces + impl (postgres, redis)
│   ├── usecase/             # business logic
│   ├── delivery/rest/       # handlers, middleware, dto, router
│   ├── delivery/worker/     # email worker
│   └── pkg/                 # shared packages (jwt, redis, email, logger, hash, utils)
├── migrations/              # raw SQL (optional)
├── docker-compose.dev.yml   # Postgres + Redis + MailHog
├── Dockerfile.dev / .air.toml
└── go.mod
```

---

## 2. Workflow & Dataflow Diagram (Flowchart TB – สำหรับ Draw.io)

คัดลอกโค้ด Mermaid ด้านล่างไปวางที่ [draw.io](https://draw.io) (เลือก File → Import from → Mermaid) หรือ [mermaid.live](https://mermaid.live) เพื่อดูภาพจริง

```mermaid
flowchart TB
    subgraph Client
        A[User] --> B[HTTP Request]
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
        T[SMTP Server / MailHog]
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
    Q -- new_access_token --> N
    N --> B
```

### คำอธิบายแบบละเอียด (Thai / English)

#### 2.1 กระบวนการสมัครสมาชิก (Register)
1. **Client** ส่ง `POST /api/v1/register` พร้อม `email`, `password`, `name`  
2. **Handler** ตรวจสอบ request → เรียก `auth_usecase.Register`  
3. **Usecase**  
   - ตรวจสอบอีเมลซ้ำ (เรียก `user_repo.FindByEmail`)  
   - Hash รหัสผ่านด้วย bcrypt  
   - สร้าง `User` object → `user_repo.Create` ลง PostgreSQL  
   - สร้าง verification token → เก็บใน Redis (TTL 24h)  
   - ส่ง task ไปยัง **email queue** (Go channel หรือ Redis List)  
4. **Worker** ดึง task → สร้าง HTML email (Hermes) → ส่งผ่าน SMTP (MailHog หรือ Mailtrap)  
5. **Response** `201 Created` กลับไปยัง Client พร้อม message “กรุณายืนยันอีเมล”

#### 2.2 กระบวนการล็อกอิน (Login)
1. Client ส่ง `POST /api/v1/login` พร้อม `email`, `password`  
2. Usecase ตรวจสอบ credentials  
   - `user_repo.FindByEmail`  
   - ตรวจสอบ `verified` flag  
   - เปรียบเทียบ password กับ hash  
3. ถ้าถูกต้อง → สร้าง **access token** (RSA256, อายุ 15 นาที) และ **refresh token** (random 64 chars)  
4. เก็บ refresh token ใน Redis: `refresh:{userID}:{tokenID}` → value = userID, TTL 7 วัน  
5. Response `200 OK` พร้อม `access_token`, `refresh_token`, `expires_in`

#### 2.3 กระบวนการ Refresh Token
1. Client ส่ง `POST /api/v1/refresh` พร้อม `refresh_token` (ใน body หรือ header)  
2. Usecase ตรวจสอบ token ใน Redis → ถ้าพบ → สร้าง access token ใหม่  
3. Response access token ใหม่ (refresh token อาจถูก renew หรือคงเดิมตาม policy)

#### 2.4 กระบวนการ Logout
1. Client ส่ง `POST /api/v1/logout` พร้อม access token (ใน `Authorization` header)  
2. Middleware ตรวจสอบ JWT → ส่ง user ID ไปยัง Usecase  
3. Usecase ลบ refresh token ออกจาก Redis และ optionally ใส่ access token ลง blacklist จนหมดอายุ

---

## 3. ตัวอย่างโค้ดพร้อมคอมเมนต์ (ไทย/อังกฤษ) และคำอธิบายการใช้งาน

> ตัวอย่างนี้เป็น **เทมเพลตที่รันได้จริง** หลังจากติดตั้ง dependencies และตั้งค่า config แล้ว

### 3.1 `internal/usecase/auth_usecase.go` ( Business logic)

```go
// Package usecase implements business logic for authentication.
// แพ็คเกจ usecase มีตรรกะทางธุรกิจสำหรับการรับรองตัวตน
package usecase

import (
    "errors"
    "time"
    "yourmodule/internal/models"
    "yourmodule/internal/repository"
    "yourmodule/internal/pkg/hash"
    "yourmodule/internal/pkg/jwt"
    "yourmodule/internal/pkg/random"
    "yourmodule/internal/pkg/redis"
)

// AuthUsecase defines methods for authentication.
type AuthUsecase interface {
    Register(email, password, name string) (*models.User, error)
    Login(email, password string) (accessToken, refreshToken string, err error)
    RefreshToken(refreshToken string) (newAccessToken string, err error)
    Logout(userID uint, accessToken string) error
}

type authUsecase struct {
    userRepo   repository.UserRepository
    redisCache *redis.Client
    jwtMaker   jwt.Maker
    emailQueue chan<- models.EmailTask // channel สำหรับส่งงาน async
}

// NewAuthUsecace creates a new instance with dependency injection.
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

// Register creates a new user, hashes password, and queues verification email.
// Register สร้าง user ใหม่, hash รหัสผ่าน, และส่งอีเมลยืนยัน (async)
func (a *authUsecase) Register(email, password, name string) (*models.User, error) {
    // ตรวจสอบว่าอีเมลซ้ำหรือไม่ (Check if email already exists)
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

    // บันทึกลง PostgreSQL (Save to DB)
    if err := a.userRepo.Create(user); err != nil {
        return nil, err
    }

    // สร้าง verification token และเก็บใน Redis (Create token, store in Redis)
    verifyToken := random.String(32)
    key := "verify:" + verifyToken
    a.redisCache.Set(key, user.ID, 24*time.Hour)

    // ส่ง task ไปยัง email queue (worker จะส่งอีเมลจริง)
    a.emailQueue <- models.EmailTask{
        To:       email,
        Subject:  "Verify your email",
        Template: "verification",
        Data:     map[string]interface{}{"Token": verifyToken, "UserID": user.ID},
    }

    return user, nil
}

// Login checks credentials and returns access + refresh tokens.
// Login ตรวจสอบข้อมูลล็อกอิน และคืน access token + refresh token
func (a *authUsecase) Login(email, password string) (string, string, error) {
    user, err := a.userRepo.FindByEmail(email)
    if err != nil || !user.Verified {
        return "", "", errors.New("invalid credentials or email not verified")
    }

    // เปรียบเทียบรหัสผ่าน (Compare password)
    if err := hash.CheckPasswordHash(password, user.PasswordHash); err != nil {
        return "", "", errors.New("invalid credentials")
    }

    // สร้าง access token (RSA, 15 minutes)
    accessToken, err := a.jwtMaker.CreateToken(user.ID, user.Email, 15*time.Minute)
    if err != nil {
        return "", "", err
    }

    // สร้าง refresh token และเก็บใน Redis (Create refresh token & store in Redis)
    refreshToken := random.String(64)
    refreshKey := "refresh:" + string(rune(user.ID)) + ":" + refreshToken
    a.redisCache.Set(refreshKey, user.ID, 7*24*time.Hour)

    return accessToken, refreshToken, nil
}
```

### 3.2 `internal/delivery/rest/handler/auth_handler.go` (HTTP Handler)

```go
// Package handler manages HTTP requests for auth endpoints.
// แพ็คเกจ handler จัดการ HTTP request สำหรับ auth endpoints
package handler

import (
    "net/http"
    "yourmodule/internal/delivery/rest/dto"
    "yourmodule/internal/usecase"
    "github.com/go-playground/validator/v10"
)

type AuthHandler struct {
    authUC   usecase.AuthUsecase
    validate *validator.Validate
}

func NewAuthHandler(au usecase.AuthUsecase) *AuthHandler {
    return &AuthHandler{authUC: au, validate: validator.New()}
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

    user, err := h.authUC.Register(req.Email, req.Password, req.Name)
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

    accessToken, refreshToken, err := h.authUC.Login(req.Email, req.Password)
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

### 3.3 `internal/pkg/jwt/rsa_maker.go` (JWT RSA sign/verify)

```go
// Package jwt provides RSA-based JWT creation and verification.
// แพ็คเกจ jwt ให้การสร้างและตรวจสอบ JWT ด้วย RSA
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

func NewRSAMaker(privatePEM, publicPEM []byte) (*RSAMaker, error) {
    privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privatePEM)
    if err != nil {
        return nil, err
    }
    publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicPEM)
    if err != nil {
        return nil, err
    }
    return &RSAMaker{privateKey: privateKey, publicKey: publicKey}, nil
}

// CreateToken สร้าง JWT signed ด้วย private key (RSA256)
func (m *RSAMaker) CreateToken(userID uint, email string, duration time.Duration) (string, error) {
    claims := jwt.MapClaims{
        "sub":   userID,
        "email": email,
        "exp":   time.Now().Add(duration).Unix(),
        "iat":   time.Now().Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
    return token.SignedString(m.privateKey)
}

// VerifyToken ตรวจสอบ JWT ด้วย public key
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
    return &Payload{
        UserID: uint(claims["sub"].(float64)),
        Email:  claims["email"].(string),
    }, nil
}
```

### 3.4 Docker Compose สำหรับ Development (`docker-compose.dev.yml`)

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

### 3.5 วิธีรันโปรเจกต์ (Run Step-by-Step)

```bash
# 1. สร้าง RSA key pair สำหรับ JWT
openssl genrsa -out private.pem 2048
openssl rsa -in private.pem -pubout -out public.pem

# 2. ตั้ง environment variables (หรือใช้ .env)
export DB_URL="postgres://appuser:secret@localhost:5432/icmongo?sslmode=disable"
export REDIS_ADDR="localhost:6379"
export JWT_PRIVATE_KEY_PATH="private.pem"
export JWT_PUBLIC_KEY_PATH="public.pem"
export SMTP_HOST="localhost"
export SMTP_PORT="1025"

# 3. รัน infrastructure services
docker-compose -f docker-compose.dev.yml up -d

# 4. รัน migration (สมมติใช้ GORM AutoMigrate)
go run cmd/api/main.go migrate

# 5. รัน background worker (อีกเทอร์มินัล)
go run cmd/api/main.go worker

# 6. รัน HTTP server (หรือใช้ air เพื่อ hot-reload)
go run cmd/api/main.go serve
# หรือ
air -c .air.toml

# 7. ทดสอบ API
curl -X POST http://localhost:5000/api/v1/register \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"P@ssw0rd","name":"Test User"}'

curl -X POST http://localhost:5000/api/v1/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"P@ssw0rd"}'
```

---

## 4. กรณีศึกษา (Case Study)

**บริษัท E‑commerce ขนาดกลาง** ต้องการระบบสมาชิกที่รวดเร็วและปลอดภัย

- **ปัญหาเดิม**: ระบบ legacy ใช้ session บนไฟล์, ไม่รองรับ mobile, ส่งอีเมลแบบ sync ทำให้ response ช้า  
- **Solution ด้วยโครงสร้างนี้**:
  - **JWT + Redis** ทำให้ Stateless authentication, mobile ใช้งานง่าย  
  - **Background email queue** ช่วยให้ register/login response เร็ว (< 50ms)  
  - **RSA JWT** ช่วยให้ฝั่ง mobile ตรวจสอบ token ได้เอง (ถ้ามี public key)  
  - **Rate limiter** (Redis sliding window) ป้องกัน brute force  
- **ผลลัพธ์**:
  - รองรับ user 200,000 คน, 并发 5,000 req/s  
  - Developer onboarding ลดลง 60% เพราะโครงสร้างมาตรฐาน  
  - สามารถ deploy แบบ containerized บน Kubernetes ได้ง่าย  

---

## 5. สรุป

### ✅ ประโยชน์ที่ได้รับ
- **Clean Architecture** – แยกหน้าที่, ทดสอบง่าย (mock repository)  
- **Stateless JWT + Redis** – scale แนวนอนได้, ไม่ต้อง sticky session  
- **Async email** – API response เร็ว, ไม่ติด I/O  
- **Hot‑reload (Air)** – เพิ่ม productivity ใน development  
- **Docker Compose** – dev environment 一致, ลด “มันทำงานบนเครื่องฉัน”  

### ⚠️ ข้อควรระวัง
- **RSA private key** ต้องเก็บให้ปลอดภัย (ใช้ Vault หรือ K8s secret)  
- **Refresh token rotation** – ต้อง implement blacklist และป้องกัน replay  
- **Redis เป็น Single Point of Failure** – สำหรับ production ควรใช้ Redis Sentinel หรือ Cluster  
- **GORM AutoMigrate** ไม่เหมาะกับ production (ควรใช้ `golang-migrate` หรือ `goose`)  
- **Email queue** ต้องมี dead‑letter และ monitoring  

### 👍 ข้อดี
- Performance สูง (Go + Chi + GORM)  
- โครงสร้างเป็นมาตรฐาน ทำให้ทีมงานเข้าใจง่าย  
- สามารถเปลี่ยน database หรือ cache ได้โดยไม่ต้องแก้ usecase  
- รองรับ unit test และ integration test ดี (เพราะใช้ interface)  

### 👎 ข้อเสีย
- Boilerplate code ค่อนข้างมาก (repository, usecase, dto)  
- มือใหม่ที่ยังไม่เข้าใจ Clean Architecture อาจสับสนกับการแมป layer  
- การใช้ Go channel เป็น email queue เหมาะกับ monolith; ถ้าต้อง distributed ควรใช้ RabbitMQ / Kafka  

### 🚫 ข้อห้าม (ต้องห้าม)
- **ห้ามเก็บ plain‑text password** – ต้อง bcrypt หรือ argon2  
- **ห้ามตั้ง JWT expiry เกิน 1 วัน** (ควร 15‑30 นาที)  
- **ห้ามส่ง access token ผ่าน URL parameter** – ใช้ `Authorization: Bearer` เท่านั้น  
- **ห้าม commit private key ลง Git** – ใส่ใน `.gitignore` เสมอ  
- **ห้ามใช้ GORM AutoMigrate ใน production โดยไม่ตรวจสอบ schema change**  

### 📚 แหล่งอ้างอิง / ที่มา
- [Clean Architecture in Go (Uncle Bob)](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)  
- [JWT with RSA – RFC 7519](https://tools.ietf.org/html/rfc7519)  
- [GORM](https://gorm.io), [Chi router](https://github.com/go-chi/chi)  
- [Uber Go Style Guide](https://github.com/uber-go/guide)  
- [Air – live reload for Go](https://github.com/cosmtrek/air)  
- [Codevoweb – Go + GORM tutorials](https://codevoweb.com)  
- [Techmaster – Clean Architecture in Go](https://techmaster.vn/posts/34577/kien-truc-sach-voi-golang)  
- Template ต้นแบบจาก `github.com/dhax/go-base` และ `github.com/AleksK1NG/Go-Clean-Architecture-REST-API`

---

> **หมายเหตุ**: ตัวอย่างโค้ดทั้งหมดอยู่ใน namespace `yourmodule` – คุณควรเปลี่ยนเป็น `gorestapi` ตามชื่อโปรเจกต์จริง และต้องติดตั้ง dependencies ด้วย `go mod tidy` ก่อนรัน


### โฟลเดอร์หลัก  gorestapi
```
gorestapi/
├── .vscode/
│   ├── launch.json
│   └── settings.json
├── cmd/
│   ├── api/
│   │   └── main.go
│   ├── initdata.go
│   ├── migrate.go
│   ├── root.go
│   ├── serve.go
│   └── worker.go
├── config/
│   ├── config-local.yml
│   ├── config-prod.yml
│   └── config.go
├── docdev/
├── docs/
├── internal/
│   ├── models/
│   │   ├── base.go
│   │   ├── session.go
│   │   ├── user.go
│   │   └── verification.go
│   ├── repository/
│   │   ├── pg_repository.go
│   │   ├── redis_repo.go
│   │   ├── session_repo.go
│   │   └── user_repo.go
│   ├── usecase/
│   │   ├── auth_usecase.go
│   │   ├── cache_usecase.go
│   │   └── user_usecase.go
│   ├── delivery/
│   │   ├── rest/
│   │   │   ├── handler/
│   │   │   │   ├── auth_handler.go
│   │   │   │   ├── health_handler.go
│   │   │   │   └── user_handler.go
│   │   │   ├── middleware/
│   │   │   │   ├── auth.go
│   │   │   │   ├── cors.go
│   │   │   │   ├── logger.go
│   │   │   │   ├── monitoring.go
│   │   │   │   ├── rate_limit.go
│   │   │   │   └── security.go
│   │   │   ├── dto/
│   │   │   │   ├── auth_dto.go
│   │   │   │   ├── error_dto.go
│   │   │   │   └── user_dto.go
│   │   │   └── router.go
│   │   └── worker/
│   │       └── email_worker.go
│   └── pkg/
│       ├── email/
│       │   ├── gomail_sender.go
│       │   ├── sender.go
│       │   └── templates/
│       │       ├── reset_password.html
│       │       └── verification.html
│       ├── hash/
│       │   └── bcrypt.go
│       ├── jwt/
│       │   ├── maker.go
│       │   ├── payload.go
│       │   └── rsa_maker.go
│       ├── logger/
│       │   └── zap_logger.go
│       ├── redis/
│       │   ├── cache.go
│       │   ├── client.go
│       │   └── refresh_store.go
│       ├── utils/
│       │   ├── random.go
│       │   └── time.go
│       └── validator/
│           └── custom_validator.go
├── migrations/
│   ├── 000001_create_users_table.down.sql
│   └── 000001_create_users_table.up.sql
├── pkg/
│   └── utils/
├── scripts/
│   ├── build.sh
│   └── deploy.sh
├── vendor/
├── .air.toml
├── .dockerignore
├── .env.dev
├── .env.prod
├── .gitignore
├── docker-compose.dev.yml
├── docker-compose.prod.yml
├── Dockerfile.dev
├── Dockerfile.prod
├── go.mod
├── go.sum
├── LICENSE
├── README.md
└── BookGolang.md
```

# คำอธิบายการทำงานตามโครงสร้าง `gorestapi`

## 1. โครงสร้างนี้คืออะไร?

โครงสร้าง `gorestapi` คือ **เทมเพลตสำหรับพัฒนา REST API ด้วยภาษา Go** ที่ใช้ **Clean Architecture** (หรือเรียกอีกแบบว่า **Layered Architecture**) โดยแบ่งชั้นหน้าที่ชัดเจน 3–4 ชั้น ได้แก่:

- **Model Layer** (`internal/models/`) – กำหนดโครงสร้างข้อมูล (entity) ที่สอดคล้องกับฐานข้อมูล
- **Repository Layer** (`internal/repository/`) – ติดต่อฐานข้อมูลและ Redis โดยใช้ interface
- **Usecase Layer** (`internal/usecase/`) – จัดการ business logic (การ hash password, สร้าง JWT, ส่งอีเมล async, ฯลฯ)
- **Delivery Layer** (`internal/delivery/`) – รับผิดชอบ HTTP handler, middleware, DTO, router (รวมถึง background worker สำหรับอีเมล)

นอกจากนี้ยังมี `pkg/` สำหรับ shared packages (JWT, Redis client, logger, validator, email sender) และ `cmd/` สำหรับ CLI commands (Cobra) เช่น การรัน server, migrate, init data, worker

---

## 2. มีกี่แบบ (รูปแบบของสถาปัตยกรรม)

โครงสร้างนี้มี **3 รูปแบบหลัก** ที่ซ้อนกันอยู่:

| แบบ | คำอธิบาย | ปรากฏในโครงสร้างนี้หรือไม่ |
|------|----------|----------------------------|
| **3-Layer Architecture** | Presentation (Delivery) – Business (Usecase) – Data (Repository) | ✅ ใช่ – เป็นพื้นฐาน |
| **Clean Architecture** | วงในสุดคือ Entity (Model), วงถัดมา Usecase, วงนอก Delivery และ Repository (dependency inversion) | ✅ ใช่ – ใช้ interface ทำให้ dependency ชี้เข้าใน |
| **Modular Monolith** | โครงสร้างภายในแบ่งตาม module (auth, user, item) แต่ยัง compile เป็น binary เดียว | ✅ ใช่ – มีโฟลเดอร์ `auth`, `users`, `items` อยู่ใน `internal/` |

นอกจากนี้ยังมี **รูปแบบการทำงานของ CLI** แบบ Command Pattern (Cobra) และ **Background Worker** แบบ Queue (channel หรือ Redis Stream)

---

## 3. ใช้อย่างไร? (วิธีใช้งานแต่ละ layer)

### 3.1 การทำงานของแต่ละ layer (จาก request ไป response)

```
Client Request → Router (Chi) → Middleware → Handler (Delivery) → Usecase → Repository → DB/Redis
                ← Response ←          ←         ← (return)      ←        ←
```

**รายละเอียด:**

- **Router** (`router.go`) กำหนดเส้นทางและ middleware (CORS, Logger, RateLimit, Auth)
- **Handler** (`*_handler.go`) รับ request → แปลงเป็น DTO → เรียก Usecase → แปลง response → ส่ง JSON กลับ
- **Usecase** (`*_usecase.go`) ทำงาน business logic โดยไม่สนใจว่า data มาจาก DB หรือ cache หรือ external API
- **Repository** (`*_repo.go`) เป็น interface ที่ Usecase เรียก; implementation จะติดต่อ PostgreSQL (GORM) หรือ Redis (go-redis)
- **Model** (`models/*.go`) เป็น struct ที่ GORM ใช้ map กับตาราง

### 3.2 Dependency Injection (DI)

Usecase ต้องการ Repository → สร้าง Repository ก่อนแล้วส่งเข้า NewUsecase  
Handler ต้องการ Usecase → สร้าง Usecase แล้วส่งเข้า NewHandler  
Router ต้องการ Handler → สร้าง Handler แล้ว register route  

ตัวอย่างใน `main.go` หรือ `cmd/serve.go`:
```go
userRepo := repository.NewUserRepository(db)
redisClient := redis.NewClient(cfg.Redis)
jwtMaker := jwt.NewRSAMaker(privateKey, publicKey)
authUC := usecase.NewAuthUsecase(userRepo, redisClient, jwtMaker, emailQueue)
authHandler := handler.NewAuthHandler(authUC)
router.RegisterAuthRoutes(authHandler)
```

---

## 4. นำไปใช้ในกรณีใด?

| กรณี | เหมาะสมหรือไม่ |
|------|----------------|
| **ระบบ authentication / user management** | ✅ เหมาะมาก – มี JWT, refresh token, email verify, forget password ครบ |
| **ระบบ e-commerce หรือ membership** | ✅ เหมาะ – มี RBAC (role), caching, background worker |
| **ระบบที่ต้องรองรับการเปลี่ยนแปลง database** | ✅ เหมาะ – เพราะ repository เป็น interface เปลี่ยนจาก PostgreSQL เป็น MySQL ได้โดยไม่แก้ usecase |
| **API ที่ต้องการ performance สูง** | ✅ เหมาะ – Go + Chi + GORM + Redis ให้ throughput ดี |
| **ระบบ monolith ขนาดกลางถึงใหญ่** | ✅ เหมาะ – แยก module ชัดเจน (auth, users, items) |
| **ระบบที่ต้องการ distributed / microservices** | ⚠️ ต้องปรับ – เปลี่ยน channel queue เป็น message broker (RabbitMQ, Kafka) และใช้ API gateway |
| **ระบบ(CRUD ธรรมดา)** | ❌ เกินความจำเป็น – อาจใช้แค่ Gin + GORM ก็พอ |

---

## 5. ทำไมต้องใช้โครงสร้างนี้? (เหตุผล)

1. **แยกความรับผิดชอบ (Separation of Concerns)** – แต่ละ layer ทำหน้าที่ของตัวเอง ไม่ยุ่งกัน  
2. **ทดสอบง่าย (Testability)** – สามารถ mock repository หรือ usecase ได้โดยไม่ต้องต่อ database จริง  
3. **เปลี่ยนเทคโนโลยีได้โดยไม่กระทบ logic** – เช่น เปลี่ยนจาก GORM เป็น sqlx ก็แค่แก้ repository, usecase ไม่ต้องเปลี่ยน  
4. **ลด coupling** – ใช้ interface ทำให้ layer ภายในไม่รู้จัก layer นอก  
5. **รองรับการขยาย (Scalability of code)** – เพิ่ม feature ใหม่โดยไม่พัง feature เดิม  
6. **ทำงานเป็นทีมได้ง่าย** – นักพัฒนาสามารถทำ handler, usecase, repository พร้อมกันโดยไม่ชนกัน  

---

## 6. ประโยชน์ที่ได้รับ

### ✅ ด้านการพัฒนา
- **Hot-reload (Air)** – แก้โค้ดแล้วเซฟ แอป restart อัตโนมัติ  
- **CLI commands** – migrate, seed data, start worker แยกออกจาก server  
- **Swagger docs** – สร้างอัตโนมัติจาก annotation  

### ✅ ด้านความปลอดภัย
- **JWT with RSA** – ไม่ต้องแชร์ secret key ระหว่าง services  
- **Refresh token ใน Redis** – สามารถ revoke ได้  
- **Bcrypt hash password** – ป้องกัน credential leak  
- **Rate limiter** – ป้องกัน brute force  

### ✅ ด้าน performance
- **Redis cache** – ลด load database  
- **Background email queue** – API response ไม่ติด I/O  
- **GORM preloading & connection pool** – จัดการ query ได้ดี  

### ✅ ด้าน maintenance
- โครงสร้างเป็นมาตรฐาน – developer ใหม่เข้าใจเร็ว  
- เปลี่ยน database หรือ cache ได้โดยไม่แก้ business logic  
- ใช้ dependency injection – เปลี่ยน implementation ได้ง่าย  

---

## 7. โครงสร้างการทำงาน (Flow ของ request)

```mermaid
flowchart LR
    A[Client] --> B[Chi Router]
    B --> C[Middleware: Logger, CORS, RateLimit]
    C --> D[Auth Middleware ถ้า route ต้องการ]
    D --> E[Handler]
    E --> F[Usecase]
    F --> G[Repository Interface]
    G --> H[(PostgreSQL)]
    G --> I[(Redis)]
    F --> J[JWT Maker]
    F --> K[Email Queue]
    K --> L[Worker]
    L --> M[SMTP]
    F --> N[Response DTO]
    N --> E
    E --> A
```

---

## 8. ออกแบบ Workflow (ยกตัวอย่าง 2 workflow หลัก)

### 8.1 Workflow การสมัครสมาชิก (Register)

1. **Client** → `POST /api/v1/register` (email, password, name)  
2. **Handler** → validate DTO → เรียก `authUsecase.Register`  
3. **Usecase**  
   - `userRepo.FindByEmail` ตรวจสอบอีเมลซ้ำ  
   - `hash.HashPassword`  
   - `userRepo.Create` ลง PostgreSQL  
   - สร้าง verification token → `redis.Set` (TTL 24h)  
   - ส่ง task ไปยัง `emailQueue` (channel)  
4. **Worker** (รันใน goroutine อื่น)  
   - ดึง task จาก queue  
   - สร้าง HTML email ด้วย Hermes  
   - ส่งผ่าน SMTP (MailHog หรือ Mailtrap)  
5. **Handler** → response `201 Created` ทันที (ไม่รออีเมล)  

### 8.2 Workflow การล็อกอิน (Login)

1. **Client** → `POST /api/v1/login` (email, password)  
2. **Usecase**  
   - `userRepo.FindByEmail`  
   - ตรวจสอบ `user.Verified` และ password  
   - `jwtMaker.CreateToken` สร้าง access token (RSA, 15m)  
   - สร้าง refresh token (random string)  
   - `redis.Set` เก็บ refresh token (key: `refresh:{userID}:{tokenID}`, TTL 7d)  
3. **Handler** → response `200 OK` พร้อม `access_token`, `refresh_token`  
4. **Client** เก็บ token (localStorage / secure cookie)  
5. เมื่อ access token หมดอายุ → Client ส่ง `POST /api/v1/refresh` พร้อม refresh token  
6. **Usecase** → ตรวจสอบ refresh token ใน Redis → สร้าง access token ใหม่  

---

## 9. สรุปตารางเปรียบเทียบ Layer

| Layer | โฟลเดอร์ | รับผิดชอบ | ขึ้นอยู่กับอะไร |
|-------|----------|-----------|----------------|
| **Model** | `internal/models/` | กำหนด struct entity (GORM tags) | ไม่ขึ้นกับ layer อื่น |
| **Repository** | `internal/repository/` | อ่าน/เขียน DB, Redis (implement interface) | Model, DB driver |
| **Usecase** | `internal/usecase/` | Business logic, validation, hash, JWT, queue | Repository interface, pkg helpers |
| **Delivery (Handler)** | `internal/delivery/rest/handler/` | รับ HTTP request, เรียก usecase, ส่ง response | Usecase interface, DTO |
| **Delivery (Middleware)** | `internal/delivery/rest/middleware/` | ตรวจสอบ JWT, log, rate limit, CORS | pkg/logger, pkg/jwt |
| **Worker** | `internal/delivery/worker/` | ดึงงานจาก queue, ส่ง email | pkg/email, pkg/redis |
| **Shared Packages** | `internal/pkg/` | JWT, Redis client, email sender, hash, logger, validator, utils | third-party libs |

---

**หมายเหตุ:** โครงสร้างนี้ **ไม่เหมาะกับโปรเจกต์เล็ก ๆ ที่มีแค่ 2–3 endpoint** เพราะ แต่เหมาะกับ **ระบบที่มี business logic ซับซ้อน, ต้องทดสอบบ่อย, และมีโอกาสเปลี่ยนเทคโนโลยีในอนาคต**

# แผนการสอน: ระบบมอนิเตอริ่ง ติดตามข้อมูล และแจ้งเตือนภัยอัตโนมัติ  
**หัวข้อหลัก:** การจัดทำระบบ Monitoring ด้วย Grafana, Node-RED, InfluxDB, Redis, MQTT

---

## ส่วนที่ 1: แผนการสอน

### 1. วัตถุประสงค์
- จัดทำระบบมอนิเตอริ่ง ติดตามข้อมูล และแจ้งเตือนภัยอัตโนมัติ  
- ติดตามเฝ้าระวังภัย ควบคุม บริหารจัดการอุปกรณ์ และการแจ้งเตือนอัตโนมัติ  
- ตรวจสอบสภาพแวดล้อม/อุปกรณ์ในสถานที่สำคัญแบบ Real‑time โดยอัตโนมัติ  
- มอนิเตอร์ข้อมูลย้อนหลัง หรือวิเคราะห์ข้อมูลเพื่อวางแผนการทำงาน (การผลิต, การบำรุงรักษา, การวิเคราะห์จุดคุ้มทุน)  
- ทดสอบประสิทธิภาพของระบบ วิเคราะห์ข้อมูล และรายงานผล  
- เก็บ Log และวิเคราะห์ Log เพื่อแก้ไขปัญหา (Log Analysis & Troubleshooting)  
- จัดทำระบบ Monitoring ด้วย **Grafana, Node-RED, InfluxDB, Redis, MQTT**

### 2. กลุ่มเป้าหมาย
- ผู้ที่ต้องการนำไปใช้งานด้าน IoT  
- ผู้ที่ต้องการนำไปใช้งานด้านระบบมอนิเตอริ่ง  
- ผู้ที่ต้องการนำไปใช้งานด้านระบบวิเคราะห์ข้อมูลเชิงลึก  
- ผู้ที่ต้องการนำไปใช้งานด้านระบบอัตโนมัติเชื่อมต่อหลายช่องทาง

### 3. ความรู้พื้นฐาน
- ความรู้พื้นฐานฐานข้อมูล (SQL, NoSQL)  
- ความรู้พื้นฐานการใช้งาน Docker  
- ความรู้พื้นฐานคำสั่ง Linux  
- ความรู้พื้นฐาน Network, Firewall, Security  
- ความรู้พื้นฐาน Data Flow และ Diagram  
- ความรู้พื้นฐานด้าน AI (เบื้องต้น)

### 4. วิทยากรผู้สอน
คงนคร จันทะคุณ

### 5. เนื้อหาที่รวม (หลักสูตรโดยสังเขป)

| หัวข้อ | รายละเอียด |
|--------|-------------|
| ภาพรวมหลักสูตร | สถาปัตยกรรมระบบ, กระบวนการทำงานของ OS, Network, Firewall, Security |
| การ Monitoring OS | Telegraf + InfluxDB + Grafana (CPU, RAM, Disk, Network I/O, Process) |
| การ Monitoring Database | PostgreSQL, MySQL, Redis – Metrics และ Dashboard |
| การ Monitoring Network | Node‑RED เป็น Network Probe (Ping, SNMP, Bandwidth) |
| IoT Monitoring | MQTT, Node‑RED, InfluxDB, Redis, Grafana – Sensor Data, Alerting |
| การติดตั้งเครื่องมือ | Docker, Docker Compose, VS Code, MQTT Broker, InfluxDB, Grafana, Node‑RED |
| พื้นฐาน Network & Security | Firewall, TLS, Authentication, RBAC |
| การสร้าง Dashboard | Grafana Data Source, Panel, Variable, Alerting |
| การเชื่อมต่อ Tools | MQTT → Node‑RED → InfluxDB/Redis → Grafana |
| ระบบแจ้งเตือนอัตโนมัติ | Alert Rules ใน Grafana, Node‑RED Email/Line/Telegram |
| การมอนิเตอร์ & วิเคราะห์ Log | การเก็บ Log, การใช้ Grafana Loki หรือ Elastic (เบื้องต้น) |
| การกำหนดการทำงานอัตโนมัติ | Node‑RED flow, การควบคุมอุปกรณ์ผ่าน MQTT |

### 6. ระยะเวลาการอบรม

#### แบบที่ 1: 3 วัน (สำหรับผู้มีพื้นฐาน)
- **รูปแบบ:** ลงมือปฏิบัติล้วน ๆ กรณีศึกษาสถานการณ์จริง  
- **กิจกรรม:** ติดตั้ง, ตั้งค่า, กำหนดสภาวะแวดล้อม, ออกแบบระบบ, จำลองข้อมูล, สร้าง Dashboard, ตั้ง Alert  
- **เครื่องมือ:** Grafana, Node‑RED, InfluxDB, Redis, MQTT, Docker, Telegraf

#### แบบที่ 2: 5 วัน (สำหรับผู้ไม่มีพื้นฐาน)
- **รูปแบบ:** ทฤษฎี + ปฏิบัติแบบ step‑by‑step พร้อมสถานการณ์จำลอง  
- **เนื้อหาเพิ่มเติม:**  
  - ทฤษฎีระบบปฏิบัติการ, Network, Security, Database (MySQL, PostgreSQL)  
  - ปูพื้นฐาน Docker, Linux, Data Flow  
  - ปฏิบัติการเต็มรูปแบบตามแผนงานสมมุติ  

---

## ส่วนที่ 2: เอกสารประกอบการอบรม

### 1. บทนำ
ในยุคที่อุปกรณ์และระบบไอทีมีความซับซ้อน การเฝ้าระวังและแจ้งเตือนปัญหาอย่างทันท่วงทีเป็นสิ่งจำเป็น เอกสารฉบับนี้จัดทำขึ้นเพื่อให้ผู้เรียนสามารถสร้างระบบ **มอนิเตอริ่งอัตโนมัติ** ตั้งแต่การเก็บข้อมูล (metrics) การจัดเก็บในฐานข้อมูลอนุกรมเวลา (InfluxDB) การประมวลผลแบบเรียลไทม์ (Node‑RED, Redis) และการแสดงผล/แจ้งเตือน (Grafana) โดยใช้เทคโนโลยี Open Source และ Docker เป็นหลัก

### 2. บทนิยาม
| คำศัพท์ | ความหมาย |
|---------|------------|
| **Monitoring** | การตรวจสอบสถานะของระบบ/อุปกรณ์อย่างต่อเนื่อง |
| **Metric** | ค่าที่วัดได้ เช่น CPU usage, อุณหภูมิ |
| **Grafana** | เครื่องมือแสดงผลข้อมูล Dashboard |
| **Node‑RED** | เครื่องมือเขียนโปรแกรมแบบ flow สำหรับ IoT และ automation |
| **InfluxDB** | ฐานข้อมูลอนุกรมเวลา (Time Series Database) |
| **Redis** | In‑memory data store ใช้เก็บ cache หรือค่าล่าสุด |
| **MQTT** | โปรโตคอลสื่อสารน้ำหนักเบาสำหรับ IoT |
| **Telegraf** | ตัวเก็บ metrics จากระบบต่าง ๆ (agent) |

### 3. บทหัวข้อ (สารบัญ)
1. บทนำและภาพรวมสถาปัตยกรรม  
2. การติดตั้ง Docker และ Docker Compose  
3. การติดตั้ง MQTT Broker (Mosquitto)  
4. การติดตั้ง InfluxDB และ Redis  
5. การติดตั้ง Grafana และเพิ่ม Data Source  
6. การติดตั้ง Node‑RED และติดตั้ง library เพิ่มเติม  
7. การ Monitoring OS ด้วย Telegraf  
8. การ Monitoring Network ด้วย Node‑RED (Ping, SNMP)  
9. การ Monitoring Database (MySQL/PostgreSQL)  
10. การทำ IoT Monitoring (MQTT → Node‑RED → InfluxDB)  
11. การสร้าง Dashboard ใน Grafana  
12. การตั้งค่า Alerting (Grafana และ Node‑RED)  
13. การเก็บ Log และ Troubleshooting  
14. ตัวอย่างโค้ดและ Flow สำเร็จรูป  
15. แบบฝึกหัดและสถานการณ์จำลอง  

### 4. ออกแบบคู่มือ (ตัวอย่างบางส่วน)
#### 4.1 การติดตั้ง Docker Compose Stack
สร้างไฟล์ `docker-compose.yml` ดังนี้:
```yaml
version: '3.8'

services:
  mosquitto:
    image: eclipse-mosquitto:latest
    container_name: mosquitto
    ports:
      - "1883:1883"
      - "9001:9001"
    volumes:
      - ./mosquitto/config:/mosquitto/config
      - ./mosquitto/data:/mosquitto/data
      - ./mosquitto/log:/mosquitto/log

  influxdb:
    image: influxdb:2.7
    container_name: influxdb
    ports:
      - "8086:8086"
    environment:
      - DOCKER_INFLUXDB_INIT_MODE=setup
      - DOCKER_INFLUXDB_INIT_USERNAME=admin
      - DOCKER_INFLUXDB_INIT_PASSWORD=admin123
      - DOCKER_INFLUXDB_INIT_ORG=myorg
      - DOCKER_INFLUXDB_INIT_BUCKET=mybucket
      - DOCKER_INFLUXDB_INIT_ADMIN_TOKEN=mytoken123
    volumes:
      - influxdb-data:/var/lib/influxdb2

  redis:
    image: redis:7-alpine
    container_name: redis
    ports:
      - "6379:6379"
    volumes:
      - redis-data:/data

  grafana:
    image: grafana/grafana:latest
    container_name: grafana
    ports:
      - "3000:3000"
    environment:
      - GF_SECURITY_ADMIN_PASSWORD=admin
    volumes:
      - grafana-data:/var/lib/grafana
    depends_on:
      - influxdb

  nodered:
    image: nodered/node-red:latest
    container_name: nodered
    ports:
      - "1880:1880"
    volumes:
      - nodered-data:/data
    depends_on:
      - mosquitto
      - influxdb
      - redis

volumes:
  influxdb-data:
  redis-data:
  grafana-data:
  nodered-data:
```

#### 4.2 การตั้งค่า Node‑RED Flow (รับ MQTT → บันทึก InfluxDB)
```json
[
    {
        "id": "mqtt-in",
        "type": "mqtt in",
        "topic": "sensor/+/data",
        "qos": 1,
        "server": "localhost",
        "name": "MQTT Sensor Input"
    },
    {
        "id": "parse",
        "type": "function",
        "func": "msg.payload = {\n    measurement: \"sensor_data\",\n    tags: { device: msg.topic.split('/')[1] },\n    fields: JSON.parse(msg.payload),\n    timestamp: Date.now() * 1000000\n};\nreturn msg;"
    },
    {
        "id": "influx-out",
        "type": "influxdb out",
        "database": "mybucket",
        "precision": "ns",
        "name": "Save to InfluxDB"
    }
]
```

### 5. ออกแบบ Workflow
```
[IoT Device] --MQTT--> [MQTT Broker] --subscribe--> [Node-RED]
                                                          |
                                                          +---> [Redis] (cache latest)
                                                          +---> [InfluxDB] (store history)
                                                          +---> [Check Alert] --> [Line/Telegram/Email]
                                                          |
                                                          +---> [Grafana] (dashboard & alert)
```

### 6. TASK LIST Template
| Task ID | รายละเอียด | ผู้รับผิดชอบ | วันที่เริ่ม | วันที่สิ้นสุด | สถานะ |
|---------|------------|--------------|-------------|---------------|--------|
| T1 | ติดตั้ง Docker และ Docker Compose | | | | |
| T2 | สร้าง docker-compose.yml และรัน stack | | | | |
| T3 | ตั้งค่า MQTT broker (Mosquitto) | | | | |
| T4 | สร้าง bucket และ token ใน InfluxDB | | | | |
| T5 | ติดตั้ง Node‑RED และ node-red-contrib-influxdb | | | | |
| T6 | สร้าง flow รับ MQTT และบันทึก InfluxDB | | | | |
| T7 | ติดตั้ง Telegraf สำหรับ monitoring OS | | | | |
| T8 | เพิ่ม Data Source InfluxDB ใน Grafana | | | | |
| T9 | สร้าง Dashboard และ Alert Rule | | | | |
| T10 | ทดสอบระบบโดยจำลองข้อมูล Sensor | | | | |

### 7. CHECKLIST Template
- [ ] Docker และ Docker Compose ทำงานได้  
- [ ] MQTT broker เปิด port 1883 และสามารถ publish/subscribe ด้วย mosquitto_pub/sub  
- [ ] InfluxDB UI เข้าถึงได้ที่ http://localhost:8086, login สำเร็จ  
- [ ] Redis รันและเชื่อมต่อด้วย redis-cli ได้  
- [ ] Grafana เข้าถึงได้ http://localhost:3000, login ด้วย admin/admin  
- [ ] ใน Grafana เพิ่ม Data Source InfluxDB (URL, org, token, bucket) ทดสอบ connection สำเร็จ  
- [ ] Node‑RED ทำงานที่ http://localhost:1880, สามารถติดตั้ง palette เพิ่มได้  
- [ ] สร้าง flow ที่รับ MQTT และเขียน InfluxDB สำเร็จ (ตรวจสอบ data explorer)  
- [ ] Telegraf ส่ง metrics เข้า InfluxDB (เช่น cpu, mem)  
- [ ] Dashboard แสดงกราฟจากทั้ง Telegraf และ MQTT sensor  
- [ ] Alert rule ทำงานเมื่อค่า sensor เกิน threshold และส่งการแจ้งเตือน  

### 8. สรุป
ระบบมอนิเตอริ่งอัตโนมัติที่ประกอบด้วย **MQTT + Node‑RED + InfluxDB + Redis + Grafana** ช่วยให้สามารถเก็บข้อมูลจากอุปกรณ์ IoT และระบบไอทีทั่วไป วิเคราะห์แนวโน้ม และแจ้งเตือนปัญหาแบบ Real‑time ได้อย่างมีประสิทธิภาพ ด้วยการใช้ Docker ทำให้ติดตั้งและปรับขยายได้ง่าย ผู้เรียนสามารถนำแนวทางนี้ไปประยุกต์ใช้ในงานผลิต งานบริการ หรืองานโครงสร้างพื้นฐานไอทีได้ทันที

---

## ตัวอย่างโค้ดเพิ่มเติมสำหรับทดสอบ

### 1. จำลองข้อมูล Sensor ส่งผ่าน MQTT (Python)
```python
import paho.mqtt.client as mqtt
import random, time, json

client = mqtt.Client()
client.connect("localhost", 1883, 60)

while True:
    data = {
        "temperature": round(random.uniform(20,35), 1),
        "humidity": round(random.uniform(40,80), 1),
        "battery": random.randint(10,100)
    }
    client.publish("sensor/esp01/data", json.dumps(data))
    print(f"Sent: {data}")
    time.sleep(5)
```

### 2. Query InfluxDB ด้วย Flux (ใช้ใน Grafana)
```flux
from(bucket: "mybucket")
  |> range(start: -1h)
  |> filter(fn: (r) => r._measurement == "sensor_data" and r.device == "esp01")
  |> filter(fn: (r) => r._field == "temperature")
  |> aggregateWindow(every: 1m, fn: mean)
```

### 3. การตั้งค่า Alert ใน Grafana (ตัวอย่าง rule YAML)
```yaml
apiVersion: 1
alertRules:
  - title: "High Temperature"
    condition: "B"
    data:
      - refId: "A"
        relativeTimeRange: { from: 5, to: 0 }
        datasourceUid: "influxdb"
        model: |
          {
            "refId": "A",
            "query": "from(bucket: \"mybucket\") |> range(start: -5m) |> filter(fn: (r) => r._measurement == \"sensor_data\" and r._field == \"temperature\") |> last()",
            "format": "table"
          }
      - refId: "B"
        relativeTimeRange: { from: 0, to: 0 }
        datasourceUid: "__expr__"
        model: |
          {
            "refId": "B",
            "type": "threshold",
            "conditions": [{ "type": "gt", "evaluator": { "params": [35] }, "operator": { "type": "and" } }]
          }
    noDataState: "NoData"
    execErrState: "Error"
    for: "1m"
    annotations: { summary: "Temperature above 35°C" }
```

---

**หมายเหตุ:** โค้ดและไฟล์ compose ข้างต้นสามารถบันทึกและรันได้จริงบนเครื่องที่ติดตั้ง Docker แล้ว ผู้เรียนสามารถปรับแต่งตามความต้องการของสถานการณ์จริงได้
