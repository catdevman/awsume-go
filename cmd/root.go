/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
	"path/filepath"
	"plugin"

	"github.com/catdevman/awsume-go/pkg/hooks"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "awsume-go",
	Short: "Awsume - A cli that makes using AWS IAM credentials easy",
	Long:  "Thank you for using AWSume! Check us out at https://trek10.com",
	Run:   func(cmd *cobra.Command, args []string) {},
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
	cobra.OnInitialize(initConfig)

	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config-file", "", "config file (default is $HOME/.awsume/config.yaml)")

	plugins, err := filepath.Glob("./plugins/*.so") // config directory plugins and local plugins in the future
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

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory in ".awsume" with the name "config.yaml".
		viper.AddConfigPath(home + string(os.PathSeparator) + ".awsume")
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
	}
}
