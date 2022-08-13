package cmd

import (
	"github.com/spf13/cobra"
	"fmt")
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
	  fmt.Println(err)
	}
  }