package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-fly",
	Short: "go-fly",
	Long:  `简洁快速的GO语言WEB在線客服 https://gofly.sopans.com`,
	Args:  args,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func args(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {

		return errors.New("至少需要一个參數!")
	}
	return nil
}
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func init() {
	rootCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(installCmd)
	rootCmd.AddCommand(stopCmd)
}
