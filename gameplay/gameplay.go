package gameplay

type Gameplay interface {
	SetBet(bet int) (err error)
	Deal() (err error)
	Hit() (err error)
	Stand() (err error)
}
