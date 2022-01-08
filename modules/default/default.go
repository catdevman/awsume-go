package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

type ModuleStruct struct {
	Cmd    *cobra.Command
	Logger *log.Logger
	Name   string "default"
}

var outputProfileFlag string
var cleanFlag bool
var refreshFlag bool
var showCommandsFlag bool
var unsetFlag bool
var autoRefreshFlag bool
var killRefresherFlag bool
var listProfilesFlag bool
var refreshAutoCompleteFlag bool
var roleARNFlag string
var principalARNFlag string
var sourceProfileFlag string
var externalIdFlag string
var mfaTokenFlag string
var regionFlag string
var sessionNameFlag string
var roleDurationFlag int64
var withSamlFlag bool
var withWebIdentityFlag bool
var jsonFlag string
var credentialsFileFlag string
var configFileFlag string
var listPluginsFlag bool
var infoFlag bool
var debugFlag bool

func New(c *cobra.Command, logger *log.Logger) (interface{}, error) {
	return ModuleStruct{Cmd: c, Logger: logger}, nil
}

func (s ModuleStruct) PluginName() string {
	return s.Name
}

func (s ModuleStruct) AddArguments() {
	s.Cmd.PersistentFlags().StringVarP(&outputProfileFlag, "output-profile", "o", "", "A profile to output credentials to")
	s.Cmd.PersistentFlags().BoolVarP(&cleanFlag, "clean", "", false, "Clean expired output profiles")
	s.Cmd.PersistentFlags().BoolVarP(&refreshFlag, "refresh", "r", false, "Force refresh credentials")
	s.Cmd.PersistentFlags().BoolVarP(&showCommandsFlag, "show-commands", "s", false, "Show the commands to set the credentials")
	s.Cmd.PersistentFlags().BoolVarP(&unsetFlag, "unset", "u", false, "Unset your aws environment variables")
	s.Cmd.PersistentFlags().BoolVarP(&autoRefreshFlag, "auto-refresh", "a", false, "Auto refresh credentials")
	s.Cmd.PersistentFlags().BoolVarP(&killRefresherFlag, "kill-refresher", "k", false, "Kill autoawsume")
	s.Cmd.PersistentFlags().BoolVarP(&listProfilesFlag, "list-profiles", "l", false, "List profiles, for now always \"more\" this is just boolean flag")
	s.Cmd.PersistentFlags().BoolVarP(&refreshAutoCompleteFlag, "refresh-autocomplete", "", false, "Refresh all plugin autocommplete profiles")
	s.Cmd.PersistentFlags().StringVarP(&roleARNFlag, "role-arn", "", "", "Role ARN or <partiiton>:<account_id>:<role_name>")
	s.Cmd.PersistentFlags().StringVarP(&principalARNFlag, "prinipal-arn", "", "", "Principal ARN or <partition>:<account_id>:<provider_name>")
	s.Cmd.PersistentFlags().StringVarP(&sourceProfileFlag, "source-profile", "", "", "source_profile to use (role-arn only)")
	s.Cmd.PersistentFlags().StringVarP(&externalIdFlag, "external-id", "", "", "External ID to pass to the assume_role")
	s.Cmd.PersistentFlags().StringVarP(&mfaTokenFlag, "mfa-token", "", "", "You mfa token")
	s.Cmd.PersistentFlags().StringVarP(&regionFlag, "region", "", "", "The region you want to awsume into")
	s.Cmd.PersistentFlags().StringVarP(&sessionNameFlag, "session-name", "", "", "Set a custom role session name")
	s.Cmd.PersistentFlags().Int64VarP(&roleDurationFlag, "role-duration", "", 0, "Seconds to get role creds for, 0 means the assume_role call will be made without role duration")
	s.Cmd.PersistentFlags().BoolVarP(&withSamlFlag, "with-saml", "", false, "Use saml (requires plugin)")
	s.Cmd.PersistentFlags().BoolVarP(&withWebIdentityFlag, "with-web-identity", "", false, "Use web identity (requires plugin")
	s.Cmd.PersistentFlags().StringVarP(&jsonFlag, "json", "", "", "Use json credentials")
	s.Cmd.PersistentFlags().StringVarP(&credentialsFileFlag, "credentials-file", "", "", "Target a shared credentials file")
	s.Cmd.PersistentFlags().StringVarP(&configFileFlag, "config-file", "", "", "Target a config file")
	s.Cmd.PersistentFlags().BoolVarP(&listPluginsFlag, "list-plugins", "", false, "List installed plugins")
	s.Cmd.PersistentFlags().BoolVarP(&infoFlag, "info", "", false, "Print any info logs to stderr")
	s.Cmd.PersistentFlags().BoolVarP(&debugFlag, "debug", "", false, "Print any debug logs to stderr")
}

func (s ModuleStruct) PostAddArguments() {
	if withSamlFlag && withWebIdentityFlag {
		fmt.Println("Can not have both --with-saml and --with-web-identity")
		os.Exit(1)
	}
}
