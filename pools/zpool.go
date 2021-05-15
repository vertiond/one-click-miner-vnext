package pools

import (
	"fmt"
	"time"

	"github.com/vertiond/verthash-one-click-miner/util"
)

var _ Pool = &zpool{}

type zpool struct {
	LastFetchedPayout time.Time
	LastPayout        uint64
}

func Newzpool() *zpool {
	return &zpool{}
}

func (p *zpool) GetPendingPayout(addr string) uint64 {
	jsonPayload := map[string]interface{}{}
	err := util.GetJson(fmt.Sprintf("https://zpool.ca/api/wallet?address=%s", addr), &jsonPayload)
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

func (p *zpool) GetStratumUrl() string {
	return "stratum+tcp://verthash.mine.zpool.ca:6144"
}

func (p *zpool) GetPassword() string {
	return "c=VTC,zap=VTC"
}

func (p *zpool) GetID() int {
	return 7
}

func (p *zpool) GetName() string {
	return "zpool.ca"
}

func (p *zpool) GetFee() float64 {
	return 0.50
}

func (p *zpool) OpenBrowserPayoutInfo(addr string) {
	util.OpenBrowser(fmt.Sprintf("https://zpool.ca/wallet/%s", addr))
}
