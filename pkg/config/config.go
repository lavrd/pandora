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
				Credentials string
			}
			Subjects struct {
				Credentials string
			}
		}
		Dashboard struct {
			Template string
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
		Membership struct {
			Endpoint string
		}
		Master struct {
			Endpoint string
		}
		Node struct {
			Endpoint  string
			Email     string
			FullName  string `mapstructure:"full_name"`
			SecretKey string `mapstructure:"secret_key"`
		}
		Discovery struct {
			Endpoint string
		}
	}{}
)
