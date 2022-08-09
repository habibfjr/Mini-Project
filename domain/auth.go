package domain

import (
	"database/sql"

	"github.com/golang-jwt/jwt/v4"
)

type AccessTokenClaims struct {
	Username  string        `json:"username"`
	CompanyID sql.NullInt32 `json:"company_id"`
	Role      string        `json:"role"`
	jwt.StandardClaims
}
