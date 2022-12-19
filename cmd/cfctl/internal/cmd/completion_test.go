package cmd

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
)

func TestGetOuterKeys(t *testing.T) {
	viper.SetConfigFile("../testdata/hyperchain_rbft_k8s.yaml")

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println(err)
		return
	}

	keys := getOuterKeys("")
	t.Log(keys)

	keys = getOuterKeys("s")
	t.Log(keys)

	keys = getOuterKeys("spec")
	t.Log(keys)

	keys = getOuterKeys("spec.te")
	t.Log(keys)

	keys = getOuterKeys("spec.so")
	t.Log(keys)

	keys = getOuterKeys("spec.p")
	t.Log(keys)
}
