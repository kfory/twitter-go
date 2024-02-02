package models

type Secret struct {
	Host     string `json:"host"` //ALT + 96 para battis
	Username string `json:"username"`
	Password string `json:"password"`
	JWTsign  string `json:"jwtsign"`
	Database string `json:"database"`
}
