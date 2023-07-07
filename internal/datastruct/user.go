package datastruct

import "github.com/google/uuid"

type UserJWT struct {
	Email string    `json:"email"`
	ID    uuid.UUID `json:"id"`
}
