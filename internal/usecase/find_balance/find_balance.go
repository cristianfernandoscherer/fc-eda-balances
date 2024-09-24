package find_balance

import (
	"github.com.br/cristian.scherer/eda-balance/internal/database"
)

type FindBalanceInputDTO struct {
	AccountID string
}

type FindBalanceOutputDTO struct {
	ID     string
	Amount float64
}

type FindBalanceUseCase struct {
	BallanceDB database.BalanceDB
}

func NewFindBalanceUseCase(ballanceDB database.BalanceDB) *FindBalanceUseCase {
	return &FindBalanceUseCase{
		BallanceDB: ballanceDB,
	}
}

func (usecase *FindBalanceUseCase) Executa(input FindBalanceInputDTO) (*FindBalanceOutputDTO, error) {

	err, balance := usecase.BallanceDB.FindByAccountID(input.AccountID)
	if err != nil {
		return nil, err
	}

	return &FindBalanceOutputDTO{
		ID:     input.AccountID,
		Amount: balance.Amount,
	}, nil
}
