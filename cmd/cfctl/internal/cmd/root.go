/*
Copyright © 2022 jaronnie jaron@jaronnie.com

*/

package cmd

import (
	"bytes"
	"io"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	ConfigFile string
	FileType   string // only when used pipe mode will be used
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cfctl",
	Short: "cfctl means config customize by command line",
	Long:  `cfctl means config customize by command line.`,
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cc.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringVarP(&ConfigFile, "config", "f", "", "config file path")
	rootCmd.PersistentFlags().StringVarP(&FileType, "type", "p", "", "specify config file type")
}

func tryReadConfig() (err error) {
	switch {
	case ConfigFile != "" && filepath.Ext(ConfigFile) != "":
		viper.SetConfigFile(ConfigFile)
		return viper.ReadInConfig()
	case ConfigFile != "" && filepath.Ext(ConfigFile) == "" && FileType != "":
		viper.SetConfigFile(ConfigFile)
		viper.SetConfigType(FileType)
		return viper.ReadInConfig()
	case ConfigFile != "" && filepath.Ext(ConfigFile) == "" && FileType == "":
		viper.SetConfigFile(ConfigFile)
		// try set config type
		for _, v := range viper.SupportedExts {
			viper.SetConfigType(v)
			if err = viper.ReadInConfig(); err == nil {
				return nil
			}
		}
		return err
	}

	// os stdin
	if stat, err := os.Stdin.Stat(); err != nil {
		return err
	} else if stat.Size() == 0 {
		return errors.Errorf("please specify config file or set os stdin")
	}

	// read stdin bytes
	stdinBytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		return err
	}

	if FileType != "" {
		viper.SetConfigType(FileType)
		return viper.ReadConfig(bytes.NewBuffer(stdinBytes))
	}

	for _, v := range viper.SupportedExts {
		viper.SetConfigType(v)
		if err = viper.ReadConfig(bytes.NewBuffer(stdinBytes)); err == nil {
			return nil
		}
	}
	return err
}
