package requests

import (
	"gocapitalist/enums"
	"time"
)

type DocumentHistory struct {
	PeriodFrom      time.Time
	PeriodTo        time.Time
	DocumentState   enums.DocumentState
	Limit           int
	Page            int
	Account         string
	ExternalAccount int
}
