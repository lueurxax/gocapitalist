package importBatchAdvanced

import (
	"gocapitalist/internal"
	"gocapitalist/requests"
	"gocapitalist/responses"
)

type ImportBatchAdvanced struct {
	*internal.BaseClient
}

// https://www.admitad.com/en/developers/doc/api_ru/methods/public/currency_exchange_rate/
func (b *ImportBatchAdvanced) Import(request requests.ImportBatchAdvanced) (*responses.ImportBatchAdvanced, error) {
	data, errResponse := new(responses.ImportBatchAdvanced), new(responses.ErrorResponse)

	httpParams, logParams, err := request.Params()
	if err != nil {
		return nil, err
	}
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
