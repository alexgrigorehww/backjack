package wallet

type Wallet struct {
	amount int
}

func (w *Wallet) lostMoney(bet int) int {
	return w.amount - bet
}

func (w *Wallet) wonMoney(bet int) int {
	return w.amount + bet
}
