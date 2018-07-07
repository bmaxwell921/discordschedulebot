package main

import (
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	tokenPtr := flag.String("t", "", "Bot Token")

	flag.Parse()

	fmt.Println("token was: ", *tokenPtr)

	dg, err := discordgo.New("Bot " + *tokenPtr)

	if (err != nil) {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	dg.AddHandler(newMessage)

	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Main is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

func newMessage(session *discordgo.Session, message *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if message.Author.ID == session.State.User.ID {
		return
	}

	if !strings.HasPrefix(message.Content, "!hello") {
		fmt.Println("I don't care about this message")
		return
	}

	session.ChannelMessageSend(message.ChannelID, "I'm reading you loud and clear from compute engine")
}
