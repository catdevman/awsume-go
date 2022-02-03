package shared

import (
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
