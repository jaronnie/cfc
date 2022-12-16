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
	InvalidSetObjectArgs = errors.New("invalid set_object cmd args")
)

// setObjectCmd represents the set_object command
var setObjectCmd = &cobra.Command{
	Use:               "set_object",
	Short:             "set_object config key with object type",
	Long:              `set_object config key with object type.`,
	ValidArgsFunction: ValidArgsFunction,
	RunE:              setObject,
}

func setObject(cmd *cobra.Command, args []string) error {
	if len(args) != 2 {
		return InvalidSetObjectArgs
	}
	cmd.SilenceUsage = true

	key := args[0]
	value := args[1]

	if err := tryReadConfig(); err != nil {
		return err
	}

	castValue, err := cast.ToStringMapE(value)
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
	rootCmd.AddCommand(setObjectCmd)
}
