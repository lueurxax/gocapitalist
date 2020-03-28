package сurrencyRates

import (
	"github.com/hugmouse/gocapitalist/internal"
	"github.com/hugmouse/gocapitalist/responses"
)

type Currency struct {
	*internal.BaseClient
}

func (c *Currency) Params() (map[string]string, map[string]interface{}) {
	params := map[string]string{}
	logParams := map[string]interface{}{}
	if c == nil {
		return params, logParams
	}

	logParams["operation"] = "currency_rates"
	params["operation"] = "currency_rates"

	if c.Auth.Login != "" {
		logParams["login"] = c.Auth.Login
		params["login"] = c.Auth.Login
	}

	if c.Auth.Token != "" {
		logParams["token"] = c.Auth.Token
		params["token"] = c.Auth.Token
	}

	if len(c.Auth.EncryptedPassword) > 0 {
		logParams["encrypted_password"] = c.Auth.EncryptedPassword
		params["encrypted_password"] = c.Auth.EncryptedPassword
	}

	return params, logParams
}

// https://www.admitad.com/en/developers/doc/api_ru/methods/public/currency_exchange_rate/
func (c *Currency) Get() (*responses.Currency, error) {
	data, errResponse := new(responses.Currency), new(responses.ErrorResponse)

	httpParams, logParams := c.Params()

	c.Logger.Debug("make request", "currency_rates", logParams, nil)

	resp, err := c.R().
		SetFormData(httpParams).
		EnableTrace().
		SetResult(data).
		SetError(errResponse).
		SetHeader("x-response-format", "json").
		Post("/")

	if err != nil {
		c.Logger.Error("http error", "currency_rates", logParams, err)
		return nil, err
	}

	c.Metrics.Collect("currency_rates", resp.StatusCode(), errResponse.Code, resp.Time())

	if resp.Error() != nil {
		c.Logger.Error("app error", "currency_rates", errResponse.ErrLogParams(logParams), errResponse)
		return nil, errResponse
	}

	c.Logger.Debug("success request", "currency_rates", logParams, nil)

	return data, nil

}
