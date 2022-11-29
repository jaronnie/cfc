/*
Copyright © 2022 jaronnie jaron@jaronnie.com

*/

package cmd

import (
	"errors"
	"strings"

	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	InvalidSetIntsArgs = errors.New("invalid set_ints cmd args")
)

// setIntsCmd represents the set_ints command
var setIntsCmd = &cobra.Command{
	Use:   "set_ints",
	Short: "set_ints config key with int array type",
	Long:  `set_ints config key with int array type.`,
	RunE:  setInts,
}

func setInts(cmd *cobra.Command, args []string) error {
	if len(args) != 2 {
		return InvalidSetIntsArgs
	}
	cmd.SilenceUsage = true

	key := args[0]
	value := args[1]

	value = strings.ReplaceAll(value, " ", "")

	split := strings.Split(value, ",")

	castValue, err := cast.ToIntSliceE(split)
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
	rootCmd.AddCommand(setIntsCmd)
}
