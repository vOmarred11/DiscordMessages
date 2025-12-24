package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func main() {
	botToken := "BOT_TOKEN"
	channelID := "CHANNEL_ID"
	dg, err := discordgo.New("Bot " + botToken)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = dg.Open()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(dg *discordgo.Session) {
		err := dg.Close()
		if err != nil {
			panic(err)
		}
	}(dg)
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	_, err = dg.ChannelMessageSend(channelID, text)
	if err != nil {
		fmt.Println(err)
	}
}
