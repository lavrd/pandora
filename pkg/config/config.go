package config

var (
	// Viper viper config
	Viper = &struct {
		Runtime struct {
			Verbose bool
			Clean   bool
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
			PSSMessage string
			JWTKey     string
		}
		Node struct {
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
