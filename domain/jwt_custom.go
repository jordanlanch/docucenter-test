package domain

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

// JwtCustomClaims representa los claims personalizados que se incluyen en el JWT de la plataforma .
type JwtCustomClaims struct {
	Email string    `json:"email"`
	ID    uuid.UUID `json:"id"`
	jwt.StandardClaims
}

// JwtCustomRefreshClaims representa los claims personalizados que se incluyen en el JWT de actualización en la plataforma .
type JwtCustomRefreshClaims struct {
	ID uuid.UUID `json:"id"`
	jwt.StandardClaims
}
