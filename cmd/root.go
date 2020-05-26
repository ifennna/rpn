package cmd

import (
	"fmt"
	"noculture/rpn/core"
	"os"

	"github.com/spf13/cobra"
)

var interactive = false
var root = &cobra.Command{
	Use:   "rpn",
	Short: "A reverse polish notation calculator",
	Long:  `rpn is a cli tool that brings the power and flexibility of Reverse Polish Notation to your terminal.`,
	Run: func(cmd *cobra.Command, args []string) {
		if interactive {
			core.Repl()
		} else {
			core.Calculate(args)
		}
	},
}

func Register(cmd *cobra.Command) {
	root.AddCommand(cmd)
}

func Execute() {
	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	root.Flags().BoolVarP(&interactive, "interactive", "i", false, "Lauch interactive mode")
}
