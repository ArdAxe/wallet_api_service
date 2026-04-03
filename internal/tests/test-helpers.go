package tests

import (
	"database/sql"
	"github.com/google/uuid"
	"log"
	"os"
	"testing"
	"wallet_api_service/internal/repository"

	_ "github.com/lib/pq"
)

func SetupTestRepo(t *testing.T) *repository.WalletRepository {
	dbURL := os.Getenv("TEST_DB_URL")
	if dbURL == "" {
		dbURL = "postgres://postgres:postgres@localhost:5432/wallet_test?sslmode=disable"
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		t.Fatal(err)
	}

	// очищаем таблицу перед тестом
	_, err = db.Exec(`TRUNCATE wallets`)
	if err != nil {
		t.Fatal(err)
	}

	return repository.New(db)
}

func CreateWallet(repo *repository.WalletRepository, balance int64) uuid.UUID {
	id := uuid.New()

	_, err := repo.Db.Exec(`
        INSERT INTO wallets (id, balance)
        VALUES ($1, $2)
    `, id, balance)

	if err != nil {
		log.Fatal(err)
	}

	return id
}
