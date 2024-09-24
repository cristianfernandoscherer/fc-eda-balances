package entity

import (
	"github.com/google/uuid"
)

type Balance struct {
	ID      string  `json:"id"`
	Account string  `json:"account"`
	Amount  float64 `json:"amount"`
}

func NewBalance(
	account string,
	amount float64,
) (*Balance, error) {
	client := &Balance{
		ID:      uuid.New().String(),
		Account: account,
		Amount:  amount,
	}
	return client, nil
}
