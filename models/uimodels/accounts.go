package uimodels

import (
	"strconv"
	models "zonaquant/models/dbmodels"
)

type UIuser struct {
	FullName string
	Username string
	ID       uint64
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
		ID:       uint64(u.ID),
		Role:     u.Role.Role,
	}

	for _, account := range u.Accounts {
		uiuser.Accounts = append(uiuser.Accounts, UIaccount{
			AccountID:     strconv.FormatInt(int64(account.ID), 10),
			AccountNumber: account.AccountNumber,
			AccountType:   account.AccountType.AccountType,
			Balance:       balanceToString(account.Balance),
		})
	}

	return uiuser
}

func balanceToString(balance float64) string {
	if balance >= 0 {
		return "$" + strconv.FormatFloat(balance, 'f', 2, 64)
	} else {
		return "-$" + strconv.FormatFloat(-balance, 'f', 2, 64)
	}
}
