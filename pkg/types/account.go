package types

// Session
type Session struct {
	Token string `json:"token"`
}

// Account
type Account struct {
	Email     string         `json:"email"`
	Type      int            `json:"type"`
	PublicKey string         `json:"public_key"`
	Secure    *AccountSecure `json:"secure,omitempty"`
}

type AccountSecure struct {
	Password   string `json:"password,omitempty"`
	PrivateKey string `json:"private_key"`
}

// Public returns public account info
func (acc *Account) Public() *Account {
	acc.Secure = nil
	return acc
}

// AccountRecovery
type AccountRecovery struct {
	Password string `json:"password"`
}
