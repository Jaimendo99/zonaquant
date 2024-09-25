package networkmodels

import (
	models "zonaquant/models/dbmodels"
)

type CreateUserPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
	RoleID   int    `json:"roleId"`
}

func (p *CreateUserPayload) ToDb() models.User {
	return models.User{
		Username: p.Username,
		Password: p.Password,
		RoleID:   uint(p.RoleID),
	}
}
