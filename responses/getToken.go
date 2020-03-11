package responses

type Token struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    Data   `json:"data,omitempty"`
}

type Data struct {
	Token                string `json:"token"`
	Modulus              string `json:"modulus"`
	Exponent             string `json:"exponent"`
	RsaPublicKeyPkcs1Pem string `json:"rsa_public_key_pkcs1_pem"`
}
