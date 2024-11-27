package kafka

import (
	"encoding/json"
	"fc-eda/internal/database"
	"fc-eda/internal/usecase/create_account"
	"fc-eda/pkg/kafka"
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type BalanceDTO struct {
	Name    string `json:"Name"`
	Payload struct {
		AccountIDFrom        string  `json:"account_id_from"`
		AccountIDTo          string  `json:"account_id_to"`
		BalanceAccountIDFrom float64 `json:"balance_account_id_from"`
		BalanceAccountIDTo   float64 `json:"balance_account_id_to"`
	} `json:"Payload"`
}

type Consumer struct {
	ConfigMap *ckafka.ConfigMap
	AccountDb *database.AccountDB
}

func NewConsumer(configMap *ckafka.ConfigMap, accountDb *database.AccountDB) *Consumer {
	return &Consumer{
		ConfigMap: configMap,
		AccountDb: accountDb,
	}
}

func (k *Consumer) Consume() {
	updateBalanceAccountUseCase := create_account.NewCreateAccountUseCase(k.AccountDb)

	topics := []string{"balances"}
	kafkaConsumer := kafka.NewConsumer(k.ConfigMap, topics)

	balanceChan := make(chan *ckafka.Message)

	go kafkaConsumer.Consume(balanceChan)
	go func() {
		for msg := range balanceChan {
			var balanceUpdatedInputDTO BalanceDTO
			if err := json.Unmarshal([]byte(msg.Value), &balanceUpdatedInputDTO); err != nil {
				fmt.Printf("Error decoding Kafka message: %v\n", err)
			} else {
				inputFrom := create_account.CreateAccountInputDTO{
					ID:      balanceUpdatedInputDTO.Payload.AccountIDFrom,
					Balance: balanceUpdatedInputDTO.Payload.BalanceAccountIDFrom,
				}
				err = updateBalanceAccountUseCase.Execute(inputFrom)
				if err != nil {
					fmt.Printf("Error on update balance account %s: ,%v\n", inputFrom.ID, err)
				}
				inputTo := create_account.CreateAccountInputDTO{
					ID:      balanceUpdatedInputDTO.Payload.AccountIDTo,
					Balance: balanceUpdatedInputDTO.Payload.BalanceAccountIDTo,
				}
				err = updateBalanceAccountUseCase.Execute(inputTo)
				if err != nil {
					fmt.Printf("Error on update balance account %s: ,%v\n", inputTo.ID, err)
				}
			}
		}
	}()
}
