package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"wallet_api_service/internal/handler"
	"wallet_api_service/internal/repository"
	"wallet_api_service/internal/service"
)

func main() {
	db, err := sql.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.New(db)
	svc := service.New(repo)
	h := handler.New(svc)

	http.HandleFunc("/api/v1/wallet", h.Process)
	http.HandleFunc("/api/v1/wallets/", h.GetBalance)

	log.Println("started :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
