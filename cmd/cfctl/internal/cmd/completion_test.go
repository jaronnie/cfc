package cmd

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
)

func TestGetOuterKeys(t *testing.T) {
	viper.SetConfigFile("/Users/jaronnie/Desktop/jaronnie/git.hyperchain.cn/blocface/bricklayer/_example/brick/console/chain/hyperchain_rbft_k8s.yaml")

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println(err)
		return
	}

	keys := getOuterKeys("spec.template.alliancechaininfo.machines.0.nodes")
	fmt.Println(keys)
}
