package main

import (
	"database/sql"
	"fmt"

	"github.com.br/cristian.scherer/eda-balance/internal/database"
	"github.com.br/cristian.scherer/eda-balance/internal/usecase/find_balance"
	"github.com.br/cristian.scherer/eda-balance/internal/web"
	"github.com.br/cristian.scherer/eda-balance/internal/web/webserver"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "balance", "balance", "mysqlbalance", "3306", "balance"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	balanceDb := database.NewBalanceDB(db)
	ucFind := find_balance.NewFindBalanceUseCase(*balanceDb)
	clientHandler := web.NewWebAccountHandler(*ucFind)

	webserver := webserver.NewWebServer(":3000")
	webserver.AddHandler("/balances/{account_id}", clientHandler.FindBalance)
	fmt.Println("Server is running")
	webserver.Start()
}
