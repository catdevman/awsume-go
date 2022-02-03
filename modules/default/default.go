package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/bigkevmcd/go-configparser"
	"github.com/catdevman/awsume-go/shared"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type ModuleStruct struct {
	Cmd    *cobra.Command
	Config *viper.Viper
	Logger *zerolog.Logger
	Name   string
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

func New(a *shared.Awsume) (interface{}, error) {
	return ModuleStruct{Cmd: a.Cmd, Config: a.Config, Logger: a.Logger}, nil
}

func (s ModuleStruct) PluginName() string {
	return "default"
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
	s.Cmd.PersistentFlags().BoolVarP(&withWebIdentityFlag, "with-web-identity", "", false, "Use web identity (requires plugin)")
	s.Cmd.PersistentFlags().StringVarP(&jsonFlag, "json", "", "", "Use json credentials")
	s.Cmd.PersistentFlags().StringVarP(&credentialsFileFlag, "credentials-file", "", "", "Target a shared credentials file")
	s.Cmd.PersistentFlags().StringVarP(&configFileFlag, "config-file", "", "", "Target a config file")
	s.Cmd.PersistentFlags().BoolVarP(&listPluginsFlag, "list-plugins", "", false, "List installed plugins")
	s.Cmd.PersistentFlags().BoolVarP(&infoFlag, "info", "", false, "Print any info logs to stderr")
	s.Cmd.PersistentFlags().BoolVarP(&debugFlag, "debug", "", false, "Print any debug logs to stderr")
}

func (s ModuleStruct) PostAddArguments() {
	if debugFlag {
		s.Logger.Debug().Msg("Debug is set to true")
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		s.Logger.Level(zerolog.DebugLevel)
	}

	if withSamlFlag && withWebIdentityFlag {
		s.Logger.Debug().Msg("Can not have both --with-saml and --with-web-identity")
		os.Exit(1)
	}
	if autoRefreshFlag {
		if roleARNFlag != "" {
			s.Logger.Debug().Msg("Cannot use autoawsume with a given role_arn")
			os.Exit(1)
		}
		if jsonFlag != "" {
			s.Logger.Debug().Msg("Cannot use autoawsume with json")
			os.Exit(1)
		}
	}
	if cleanFlag {
		s.runClean()
	}
}

func (s ModuleStruct) runClean() {
	_, credentialsFile := s.getAWSFiles()
	s.removeExpiredOutputProfiles(credentialsFile)
	os.Exit(0)
}

func (s ModuleStruct) getAWSFiles() (configFile, credentialsFile string) {

	if os.Getenv("AWS_CONFIG_FILE") != "" {
		s.Logger.Debug().Msg("Setting config file from env var")
		configFile = os.Getenv("AWS_CONFIG_FILE")
	} else if configFileFlag != "" {
		s.Logger.Debug().Msg("Setting config file from config-file CLI parameter")
		configFile = configFileFlag
	} else if file, ok := s.Config.Get("config-file").(string); ok {
		s.Logger.Debug().Msg("Setting config file from config-file in config file")
		configFile = file
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			s.Logger.Debug().Msg("Had issue with getting home directory")
		}
		s.Logger.Debug().Msg("Setting config file from default location for aws")
		configFile = strings.Join([]string{home, ".aws", "config"}, string(os.PathSeparator))
	}

	if os.Getenv("AWS_SHARED_CREDENTIALS_FILE") != "" {
		s.Logger.Debug().Msg("Setting config file from env var")
		credentialsFile = os.Getenv("AWS_SHARED_CREDENTIALS_FILE")
	} else if credentialsFileFlag != "" {
		s.Logger.Debug().Msg("Setting credentials file from credentials-file CLI parameter")
		credentialsFile = credentialsFileFlag
	} else if file, ok := s.Config.Get("credentials-file").(string); ok {
		s.Logger.Debug().Msg("Setting credentials file from credentials-file in config file")
		credentialsFile = file
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			s.Logger.Error().Msg("Had issue with getting home directory")
		}
		s.Logger.Debug().Msg("Setting credentials file from default local for aws")
		credentialsFile = strings.Join([]string{home, ".aws", "credentials"}, string(os.PathSeparator))
	}

	s.Logger.Debug().Msg(configFile)
	s.Logger.Debug().Msg(credentialsFile)

	return configFile, credentialsFile
}

func (s ModuleStruct) removeExpiredOutputProfiles(fileName string) {
	s.Logger.Debug().Msg("Removing Expired Output Profiles from " + fileName)
	var profileConfig = viper.New()
	profileConfig.SetConfigFile(fileName)
	profileConfig.SetConfigType("ini")
	err := profileConfig.ReadInConfig() // Find and read the config file
	if err != nil {                     // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	// fmt.Println(profileConfig)
	// TODO
	// Skip if it has `manager` that is not awsume
	// Skip if it has `autoawsume`
	// If you find something with expiration that is older than now then find all keys with that "prefix" and remove them
}

func (s ModuleStruct) PreCollectProfiles() {
	s.Logger.Debug().Msg("In Pre Collect Profiles")
}

func (s ModuleStruct) CollectProfiles() shared.Profiles {
	configFile, credentialsFile := s.getAWSFiles()
	configs, err := configparser.NewConfigParserFromFile(configFile)
	profiles := shared.Profiles{}
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	creds, err := configparser.NewConfigParserFromFile(credentialsFile)
	if err != nil {
		panic(fmt.Errorf("Fatal error credentials fiie: %w \n", err))
	}
	for _, cred := range creds.Sections() {
		if _, ok := profiles[cred]; ok {
			// TODO Profile already exists now what
		} else {
			c, _ := creds.Items(cred)
			p := shared.Profile{}
			if val, ok := c["mfa_serial"]; ok {
				p.MfaSerial = val
			}
			if val, ok := c["region"]; ok {
				p.Region = val
			} else {
				// TODO is this bad??
				p.Region = "us-east-1"
			}

			if val, ok := c["aws_access_key_id"]; ok {
				p.AwsAccessKeyId = val
			}

			if val, ok := c["aws_secret_access_key"]; ok {
				p.AwsSecretAccessKey = val
			}
			profiles[cred] = p
		}
	}

	for _, config := range configs.Sections() {
		shortName := strings.Replace(config, "profile ", "", -1)

		if _, ok := profiles[shortName]; ok {
			c, _ := configs.Items(config)
			if val, ok := c["region"]; ok {
				p := profiles[shortName]
				p.Region = val
				profiles[shortName] = p
			}
			if val, ok := c["output"]; ok {
				p := profiles[shortName]
				p.Output = val
				profiles[shortName] = p
			}
		} else {
			//TODO no profile exists with this name yet
		}
	}
	s.Logger.Debug().Msg("In Collect Profiles")
	return profiles
}

func (s ModuleStruct) PostCollectProfiles() {
	s.Logger.Debug().Msg("In Post Collect Profiles")
}
