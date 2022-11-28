/*
Copyright © 2022 jaronnie git.hyperchain.cn

*/

package cmd

import (
	"errors"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/jaronnie/cfc/cmd/cfctl/internal/extend/castEx"
)

var (
	InvalidSetFloatsArgs = errors.New("invalid set_strings cmd args")
)

// setFloatsCmd represents the set_floats command
var setFloatsCmd = &cobra.Command{
	Use:   "set_floats",
	Short: "set_floats config key",
	Long:  `set_floats config key.`,
	RunE:  setFloats,
}

func setFloats(cmd *cobra.Command, args []string) error {
	if len(args) != 2 {
		return InvalidSetFloatsArgs
	}
	cmd.SilenceUsage = true

	key := args[0]
	value := args[1]

	value = strings.ReplaceAll(value, " ", "")

	split := strings.Split(value, ",")

	castValue, err := castEx.ToFloat64SliceE(split)
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
	rootCmd.AddCommand(setFloatsCmd)
}
