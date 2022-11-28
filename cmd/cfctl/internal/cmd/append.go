/*
Copyright © 2022 jaronnie git.hyperchain.cn

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
	Short: "append config key returns value",
	Long:  `append config key returns value.`,
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

			switch v[0].(type) {
			case string:
				v = append(v, cast.ToString(value))
			case int64:
				v = append(v, cast.ToInt64(value))
			case float64:
				v = append(v, cast.ToFloat64(value))
			case map[string]interface{}:
				v = append(v, cast.ToStringMap(value))
			}

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
