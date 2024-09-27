package networkmodels

import (
	models "zonaquant/models/dbmodels"

	"github.com/golang-jwt/jwt/v5"
)

type CreateUserPayload struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	RoleID   int    `json:"roleId" form:"roleId"`
}

func (p *CreateUserPayload) ToDb() models.User {
	return models.User{
		Username: p.Username,
		Password: p.Password,
		RoleID:   uint(p.RoleID),
	}
}

type LoginPayload struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type JwtClaims struct {
	Sub  uint `json:"sub"`
	Role int  `json:"role"`
	jwt.RegisteredClaims
}
