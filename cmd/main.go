package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com.br/cristian.scherer/eda-balance/internal/database"
	register_balance "github.com.br/cristian.scherer/eda-balance/internal/usecase/register_ballance"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	_ "github.com/go-sql-driver/mysql"
)

type PayloadKafka struct {
	AccountIdFrom        string `json:"account_id_from"`
	AccountIdTo          string `json:"account_id_to"`
	BalanceAccountIdFrom int    `json:"balance_account_id_from"`
	BalanceAccountIdTo   int    `json:"balance_account_id_to"`
}

type kafkaMessage struct {
	Name    string       `json:"Name"`
	Payload PayloadKafka `json:"Payload"`
}

func main() {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "wallet",
	})

	err = c.SubscribeTopics([]string{"balances", "^aRegex.*[Tt]opic"}, nil)

	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "balance", "balance", "mysqlbalance", "3306", "balance"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	run := true

	for run {
		msg, err := c.ReadMessage(time.Second)
		if err == nil {
			fmt.Println(string(msg.Value))

			objKafka := kafkaMessage{}

			err := json.Unmarshal(msg.Value, &objKafka)
			if err != nil {
				panic(err)
			}

			balanceDb := database.NewBalanceDB(db)
			uc := register_balance.NewRegisterBallanceUseCase(*balanceDb)

			//Atualiza conta para
			newRecord := &register_balance.CreateBallanceInputDTO{
				Account: objKafka.Payload.AccountIdTo,
				Amount:  float64(objKafka.Payload.BalanceAccountIdTo),
			}
			_, err = uc.Executa(*newRecord)

			if err == nil {
				//Atualiza conta de
				newRecord = &register_balance.CreateBallanceInputDTO{
					Account: objKafka.Payload.AccountIdFrom,
					Amount:  float64(objKafka.Payload.BalanceAccountIdFrom),
				}
				_, err = uc.Executa(*newRecord)

				if err != nil {
					fmt.Printf("Não foi possível atualizar o valor: %v\n", err.Error())
				}
			} else {
				fmt.Printf("Não foi possível atualizar o valor: %v\n", err.Error())
			}
		} else if !err.(kafka.Error).IsTimeout() {

			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	c.Close()
}
