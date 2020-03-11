package requests

type Token struct {
	Login string
}

func (r *Token) Params() (map[string]string, map[string]interface{}) {
	params := map[string]string{}
	logParams := map[string]interface{}{}
	if r == nil {
		return params, logParams
	}

	logParams["operation"] = "get_token"
	params["operation"] = "get_token"

	if r.Login != "" {
		logParams["login"] = r.Login
		params["login"] = r.Login
	}

	return params, logParams
}
