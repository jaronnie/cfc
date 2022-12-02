/*
Copyright © 2022 jaronnie jaron@jaronnie.com

*/

package cmd

import (
	"os"

	"github.com/spf13/cobra"
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

func init() {
	rootCmd.AddCommand(completionCmd)
}
