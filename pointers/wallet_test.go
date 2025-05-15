package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		got := wallet.Balance()

		want := Bitcoin(10)

		if got != want {
			t.Errorf("want %s but got %s", want, got)
		}
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{Bitcoin(30)}
		wallet.Withdraw(Bitcoin(10))
		got := wallet.Balance()

		want := Bitcoin(20)

		if got != want {
			t.Errorf("want %s but got %s", want, got)
		}
	})
}
