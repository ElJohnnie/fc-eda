package main

import (
	"database/sql"
	"fc-eda/web"
	"fmt"
	"log"
	"net/http"

	"fc-eda/internal/database"
	"fc-eda/internal/kafka"
	"fc-eda/internal/usecase/get_account"

	_ "github.com/go-sql-driver/mysql"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "balances-mysql", "3306", "balance"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	accountDb := database.NewAccountDB(db)

	configMap := ckafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "wallet",
	}

	kafkaConsumer := kafka.NewConsumer(&configMap, accountDb)
	kafkaConsumer.Consume()

	getAccountUseCase := get_account.NewGetAccountOutputDTO(accountDb)

	accountHandler := web.NewWebAccountHandler(*getAccountUseCase)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/balances/{id}", accountHandler.GetAccount)

	fmt.Println("Balance is running on port :3003")
	log.Fatal(http.ListenAndServe(":3003", r))
}
