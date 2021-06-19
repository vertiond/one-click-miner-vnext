package payouts

var _ Payout = &XMRPayout{}

type XMRPayout struct{}

func NewXMRPayout() *XMRPayout {
	return &XMRPayout{}
}

func (p *XMRPayout) GetID() int {
	return 12
}

func (p *XMRPayout) GetName() string {
	return "Monero"
}

func (p *XMRPayout) GetTicker() string {
	return "XMR"
}

func (p *XMRPayout) GetCoingeckoExchange() string {
	return "binance"
}
