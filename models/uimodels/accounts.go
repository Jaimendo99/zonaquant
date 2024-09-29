package uimodels

import (
	"strconv"
	models "zonaquant/models/dbmodels"
)

type UIuser struct {
	FullName string
	Username string
	Role     string
	Accounts []UIaccount
}

type UIaccount struct {
	AccountID     string
	AccountNumber string
	AccountType   string
	Balance       string
}

func UserToUI(u *models.User) UIuser {
	uiuser := UIuser{
		FullName: u.Name + " " + u.LastName,
		Username: u.Username,
		Role:     u.Role.Role,
	}

	for _, account := range u.Accounts {
		uiuser.Accounts = append(uiuser.Accounts, UIaccount{
			AccountID:     strconv.FormatInt(int64(account.ID), 10),
			AccountNumber: account.AccountNumber,
			AccountType:   account.AccountType.AccountType,
			Balance:       "$" + strconv.FormatFloat(float64(account.Balance), 'f', 2, 64),
		})
	}

	return uiuser
}
