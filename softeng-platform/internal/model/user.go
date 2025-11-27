package model

import (
	"time"
)

type User struct {
	ID          int       `json:"id" db:"id"`
	Username    string    `json:"username" db:"username"`
	Nickname    string    `json:"nickname" db:"nickname"`
	Email       string    `json:"email" db:"email"`
	Password    string    `json:"-" db:"password"`
	Avatar      string    `json:"avater" db:"avatar"`
	Description string    `json:"description" db:"description"`
	FacePhoto   string    `json:"face_photo" db:"face_photo"`
	Role        string    `json:"role" db:"role"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type RegisterRequest struct {
	Username        string `json:"username" binding:"required"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,min=6"`
	EmailPassword   string `json:"email_password" binding:"required"`
	CertifyPassword string `json:"certify_password" binding:"required"`
}

type LoginRequest struct {
	UsernameOrEmail string `json:"username_or_email" binding:"required"`
	Password        string `json:"password" binding:"required"`
}

type UpdateProfileRequest struct {
	Nickname    string `json:"nickname"`
	Avatar      string `json:"avater"`
	Description string `json:"description"`
	FacePhoto   string `json:"face_photo"`
}
