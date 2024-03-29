/*
Copyright © 2022 Lucas Pearson <catdevman@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/catdevman/awsume-go/proto"
	"github.com/catdevman/awsume-go/shared"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile  string
	plugins  []*plugin.Client
	profiles shared.Profiles
	logger   hclog.Logger
)

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
		getProfileNames(plugins)
		time.Sleep(time.Millisecond * 100)
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

	pluginFiles, err := filepath.Glob(strings.Join([]string{home, ".awsume", "plugins", "*"}, string(os.PathSeparator))) // config directory plugins and local plugins in the future
	if err != nil {
		panic(err)
	}
	logger = hclog.New(&hclog.LoggerOptions{
		Name:   "awsume",
		Output: os.Stderr,
		Level:  hclog.Debug,
	})

	// loop through files and boot them up
	for _, filename := range pluginFiles {
		client := plugin.NewClient(&plugin.ClientConfig{
			HandshakeConfig:  shared.Handshake,
			Plugins:          shared.PluginMap,
			Cmd:              exec.Command("sh", "-c", filename), // TODO: This seems heavy handed can we load these paths from the conf file
			AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
			Logger:           logger,
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
	for _, p := range plugs {
		client, err := p.Client()
		if err != nil {
			fmt.Println("Error:", err.Error())
			continue
		}
		raw, err := client.Dispense("arguments_grpc")
		if err != nil {
			fmt.Println("Error:", err.Error())
			continue
		}
		argsplugin := raw.(shared.ArgumentsService)
		argsplugin.Pre()
	}
}

func handleArgs(plugs []*plugin.Client) {
	for _, p := range plugs {
		client, err := p.Client()
		if err != nil {
			fmt.Println("Error:", err.Error())
			continue
		}
		raw, err := client.Dispense("arguments_grpc")
		if err != nil {
			fmt.Println("Error:", err.Error())
			continue
		}
		argsplugin := raw.(shared.ArgumentsService)
		args, err := argsplugin.Get()
		// TODO: handle err from rpc call
		// TODO: handle if plugins have the same flag
		for _, a := range args.Arguments {
			switch a.Type {
			case "string":
				rootCmd.PersistentFlags().String(a.Name, a.Value, a.Flag)
			case "int", "int32", "int64":
				// TODO: handle err when converting string to int
				v, _ := strconv.Atoi(a.Value)
				rootCmd.PersistentFlags().Int(a.Name, v, a.Flag)
			case "bool":
				// TODO: handle error when converting string to bool
				v, _ := strconv.ParseBool(a.Value)
				rootCmd.PersistentFlags().Bool(a.Name, v, a.Flag)
			}
		}
	}
}

func handlePostArgs(plugs []*plugin.Client) {
	for _, p := range plugs {
		client, err := p.Client()
		if err != nil {
			fmt.Println("Error:", err.Error())
			continue
		}
		raw, err := client.Dispense("arguments_grpc")
		if err != nil {
			fmt.Println("Error:", err.Error())
			continue
		}
		argsplugin := raw.(shared.ArgumentsService)
		argsplugin.Post(&proto.ArgumentsMsg{})
	}
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
		profileplugin := raw.(shared.ProfilesService)
		profileplugin.Pre()
	}
}

func handleCollectProfiles(plugs []*plugin.Client) {
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
		profileplugin := raw.(shared.ProfilesService)
		profileMsg, err := profileplugin.Get()
		if err != nil {
			continue
		}
		if profiles == nil {
			profiles = shared.Profiles{}
		}
		for _, p := range profileMsg.Profiles {
			fmt.Println(p)
			// profiles = profiles.Add(prs)
		}
	}
}

func handlePostCollectProfiles(plugs []*plugin.Client) {
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
		profileplugin := raw.(shared.ProfilesService)
		// profileplugin.Post(profiles)
		profileplugin.Post(&proto.ProfilesMsg{
			Profiles: []*proto.Profile{
				{
					RoleArn:            "aws:arn:iam:thing",
					AwsAccessKeyId:     "fake",
					AwsAccessKeySecret: "fake",
					Region:             "us-east-1",
				},
			},
		})
	}
}

func getCredentials(plugs []*plugin.Client) {
	handlePreGetCredentials(plugs)
	handleGetCredentials(plugs)
	handlePostGetCredentials(plugs)
}

func handlePreGetCredentials(plugs []*plugin.Client) {
	for _, p := range plugs {
		client, err := p.Client()
		if err != nil {
			fmt.Println("Error:", err.Error())
			continue
		}
		raw, err := client.Dispense("credentials_grpc")
		if err != nil {
			fmt.Println("Error:", err.Error())
			continue
		}
		profileplugin := raw.(shared.CredentialsService)
		profileplugin.Pre()
	}
}

func handleGetCredentials(plugs []*plugin.Client) {
	for _, p := range plugs {
		client, err := p.Client()
		if err != nil {
			fmt.Println("Error:", err.Error())
			continue
		}
		raw, err := client.Dispense("credentials_grpc")
		if err != nil {
			fmt.Println("Error:", err.Error())
			continue
		}
		profileplugin := raw.(shared.CredentialsService)
		profileplugin.Get()
	}
}

func handlePostGetCredentials(plugs []*plugin.Client) {
	for _, p := range plugs {
		client, err := p.Client()
		if err != nil {
			fmt.Println("Error:", err.Error())
			continue
		}
		raw, err := client.Dispense("credentials_grpc")
		if err != nil {
			fmt.Println("Error:", err.Error())
			continue
		}
		profileplugin := raw.(shared.CredentialsService)
		profileplugin.Post()
	}
}

func getProfileNames(plugs []*plugin.Client) {
	for _, p := range plugs {
		client, err := p.Client()
		if err != nil {
			fmt.Println("Error:", err.Error())
			continue
		}
		raw, err := client.Dispense("profile_names_grpc")
		if err != nil {
			fmt.Println("Error:", err.Error())
			continue
		}
		profileplugin := raw.(shared.ProfileNamesService)
		profileplugin.Get()
	}
}
