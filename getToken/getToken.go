package getToken

import (
	"gocapitalist/internal"
)

type GetToken struct {
	*internal.BaseClient
}

// https://capitalist.net/developers/api/page/get_token
// func (r GetToken) Get(request requests.Token) (*responses.Token, error) {
// 	data, errResponse := new(responses.Token), new(responses.ErrorResponse)
//
// 	httpParams, logParams := request.Params()
//
// 	r.Logger.Debug("make request", "get_token", logParams, nil)
//
// 	resp, err := r.R().
// 		SetFormData(httpParams).
// 		EnableTrace().
// 		SetResult(data).
// 		SetError(errResponse).
// 		SetHeader("x-response-format", "json").
// 		Post("/")
//
// 	if err != nil {
// 		r.Logger.Error("http error", "get_token", logParams, err)
// 		return nil, err
// 	}
//
// 	r.Metrics.Collect("get_token", resp.StatusCode(), errResponse.Code, resp.Time())
//
// 	if resp.Error() != nil {
// 		r.Logger.Error("app error", "get_token", errResponse.ErrLogParams(logParams), errResponse)
// 		return nil, errResponse
// 	}
//
// 	if data.Code != 0 {
// 		r.Logger.Error("api error", "get_token", errResponse.ErrLogParams(logParams), errResponse)
// 		return nil, errors.New(data.Message)
// 	}
//
// 	r.Logger.Debug("success request", "get_token", logParams, nil)
//
// 	return data, nil
// }
