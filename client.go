package gocapitalist

import (
	"gocapitalist/defaults"
	"gocapitalist/internal"
	"gocapitalist/сurrencyRates"
)

type AClient struct {
	bc *internal.BaseClient
	сurrencyRates.Currency
}

func New(url string, logger internal.Logger, mc internal.MetricsCollector) *AClient {
	if logger == nil {
		logger = defaults.NewLogrus()
	}
	if mc == nil {
		mc = &defaults.EmptyMetrics{}
	}
	bc := internal.NewBaseClient(url, logger, mc)
	return &AClient{
		Currency: сurrencyRates.Currency{bc},
		bc:       bc}
}

func (ac *AClient) SetAuth(login, password string) error {
	return ac.bc.SetAuth(login, password)
}
