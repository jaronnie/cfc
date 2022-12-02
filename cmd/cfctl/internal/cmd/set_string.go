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
	InvalidSetStringArgs = errors.New("invalid set_string cmd args")
)

// setStringCmd represents the set command
var setStringCmd = &cobra.Command{
	Use:   "set_string",
	Short: "set_string config key with string type",
	Long:  `set_String config key with string type.`,
	RunE:  setString,
}

func setString(cmd *cobra.Command, args []string) error {
	if len(args) != 2 {
		return InvalidSetStringArgs
	}
	cmd.SilenceUsage = true

	key := args[0]
	value := args[1]

	if err := tryReadConfig(); err != nil {
		return err
	}

	castValue, err := cast.ToStringE(value)
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
	rootCmd.AddCommand(setStringCmd)
}
