package responses

type DocumentHistory struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
	Data    DataDoc `json:"data"`
}
type Pages struct {
	ItemCount        int `json:"itemCount"`
	PageCount        int `json:"pageCount"`
	PageSize         int `json:"pageSize"`
	CurrentPage      int `json:"currentPage"`
	CurrentPageCount int `json:"currentPageCount"`
}
type History struct {
	Number         int         `json:"number"`
	ID             int         `json:"id"`
	Batch          interface{} `json:"batch"`
	Timestamp      int         `json:"timestamp"`
	State          string      `json:"state"`
	StateLocalized string      `json:"stateLocalized"`
	Channel        string      `json:"channel"`
	Correspondent  string      `json:"correspondent"`
	Comment        string      `json:"comment"`
	Account        string      `json:"account"`
	Amount         int         `json:"amount"`
	Tax            interface{} `json:"tax"`
	Currency       string      `json:"currency"`
	DestAmount     interface{} `json:"destAmount"`
	DestCurrency   interface{} `json:"destCurrency"`
	Description    string      `json:"description"`
	Outgoing       bool        `json:"outgoing"`
	SelfExchange   bool        `json:"selfExchange"`
}
type DataDoc struct {
	Pages   Pages     `json:"pages"`
	History []History `json:"history"`
}
