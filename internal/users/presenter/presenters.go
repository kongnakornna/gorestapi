package presenter

import (
	"time"

	"github.com/google/uuid"
)

// UserCreate – used when creating a new user
type UserCreate struct {
	Email           string `json:"email" validate:"required,email" example:"kongnakornna@gmail.com"`
	Password        string `json:"password" validate:"required,min=8" example:"password"`
	ConfirmPassword string `json:"confirm_password" validate:"required,min=8" example:"password"`
	RoleID          int    `json:"role_id" validate:"required,min=1" example:"2"`
	Firstname       string `json:"firstname,omitempty" example:"Kongnakorn"`
	Lastname        string `json:"lastname,omitempty" example:"Jantakun"`
	Fullname        string `json:"fullname,omitempty" example:"Kongnakorn Jantakun"`
	MobileNumber    string `json:"mobile_number,omitempty" example:"0812345678"`
	PhoneNumber     string `json:"phone_number,omitempty" example:"021234567"`
	LineID          string `json:"line_id,omitempty" example:"kongnakorn_line"`
	LocationID      string `json:"location_id,omitempty" example:"loc_001"`
}

// UserUpdate – used when updating a user (all fields optional)
type UserUpdate struct {
	Firstname    *string `json:"firstname,omitempty"`
	Lastname     *string `json:"lastname,omitempty"`
	Fullname     *string `json:"fullname,omitempty"`
	MobileNumber *string `json:"mobile_number,omitempty"`
	PhoneNumber  *string `json:"phone_number,omitempty"`
	LineID       *string `json:"line_id,omitempty"`
	LocationID   *string `json:"location_id,omitempty"`
}

// UserUpdateRole – used by admin to change user role
type UserUpdateRole struct {
	RoleID int `json:"role_id" validate:"required"`
}

// UserUpdatePassword – used when changing password
type UserUpdatePassword struct {
	OldPassword     string `json:"old_password" validate:"required,min=8"`
	NewPassword     string `json:"new_password" validate:"required,min=8"`
	ConfirmPassword string `json:"confirm_password" validate:"required,min=8"`
}

// UserResponse – full user data returned by API
type UserResponse struct {
	ID           uuid.UUID `json:"id"`
	Email        string    `json:"email"`
	RoleID       int       `json:"role_id"`
	Firstname    *string   `json:"firstname,omitempty"`
	Lastname     *string   `json:"lastname,omitempty"`
	Fullname     *string   `json:"fullname,omitempty"`
	MobileNumber *string   `json:"mobile_number,omitempty"`
	PhoneNumber  *string   `json:"phone_number,omitempty"`
	LineID       *string   `json:"line_id,omitempty"`
	LocationID   *string   `json:"location_id,omitempty"`
	Status       int16     `json:"status"` // 1 active, 0 inactive
	IsSuperUser  bool      `json:"is_superuser"`
	Verified     bool      `json:"verified"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// Auth related DTOs (unchanged)
type UserSignIn struct {
	Email    string `json:"email" validate:"required" example:"kongnakornna@gmail.com"`
	Password string `json:"password" validate:"required,min=8" example:"password"`
}

type Token struct {
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	TokenType    string `json:"token_type,omitempty"`
}

type PublicKey struct {
	PublicKeyAccessToken  string `json:"public_key_access_token,omitempty"`
	PublicKeyRefreshToken string `json:"public_key_refresh_token,omitempty"`
}

type ForgotPassword struct {
	Email string `json:"email" validate:"required" example:"kongnakornna@gmail.com"`
}

type ResetPassword struct {
	NewPassword     string `json:"new_password" validate:"required,min=8" example:"password"`
	ConfirmPassword string `json:"confirm_password" validate:"required,min=8" example:"password"`
}
