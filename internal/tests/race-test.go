package tests

import (
	"context"
	"sync"
	"testing"
)

func TestConcurrentWithdraw(t *testing.T) {
	repo := SetupTestRepo(t)

	id := CreateWallet(repo, 1000)

	wg := sync.WaitGroup{}
	n := 100

	for i := 0; i < n; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			_, _ = repo.Withdraw(context.Background(), id, 10)
		}()
	}

	wg.Wait()

	balance, err := repo.GetBalance(context.Background(), id)
	if err != nil {
		t.Fatal(err)
	}

	if balance != 0 {
		t.Fatalf("expected 0, got %d", balance)
	}
}
