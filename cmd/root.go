/*
Copyright © 2022 Lucas Pearson <catdevman@gmail.com>

*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/catdevman/awsume-go/shared"
	"github.com/hashicorp/go-plugin"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var plugins []*plugin.Client
var logger zerolog.Logger
var profiles shared.Profiles

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "awsume",
	Version: "0.0.1",
	Short:   "Awsume - A cli that makes using AWS IAM credentials easy",
	Long:    "Thank you for using AWSume! Check us out at https://trek10.com",
	Run: func(cmd *cobra.Command, args []string) {
		handlePostArgs(plugins)
		getProfiles(plugins)
		getCredentials(plugins)
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

	//zerolog.SetGlobalLevel(zerolog.InfoLevel)
	//logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).With().Timestamp().Logger()

	pluginFiles, err := filepath.Glob(strings.Join([]string{home, ".awsume", "plugins", "*"}, string(os.PathSeparator))) // config directory plugins and local plugins in the future
	if err != nil {
		panic(err)
	}

	// We should no longer need a share app to pass around we will use messages over grpc to communicate state back to the main process
	//awsume := shared.Awsume{Cmd: rootCmd, Config: viper.GetViper(), Logger: &logger}
	// loop through files and boot them up
	// TODO: How do I know what to dispense from the RPC client... can they all have the same name like `awsume-plugin`?
	for _, filename := range pluginFiles {
		client := plugin.NewClient(&plugin.ClientConfig{
			HandshakeConfig:  shared.Handshake,
			Plugins:          shared.PluginMap,
			Cmd:              exec.Command("sh", "-c", filename), //TODO: This seems heavy handed can we load these paths from the conf file
			AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
		})

		plugins = append(plugins, client)
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

func handlePreArgs(plugs []*plugin.Client) {

}

func handleArgs(plugs []*plugin.Client) {

}

func handlePostArgs(plugs []*plugin.Client) {
}

func getProfiles(plugs []*plugin.Client) {
	handlePreCollectProfiles(plugs)
	handleCollectProfiles(plugs)
	handlePostCollectProfiles(plugs)
}

func handlePreCollectProfiles(plugs []*plugin.Client) {
	for _, p := range plugs {
		client, err := p.Client()
		if err != nil {
			fmt.Println("Error:", err.Error())
			continue
		}
		raw, err := client.Dispense("profiles_grpc")
		if err != nil {
			fmt.Println("Error:", err.Error())
			continue
		}
		pregetprofileplugin := raw.(shared.ProfilesService)
		pregetprofileplugin.Pre()
	}
}

func handleCollectProfiles(plugs []*plugin.Client) {
	//	for _, p := range plugs {
	//		getprofileplugin, ok := &p.(hooks.CollectProfilesHook)
	//		if ok {
	//			// TODO: mutex lock and use go routines
	//			// or make a profiles channel give to CollecctProfiles
	//			prs := getprofileplugin.CollectProfiles()
	//			if profiles == nil {
	//				profiles = shared.Profiles{}
	//			}
	//			fmt.Println(prs)
	//			profiles = profiles.Add(prs)
	//		}
	//	}
}

func handlePostCollectProfiles(plugs []*plugin.Client) {
}

func getCredentials(plugs []*plugin.Client) {
	handlePreGetCredentials(plugs)
	handleGetCredentials(plugs)
	handlePostGetCredentials(plugs)
}

func handlePreGetCredentials(plugs []*plugin.Client) {

}

func handleGetCredentials(plugs []*plugin.Client) {

}

func handlePostGetCredentials(plugs []*plugin.Client) {

}
