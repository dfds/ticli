package configuration

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var TestConfigurationCmd = &cobra.Command{
	Use:   "testconfig",
	Short: "Print configuration for debug purposes",
	Run: func(cmd *cobra.Command, args []string) {
		testConfigurationCmdFunction(cmd, args)
	},
}

func testConfigurationCmdFunction(cmd *cobra.Command, args []string) {
	keys := viper.GetViper().AllKeys()
	fmt.Println(">> Viper Config Values:")
	for _, key := range keys {
		fmt.Printf("%s: %s\n", key, viper.Get(key))
	}

	fmt.Println(">> Arguments:")
	for i, arg := range args {
		fmt.Printf("%d: %s\n", i, arg)
	}

	fmt.Println(">> Flags:")
	cmd.Flags().VisitAll(func(flag *pflag.Flag) {
		fmt.Printf("%s: %s\n", flag.Name, flag.Value)
	})
}
