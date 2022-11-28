/*
Copyright © 2022 jaronnie git.hyperchain.cn

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
	InvalidSetStringsArgs = errors.New("invalid set_strings cmd args")
)

// setStringsCmd represents the set_strings command
var setStringsCmd = &cobra.Command{
	Use:   "set_strings",
	Short: "set_strings config key",
	Long:  `set_strings config key.`,
	RunE:  setStrings,
}

func setStrings(cmd *cobra.Command, args []string) error {
	if len(args) != 2 {
		return InvalidSetStringsArgs
	}
	cmd.SilenceUsage = true

	key := args[0]
	value := args[1]

	value = strings.ReplaceAll(value, " ", "")

	split := strings.Split(value, ",")

	castValue, err := cast.ToStringSliceE(split)
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
	rootCmd.AddCommand(setStringsCmd)
}
