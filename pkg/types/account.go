package types

type (
	Session struct {
		Token string `json:"token"`
	}

	Account struct {
		Key       string         `json:"_key"`
		PublicKey string         `json:"public_key"`
		Secure    *AccountSecure `json:"secure,omitempty"`
		Meta      *AccountMeta   `json:"meta"`
	}

	AccountMeta struct {
		Email    string `json:"email"`
		FullName string `json:"full_name"`
	}

	AccountSecure struct {
		PrivateKey []byte `json:"private_key"`
	}
)

// Public returns public account info
func (acc *Account) Public() *Account {
	acc.Secure = nil
	return acc
}
