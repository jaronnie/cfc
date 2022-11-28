/*
Copyright © 2022 jaronnie git.hyperchain.cn

*/

package cmd

import (
	"errors"

	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	InvalidSetFloatArgs = errors.New("invalid set_float cmd args")
)

// setFloatCmd represents the set_float command
var setFloatCmd = &cobra.Command{
	Use:   "set_float",
	Short: "set_float config key",
	Long:  `set_float config key.`,
	RunE:  setFloat,
}

func setFloat(cmd *cobra.Command, args []string) error {
	if len(args) != 2 {
		return InvalidSetFloatArgs
	}
	cmd.SilenceUsage = true

	key := args[0]
	value := args[1]

	castValue, err := cast.ToFloat64E(value)
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
	rootCmd.AddCommand(setFloatCmd)
}
