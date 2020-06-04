package wallet

import (
	"testing"
)

func TestWallet_LostMoneyValue(t *testing.T) {
	if (&Wallet{amount: 100}).lostMoney(50) != 50 {
		t.Error("Remaining amount should be 50")
	}
}
