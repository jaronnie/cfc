/*
Copyright © 2022 jaronnie jaron@jaronnie.com

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
	InvalidSetObjectsArgs = errors.New("invalid set_objects cmd args")
)

// setObjectsCmd represents the set_object command
var setObjectsCmd = &cobra.Command{
	Use:   "set_objects",
	Short: "set_objects config key with object array type",
	Long:  `set_objects config key with object array type.`,
	RunE:  setObjects,
}

func setObjects(cmd *cobra.Command, args []string) error {
	if len(args) != 2 {
		return InvalidSetObjectsArgs
	}
	cmd.SilenceUsage = true

	key := args[0]
	value := args[1]

	value = strings.ReplaceAll(value, " ", "")

	split := strings.Split(value, ",")

	castValue, err := castEx.ToStringMapSliceE(split)
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
	rootCmd.AddCommand(setObjectsCmd)
}
