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
				Templates struct {
					Credentials string
				}
				Subjects struct {
					Credentials string
				}
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
			Endpoint string
		}
		Node struct {
			Dashboard struct {
				Template string
			}
			Meta struct {
				Email     string
				Name      string
				SecretKey string `mapstructure:"secret_key"`
			}
			Port int
		}
		Discovery struct {
			Broker struct {
				Endpoint string
				User     string
				Password string
			}
			Endpoint string
		}
	}{}
)
