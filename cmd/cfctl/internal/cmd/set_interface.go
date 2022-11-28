package cmd

import (
	"strings"

	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

func setInterface(key string, castValue interface{}) error {
	// Determine whether to set index value
	if isSetIndexValue(key) {
		if err := viper.SetIndex(key, castValue); err != nil {
			return err
		}
	} else {
		viper.Set(key, castValue)
	}
	return nil
}

func isSetIndexValue(key string) bool {
	// defaultKeyDelimiter is .
	split := strings.Split(key, ".")
	for i, v := range split {
		if _, err := cast.ToIntE(v); err == nil {
			// cast to int successfully
			// Determine whether the key is an array
			join := strings.Join(split[:i], ".")
			switch viper.Get(join).(type) {
			case []interface{}:
				return true
			}
		}
	}
	return false
}
