package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:   "rpn",
	Short: "A reverse polish notation calculator",
	Long: `rpn is a cli tool that brings the power and flexibility of 
	Reverse Polish Notation to your terminal.`,
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
