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
	InvalidSetIntArgs = errors.New("invalid set_int cmd args")
)

// setIntCmd represents the set_int command
var setIntCmd = &cobra.Command{
	Use:   "set_int",
	Short: "set_int config key",
	Long:  `set_int config key.`,
	RunE:  setInt,
}

func setInt(cmd *cobra.Command, args []string) error {
	if len(args) != 2 {
		return InvalidSetIntArgs
	}
	cmd.SilenceUsage = true

	key := args[0]
	value := args[1]

	castValue, err := cast.ToIntE(value)
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
	rootCmd.AddCommand(setIntCmd)
}
