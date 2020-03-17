package gocapitalist

import (
	"gocapitalist/API/addPaymentNotification"
	"gocapitalist/API/createAccount"
	"gocapitalist/API/getAccounts"
	"gocapitalist/API/getBatchInfo"
	"gocapitalist/API/getCashinRequisites"
	"gocapitalist/API/getDocumentHistoryExt"
	"gocapitalist/API/importBatchAdvanced"
	"gocapitalist/API/isVerifiedAccount"
	"gocapitalist/API/passwordRecovery"
	"gocapitalist/API/passwordRecoveryGenerateCode"
	"gocapitalist/API/profileChangeEmail"
	"gocapitalist/API/profileGetVerificationCode"
	"gocapitalist/API/registerEmailConfirm"
	"gocapitalist/API/registerInvitee"
	"gocapitalist/defaults"
	"gocapitalist/internal"
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
	passwordRecovery.PasswordRecovery
	passwordRecoveryGenerateCode.PasswordRecoveryGenerateCode
	profileChangeEmail.ProfileChangeEmail
	profileGetVerificationCode.ProfileGetVerificationCode
	registerEmailConfirm.RegisterEmailConfirm
	registerInvitee.RegisterInvitee
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
		Currency:                     сurrencyRates.Currency{BaseClient: bc},
		BatchInfo:                    getBatchInfo.BatchInfo{BaseClient: bc},
		ImportBatchAdvanced:          importBatchAdvanced.ImportBatchAdvanced{BaseClient: bc},
		GetDocumentHistoryExt:        getDocumentHistoryExt.GetDocumentHistoryExt{BaseClient: bc},
		AddPaymentNotification:       addPaymentNotification.AddPaymentNotification{BaseClient: bc},
		IsVerifiedAccount:            isVerifiedAccount.IsVerifiedAccount{BaseClient: bc},
		GetAccounts:                  getAccounts.GetAccounts{BaseClient: bc},
		CreateAccount:                createAccount.CreateAccount{BaseClient: bc},
		GetCashinRequisites:          getCashinRequisites.GetCashinRequisites{BaseClient: bc},
		PasswordRecovery:             passwordRecovery.PasswordRecovery{BaseClient: bc},
		PasswordRecoveryGenerateCode: passwordRecoveryGenerateCode.PasswordRecoveryGenerateCode{BaseClient: bc},
		ProfileChangeEmail:           profileChangeEmail.ProfileChangeEmail{BaseClient: bc},
		ProfileGetVerificationCode:   profileGetVerificationCode.ProfileGetVerificationCode{BaseClient: bc},
		RegisterEmailConfirm:         registerEmailConfirm.RegisterEmailConfirm{BaseClient: bc},
		RegisterInvitee:              registerInvitee.RegisterInvitee{BaseClient: bc},
		bc:                           bc}
}

func (ac *AClient) SetAuth(login, password string) error {
	return ac.bc.SetAuth(login, password)
}
