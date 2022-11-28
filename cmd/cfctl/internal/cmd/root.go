/*
Copyright © 2022 jaronnie git.hyperchain.cn

*/

package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	ConfigFile string
	FileType   string // only when used pipe mode will be used
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cfc",
	Short: "cfc means config customize",
	Long:  `cfc means config customize.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cc.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringVarP(&ConfigFile, "config", "f", "", "config file path")
	rootCmd.PersistentFlags().StringVarP(&FileType, "type", "p", "toml", "only take effect reading from pipeline, default toml type")
}

func initConfig() {
	if ConfigFile != "" {
		viper.SetConfigFile(ConfigFile)
		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}
	} else if os.Stdin != nil {
		viper.SetConfigType(FileType)
		if err := viper.ReadConfig(os.Stdin); err != nil {
			panic(err)
		}
	}
}
