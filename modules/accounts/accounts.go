package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/catdevman/awsume-go/shared"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var VALID_TREK10_ROLES = [5]string{
	"readonly",
	"working",
	"poweruser",
	"manager",
	"kernel",
}

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

type ListAccountsResponse struct {
	Data struct {
		ListAccounts struct {
			NextToken string    `json:"nextToken"`
			Items     []Account `json:"items"`
		} `json:"listAccounts"`
	} `json:"data"`
}

type Account struct {
	AccountId     string `json:"accountId"`
	Alias         string `json:"alias"`
	AccountStatus string `json:"accountStatus"`
}

type ModuleStruct struct {
	Cmd    *cobra.Command
	Config *viper.Viper
	Logger *zerolog.Logger
	Name   string
}

func New(a *shared.Awsume) (interface{}, error) {
	return ModuleStruct{Cmd: a.Cmd, Config: a.Config, Logger: a.Logger, Name: "accounts"}, nil

}

func (s ModuleStruct) PluginName() string {
	return s.Name
}

func (s ModuleStruct) AddArguments() {
}

func (s ModuleStruct) PostAddArguments() {
}

func (s ModuleStruct) PreCollectProfiles() {
	s.Logger.Info().Msg("In Pre Collect Profiles")
}

func (s ModuleStruct) CollectProfiles() shared.Profiles {
	s.Logger.Debug().Msg("In Collect Profiles")
	profiles := shared.Profiles{}
	accessToken, err := s.getAccessToken()
	if err != nil {
		s.Logger.Error().Msg(fmt.Sprintf("An Error Occured %v", err))
	}
	jsonData := map[string]string{
		"query": `
            {
 listAccounts(limit:1000, filter:{accountStatus:{ne:"INACTIVE"}}) { nextToken items { accountId alias accountStatus } }
            }
        `,
	}
	jsonValue, _ := json.Marshal(jsonData)
	request, err := http.NewRequest("POST", os.Getenv("GRAPHQL_ENDPOINT"), bytes.NewBuffer(jsonValue))
	client := &http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: 3 * time.Second,
			}).DialContext,
			TLSHandshakeTimeout: 10 * time.Second,
		},
	}
	request.Header.Set("Authorization", accessToken)
	response, err := client.Do(request)
	defer response.Body.Close()
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	var listAccountsResponse ListAccountsResponse
	err = json.Unmarshal(data, &listAccountsResponse)
	if err != nil {
	}
	//listProfilesFlag := s.Cmd.PersistentFlags().Lookup("list-profiles")
	flag := true
	profiles, err = getAccountsFromArgonaut(listAccountsResponse, flag)
	if err != nil {
	}

	//fmt.Println(fmt.Sprintf("%+v", listProfilesFlag))
	return profiles

}

func (s ModuleStruct) PostCollectProfiles() {
	s.Logger.Info().Msg("In Post Collect Profiles")
}

func (s ModuleStruct) getAccessToken() (string, error) {
	postBody, _ := json.Marshal(map[string]string{
		"client_id":     os.Getenv("AUTH0_CLIENT_ID"),
		"client_secret": os.Getenv("AUTH0_CLIENT_SECRET"),
		"audience":      "https://harbor.trek10.com",
		"grant_type":    "client_credentials",
	})
	responseBody := bytes.NewBuffer(postBody)
	c := &http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: 3 * time.Second,
			}).DialContext,
			TLSHandshakeTimeout: 10 * time.Second,
		},
	}
	resp, err := c.Post("https://trek10.auth0.com/oauth/token", "application/json", responseBody)
	//Handle Error
	if err != nil {
		s.Logger.Error().Msg(fmt.Sprintf("An Error Occured %v", err))
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		s.Logger.Error().Msg(err.Error())
	}
	sb := string(body)

	var accessTokenResponse AccessTokenResponse
	err = json.Unmarshal([]byte(sb), &accessTokenResponse)
	if err != nil {
		s.Logger.Error().Msg("Could not unmarshal json from access token")
		return "", err
	}
	return accessTokenResponse.AccessToken, nil
}

func getAccountsFromArgonaut(ar ListAccountsResponse, flag bool) (shared.Profiles, error) {
	ps := shared.Profiles{}
	p := shared.Profile{}
	for _, pr := range ar.Data.ListAccounts.Items {
		p = shared.Profile{}
		if flag {
			p.RoleArn = fmt.Sprintf("arn:aws:iam::%s:role/trek10-%s", pr.AccountId, "role")
			p.SourceProfile = "default" //TODO get from config
			p.MfaSerial = "arn:aws:iam::800094578424:mfa/lpearson"
			p.Region = "None"
			ps[pr.Alias] = p
		} else {
			for _, v := range VALID_TREK10_ROLES {
				p = shared.Profile{}
				fullProfile := fmt.Sprintf("%s-%s", pr.Alias, v)
				p.RoleArn = fmt.Sprintf("arn:aws:iam::%s:role/trek10-%s", pr.AccountId, "role")
				p.SourceProfile = "default" //TODO get from config
				p.MfaSerial = "arn:aws:iam::800094578424:mfa/lpearson"
				p.Region = "None"
				ps[fullProfile] = p
			}
		}
	}
	return ps, nil
}
