package networkmodels

import (
	models "zonaquant/models/dbmodels"

	"github.com/golang-jwt/jwt/v5"
)

type CreateUserPayload struct {
	Name     string `json:"name" form:"name"`
	LastName string `json:"lastName" form:"lastName"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	RoleID   int    `json:"roleId" form:"roleId"`
	Box      bool   `json:"box" form:"box"`
}

func (p *CreateUserPayload) ToDb() models.User {
	return models.User{
		Name:     p.Name,
		LastName: p.LastName,
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
