package wallet

type Wallet struct {
	amount int
}

type SerializableWallet struct {
	Amount int
}

func (w *Wallet) SetAmount(amount int) {
	w.amount = amount
}

func (w *Wallet) GetAmount() int {
	return w.amount
}

func (w *Wallet) LostMoney(bet int) int {
	w.amount -= bet
	return w.amount
}

func (w *Wallet) WonMoney(bet int) int {
	w.amount += bet
	return w.amount
}

func (w *Wallet) GetSerializable() *SerializableWallet {
	serializableWallet := SerializableWallet{
		Amount: w.GetAmount(),
	}
	return &serializableWallet
}

func (w *SerializableWallet) Deserialize() *Wallet {
	serializableWallet := Wallet{
		amount: w.Amount,
	}
	return &serializableWallet
}
