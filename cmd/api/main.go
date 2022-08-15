package api

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
)

func initialize() *echo.Echo {
	e := echo.New()
	return e
}

func Register(root *cobra.Command) {
	root.AddCommand(
		&cobra.Command{
			Use:   "api",
			Short: "SHOP API Component",
			Run: func(cmd *cobra.Command, args []string) {
				initialize()
			},
		},
	)
}
