/*
Copyright © 2022 jaronnie jaron@jaronnie.com

*/

package cmd

import (
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// appendCmd represents the append command
var appendCmd = &cobra.Command{
	Use:   "append",
	Short: "append config key, the key value must by array",
	Long:  `append config key, the key value must by array.`,
	RunE:  appendEx,
}

func appendEx(cmd *cobra.Command, args []string) error {
	cmd.SilenceUsage = true
	if len(args) == 0 {
		return EmptyKey
	}
	key := args[0]
	value := args[1]

	if b := viper.IsSet(key); b {
		getValue := viper.Get(key)
		switch v := getValue.(type) {
		case []interface{}:
			if len(v) == 0 {
				// 空 slice, 不支持
				return errors.Errorf("empty slice do not support append operation." +
					"\nYou should use set_strings or set_ints or set_floats or set_objects to initialize")
			}

			var castValue interface{}
			var err error

			switch v[0].(type) {
			case string:
				if castValue, err = cast.ToStringE(value); err != nil {
					return err
				}
			case int64:
				if castValue, err = cast.ToInt64E(value); err != nil {
					return err
				}
			case float64:
				if castValue, err = cast.ToFloat64E(value); err != nil {
					return err
				}
			case map[string]interface{}:
				if castValue, err = cast.ToStringMapE(value); err != nil {
					return err
				}
			}

			v = append(v, castValue)

			if err := setInterface(key, v); err != nil {
				return err
			}
			if err := viper.WriteConfigAs(ConfigFile); err != nil {
				return err
			}
		default:
			return errors.Errorf("only array keys support append")
		}
	} else {
		return errors.Errorf("key [%s] not exist", key)
	}
	return nil
}

func init() {
	rootCmd.AddCommand(appendCmd)
}
