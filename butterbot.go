package main

import (
	"butterbot/cmd"
	"butterbot/service/discord"
	"butterbot/service/whalealert"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var (
	ctx context.Context
)

func main() {

	cmd.RootCmd.Execute()
	ctx = context.Background()
	//fmt.Println("discord token: " + DiscordToken)
	whalealert.Serve(ctx, cmd.WhaleApikey, discord.WhaleAlertSendDiscordMessage)
	discord.Serve(ctx, cmd.DiscordToken)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	discord.OnDisconnect()
	fmt.Println("Discord endpoint is now shutting down.")
	discord.Bot.Close()
}
