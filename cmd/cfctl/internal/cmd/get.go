/*
Copyright © 2022 jaronnie git.hyperchain.cn

*/

package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
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

	fmt.Print(viper.Get(key))
	return nil
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
