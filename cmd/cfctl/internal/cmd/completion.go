/*
Copyright © 2022 jaronnie jaron@jaronnie.com

*/

package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/jaronnie/cfc/cmd/cfctl/internal/utilx"
)

var completionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish|powershell]",
	Short: "Generate completion script",
	Long: `To load completions:

Bash:

$ source <(cfctl completion bash)

# To load completions for each session, execute once:
Linux:
  $ cfctl completion bash > /etc/bash_completion.d/cfctl
MacOS:
  $ cfctl completion bash > /usr/local/etc/bash_completion.d/cfctl

Zsh:

# If shell completion is not already enabled in your environment you will need
# to enable it.  You can execute the following once:

$ echo "autoload -U compinit; compinit" >> ~/.zshrc

# To load completions for each session, execute once:
$ cfctl completion zsh > "${fpath[1]}/_cfctl"

# You will need to start a new shell for this setup to take effect.

Fish:

$ cfctl completion fish | source

# To load completions for each session, execute once:
$ cfctl completion fish > ~/.config/fish/completions/cfctl.fish
`,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			_ = cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			_ = cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			_ = cmd.Root().GenFishCompletion(os.Stdout, true)
		case "powershell":
			_ = cmd.Root().GenPowerShellCompletion(os.Stdout)
		}
	},
}

func ValidArgsFunction(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	err := tryReadConfig()
	if err != nil {
		return nil, cobra.ShellCompDirectiveDefault
	}

	keys := viper.AllKeys()

	cobra.CompDebugln(strings.Join(keys, ","), true)

	choiceKeys := make([]string, 0)

	switch cmd.Name() {
	case "set_string":
		for _, b := range keys {
			if _, ok := viper.Get(b).(string); ok {
				choiceKeys = append(choiceKeys, b)
			}
		}
	case "set_int":
		for _, b := range keys {
			if _, ok := viper.Get(b).(int); ok {
				choiceKeys = append(choiceKeys, b)
			}
		}
	default:
		cobra.CompDebugln(toComplete, true)
		choiceKeys = getOuterKeys(toComplete)
	}

	return choiceKeys, cobra.ShellCompDirectiveNoSpace
}

func getOuterKeys(key string) []string {
	var subViper *viper.Viper

	outerKeys := make([]string, 0)
	if key != "" {
		subViper = viper.Sub(strings.TrimRight(key, "."))
		if key != "" && !strings.HasSuffix(key, ".") {
			return nil
		}
	} else {
		subViper = viper.GetViper()

	}

	if subViper == nil {
		return nil
	}

	for _, v := range subViper.AllKeys() {
		split := strings.Split(v, ".")

		if len(split) > 0 {
			outerKeys = append(outerKeys, key+split[0])
		} else {
			outerKeys = append(outerKeys, v)
		}

	}

	cobra.CompDebugln(strings.Join(outerKeys, ","), true)

	return utilx.RemoveDuplicateElement(outerKeys)
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
