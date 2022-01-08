/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"plugin"
	"strings"

	"github.com/catdevman/awsume-go/pkg/hooks"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "awsume-go",
	Short: "Awsume - A cli that makes using AWS IAM credentials easy",
	Long:  `Thank you for using AWSume! Check us out at https://trek10.com`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(strings.Join(args, " "))
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.awsume-go.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	plugins, err := filepath.Glob("./plugins/*.so")
	if err != nil {
		panic(err)
	}

	for _, filename := range plugins {
		p, err := plugin.Open(filename)
		if err != nil {
			panic(err)
		}

		symbol, err := p.Lookup("New")
		if err != nil {
			panic(err)
		}

		plug, err := symbol.(func(*cobra.Command) (interface{}, error))(rootCmd)
		if err != nil {
			panic(err)
		}
		plugprearg, ok := plug.(hooks.AddArgumentsHook)
		if ok {
			plugprearg.AddArguments()
		}

		plugpost, ok := plug.(hooks.PostAddArgumentsHook)
		if ok {
			plugpost.PostAddArguments()
		}
	}

}
