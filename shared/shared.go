package shared

import (
	"strings"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Awsume struct {
	Cmd    *cobra.Command
	Config *viper.Viper
	Logger *zerolog.Logger
}

type Profile struct {
	RoleArn            string
	SourceProfile      string
	AwsAccessKeyId     string
	AwsSecretAccessKey string
	MfaSerial          string
	Region             string
	Output             string
}

type Profiles map[string]Profile

func (p Profiles) Add(prs Profiles) Profiles {
	for key, pr := range prs {
		if _, ok := p[key]; !ok {
			p[key] = pr
		}
	}

	return p
}

func (p Profile) GetAccountId() string {
	//TODO Add SSO Profile stuff
	if p.RoleArn != "" {
		return strings.Split(p.RoleArn, ":")[4]
	}

	if p.SourceProfile != "" {
		return strings.Split(p.SourceProfile, ":")[4]
	}

	if p.MfaSerial != "" {
		return strings.Split(p.MfaSerial, ":")[4]
	}

	return ""
}
