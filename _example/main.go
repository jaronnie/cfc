package main

import (
	"fmt"
	"github.com/jaronnie/cfc/encoding"
	"github.com/jaronnie/cfc/encoding/toml"
	"github.com/spf13/viper"
)

func main() {
	x := viper.New()
	x.SetConfigFile("./test.toml")
	err := x.ReadInConfig()
	if err != nil {
		panic(err)
		return
	}
	x.Set("name", "value2")
	x.Set("pool.size", 120)

	settings := x.AllSettings()

	registry := encoding.NewEncoderRegistry()
	err = registry.RegisterEncoder("toml", toml.Codec{})
	if err != nil {
		fmt.Println(err)
		return
	}

	encode, err := registry.Encode("toml", settings)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(encode))
}
