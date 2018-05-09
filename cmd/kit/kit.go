package main

import (
	"path/filepath"
	"strings"

	"github.com/spacelavr/pandora/pkg/config"
	"github.com/spacelavr/pandora/pkg/discovery"
	"github.com/spacelavr/pandora/pkg/master"
	"github.com/spacelavr/pandora/pkg/membership"
	"github.com/spacelavr/pandora/pkg/node"
	"github.com/spacelavr/pandora/pkg/tracker"
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

			if err := viper.Unmarshal(config.Viper); err != nil {
				log.Fatal(err)
			}

			log.SetVerbose(config.Viper.Runtime.Verbose)
		},

		Run: func(cmd *cobra.Command, args []string) {
			var (
				done    = make(chan bool)
				apps    = make(chan bool)
				wait    = 0
				daemons = map[string]func() bool{
					"node":       node.Daemon,
					"master":     master.Daemon,
					"discovery":  discovery.Daemon,
					"membership": membership.Daemon,
				}
			)

			components := []string{"node", "master", "discovery", "membership"}

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
							return
						}
					}
				}
			}()

			<-done
		},
	}
)

func init() {
	CLI.Flags().StringVarP(&cfg, "config", "c", "./contrib/config.yml", "/path/to/config.yml")
}

func main() {
	if err := CLI.Execute(); err != nil {
		log.Fatal(err)
	}
}
