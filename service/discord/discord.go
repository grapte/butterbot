package discord

import (
	"butterbot/types"
	"context"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"time"
)

var (
	ctx       context.Context
	Bot       *discordgo.Session
	token     string
	channelID string = ""
)

//
// todo: expand text events from MessageCreate back into the event loop as more user cmds are created.
// todo: manage active bot channels
//

func Serve(c context.Context, t string) {
	ctx, token = c, t
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Service session,", err)
	}

	Bot = dg
	dg.AddHandler(onConnect)
	//dg.AddHandler(onDisconnect)
	//dg.AddHandler(pingpongCreate)
	//dg.AddHandler(ProcessMessages)

	dg.Identify.Intents = discordgo.IntentGuildMessages

	// Open a websocket connection to Service and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection to discord,", err)
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Discord endpoint is now running.")
	select {
	case <-ctx.Done():
		return // exit
	default:

	}

}

func onConnect(s *discordgo.Session, con *discordgo.Connect) {
	s.ChannelMessageSend(channelID, "Bot is Connected")
}

func onDisconnect(s *discordgo.Session, dcon *discordgo.Disconnect) {
	OnDisconnect()
}

func OnDisconnect() {
	Bot.ChannelMessageSend(channelID, "Bot is Disconnected")
}

func pingpongCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the Bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}

func WhaleAlertSendDiscordMessage(rec *types.RecordsResponse) {
	var emb []*discordgo.MessageEmbed
	for i, r := range rec.Transactions {
		emb = append(emb, &discordgo.MessageEmbed{
			Type:        "rich",
			Title:       fmt.Sprintf("USD %20.f | %s %20.f", r.AmountUsd, r.Symbol, r.Amount),
			Description: fmt.Sprintf("Hash: %s", r.Hash),
			Color:       0x00FFFF,
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:  fmt.Sprintf("From: %s", r.From.OwnerType),
					Value: r.From.Address,
				},
				{
					Name:  fmt.Sprintf("To: %s %s", r.To.OwnerType, r.To.Owner),
					Value: r.To.Address,
				},
			},
			Author: &discordgo.MessageEmbedAuthor{
				Name:    fmt.Sprintf("Blockchain: %s", r.Blockchain),
				IconURL: "https://cdn-icons-png.flaticon.com/512/5170/5170907.png",
			},
			Footer: &discordgo.MessageEmbedFooter{
				Text: fmt.Sprintf("Date: %s", time.Unix(int64(r.Timestamp), 0).Format("2006-01-02 15:04:05")),
			},
		})
		// send every 10th message due to discord embed limit for a single message
		if i%10 == 0 {
			msg := &discordgo.MessageSend{
				Embeds: emb,
			}
			_, err := Bot.ChannelMessageSendComplex(channelID, msg)
			if err != nil {
				fmt.Println(err)
			}
			emb = []*discordgo.MessageEmbed{}
		}
	}
	if len(emb) > 0 {
		msg := &discordgo.MessageSend{
			Embeds: emb,
		}
		_, err := Bot.ChannelMessageSendComplex(channelID, msg)
		if err != nil {
			fmt.Println(err)
		}
	}
}
