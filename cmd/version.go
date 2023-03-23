// Copyright 2022 negineri.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Version  = "none"
	Revision = "none"
)

func newVersionCmd() *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Show version information",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("version called")
			fmt.Printf("Version:    %v\n", Version)
			fmt.Printf("Git commit: %v\n", Revision)
			return nil
		},
	}
	return versionCmd
}
