package cmd

import (
	"espad/back/uppash/internal/adapters/driving"
	"github.com/spf13/cobra"
)

var (
	rootcmd = &cobra.Command{
		Use: "uploader",
		Short: "upload files",
		Long: "upload files to minio server using tus service",
		Run: func(cmd *cobra.Command, args []string) {
			driving.StartEngine()
		},
	}
)

func Execute() error {
	return rootcmd.Execute()
}
