package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	DiscordToken string
	WhaleApikey  string
)

// RootCmd is the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "butterbot",
	Short: "A flavored discord bot",
	Long: `Butterbot is a discord bot that does things.
	Currently it can listen to the whale alert api and post transactions to discord.`,
	Run: run,
}

func init() {
	RootCmd.Flags().StringVarP(&DiscordToken, "discord_token", "d", "", "Discord token")
	RootCmd.Flags().StringVarP(&WhaleApikey, "whale_apikey", "w", "", "Whale alert apikey")
	viper.BindPFlag("discord_token", RootCmd.Flags().Lookup("token"))
	viper.BindPFlag("whale_api_key", RootCmd.Flags().Lookup("whale_apikey"))
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.ReadInConfig()

	if DiscordToken == "" {
		fmt.Println("Loading discord token from config")
		DiscordToken = viper.Get("discord_token").(string)
	} else {
		fmt.Errorf("discord token not found")
	}

	if WhaleApikey == "" {
		fmt.Println("Loading whale alert apikey from config")
		WhaleApikey = viper.Get("whale_api_key").(string)
	} else {
		fmt.Errorf("whalealert api key not found")
	}
}

func run(cmd *cobra.Command, args []string) {

}
