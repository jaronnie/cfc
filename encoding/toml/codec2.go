//go:build !viper_toml1
// +build !viper_toml1

package toml

import (
	toml "github.com/pelletier/go-toml/v2"
)

// Codec implements the encoding.Encoder and encoding.Decoder interfaces for TOML encoding.
type Codec struct{}

func (Codec) Encode(v map[string]interface{}) ([]byte, error) {
	return toml.Marshal(v)
}

func (Codec) Decode(b []byte, v map[string]interface{}) error {
	return toml.Unmarshal(b, &v)
}
