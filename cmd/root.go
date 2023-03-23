// Copyright 2022 negineri.
// SPDX-License-Identifier: Apache-2.0

// Package cmd is provided by cobra
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	configfileDefault = "~/.hashibiro.yaml"
)

var cfgFile string

func newRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "hashibiro",
		Short: "alive monitor",
		RunE: func(cmd *cobra.Command, args []string) error {
			version, err := cmd.Flags().GetBool("version")
			if err != nil {
				return err
			}
			if version {
				fmt.Println(Version)
				return nil
			}
			return nil
		},
	}

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ~/.hashibiro.yaml)")

	rootCmd.Flags().BoolP("version", "v", false, "Show version")
	rootCmd.AddCommand(newVersionCmd())
	return rootCmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := newRootCmd().Execute()
	if err != nil {
		os.Exit(1)
	}
}

// Initialize is renaming of init()
// Rewrite in a different form in the future.
// This is called by main.main() before Execute().
func Initialize() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	if cfgFile == "" {
		cfgFile = configfileDefault
	}
	if strings.HasPrefix(cfgFile, "~") {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		cfgFile = strings.Replace(cfgFile, "~", home, 1)
	}
	viper.SetConfigFile(cfgFile)

	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
