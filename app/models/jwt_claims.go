package models

import "github.com/golang-jwt/jwt/v4"

type JwtClaims struct {
	Subject  string `json:"sub"`
	Expiry   int64  `json:"exp"`
	IssuedAt int64  `json:"iat"`
	Issuer   string `json:"iss"`
	jwt.Claims
}
