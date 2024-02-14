package cli

import (
	"os"

	"github.com/spf13/cobra"
)

func Completion() *cobra.Command {
	completionCmd := &cobra.Command{
		Use:   "completion [bash|zsh|fish|powershell]",
		Short: "Generate completion script",
		Long: `To load completions:
Bash:
  $ source <(twigo completion bash)
  # To load completions for each session, execute once:
  # Linux:
  $ twigo completion bash > /etc/bash_completion.d/twigo
  # macOS:
  $ twigo completion bash > /usr/local/etc/bash_completion.d/twigo
Zsh:
  # If shell completion is not already enabled in your environment,
  # you will need to enable it.  You can execute the following once:
  $ echo "autoload -U compinit; compinit" >> ~/.zshrc
  # To load completions for each session, execute once:
  $ twigo completion zsh > "${fpath[1]}/_twigo"
  # You will need to start a new shell for this setup to take effect.
fish:
  $ twigo completion fish | source
  # To load completions for each session, execute once:
  $ twigo completion fish > ~/.config/fish/completions/twigo.fish
PowerShell:
  PS> twigo completion powershell | Out-String | Invoke-Expression
  # To load completions for every new session, run:
  PS> twigo completion powershell > twigo.ps1
  # and source this file from your PowerShell profile.
`,
		DisableFlagsInUseLine: true,
		ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
		Args:                  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		Run: func(cmd *cobra.Command, args []string) {
			switch args[0] {
			case "bash":
				_ = cmd.Root().GenBashCompletion(os.Stdout)
			case "zsh":
				_ = cmd.Root().GenZshCompletion(os.Stdout)
			case "fish":
				_ = cmd.Root().GenFishCompletion(os.Stdout, true)
			case "powershell":
				_ = cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
			}
		},
	}

	return completionCmd
}
