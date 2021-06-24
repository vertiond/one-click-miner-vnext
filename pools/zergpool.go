package pools

import (
	"fmt"
	"time"

	"github.com/vertiond/verthash-one-click-miner/payouts"
	"github.com/vertiond/verthash-one-click-miner/util"
)

var _ Pool = &Zergpool{}

type Zergpool struct {
	LastFetchedPayout time.Time
	LastPayout        uint64
}

func NewZergpool() *Zergpool {
	return &Zergpool{}
}

func (p *Zergpool) GetPayouts(testnet bool) []payouts.Payout {
	if testnet {
		return []payouts.Payout{
			payouts.NewVTCPayout(),
		}
	}
	return []payouts.Payout{
		payouts.NewDOGEPayout(),
		payouts.NewVTCPayout(),
		payouts.NewBTCPayout(),
		payouts.NewBCHPayout(),
		payouts.NewDASHPayout(),
		payouts.NewDGBPayout(),
		payouts.NewETHPayout(),
		payouts.NewFIROPayout(),
		payouts.NewGRSPayout(),
		payouts.NewLTCPayout(),
		payouts.NewXMRPayout(),
		payouts.NewRVNPayout(),
	}
}

func (p *Zergpool) GetPendingPayout(addr string) uint64 {
	jsonPayload := map[string]interface{}{}
	err := util.GetJson(fmt.Sprintf("http://api.zergpool.com:8080/api/walletEx?address=%s", addr), &jsonPayload)
	if err != nil {
		return 0
	}
	vtc, ok := jsonPayload["unpaid"].(float64)
	if !ok {
		return 0
	}
	vtc *= 100000000
	return uint64(vtc)
}

func (p *Zergpool) GetStratumUrl() string {
	return "stratum+tcp://verthash.mine.zergpool.com:4534"
}

func (p *Zergpool) GetPassword(payoutTicker string) string {
	return fmt.Sprintf("c=%s,mc=VTC", payoutTicker)
}

func (p *Zergpool) GetID() int {
	return 5
}

func (p *Zergpool) GetName() string {
	return "Zergpool.com"
}

func (p *Zergpool) GetFee() float64 {
	return 0.50
}

func (p *Zergpool) OpenBrowserPayoutInfo(addr string) {
	util.OpenBrowser(fmt.Sprintf("https://zergpool.com/?address=%s", addr))
}
