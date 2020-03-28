package gocapitalist

import (
	"github.com/lueurxax/gocapitalist/API/addPaymentNotification"
	"github.com/lueurxax/gocapitalist/API/createAccount"
	"github.com/lueurxax/gocapitalist/API/getAccounts"
	"github.com/lueurxax/gocapitalist/API/getBatchInfo"
	"github.com/lueurxax/gocapitalist/API/getCashinRequisites"
	"github.com/lueurxax/gocapitalist/API/getDocumentHistoryExt"
	"github.com/lueurxax/gocapitalist/API/importBatchAdvanced"
	"github.com/lueurxax/gocapitalist/API/isVerifiedAccount"
	"github.com/lueurxax/gocapitalist/API/passwordRecovery"
	"github.com/lueurxax/gocapitalist/API/passwordRecoveryGenerateCode"
	"github.com/lueurxax/gocapitalist/API/profileChangeEmail"
	"github.com/lueurxax/gocapitalist/API/profileGetVerificationCode"
	"github.com/lueurxax/gocapitalist/API/registerEmailConfirm"
	"github.com/lueurxax/gocapitalist/API/registerInvitee"
	"github.com/lueurxax/gocapitalist/API/сurrencyRates"
	"github.com/lueurxax/gocapitalist/defaults"
	"github.com/lueurxax/gocapitalist/internal"
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
