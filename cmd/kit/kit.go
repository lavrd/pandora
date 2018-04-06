package main

import (
	"path/filepath"
	"strings"

	"github.com/spacelavr/pandora/pkg/api"
	"github.com/spacelavr/pandora/pkg/core"
	"github.com/spacelavr/pandora/pkg/discovery"
	"github.com/spacelavr/pandora/pkg/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	config string

	CLI = &cobra.Command{
		PersistentPreRun: func(cmd *cobra.Command, args []string) {

			abs, err := filepath.Abs(config)
			if err != nil {
				log.Fatal(err)
			}

			// get the config name
			base := filepath.Base(abs)
			// get the path
			path := filepath.Dir(abs)

			viper.SetConfigName(strings.Split(base, ".")[0])
			viper.AddConfigPath(path)

			if err := viper.ReadInConfig(); err != nil {
				log.Fatal(err)
			}

			log.SetVerbose(viper.GetBool("verbose"))
		},

		Run: func(cmd *cobra.Command, args []string) {

			var (
				done    = make(chan bool, 1)
				apps    = make(chan bool)
				wait    = 0
				daemons = map[string]func() bool{
					"api":       api.Daemon,
					"core":      core.Daemon,
					"discovery": discovery.Daemon,
				}
			)

			components := []string{"api", "core", "discovery"}

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
	CLI.Flags().StringVarP(&config, "config", "c", "./contrib/config.yml", "/path/to/config.yml")
}

func main() {
	if err := CLI.Execute(); err != nil {
		log.Fatal(err)
	}
}
