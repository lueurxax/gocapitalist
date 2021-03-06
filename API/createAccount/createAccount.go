package createAccount

import (
	"github.com/lueurxax/gocapitalist/internal"
	"github.com/lueurxax/gocapitalist/requests"
	"github.com/lueurxax/gocapitalist/responses"
)

type CreateAccount struct {
	*internal.BaseClient
}

// https://capitalist.net/developers/api/page/create_account
func (b *CreateAccount) New(request requests.CreateAccount) (*responses.CreateAccount, error) {
	data, errResponse := new(responses.CreateAccount), new(responses.ErrorResponse)

	httpParams, logParams := request.Params()
	for k, v := range b.Auth.ParamsForAuth {
		httpParams[k] = v
	}

	b.Logger.Debug("make request", httpParams["operation"], logParams, nil)

	resp, err := b.R().
		SetFormData(httpParams).
		EnableTrace().
		SetResult(data).
		SetError(errResponse).
		SetHeader("x-response-format", "json").
		Post("/")

	if err != nil {
		b.Logger.Error("http error", httpParams["operation"], logParams, err)
		return nil, err
	}

	b.Metrics.Collect(httpParams["operation"], resp.StatusCode(), errResponse.Code, resp.Time())

	if resp.Error() != nil {
		b.Logger.Error("app error", httpParams["operation"], errResponse.ErrLogParams(logParams), errResponse)
		return nil, errResponse
	}

	b.Logger.Debug("success request", httpParams["operation"], logParams, nil)

	return data, nil

}
