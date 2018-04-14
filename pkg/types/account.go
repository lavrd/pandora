package types

// Session
type Session struct {
	Token string `json:"token"`
}

// Account
type Account struct {
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

// Public get only public account info
func (acc *Account) Public() *Account {
	acc.Password = ""
	return acc
}

// AccountRecovery
type AccountRecovery struct {
	Password string `json:"password"`
}
