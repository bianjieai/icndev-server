package cmd

import (
	"github.com/bianjieai/icndev-server/configs"
	"github.com/bianjieai/icndev-server/internal/app"
	"github.com/bianjieai/icndev-server/internal/app/constant"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

var (
	localConfig string
	startCmd    = &cobra.Command{
		Use:   "start",
		Short: "Start Icndev Server.",
		Run: func(cmd *cobra.Command, args []string) {
			online()
		},
	}
	testCmd = &cobra.Command{ // test
		Use:   "test",
		Short: "Start Test Icndev Server.",
		Run: func(cmd *cobra.Command, args []string) {
			test()
		},
	}
)

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.AddCommand(testCmd)
	testCmd.Flags().StringVarP(&localConfig, "Config", "c", "", "conf path: /opt/local.toml")
}

func test() {
	data, err := ioutil.ReadFile(localConfig)
	if err != nil {
		panic(err)
	}
	config, err := configs.ReadConfig(data)
	if err != nil {
		panic(err)
	}
	run(config)
}

func online() {
	var config *configs.Config

	filepath, found := os.LookupEnv(constant.EnvNameConfigFilePath)
	if found {
	} else {
		panic("not found CONFIG_FILE_PATH")
	}
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	config, err = configs.ReadConfig(data)
	if err != nil {
		panic(err)
	}

	run(config)
}

func run(cfg *configs.Config) {
	app.Serve(cfg)
}
