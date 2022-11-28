/*
Copyright © 2022 jaronnie jaron@jaronnie.com

*/

package cmd

import (
	"errors"

	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	InvalidSetBoolArgs = errors.New("invalid set_bool cmd args")
)

// setBoolCmd represents the set_bool command
var setBoolCmd = &cobra.Command{
	Use:   "set_bool",
	Short: "set_bool config key",
	Long:  `set_bool config key.`,
	RunE:  setBool,
}

func setBool(cmd *cobra.Command, args []string) error {
	if len(args) != 2 {
		return InvalidSetBoolArgs
	}
	cmd.SilenceUsage = true

	key := args[0]
	value := args[1]

	castValue, err := cast.ToBoolE(value)
	if err != nil {
		return err
	}

	if err := setInterface(key, castValue); err != nil {
		return err
	}

	if err := viper.WriteConfigAs(ConfigFile); err != nil {
		return err
	}
	return nil
}

func init() {
	rootCmd.AddCommand(setBoolCmd)
}
