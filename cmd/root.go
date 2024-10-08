package cmd

import (
	"fmt"
	"os"

	"github.com/rchouinard/commit-assistant/assistant"
	"github.com/rchouinard/commit-assistant/git"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "commit-assistant",
	Short: "A utility to generate commit messages based on changes in a project",
	PreRun: func(cmd *cobra.Command, args []string) {
		if !git.IsGitInstalled() {
			fmt.Println("Git is not installed or not in the PATH")
			os.Exit(1)
		}

		isRepo, _ := git.IsGitRepo()
		if !isRepo {
			fmt.Println("Not a git repository")
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		stagedFiles, err := git.GetStagedFiles()
		cobra.CheckErr(err)

		if len(stagedFiles) == 0 {
			fmt.Println("No staged changes detected")
			os.Exit(1)
		}

		gitDiff, err := git.DiffFiles(stagedFiles)
		if err != nil {
			fmt.Println("Error generating diff: ", err)
			os.Exit(1)
		}

		asst := assistant.NewOllamaAssistant(assistant.Config{
			BaseURL: viper.GetString("base-url"),
			Model:   viper.GetString("model"),
		})
		resp, err := asst.GenerateMessage(cmd.Context(), gitDiff)
		cobra.CheckErr(err)

		fmt.Println(resp)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.Flags().StringP("api-key", "k", "", "")
	rootCmd.Flags().StringP("base-url", "b", "", "")
	rootCmd.Flags().StringP("model", "m", "", "")
	rootCmd.Flags().StringP("provider", "p", "ollama", "")

	viper.BindPFlag("api-key", rootCmd.Flags().Lookup("api-key"))
	viper.BindPFlag("base-url", rootCmd.Flags().Lookup("base-url"))
	viper.BindPFlag("model", rootCmd.Flags().Lookup("model"))
	viper.BindPFlag("provider", rootCmd.Flags().Lookup("provider"))
}

func initConfig() {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	viper.AddConfigPath(home)
	viper.SetConfigName(".commit-assistant")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err == nil {
		//fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
