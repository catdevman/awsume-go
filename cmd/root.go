/*
Copyright © 2022 Lucas Pearson <catdevman@gmail.com>

*/
package cmd

import (
	"os"
	"path/filepath"
	"plugin"
	"strings"

	"github.com/catdevman/awsume-go/pkg/hooks"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var plugins []interface{}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "awsume",
	Version: "0.0.1",
	Short:   "Awsume - A cli that makes using AWS IAM credentials easy",
	Long:    "Thank you for using AWSume! Check us out at https://trek10.com",
	Run: func(cmd *cobra.Command, args []string) {
		handlePostArgs(plugins)
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
	cobra.OnInitialize(initConfig)
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	pluginFiles, err := filepath.Glob(strings.Join([]string{home, ".awsume", "plugins", "*.so"}, string(os.PathSeparator))) // config directory plugins and local plugins in the future
	if err != nil {
		panic(err)
	}

	for _, filename := range pluginFiles {
		p, err := plugin.Open(filename)
		log.Info().Msg(filename)
		if err != nil {
			panic(err)
		}

		symbol, err := p.Lookup("New")
		if err != nil {
			panic(err)
		}
		plug, err := symbol.(func(*cobra.Command, *viper.Viper) (interface{}, error))(rootCmd, viper.GetViper())
		if err != nil {
			panic(err)
		}
		plugins = append(plugins, plug)
	}
	handlePreArgs(plugins)
	handleArgs(plugins)
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

func handlePreArgs(plugs []interface{}) {
	for _, p := range plugs {
		preargplugin, ok := p.(hooks.PreAddArgumentsHook)
		if ok {
			preargplugin.PreAddArguments()
		}
	}

}

func handleArgs(plugs []interface{}) {
	for _, p := range plugs {
		addargsplugin, ok := p.(hooks.AddArgumentsHook)
		if ok {
			addargsplugin.AddArguments()
		}
	}

}

func handlePostArgs(plugs []interface{}) {
	for _, p := range plugs {
		postargsplugin, ok := p.(hooks.PostAddArgumentsHook)
		if ok {
			postargsplugin.PostAddArguments()
		}
	}

}
