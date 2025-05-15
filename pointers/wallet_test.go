package pointers

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw with funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, Bitcoin(20))
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func ExampleWallet_Balance() {
	wallet := Wallet{balance: Bitcoin(20)}
	balance := wallet.Balance()

	fmt.Printf("balance is %s", balance)

	// Output: balance is 20 BTC
}

func ExampleWallet_Deposit() {
	wallet := Wallet{balance: Bitcoin(20)}
	wallet.Deposit(Bitcoin(10))
	balance := wallet.Balance()

	fmt.Printf("balance is %s", balance)

	// Output: balance is 30 BTC
}

func ExampleWallet_Withdraw() {
	wallet := Wallet{balance: Bitcoin(20)}
	wallet.Withdraw(Bitcoin(10))
	balance := wallet.Balance()

	fmt.Printf("balance is %s", balance)

	// Output: balance is 10 BTC
}
