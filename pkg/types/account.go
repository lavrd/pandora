package types

// Session
type Session struct {
	Token string `json:"token"`
}

// Account
type Account struct {
	PublicKey string         `json:"public_key"`
	Secure    *AccountSecure `json:"secure,omitempty"`
	Meta      *AccountMeta   `json:"meta"`
}

// AccountMeta
type AccountMeta struct {
	Email    string `json:"email"`
	FullName string `json:"full_name"`
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
