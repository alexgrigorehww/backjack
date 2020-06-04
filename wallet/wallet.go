package wallet

type Wallet struct {
	amount int
}

func lostMoney(w *Wallet, bet int) int {
	return w.amount - bet
}

func wonMoney(w *Wallet, bet int) int {
	return w.amount - bet
}
