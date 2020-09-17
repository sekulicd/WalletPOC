package domain

type Wallet struct {
	Seed string
}

func (w *Wallet) Validate() {
	//do some validation
}
