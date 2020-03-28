package github

import (
	"github.com/hugmouse/gocapitalist/defaults"
	"github.com/hugmouse/gocapitalist/internal"
	"github.com/hugmouse/gocapitalist/сurrencyRates"
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
		Currency: сurrencyRates.Currency{BaseClient: bc},
		bc:       bc}
}

func (ac *AClient) SetAuth(login, password string) error {
	return ac.bc.SetAuth(login, password)
}
