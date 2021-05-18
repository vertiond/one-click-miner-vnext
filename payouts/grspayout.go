package payouts

var _ Payout = &GRSPayout{}

type GRSPayout struct{}

func NewGRSPayout() *GRSPayout {
	return &GRSPayout{}
}

func (p *GRSPayout) GetID() int {
	return 10
}

func (p *GRSPayout) GetName() string {
	return "Groestlcoin"
}

func (p *GRSPayout) GetTicker() string {
	return "GRS"
}

func (p *GRSPayout) GetCoingeckoExchange() string {
	return "bittrex"
}
