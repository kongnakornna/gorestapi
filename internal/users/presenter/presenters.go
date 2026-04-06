package processor

import (
	"context"
	"encoding/json"
	"fmt"

	"gorestapi/config"
	"gorestapi/internal/processor"
	"gorestapi/internal/users"
	"gorestapi/pkg/logger"
	"gorestapi/pkg/sendEmail"

	"github.com/hibiken/asynq"
)

type userRedisTaskProcessor struct {
	processor.RedisTaskProcessor
	emailSender sendEmail.EmailSender
}

// NewUserRedisTaskProcessor creates a user-specific task processor.
func NewUserRedisTaskProcessor(server *asynq.Server, cfg *config.Config, logger logger.Logger, emailSender sendEmail.EmailSender) users.UserRedisTaskProcessor {
	return &userRedisTaskProcessor{
		RedisTaskProcessor: processor.NewRedisTaskProcessor(server, cfg, logger),
		emailSender:        emailSender,
	}
}

// ProcessTaskSendEmail handles the send-email task.
func (p *userRedisTaskProcessor) ProcessTaskSendEmail(ctx context.Context, task *asynq.Task) error {
	var payload users.PayloadSendEmail
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("failed to unmarshal payload: %w", asynq.SkipRetry)
	}

	if err := p.emailSender.SendEmail(
		ctx,
		payload.From,
		payload.To,
		payload.Subject,
		payload.BodyHtml,
		payload.BodyPlain,
	); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	p.Logger.Infof("Type: %v, Msg: email sent to %s", task.Type(), payload.To)
	return nil
}

// UserCreate ใช้สำหรับการสร้างผู้ใช้ (ลงทะเบียน)
// UserCreate is used for user registration
type UserCreate struct {
	Name            string `json:"name" validate:"required" example:"Xuan Hien"`
	Email           string `json:"email" validate:"required,email" example:"user@example.com"`
	Password        string `json:"password" validate:"required,min=8" example:"password123"`
	ConfirmPassword string `json:"confirm_password" validate:"required,min=8" example:"password123"`
}

// UserUpdate ใช้อัปเดตข้อมูลผู้ใช้
// UserUpdate is used for updating user info
type UserUpdate struct {
	Name string `json:"name" example:"New Name"`
}

// UserUpdatePassword ใช้เปลี่ยนรหัสผ่าน
// UserUpdatePassword is used for password change
type UserUpdatePassword struct {
	OldPassword     string `json:"old_password" validate:"required,min=8" example:"oldpass"`
	NewPassword     string `json:"new_password" validate:"required,min=8" example:"newpass"`
	ConfirmPassword string `json:"confirm_password" validate:"required,min=8" example:"newpass"`
}

// UserResponse สำหรับตอบกลับข้อมูลผู้ใช้
// UserResponse is the response structure for user data
type UserResponse struct {
	Id          uuid.UUID `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Email       string    `json:"email,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	IsActive    bool      `json:"is_active"`
	IsSuperUser bool      `json:"is_superuser"`
	Verified    bool      `json:"verified"`
}

// UserSignIn สำหรับ login
// UserSignIn is used for login
type UserSignIn struct {
	Email    string `json:"email" validate:"required,email" example:"user@example.com"`
	Password string `json:"password" validate:"required,min=8" example:"password123"`
}

// Token สำหรับตอบกลับ access/refresh token
// Token response structure
type Token struct {
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	TokenType    string `json:"token_type,omitempty"`
}

// RefreshRequest ใช้สำหรับ refresh token และ logout
// RefreshRequest is used for refresh and logout operations
type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

// ForgotPassword ใช้ขอรีเซ็ตรหัสผ่าน
// ForgotPassword request
type ForgotPassword struct {
	Email string `json:"email" validate:"required,email" example:"user@example.com"`
}

// ResetPassword ใช้ตั้งรหัสผ่านใหม่
// ResetPassword request
type ResetPassword struct {
	NewPassword     string `json:"new_password" validate:"required,min=8" example:"newpass123"`
	ConfirmPassword string `json:"confirm_password" validate:"required,min=8" example:"newpass123"`
}

// ResendVerificationRequest ใช้ขอส่งอีเมลยืนยันใหม่
// ResendVerificationRequest request
type ResendVerificationRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type PublicKey struct {
	PublicKeyAccessToken  string `json:"public_key_access_token,omitempty"`
	PublicKeyRefreshToken string `json:"public_key_refresh_token,omitempty"`
}

