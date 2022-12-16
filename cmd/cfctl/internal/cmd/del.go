/*
Copyright © 2022 jaronnie jaron@jaronnie.com

*/

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:               "del",
	Short:             "del config key",
	Long:              `del config key.`,
	ValidArgsFunction: ValidArgsFunction,
	RunE:              del,
}

func del(cmd *cobra.Command, args []string) error {
	cmd.SilenceUsage = true
	if len(args) == 0 {
		return EmptyKey
	}

	key := args[0]

	if err := tryReadConfig(); err != nil {
		return err
	}

	viper.UnSet(key)
	if err := viper.WriteConfigAs(ConfigFile); err != nil {
		return err
	}
	return nil
}

func init() {
	rootCmd.AddCommand(delCmd)
}
