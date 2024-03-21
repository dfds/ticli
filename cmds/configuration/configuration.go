package configuration

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func homeConfigDirectory() string {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Error while getting home directory %s", err)
	}
	homeConfigDirectory := fmt.Sprintf("%s/.mcdonald/", homedir)

	_ = os.Mkdir(homeConfigDirectory, os.ModePerm)

	return homeConfigDirectory
}

func InitializeConfiguration() {
	homeConfigDir := homeConfigDirectory()

	viper.SetConfigName("mcconfig")
	viper.AddConfigPath(homeConfigDir)

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Fatalf("Error while reading home config file %s", err)
		}
	}

	tokenViper := viper.New()
	tokenViper.SetConfigName("token")
	tokenViper.AddConfigPath(homeConfigDir)
	err = tokenViper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Fatalf("Error while reading token config file %s", err)
		}
	}

	localViper := viper.New()
	localViper.SetConfigName("mcconfig")
	localViper.AddConfigPath(".")
	err = localViper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Fatalf("Error while reading local config file %s", err)
		}
	}

	err = viper.MergeConfigMap(localViper.AllSettings())
	if err != nil {
		log.Fatalf("Error while merging config file %s", err)
	}
	err = viper.MergeConfigMap(tokenViper.AllSettings())
	if err != nil {
		log.Fatalf("Error while merging config file %s", err)
	}

}

func SetConfigurationValue(key, value string) {
	viper.Set(key, value)
}

func SaveAccessToken(accessToken string) {
	configPath := homeConfigDirectory()

	v := viper.New()
	v.Set("access-token", accessToken)

	v.WriteConfigAs(fmt.Sprintf("%s/token.json", configPath))
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
