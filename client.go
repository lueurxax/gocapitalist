package gocapitalist

import (
	"gocapitalist/defaults"
	"gocapitalist/getBatchInfo"
	"gocapitalist/importBatchAdvanced"
	"gocapitalist/internal"
	"gocapitalist/сurrencyRates"
)

type AClient struct {
	bc *internal.BaseClient
	сurrencyRates.Currency
	getBatchInfo.BatchInfo
	importBatchAdvanced.ImportBatchAdvanced
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
		Currency:            сurrencyRates.Currency{bc},
		BatchInfo:           getBatchInfo.BatchInfo{bc},
		ImportBatchAdvanced: importBatchAdvanced.ImportBatchAdvanced{bc},
		bc:                  bc}
}

func (ac *AClient) SetAuth(login, password string) error {
	return ac.bc.SetAuth(login, password)
}
