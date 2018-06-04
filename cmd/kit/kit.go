package main

import (
	"path/filepath"
	"strings"
	"time"

	"github.com/spacelavr/pandora/pkg/conf"
	"github.com/spacelavr/pandora/pkg/discovery"
	"github.com/spacelavr/pandora/pkg/master"
	"github.com/spacelavr/pandora/pkg/membership"
	"github.com/spacelavr/pandora/pkg/node"
	"github.com/spacelavr/pandora/pkg/utils/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfg string

	// CLI main command
	CLI = &cobra.Command{
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			abs, err := filepath.Abs(cfg)
			if err != nil {
				log.Fatal(err)
			}

			base := filepath.Base(abs)
			path := filepath.Dir(abs)

			viper.SetConfigName(strings.Split(base, ".")[0])
			viper.AddConfigPath(path)

			if err := viper.ReadInConfig(); err != nil {
				log.Fatal(err)
			}

			if err := viper.Unmarshal(conf.Viper); err != nil {
				log.Fatal(err)
			}

			log.Debug(conf.Viper.Membership.Mail.Token)

			log.SetVerbose(conf.Viper.Runtime.Verbose)
		},

		Run: func(cmd *cobra.Command, args []string) {
			var (
				done    = make(chan bool)
				apps    = make(chan bool)
				wait    = 0
				daemons = map[string]func() bool{
					node.NODE:             node.Daemon,
					master.MASTER:         master.Daemon,
					discovery.DISCOVERY:   discovery.Daemon,
					membership.MEMBERSHIP: membership.Daemon,
				}
			)

			components := []string{node.NODE, master.MASTER, discovery.DISCOVERY, membership.MEMBERSHIP}

			if len(args) > 0 {
				components = args
			}

			for _, app := range components {
				go func(app string) {
					if _, ok := daemons[app]; ok {
						wait++
						apps <- daemons[app]()
					}
				}(app)
			}

			go func() {
				for {
					select {
					case <-apps:
						wait--
						if wait == 0 {
							done <- true
						}
					}
				}
			}()

			<-done
		},
	}
)

func init() {
	CLI.Flags().StringVarP(&cfg, "conf", "c", "./contrib/conf.yml", "/path/to/conf.yml")
	viper.BindEnv("membership.mail.token")
}

func main() {
	if err := CLI.Execute(); err != nil {
		log.Fatal(err)
	}
}
