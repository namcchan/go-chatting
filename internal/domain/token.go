package domain

type Tokens struct {
	Access  TokenPayload `json:"access"`
	Refresh TokenPayload `json:"refresh"`
}

type TokenPayload struct {
	Value   string `json:"value"`
	Expires int64  `json:"expires"`
}
