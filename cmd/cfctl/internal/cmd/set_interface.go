package cmd

import (
	"strings"

	"github.com/spf13/cast"
	"github.com/spf13/viper"

	"github.com/jaronnie/cfc/cmd/cfctl/internal/extend/viperEx"
)

func setInterface(key string, castValue interface{}) error {
	// Determine whether to set index value
	if isSetIndexValue(key) {
		// set index value
		// see https://github.com/spf13/viper/issues/1140
		// use extend/viperEx supported by https://github.com/fluffy-bunny/viperEx
		allSettings := viper.AllSettings()
		myViperEx, err := viperEx.New(allSettings, func(ve *viperEx.ViperEx) error {
			return nil
		})
		if err != nil {
			return err
		}
		if err = myViperEx.UpdateDeepPath(key, castValue); err != nil {
			return err
		}
		if err = myViperEx.Unmarshal(&allSettings); err != nil {
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
