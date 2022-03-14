package main

import (
	"github.com/azate/giex/cmd"
	"github.com/spf13/cobra"
)

func main() {
	cobra.CheckErr(cmd.NewCmd().Execute())
}
