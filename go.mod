module github.com/catdevman/awsume-go

go 1.17

replace github.com/catdevman/awsume-go/pkg/hooks => ./pkg/hooks

require (
	github.com/catdevman/awsume-go/pkg/hooks v0.0.0-00010101000000-000000000000
	github.com/rs/zerolog v1.26.1
	github.com/spf13/cobra v1.3.0
	github.com/spf13/viper v1.10.1
)

require (
	github.com/aws/aws-sdk-go v1.42.34 // indirect
	github.com/bigkevmcd/go-configparser v0.0.0-20210106142102-909504547ead // indirect
	github.com/fsnotify/fsnotify v1.5.1 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/magiconair/properties v1.8.5 // indirect
	github.com/mitchellh/mapstructure v1.4.3 // indirect
	github.com/mvo5/goconfigparser v0.0.0-20201015074339-50f22f44deb5 // indirect
	github.com/pelletier/go-toml v1.9.4 // indirect
	github.com/spf13/afero v1.8.0 // indirect
	github.com/spf13/cast v1.4.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.2.0 // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
	golang.org/x/text v0.3.7 // indirect
	gopkg.in/ini.v1 v1.66.2 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
