package conf

var (
	// Viper conf struct
	Viper = &struct {
		Runtime struct {
			Verbose bool
			Clean   bool
		}
		TLS struct {
			Cert string
			Key  string
		}
		Membership struct {
			Mail struct {
				Send  bool
				Token string
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
			Database struct {
				FilePath string
			}
			Dashboard struct {
				Template string
			}
			Meta struct {
				Email string
				Name  string
			}
			Endpoint string
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
