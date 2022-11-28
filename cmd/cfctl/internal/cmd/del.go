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
	Use:   "del",
	Short: "del config key returns value",
	Long:  `del config key returns value.`,
	RunE:  del,
}

func del(cmd *cobra.Command, args []string) error {
	cmd.SilenceUsage = true
	if len(args) == 0 {
		return EmptyKey
	}

	key := args[0]

	viper.UnSet(key)
	if err := viper.WriteConfigAs(ConfigFile); err != nil {
		return err
	}
	return nil
}

func init() {
	rootCmd.AddCommand(delCmd)
}
