package domain

import "time"

type TokenData struct {
	access  Token
	refresh Token
}

type Token struct {
	value   string
	expires time.Time
}
