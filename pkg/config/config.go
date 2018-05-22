package config

var (
	Viper = &struct {
		Runtime struct {
			Verbose bool
			Clean   bool
		}
		Membership struct {
			Mail struct {
				Send     bool
				Token    string
				Endpoint string
				Email    string
				Name     string
			}
			Database struct {
				Endpoint string
				User     string
				Password string
				Database string
			}
			Endpoint string
		}
		Master struct {
			Database struct {
				FilePath string
			}
			Endpoint string
		}
		Node struct {
			Database struct {
				FilePath string
			}
			Dashboard struct {
				Template string
			}
			Meta struct {
				Email     string
				Name      string
				SecretKey string `mapstructure:"secret_key"`
			}
			Endpoint string
		}
		Discovery struct {
			Database struct {
				FilePath string
			}
			Broker struct {
				Endpoint string
				User     string
				Password string
			}
			Endpoint string
		}
	}{}
)
