package users

import (
	"context"
	"time"

	"icmongolang/internal"
	"icmongolang/internal/models"
)

type UserPgRepository interface {
	internal.PgRepository[models.SdUser]
	GetByEmail(ctx context.Context, email string) (*models.SdUser, error)
	UpdatePassword(ctx context.Context, exp *models.SdUser, newPassword string) (*models.SdUser, error)
	UpdateVerificationCode(ctx context.Context, exp *models.SdUser, newVerificationCode string) (*models.SdUser, error)
	UpdateVerification(ctx context.Context, exp *models.SdUser, newVerificationCode string, newVerified bool) (*models.SdUser, error)
	GetByVerificationCode(ctx context.Context, verificationCode string) (*models.SdUser, error)
	UpdatePasswordReset(ctx context.Context, exp *models.SdUser, passwordResetToken string, passwordResetAt time.Time) (*models.SdUser, error)
	GetByResetToken(ctx context.Context, resetToken string) (*models.SdUser, error)
	GetByResetTokenResetAt(ctx context.Context, resetToken string, resetAt time.Time) (*models.SdUser, error)
	UpdatePasswordResetToken(ctx context.Context, exp *models.SdUser, newPassword string, resetToken string) (*models.SdUser, error)
}
