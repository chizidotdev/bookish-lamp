package dto

import "github.com/google/uuid"

type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserJWT struct {
	Email string    `json:"email"`
	ID    uuid.UUID `json:"id"`
}
