package conf

var (
	// Conf describe conf struct
	Conf = &struct {
		Runtime struct {
			Verbose bool
			Clean   bool
		}
		TLS struct {
			Cert string
			Key  string
		}
		SendGrid struct {
			Active bool
			Token  string
		} `mapstructure:"sendgrid"`
		Arangodb struct {
			Endpoint string
			User     string
			Password string
			Database string
		}
		Membership struct {
			Endpoint string
		}
		Master struct {
			Endpoint string
		}
		Node struct {
			Leveldb   string
			Dashboard string
			Meta      struct {
				Email string
				Name  string
			}
			Endpoint string
		}
		NATS struct {
			Endpoint string
			User     string
			Password string
		}
		Discovery struct {
			Endpoint string
		}
	}{}
)
