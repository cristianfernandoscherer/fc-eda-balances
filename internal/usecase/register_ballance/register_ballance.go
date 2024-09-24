package register_balance

import (
	"github.com.br/cristian.scherer/eda-balance/internal/database"
	"github.com.br/cristian.scherer/eda-balance/internal/entity"
)

type CreateBallanceInputDTO struct {
	Account string  `json:"account"`
	Amount  float64 `json:"amount"`
}

type CreateBallanceOutputDTO struct {
	ID     string
	Amount float64
}

type RegisterBallanceUseCase struct {
	BallanceDB database.BalanceDB
}

func NewRegisterBallanceUseCase(ballanceDB database.BalanceDB) *RegisterBallanceUseCase {
	return &RegisterBallanceUseCase{
		BallanceDB: ballanceDB,
	}
}

func (registerBallance *RegisterBallanceUseCase) Executa(input CreateBallanceInputDTO) (*CreateBallanceOutputDTO, error) {

	err, balance := registerBallance.BallanceDB.FindByAccountID(input.Account)
	if err != nil {
		return nil, err
	}

	if balance.ID == "" {
		balance, _ := entity.NewBalance(
			input.Account,
			input.Amount,
		)

		err = registerBallance.BallanceDB.Save(*balance)
		if err != nil {
			return nil, err
		}

		return &CreateBallanceOutputDTO{
			ID:     balance.ID,
			Amount: balance.Amount,
		}, nil
	}

	err = registerBallance.BallanceDB.Update(input.Amount, balance.ID)
	if err != nil {
		return nil, err
	}

	return &CreateBallanceOutputDTO{
		ID:     balance.ID,
		Amount: input.Amount,
	}, nil
}
