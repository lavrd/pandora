package config

var (
	// Viper viper config
	Viper = &struct {
		Runtime struct {
			Verbose bool
			Clean   bool
		}
		Dashboard struct {
			Template string
		}
		Mail struct {
			Send     bool
			Token    string
			Endpoint string
			Email    string
			Name     string
			Templates struct {
				Account struct {
					Created  string
					Recovery string
				}
			}
			Subjects struct {
				Account struct {
					Created  string
					Recovery string
				}
			}
		}
		Validator struct {
			Tracker string
			Broker struct {
				Endpoint string
				User     string
				Password string
			}
		}
		Membership struct {
			Tracker struct {
				Endpoint string
			}
		}
		Tracker struct {
			Endpoint string
		}
		Database struct {
			Endpoint string
			User     string
			Password string
			Database string
		}
		Broker struct {
			Endpoint string
			User     string
			Password string
		}
		Secure struct {
			Key  string
			Salt string
			Jwt  string
		}
		Node struct {
			Database struct {
				Endpoint string
				User     string
				Password string
				Database string
			}
			Tracker struct {
				Endpoint string
			}
			Port     int
			FullName string
			Backup struct {
				File string
			}
			Discovery struct {
				Endpoint string
			}
		}
		Discovery struct {
			Port int
		}
		Api struct {
			Port int
		}
	}{}
)
