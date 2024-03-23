package main

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}

	assertError := func(t testing.TB, err error) {
		t.Helper()

		if err != nil {
			t.Errorf("%v", err)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(25))

		assertBalance(t, wallet, Bitcoin(25))

	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{20}
		err := wallet.Withdraw(9)

		assertError(t, err)
		assertBalance(t, wallet, Bitcoin(11))
	})
}
