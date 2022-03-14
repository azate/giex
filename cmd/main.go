package cmd

import (
	"github.com/azate/giex/internal/runner"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strings"
)

func NewCmd() *cobra.Command {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("GIEX")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	cmd := &cobra.Command{
		Use: "giex",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := runner.LoadConfig()
			if err := cfg.Check(); err != nil {
				return err
			}

			return runner.New(cfg).Go()
		},
	}

	cmd.Flags().StringP("input", "i", "domains.txt", "Path to the file with domains <one row one domain>")
	cmd.Flags().StringP("output", "o", "/tmp", "Path to the folder for saving the git configs")
	cmd.Flags().UintP("max-workers", "w", 100, "Maximum workers")
	cmd.Flags().UintP("max-tasks", "t", 200, "Maximum prepared tasks")
	cmd.Flags().StringP("proxy", "p", "", "HTTP proxy")

	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		panic(err)
	}

	return cmd
}
