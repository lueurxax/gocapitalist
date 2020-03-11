package responses

import "fmt"

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("api error: %s", e.Message)
}

func (e *ErrorResponse) ErrLogParams(params map[string]interface{}) map[string]interface{} {
	errParams := params
	errParams["error_code"] = e.Code
	return errParams
}
