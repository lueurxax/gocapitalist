package requests

type PasswordRecovery struct {
	EncryptedPassword string
	Code              string
}

func (r *PasswordRecovery) Params() (map[string]string, map[string]interface{}) {
	params := map[string]string{}
	logParams := map[string]interface{}{}
	if r == nil {
		return params, logParams
	}

	logParams["operation"] = "get_token"
	params["operation"] = "get_token"

	if r.EncryptedPassword != "" {
		logParams["encrypted_password"] = r.EncryptedPassword
		params["encrypted_password"] = r.EncryptedPassword
	}

	if r.Code != "" {
		logParams["code"] = r.Code
		params["code"] = r.Code
	}

	return params, logParams
}
