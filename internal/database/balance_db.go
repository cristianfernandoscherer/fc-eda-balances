package database

import (
	"database/sql"

	"github.com.br/cristian.scherer/eda-balance/internal/entity"
)

type BalanceDB struct {
	DB *sql.DB
}

func NewBalanceDB(db *sql.DB) *BalanceDB {
	return &BalanceDB{
		DB: db,
	}
}

func (a *BalanceDB) Save(ballance entity.Balance) error {
	stmt, err := a.DB.Prepare("INSERT INTO balances (id, account, amount) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(ballance.ID, ballance.Account, ballance.Amount)
	if err != nil {
		return err
	}
	return nil
}

func (a *BalanceDB) FindByAccountID(AccountID string) (error, *entity.Balance) {

	var balance entity.Balance

	stmt, err := a.DB.Prepare("SELECT id, amount FROM balances WHERE account = ?")
	if err != nil {
		return err, &balance
	}
	defer stmt.Close()

	row := stmt.QueryRow(AccountID)

	err = row.Scan(&balance.ID, &balance.Amount)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &balance
		}
		return err, &balance
	}

	return nil, &balance
}

func (a *BalanceDB) Update(Amount float64, ID string) error {
	stmt, err := a.DB.Prepare("UPDATE balances SET amount = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(Amount, ID)
	if err != nil {
		return err
	}
	return nil
}
