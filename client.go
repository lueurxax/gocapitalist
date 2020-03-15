package gocapitalist

import (
	"gocapitalist/addPaymentNotification"
	"gocapitalist/createAccount"
	"gocapitalist/defaults"
	"gocapitalist/getAccounts"
	"gocapitalist/getBatchInfo"
	"gocapitalist/getCashinRequisites"
	"gocapitalist/getDocumentHistoryExt"
	"gocapitalist/importBatchAdvanced"
	"gocapitalist/internal"
	"gocapitalist/isVerifiedAccount"
	"gocapitalist/сurrencyRates"
)

type AClient struct {
	bc *internal.BaseClient
	сurrencyRates.Currency
	getBatchInfo.BatchInfo
	importBatchAdvanced.ImportBatchAdvanced
	getDocumentHistoryExt.GetDocumentHistoryExt
	addPaymentNotification.AddPaymentNotification
	isVerifiedAccount.IsVerifiedAccount
	getAccounts.GetAccounts
	createAccount.CreateAccount
	getCashinRequisites.GetCashinRequisites
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
		Currency:               сurrencyRates.Currency{BaseClient: bc},
		BatchInfo:              getBatchInfo.BatchInfo{BaseClient: bc},
		ImportBatchAdvanced:    importBatchAdvanced.ImportBatchAdvanced{BaseClient: bc},
		GetDocumentHistoryExt:  getDocumentHistoryExt.GetDocumentHistoryExt{BaseClient: bc},
		AddPaymentNotification: addPaymentNotification.AddPaymentNotification{BaseClient: bc},
		IsVerifiedAccount:      isVerifiedAccount.IsVerifiedAccount{BaseClient: bc},
		GetAccounts:            getAccounts.GetAccounts{BaseClient: bc},
		CreateAccount:          createAccount.CreateAccount{BaseClient: bc},
		GetCashinRequisites:    getCashinRequisites.GetCashinRequisites{BaseClient: bc},
		bc:                     bc}
}

func (ac *AClient) SetAuth(login, password string) error {
	return ac.bc.SetAuth(login, password)
}
