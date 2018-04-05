package main

import (
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
	`github.com/spacelavr/pandora/pkg/log`
	"github.com/spacelavr/pandora/pkg/node"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	config string

	CLI = &cobra.Command{
		PersistentPreRun: func(cmd *cobra.Command, args []string) {

			abs, err := filepath.Abs(config)
			if err != nil {
				logrus.Error(err)
			}

			// get the config name
			base := filepath.Base(abs)
			// get the path
			path := filepath.Dir(abs)

			viper.SetConfigName(strings.Split(base, ".")[0])
			viper.AddConfigPath(path)

			if err := viper.ReadInConfig(); err != nil {
				logrus.Error(err)
			}

			log.Init(viper.GetBool("verbose"))
		},

		Run: func(cmd *cobra.Command, args []string) {
			node.Daemon()
		},
	}
)

func init() {
	CLI.Flags().StringVarP(&config, "config", "c", "./contrib/config.yml", "/path/to/config.yml")
}

func main() {
	if err := CLI.Execute(); err != nil {
		logrus.Error(err)
	}
}
