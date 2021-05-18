package payouts

var _ Payout = &FIROPayout{}

type FIROPayout struct{}

func NewFIROPayout() *FIROPayout {
	return &FIROPayout{}
}

func (p *FIROPayout) GetID() int {
	return 11
}

func (p *FIROPayout) GetName() string {
	return "Zcoin"
}

func (p *FIROPayout) GetTicker() string {
	return "FIRO"
}

func (p *FIROPayout) GetCoingeckoExchange() string {
	return "bittrex"
}
