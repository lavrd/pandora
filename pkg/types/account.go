package types

// Session
type Session struct {
	Token string `json:"token"`
}

// Account
type Account struct {
	// Key used for identify document in arangodb collection
	Key       string         `json:"_key,omitempty"`
	Email     string         `json:"email"`
	Type      int            `json:"type"`
	PublicKey string         `json:"public_key"`
	Secure    *AccountSecure `json:"secure,omitempty"`
}

// Public returns public account info
func (acc *Account) Public() *Account {
	acc.Secure = nil
	return acc
}

// AccountSecure
type AccountSecure struct {
	Password   string `json:"password"`
	PrivateKey string `json:"private_key"`
}

// AccountRecovery
type AccountRecovery struct {
	Password string `json:"password"`
}
