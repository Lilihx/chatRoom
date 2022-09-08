package cmd

import (
	"github.com/lilihx/chatRoom/account"
	"github.com/lilihx/chatRoom/common/log"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "chatRoom",
	Short: "chatRoom",
	Long:  "chatRoom",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Usage()
	},
}

var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "account",
	Long:  "account",
	Run: func(cmd *cobra.Command, args []string) {
		account.InitServerAndStart()
	},
}

func Execute() {
	rootCmd.AddCommand(accountCmd)
	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
	}
}
