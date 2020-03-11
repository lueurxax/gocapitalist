package responses

type Currency struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Data    DataCurrencies `json:"data,omitempty"`
}
type Buy struct {
	Target    string  `json:"target"`
	Amount    float64 `json:"amount"`
	AmountCur string  `json:"amountCur"`
}
type Sell struct {
	Target    string  `json:"target"`
	Amount    float64 `json:"amount"`
	AmountCur string  `json:"amountCur"`
}
type UahSell struct {
	TargetAmount int    `json:"targetAmount"`
	Target       string `json:"target"`
	Amount       string `json:"amount"`
	AmountCur    string `json:"amountCur"`
}
type Rates struct {
	Buy     []Buy     `json:"buy"`
	Sell    []Sell    `json:"sell"`
	UahSell []UahSell `json:"uahSell"`
}
type DataCurrencies struct {
	Rates Rates `json:"rates"`
}
