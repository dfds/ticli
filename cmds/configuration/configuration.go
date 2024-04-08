package configuration

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.dfds.cloud/ticli/cmds/outputwriter"
)

var (
	accessTokenFile = "token.json"
)

func homeConfigDirectory() string {
	homedir, err := os.UserHomeDir()
	if err != nil {
		outputwriter.GetWriter().WriteError(fmt.Errorf("error while getting home directory %s", err))
		os.Exit(1)
	}

	homeConfigDirectory := fmt.Sprintf("%s/.ticli/", homedir)

	_ = os.Mkdir(homeConfigDirectory, os.ModePerm)

	return homeConfigDirectory
}

func InitializeConfiguration() {
	homeConfigDir := homeConfigDirectory()

	viper.SetConfigName("ticli")
	viper.AddConfigPath(homeConfigDir)

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			outputwriter.GetWriter().WriteError(fmt.Errorf("error while reading home config file %s", err))
			os.Exit(1)
		}
	}

	tokenViper := viper.New()
	tokenViper.SetConfigName("token")
	tokenViper.AddConfigPath(homeConfigDir)
	err = tokenViper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			outputwriter.GetWriter().WriteError(fmt.Errorf("error while reading token config file %s", err))
			os.Exit(1)
		}
	}

	localViper := viper.New()
	localViper.SetConfigName("ticli")
	localViper.AddConfigPath(".")
	err = localViper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			outputwriter.GetWriter().WriteError(fmt.Errorf("error while reading local config file %s", err))
			os.Exit(1)
		}
	}

	err = viper.MergeConfigMap(localViper.AllSettings())
	if err != nil {
		outputwriter.GetWriter().WriteError(fmt.Errorf("error while merging config file %s", err))
		os.Exit(1)
	}
	err = viper.MergeConfigMap(tokenViper.AllSettings())
	if err != nil {
		outputwriter.GetWriter().WriteError(fmt.Errorf("error while merging config file %s", err))
		os.Exit(1)
	}

}

func SetConfigurationValue(key, value string) {
	viper.Set(key, value)
}

func SaveAccessToken(accessToken string) {
	configPath := homeConfigDirectory()

	v := viper.New()
	v.Set("access-token", accessToken)

	v.WriteConfigAs(fmt.Sprintf("%s/%s", configPath, accessTokenFile))
}

func ClearAccessToken() {
	configPath := homeConfigDirectory()
	os.Remove(fmt.Sprintf("%s/%s", configPath, accessTokenFile))
}

func BindFlag(key string, flag *pflag.Flag) {
	viper.BindPFlag(key, flag)
}

func GetString(key string) string {
	return viper.GetString(key)
}

func GetBool(key string) bool {
	return viper.GetBool(key)
}
