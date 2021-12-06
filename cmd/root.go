package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

const Version = "1.0.0"

var rootCmd = &cobra.Command{
	Use:          "god",
	Short:        "God",
	Long:         `God`,
	Version:      Version,
	SilenceUsage: true, // disable print usage if error occurs
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	// rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c",
	// 	"", "config file (default: ./configs/config.yaml)")
}

func initConfig() {
	// if err := cfg.InitConfig(cfgFile); err != nil {
	// 	panic(err)
	// }
	// cfg.WatchConfig(cfgChange)
}
