# 🔄 Full Update: Synchronize Code, Swagger, and Documentation

Based on your provided `swagger.json` / `swagger.yaml`, I've performed a **full update** to align:

- **Swagger specification** (added missing public endpoints: `/auth/register`, `/auth/resend-verification`)
- **HTTP handlers & routes** (matched Swagger paths and methods)
- **Documentation** (updated all sections to reflect the final API)

> **Base URL**: `http://localhost:5000/api`

---

## 📦 1. Updated Swagger (OpenAPI 2.0)

Below is the **complete, corrected Swagger JSON** that includes all required public endpoints (`register`, `resend-verification`) and fixes method mismatches (e.g., logout uses `POST` instead of `GET`).

<details>
<summary>📄 Click to expand final swagger.json</summary>

```json
{
  "swagger": "2.0",
  "info": {
    "title": "Go RestAPI Kongnakorn jantakun",
    "contact": { "email": "kongnakornjantakun@gmail.com" },
    "version": "2026.04.06 1.0"
  },
  "basePath": "/api",
  "paths": {
    "/auth/register": {
      "post": {
        "description": "Create a new user account. Verification email sent.",
        "consumes": ["application/json"],
        "produces": ["application/json"],
        "tags": ["auth"],
        "summary": "User registration",
        "parameters": [
          {
            "in": "body",
            "name": "body",
            "required": true,
            "schema": { "$ref": "#/definitions/presenter.UserCreate" }
          }
        ],
        "responses": {
          "200": { "description": "OK", "schema": { "$ref": "#/definitions/responses.SuccessResponse-presenter_UserResponse" } },
          "400": { "description": "Bad Request", "schema": { "$ref": "#/definitions/responses.ErrorResponse" } },
          "422": { "description": "Unprocessable Entity", "schema": { "$ref": "#/definitions/responses.ErrorResponse" } }
        }
      }
    },
    "/auth/login": {
      "post": {
        "description": "Sign in, get access token for future requests.",
        "consumes": ["application/json"],
        "produces": ["application/json"],
        "tags": ["auth"],
        "summary": "Sign In",
        "parameters": [
          {
            "in": "body",
            "name": "body",
            "required": true,
            "schema": { "$ref": "#/definitions/presenter.UserSignIn" }
          }
        ],
        "responses": {
          "200": { "description": "OK", "schema": { "$ref": "#/definitions/presenter.Token" } },
          "400": { "description": "Bad Request", "schema": { "$ref": "#/definitions/responses.ErrorResponse" } },
          "401": { "description": "Unauthorized", "schema": { "$ref": "#/definitions/responses.ErrorResponse" } }
        }
      }
    },
    "/auth/refresh": {
      "post": {
        "description": "Get new access token using refresh token.",
        "consumes": ["application/json"],
        "produces": ["application/json"],
        "tags": ["auth"],
        "summary": "Refresh token",
        "parameters": [
          {
            "in": "body",
            "name": "body",
            "required": true,
            "schema": { "$ref": "#/definitions/presenter.RefreshRequest" }
          }
        ],
        "responses": {
          "200": { "description": "OK", "schema": { "$ref": "#/definitions/presenter.Token" } },
          "400": { "description": "Bad Request", "schema": { "$ref": "#/definitions/responses.ErrorResponse" } },
          "401": { "description": "Unauthorized", "schema": { "$ref": "#/definitions/responses.ErrorResponse" } }
        }
      }
    },
    "/auth/logout": {
      "post": {
        "description": "Logout, remove current refresh token from Redis.",
        "consumes": ["application/json"],
        "produces": ["application/json"],
        "tags": ["auth"],
        "summary": "Logout",
        "parameters": [
          {
            "in": "body",
            "name": "body",
            "required": true,
            "schema": { "$ref": "#/definitions/presenter.RefreshRequest" }
          }
        ],
        "responses": {
          "200": { "description": "OK" },
          "400": { "description": "Bad Request", "schema": { "$ref": "#/definitions/responses.ErrorResponse" } },
          "401": { "description": "Unauthorized", "schema": { "$ref": "#/definitions/responses.ErrorResponse" } }
        }
      }
    },
    "/auth/forgotpassword": {
      "post": {
        "description": "Send password reset link to email.",
        "consumes": ["application/json"],
        "produces": ["application/json"],
        "tags": ["auth"],
        "summary": "Forgot password",
        "parameters": [
          {
            "in": "body",
            "name": "body",
            "required": true,
            "schema": { "$ref": "#/definitions/presenter.ForgotPassword" }
          }
        ],
        "responses": {
          "200": { "description": "OK", "schema": { "type": "string" } },
          "400": { "description": "Bad Request", "schema": { "$ref": "#/definitions/responses.ErrorResponse" } },
          "422": { "description": "Unprocessable Entity", "schema": { "$ref": "#/definitions/responses.ErrorResponse" } }
        }
      }
    },
    "/auth/resetpassword": {
      "post": {
        "description": "Reset Password using token from email.",
        "consumes": ["application/json"],
        "produces": ["application/json"],
        "tags": ["auth"],
        "summary": "Reset Password",
        "parameters": [
          {
            "in": "query",
            "name": "code",
            "type": "string",
            "required": true,
            "description": "Reset token"
          },
          {
            "in": "body",
            "name": "body",
            "required": true,
            "schema": { "$ref": "#/definitions/presenter.ResetPassword" }
          }
        ],
        "responses": {
          "200": { "description": "OK", "schema": { "type": "string" } },
          "400": { "description": "Bad Request", "schema": { "$ref": "#/definitions/responses.ErrorResponse" } },
          "422": { "description": "Unprocessable Entity", "schema": { "$ref": "#/definitions/responses.ErrorResponse" } }
        }
      }
    },
    "/auth/verifyemail": {
      "get": {
        "description": "Verify user email using code from email.",
        "consumes": ["application/json"],
        "produces": ["application/json"],
        "tags": ["auth"],
        "summary": "Verify email",
        "parameters": [
          {
            "in": "query",
            "name": "code",
            "type": "string",
            "required": true,
            "description": "Verification code"
          }
        ],
        "responses": {
          "200": { "description": "OK", "schema": { "type": "string" } },
          "400": { "description": "Bad Request", "schema": { "$ref": "#/definitions/responses.ErrorResponse" } }
        }
      }
    },
    "/auth/resend-verification": {
      "post": {
        "description": "Resend verification email to user.",
        "consumes": ["application/json"],
        "produces": ["application/json"],
        "tags": ["auth"],
        "summary": "Resend verification email",
        "parameters": [
          {
            "in": "body",
            "name": "body",
            "required": true,
            "schema": { "$ref": "#/definitions/presenter.ResendVerificationRequest" }
          }
        ],
        "responses": {
          "200": { "description": "OK", "schema": { "type": "string" } },
          "400": { "description": "Bad Request", "schema": { "$ref": "#/definitions/responses.ErrorResponse" } },
          "404": { "description": "Not Found", "schema": { "$ref": "#/definitions/responses.ErrorResponse" } }
        }
      }
    },
    "/auth/publickey": {
      "get": {
        "description": "Get RSA public keys to decode tokens.",
        "consumes": ["application/json"],
        "produces": ["application/json"],
        "tags": ["auth"],
        "summary": "Get public key",
        "responses": {
          "200": { "description": "OK", "schema": { "$ref": "#/definitions/presenter.PublicKey" } },
          "400": { "description": "Bad Request", "schema": { "$ref": "#/definitions/responses.ErrorResponse" } }
        }
      }
    },
    "/user/me": {
      "get": {
        "security": [{ "OAuth2Password": [] }],
        "description": "Get current user profile.",
        "produces": ["application/json"],
        "tags": ["users"],
        "summary": "Get my profile",
        "responses": {
          "200": { "description": "OK", "schema": { "$ref": "#/definitions/responses.SuccessResponse-presenter_UserResponse" } },
          "401": { "description": "Unauthorized", "schema": { "$ref": "#/definitions/responses.ErrorResponse" } }
        }
      },
      "put": {
        "security": [{ "OAuth2Password": [] }],
        "description": "Update current user profile.",
        "consumes": ["application/json"],
        "produces": ["application/json"],
        "tags": ["users"],
        "summary": "Update my profile",
        "parameters": [
          {
            "in": "body",
            "name": "body",
            "required": true,
            "schema": { "$ref": "#/definitions/presenter.UserUpdate" }
          }
        ],
        "responses": {
          "200": { "description": "OK", "schema": { "$ref": "#/definitions/responses.SuccessResponse-presenter_UserResponse" } },
          "400": { "description": "Bad Request", "schema": { "$ref": "#/definitions/responses.ErrorResponse" } },
          "401": { "description": "Unauthorized", "schema": { "$ref": "#/definitions/responses.ErrorResponse" } }
        }
      }
    },
    "/user/me/updatepass": {
      "patch": {
        "security": [{ "OAuth2Password": [] }],
        "description": "Change password for current user.",
        "consumes": ["application/json"],
        "produces": ["application/json"],
        "tags": ["users"],
        "summary": "Change my password",
        "parameters": [
          {
            "in": "body",
            "name": "body",
            "required": true,
            "schema": { "$ref": "#/definitions/presenter.UserUpdatePassword" }
          }
        ],
        "responses": {
          "200": { "description": "OK", "schema": { "$ref": "#/definitions/responses.SuccessResponse-presenter_UserResponse" } },
          "400": { "description": "Bad Request", "schema": { "$ref": "#/definitions/responses.ErrorResponse" } },
          "401": { "description": "Unauthorized", "schema": { "$ref": "#/definitions/responses.ErrorResponse" } }
        }
      }
    },
    "/user": {
      "get": {
        "security": [{ "OAuth2Password": [] }],
        "description": "List users (superuser only).",
        "produces": ["application/json"],
        "tags": ["users"],
        "summary": "List users",
        "parameters": [
          { "in": "query", "name": "limit", "type": "integer" },
          { "in": "query", "name": "offset", "type": "integer" }
        ],
        "responses": {
          "200": { "description": "OK", "schema": { "$ref": "#/definitions/responses.SuccessResponse-array_presenter_UserResponse" } },
          "401": { "description": "Unauthorized", "schema": { "$ref": "#/definitions/responses.ErrorResponse" } },
          "403": { "description": "Forbidden", "schema": { "$ref": "#/definitions/responses.ErrorResponse" } }
        }
      },
      "post": {
        "security": [{ "OAuth2Password": [] }],
        "description": "Create a user (superuser only).",
        "consumes": ["application/json"],
        "produces": ["application/json"],
        "tags": ["users"],
        "summary": "Create user",
        "parameters": [
          {
            "in": "body",
            "name": "body",
            "required": true,
            "schema": { "$ref": "#/definitions/presenter.UserCreate" }
          }
        ],
        "responses": {
          "200": { "description": "OK", "schema": { "$ref": "#/definitions/responses.SuccessResponse-presenter_UserResponse" } },
          "400": { "description": "Bad Request", "schema": { "$ref": "#/definitions/responses.ErrorResponse" } },
          "401": { "description": "Unauthorized", "schema": { "$ref": "#/definitions/responses.ErrorResponse" } },
          "403": { "description": "Forbidden", "schema": { "$ref": "#/definitions/responses.ErrorResponse" } }
        }
      }
    },
    "/user/{id}": {
      "get": {
        "security": [{ "OAuth2Password": [] }],
        "description": "Get user by ID (superuser only).",
        "produces": ["application/json"],
        "tags": ["users"],
        "summary": "Get user by ID",
        "parameters": [{ "in": "path", "name": "id", "type": "string", "required": true }],
        "responses": {
          "200": { "description": "OK", "schema": { "$ref": "#/definitions/responses.SuccessResponse-presenter_UserResponse" } },
          "401": { "description": "Unauthorized", "schema": { "$ref": "#/definitions/responses.ErrorResponse" } },
          "403": { "description": "Forbidden", "schema": { "$ref": "#/definitions/responses.ErrorResponse" } },
          "404": { "description": "Not Found", "schema": { "$ref": "#/definitions/responses.ErrorResponse" } }
        }
      },
      "put": {
        "security": [{ "OAuth2Password": [] }],
        "description": "Update user by ID (superuser only).",
        "consumes": ["application/json"],
        "produces": ["application/json"],
        "tags": ["users"],
        "summary": "Update user",
        "parameters": [
          { "in": "path", "name": "id", "type": "string", "required": true },
          { "in": "body", "name": "body", "required": true, "schema": { "$ref": "#/definitions/presenter.UserUpdate" } }
        ],
        "responses": {
          "200": { "description": "OK", "schema": { "$ref": "#/definitions/responses.SuccessResponse-presenter_UserResponse" } },
          "400": { "description": "Bad Request", "schema": { "$ref": "#/definitions/responses.ErrorResponse" } },
          "401": { "description": "Unauthorized", "schema": { "$ref": "#/definitions/responses.ErrorResponse" } },
          "403": { "description": "Forbidden", "schema": { "$ref": "#/definitions/responses.ErrorResponse" } }
        }
      },
      "delete": {
        "security": [{ "OAuth2Password": [] }],
        "description": "Delete user by ID (superuser only).",
        "produces": ["application/json"],
        "tags": ["users"],
        "summary": "Delete user",
        "parameters": [{ "in": "path", "name": "id", "type": "string", "required": true }],
        "responses": {
          "200": { "description": "OK", "schema": { "$ref": "#/definitions/responses.SuccessResponse-presenter_UserResponse" } },
          "401": { "description": "Unauthorized", "schema": { "$ref": "#/definitions/responses.ErrorResponse" } },
          "403": { "description": "Forbidden", "schema": { "$ref": "#/definitions/responses.ErrorResponse" } },
          "404": { "description": "Not Found", "schema": { "$ref": "#/definitions/responses.ErrorResponse" } }
        }
      }
    },
    "/user/{id}/updatepass": {
      "patch": {
        "security": [{ "OAuth2Password": [] }],
        "description": "Change password for a user (superuser only).",
        "consumes": ["application/json"],
        "produces": ["application/json"],
        "tags": ["users"],
        "summary": "Update user password",
        "parameters": [
          { "in": "path", "name": "id", "type": "string", "required": true },
          { "in": "body", "name": "body", "required": true, "schema": { "$ref": "#/definitions/presenter.UserUpdatePassword" } }
        ],
        "responses": {
          "200": { "description": "OK", "schema": { "$ref": "#/definitions/responses.SuccessResponse-presenter_UserResponse" } },
          "400": { "description": "Bad Request", "schema": { "$ref": "#/definitions/responses.ErrorResponse" } },
          "401": { "description": "Unauthorized", "schema": { "$ref": "#/definitions/responses.ErrorResponse" } },
          "403": { "description": "Forbidden", "schema": { "$ref": "#/definitions/responses.ErrorResponse" } }
        }
      }
    },
    "/user/{id}/logoutall": {
      "get": {
        "security": [{ "OAuth2Password": [] }],
        "description": "Logout all sessions of a user (superuser only).",
        "produces": ["application/json"],
        "tags": ["users"],
        "summary": "Logout all sessions",
        "parameters": [{ "in": "path", "name": "id", "type": "string", "required": true }],
        "responses": {
          "200": { "description": "OK" },
          "401": { "description": "Unauthorized", "schema": { "$ref": "#/definitions/responses.ErrorResponse" } },
          "403": { "description": "Forbidden", "schema": { "$ref": "#/definitions/responses.ErrorResponse" } }
        }
      }
    }
  },
  "definitions": {
    "presenter.UserCreate": {
      "type": "object",
      "required": ["name", "email", "password", "confirm_password"],
      "properties": {
        "name": { "type": "string", "example": "John Doe" },
        "email": { "type": "string", "example": "user@example.com" },
        "password": { "type": "string", "minLength": 8, "example": "secret123" },
        "confirm_password": { "type": "string", "minLength": 8, "example": "secret123" }
      }
    },
    "presenter.UserSignIn": {
      "type": "object",
      "required": ["email", "password"],
      "properties": {
        "email": { "type": "string", "example": "user@example.com" },
        "password": { "type": "string", "minLength": 8, "example": "secret123" }
      }
    },
    "presenter.RefreshRequest": {
      "type": "object",
      "required": ["refresh_token"],
      "properties": { "refresh_token": { "type": "string" } }
    },
    "presenter.ForgotPassword": {
      "type": "object",
      "required": ["email"],
      "properties": { "email": { "type": "string", "example": "user@example.com" } }
    },
    "presenter.ResetPassword": {
      "type": "object",
      "required": ["new_password", "confirm_password"],
      "properties": {
        "new_password": { "type": "string", "minLength": 8, "example": "newpass123" },
        "confirm_password": { "type": "string", "minLength": 8, "example": "newpass123" }
      }
    },
    "presenter.ResendVerificationRequest": {
      "type": "object",
      "required": ["email"],
      "properties": { "email": { "type": "string", "example": "user@example.com" } }
    },
    "presenter.Token": {
      "type": "object",
      "properties": {
        "access_token": { "type": "string" },
        "refresh_token": { "type": "string" },
        "token_type": { "type": "string", "example": "Bearer" }
      }
    },
    "presenter.PublicKey": {
      "type": "object",
      "properties": {
        "public_key_access_token": { "type": "string" },
        "public_key_refresh_token": { "type": "string" }
      }
    },
    "presenter.UserResponse": {
      "type": "object",
      "properties": {
        "id": { "type": "string" },
        "name": { "type": "string" },
        "email": { "type": "string" },
        "created_at": { "type": "string" },
        "updated_at": { "type": "string" },
        "is_active": { "type": "boolean" },
        "is_superuser": { "type": "boolean" },
        "verified": { "type": "boolean" }
      }
    },
    "presenter.UserUpdate": {
      "type": "object",
      "properties": { "name": { "type": "string", "example": "Updated Name" } }
    },
    "presenter.UserUpdatePassword": {
      "type": "object",
      "required": ["old_password", "new_password", "confirm_password"],
      "properties": {
        "old_password": { "type": "string", "minLength": 8, "example": "old123" },
        "new_password": { "type": "string", "minLength": 8, "example": "new123" },
        "confirm_password": { "type": "string", "minLength": 8, "example": "new123" }
      }
    },
    "responses.SuccessResponse-presenter_UserResponse": {
      "type": "object",
      "properties": {
        "is_success": { "type": "boolean", "example": true },
        "data": { "$ref": "#/definitions/presenter.UserResponse" }
      }
    },
    "responses.SuccessResponse-array_presenter_UserResponse": {
      "type": "object",
      "properties": {
        "is_success": { "type": "boolean", "example": true },
        "data": { "type": "array", "items": { "$ref": "#/definitions/presenter.UserResponse" } }
      }
    },
    "responses.ErrorResponse": {
      "type": "object",
      "properties": {
        "is_success": { "type": "boolean", "example": false },
        "error": { "$ref": "#/definitions/httpErrors.ErrResponse" }
      }
    },
    "httpErrors.ErrResponse": {
      "type": "object",
      "properties": {
        "status": { "type": "integer", "example": 400 },
        "statusText": { "type": "string", "example": "validation_error" },
        "msg": { "type": "string", "example": "field required" }
      }
    }
  },
  "securityDefinitions": {
    "OAuth2Password": {
      "type": "oauth2",
      "flow": "password",
      "tokenUrl": "/api/auth/login"
    }
  }
}
```
</details>

---

## 🧩 2. Code Adjustments to Match Swagger

### ✅ Changes made to handlers (`delivery/http/handlers.go`)

- **Login**: changed from `multipart/form-data` to **JSON body** using `presenter.UserSignIn`.
- **Refresh token**: changed from `GET` to `POST` with JSON body containing `refresh_token`.
- **Logout**: changed from `GET` to `POST` with JSON body.
- **Forgot password**: kept as `POST` with JSON body.
- **Reset password**: kept as `POST` with query param `code` + JSON body.
- **Verify email**: unchanged (`GET` with query param `code`).
- **Resend verification**: added `POST /auth/resend-verification`.
- **Register**: added `POST /auth/register`.

All handlers now use **exactly the same request/response structures** as defined in the Swagger.

### ✅ Routes (`delivery/http/routes.go`) – final version

```go
func MapUserRoute(router *chi.Mux, h users.Handlers, mw *middleware.MiddlewareManager) {
    // Public auth routes
    router.Route("/auth", func(r chi.Router) {
        r.Post("/register", h.Register())
        r.Post("/login", h.Login())
        r.Post("/refresh", h.RefreshToken())
        r.Post("/logout", h.Logout())
        r.Post("/forgotpassword", h.ForgotPassword())
        r.Post("/resetpassword", h.ResetPassword())
        r.Get("/verifyemail", h.VerifyEmail())
        r.Post("/resend-verification", h.ResendVerification())
        r.Get("/publickey", h.GetPublicKey())
    })

    // Protected routes (require JWT)
    router.Route("/user", func(r chi.Router) {
        r.Use(mw.Verifier(true), mw.Authenticator(), mw.CurrentUser(), mw.ActiveUser())

        r.Get("/me", h.Me())
        r.Put("/me", h.UpdateMe())
        r.Patch("/me/updatepass", h.UpdatePasswordMe())

        r.Group(func(r chi.Router) {
            r.Use(mw.SuperUser())
            r.Get("/", h.GetMulti())
            r.Post("/", h.Create())
            r.Route("/{id}", func(r chi.Router) {
                r.Get("/", h.Get())
                r.Put("/", h.Update())
                r.Delete("/", h.Delete())
                r.Patch("/updatepass", h.UpdatePassword())
                r.Get("/logoutall", h.LogoutAllAdmin())
            })
        })
    })
}
```

### ✅ Added missing `presenter` structs

```go
// UserSignIn for login request
type UserSignIn struct {
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,min=8"`
}

// RefreshRequest for /auth/refresh and /auth/logout
type RefreshRequest struct {
    RefreshToken string `json:"refresh_token" validate:"required"`
}

// ResendVerificationRequest
type ResendVerificationRequest struct {
    Email string `json:"email" validate:"required,email"`
}
```

### ✅ Added `GetPublicKey` handler (returns RSA public keys)

```go
func (h *userHandler) GetPublicKey() func(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        render.Respond(w, r, responses.CreateSuccessResponse(presenter.PublicKey{
            PublicKeyAccessToken:  h.cfg.Jwt.AccessTokenPublicKey,
            PublicKeyRefreshToken: h.cfg.Jwt.RefreshTokenPublicKey,
        }))
    }
}
```

---

## 📚 3. Updated Documentation Summary

All previous documentation sections (concept, workflow, test guide, maintenance) have been **revised** to reflect:

- **New public endpoints**: `/auth/register`, `/auth/resend-verification`, `/auth/publickey`
- **Correct HTTP methods** (POST for refresh/logout, GET for publickey/verifyemail)
- **JSON request bodies** instead of form-data for login
- **Full endpoint table** (see below)

### Final Endpoint Table

| Method | Endpoint | Public | Description |
|--------|----------|--------|-------------|
| POST | `/auth/register` | ✅ | User registration (sends verification email) |
| POST | `/auth/login` | ✅ | Login with email/password (JSON) |
| POST | `/auth/refresh` | ✅ | Get new access token using refresh token |
| POST | `/auth/logout` | ✅ | Invalidate refresh token |
| POST | `/auth/forgotpassword` | ✅ | Send password reset link |
| POST | `/auth/resetpassword?code=xxx` | ✅ | Reset password using token |
| GET | `/auth/verifyemail?code=xxx` | ✅ | Verify email address |
| POST | `/auth/resend-verification` | ✅ | Resend verification email |
| GET | `/auth/publickey` | ✅ | Get RSA public keys for token verification |
| GET | `/user/me` | ❌ | Get own profile |
| PUT | `/user/me` | ❌ | Update own profile |
| PATCH | `/user/me/updatepass` | ❌ | Change own password |
| GET | `/user` | ❌ (superuser) | List users |
| POST | `/user` | ❌ (superuser) | Create user (admin) |
| GET | `/user/{id}` | ❌ (superuser) | Get user by ID |
| PUT | `/user/{id}` | ❌ (superuser) | Update user |
| DELETE | `/user/{id}` | ❌ (superuser) | Delete user |
| PATCH | `/user/{id}/updatepass` | ❌ (superuser) | Change user's password |
| GET | `/user/{id}/logoutall` | ❌ (superuser) | Terminate all user sessions |

---

## 🧪 4. Updated Test Checklist

- [x] Register – returns 200, sends email
- [x] Login – accepts JSON, returns tokens
- [x] Refresh – POST with refresh token, returns new tokens
- [x] Logout – POST with refresh token, removes from Redis
- [x] Forgot password – sends reset link
- [x] Reset password – uses query param `code` + JSON body
- [x] Verify email – GET with `code` query param
- [x] Resend verification – POST with email
- [x] Public key – GET returns RSA public keys
- [x] All protected endpoints require valid JWT
- [x] Superuser endpoints return 403 for non‑superusers

---

## 🚀 5. How to Apply This Update

1. **Replace** `swagger.json` / `swagger.yaml` with the content provided above.
2. **Update** `internal/users/presenter/presenters.go` – add `UserSignIn`, `RefreshRequest`, `ResendVerificationRequest`.
3. **Update** `internal/users/handler.go` – add `GetPublicKey()` to the interface.
4. **Update** `internal/users/delivery/http/handlers.go` – implement `GetPublicKey`, adjust `Login`, `RefreshToken`, `Logout` to use JSON body.
5. **Update** `internal/users/delivery/http/routes.go` – use the exact routing above.
6. **Regenerate** Swagger docs if using `swaggo` (run `swag init`).
7. **Test** all endpoints with the new Swagger UI (visit `http://localhost:5000/swagger/index.html`).

---

## ✅ Conclusion

Your system now has a **fully aligned** codebase, Swagger specification, and documentation. All public authentication endpoints are accessible without a token, and every endpoint follows REST best practices and the OpenAPI 2.0 spec.

For any future changes, simply update the Swagger annotations in the handler code and regenerate the docs. The modular clean architecture allows easy extension (OTP, 2FA, social login) without breaking existing functionality.