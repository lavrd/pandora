package main

import (
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"pandora/pkg/conf"
	"pandora/pkg/discovery"
	"pandora/pkg/master"
	"pandora/pkg/membership"
	"pandora/pkg/node"
	"pandora/pkg/utils/errors"
	"pandora/pkg/utils/log"
)

var (
	cfg string

	// CLI main command
	CLI = &cobra.Command{
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			abs, err := filepath.Abs(cfg)
			if err != nil {
				log.Fatal(errors.WithStack(err))
			}

			base := filepath.Base(abs)
			path := filepath.Dir(abs)

			viper.SetConfigName(strings.Split(base, ".")[0])
			viper.AddConfigPath(path)

			if err := viper.ReadInConfig(); err != nil {
				log.Fatal(errors.WithStack(err))
			}

			if err := viper.Unmarshal(conf.Conf); err != nil {
				log.Fatal(errors.WithStack(err))
			}

			log.SetVerbose(conf.Conf.Runtime.Verbose)
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
}

func main() {
	if err := CLI.Execute(); err != nil {
		log.Fatal(errors.WithStack(err))
	}
}
