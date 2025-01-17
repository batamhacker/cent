/*
Copyright © 2021 xm1k3

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/xm1k3/cent/pkg/jobs"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "cent",
	Short: "Community edition nuclei templates",
	Long: `Community edition nuclei templates, a simple tool that allows you 
to organize all the Nuclei templates offered by the community in one place.

By xm1k3`,
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("path")
		name, _ := cmd.Flags().GetString("name")
		keepfolders, _ := cmd.Flags().GetBool("keepfolders")
		console, _ := cmd.Flags().GetBool("console")

		jobs.Start(path, name, keepfolders, console)
		jobs.RemoveEmptyFolders(path + "/" + name)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cent.yaml)")

	rootCmd.PersistentFlags().StringP("path", "p", "", "Path to save the templates")
	rootCmd.Flags().StringP("name", "n", "", "Name of the main folder")
	rootCmd.Flags().BoolP("keepfolders", "k", false, "Keep folders (by default it only saves yaml files)")
	rootCmd.Flags().BoolP("console", "C", false, "Print console output")

	rootCmd.MarkFlagRequired("path")
	rootCmd.MarkFlagRequired("name")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".cent")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		// fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
