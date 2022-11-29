/*
Copyright © 2022 jaronnie jaron@jaronnie.com

*/

package cmd

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/jaronnie/cfc/cmd/cfctl/internal/utilx"
	"github.com/jaronnie/cfc/encoding"
	"github.com/jaronnie/cfc/encoding/dotenv"
	"github.com/jaronnie/cfc/encoding/hcl"
	"github.com/jaronnie/cfc/encoding/ini"
	"github.com/jaronnie/cfc/encoding/javaproperties"
	"github.com/jaronnie/cfc/encoding/json"
	"github.com/jaronnie/cfc/encoding/toml"
	"github.com/jaronnie/cfc/encoding/yaml"
)

var (
	Output string

	EmptyKey = errors.New("empty key")
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get config key returns value",
	Long:  `get config key returns value.`,
	RunE:  get,
}

func get(cmd *cobra.Command, args []string) error {
	cmd.SilenceUsage = true
	if len(args) == 0 {
		return EmptyKey
	}

	key := args[0]

	var value interface{}

	if key == "*" {
		value = viper.AllSettings()
	} else {
		value = viper.Get(key)
	}

	encoderRegister, err := registerEncoder()
	if err != nil {
		return err
	}

	var outputInterface interface{}
	switch Output {
	case "default":
		outputInterface = value
	case "json":
		outputInterface, err = utilx.BeautifyJson(value)
	default:
		var unmarshalOutputBytes []byte
		unmarshalOutputBytes, err = unmarshalOutputType(encoderRegister, value)
		outputInterface = string(unmarshalOutputBytes)
	}

	if err != nil {
		fmt.Printf("key: %s, value: %v\n", key, value)
	} else {
		fmt.Println(outputInterface)
	}
	return err
}

func registerEncoder() (*encoding.EncoderRegistry, error) {
	var registry *encoding.EncoderRegistry
	var err error
	switch Output {
	case "default":
	case "json":
		registry = encoding.NewEncoderRegistry()
		err = registry.RegisterEncoder(Output, json.Codec{})
	case "yaml", "yml":
		registry = encoding.NewEncoderRegistry()
		err = registry.RegisterEncoder(Output, yaml.Codec{})
	case "toml":
		registry = encoding.NewEncoderRegistry()
		err = registry.RegisterEncoder(Output, toml.Codec{})
	case "hcl", "tfvars":
		registry = encoding.NewEncoderRegistry()
		err = registry.RegisterEncoder(Output, hcl.Codec{})
	case "ini":
		registry = encoding.NewEncoderRegistry()
		err = registry.RegisterEncoder(Output, ini.Codec{})
	case "properties", "props", "prop":
		registry = encoding.NewEncoderRegistry()
		err = registry.RegisterEncoder(Output, &javaproperties.Codec{})
	case "dotenv", "env":
		registry = encoding.NewEncoderRegistry()
		err = registry.RegisterEncoder(Output, dotenv.Codec{})
	default:
		return nil, errors.Errorf("not support output type [%s]", Output)
	}
	return registry, err
}

func unmarshalOutputType(registry *encoding.EncoderRegistry, value interface{}) ([]byte, error) {
	if _, ok := value.(map[string]interface{}); !ok {
		return nil, errors.Errorf("not support unmarshal to %s type", Output)
	}

	return registry.Encode(Output, value.(map[string]interface{}))
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.Flags().StringVarP(&Output, "output", "o", "default", "output type")
}
