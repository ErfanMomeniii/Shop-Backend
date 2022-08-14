package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "shop",
	Short: "this project is the shop backend",
	Long: `this project is the shop backend
					Complete documentation is available at https://github.com/ErfanMomeniii/Shop-Backend`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
