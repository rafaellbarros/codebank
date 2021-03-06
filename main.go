package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/rafaellbarros/codebank/domain"
	"github.com/rafaellbarros/codebank/infrastructure/repository"
	"github.com/rafaellbarros/codebank/usercase"
	"log"
)

func init() {
	fmt.Println("Hello")
}

func main() {
	db := setupDb()
	defer db.Close()

	cc := domain.NewCreditCard()
	cc.Number = "1234"
	cc.Name = "Rafael Barros"
	cc.ExpirationYear = 2021
	cc.ExpirationMonth = 7
	cc.CVV = 123
	cc.Limit = 1000
	cc.Balance = 0

	repo := repository.NewTransactionRepositoryDb(db)
	err := repo.CreateCreditCard(*cc)
	if err != nil {
		fmt.Println(err)
	}

}

func setupTransactionUseCase(db *sql.DB) usercase.UseCaseTransaction {
	transactionRepository := repository.NewTransactionRepositoryDb(db)
	useCase := usercase.NewUseCaseTransaction(transactionRepository)
	return useCase
}


func setupDb() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"db",
		"5432",
		"postgres",
		"root",
		"codebank",
	)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("error connection to database")
	}
	return db
}
