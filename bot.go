package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"main.go/config"
)

var BotID string
var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	BotID = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is Running!")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == BotID {
		return
	}

	if m.Content == "hi" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Hi, Arrpit this side.")
	}
	if m.Content == "Hi" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Hi, Arrpit this side.")
	}
	if m.Content == "What this server is about" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "This server is in development stage now.")
	}

	s.AddHandler(OnGuildMemberAdd)

}

func OnGuildMemberAdd(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	// Get all the channels in the guild
	channels, err := s.GuildChannels(m.GuildID)
	if err != nil {
		fmt.Println("Error getting channels:", err)
		return
	}

	// Loop through the channels to find the one with the desired name
	var channel *discordgo.Channel
	for _, c := range channels {
		if c.Name == "welcome" {
			channel = c
			break
		}
	}

	// Check if the channel was found
	if channel == nil {
		fmt.Println("Error finding channel:", err)
		return
	}

	// Send the welcome message to the channel
	message := fmt.Sprintf("Welcome to the server, %s!", m.User.Mention())
	_, err = s.ChannelMessageSend(channel.ID, message)
	if err != nil {
		fmt.Println("Error sending message:", err)
	}
}
