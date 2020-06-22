package wallet_test

import (
	"testing"
)

func TestWallet_LostMoneyValue(t *testing.T) {
	if (&Wallet{amount: 100}).LostMoney(50) != 50 {
		t.Error("Remaining amount should be 50")
	}
}

func TestWallet_SetAmount(t *testing.T) {
	w := &Wallet{}
	w.SetAmount(500)
	if w.GetAmount() != 500 {
		t.Error("Amount on set should be 500")
	}
}

func TestWallet_GetAmount(t *testing.T) {
	//w := &Wallet{}
	w := new(Wallet)
	w.SetAmount(500)
	if w.GetAmount() != 500 {
		t.Error("Amount on get should be 500")
	}
}
