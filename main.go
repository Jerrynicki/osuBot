package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"osuBot/commands"
	"osuBot/util"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	osuapi "github.com/thehowl/go-osuapi"
)

var (
	discordToken string
	osuToken     string
	osuClient    osuapi.Client
	prefix       = "osu!"
)

func init() {
	token, err := ioutil.ReadFile("discord_api_token")
	util.CheckForErrors(err)

	discordToken = string(token)

}

func main() {
	bot, err := discordgo.New("Bot " + discordToken)

	bot.AddHandler(messageCreate)

	err = bot.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	bot.Close()
}

func messageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID {
		return
	}

	if !strings.HasPrefix(message.Message.Content, prefix) {
		return
	}

	seperated := strings.Split(message.Message.Content, " ")
	command := seperated[0][len(prefix):] // the raw command name without the prefix
	args := seperated[1:]                 // the arguments, seperated by a space, in an array

	// TODO: implement some kind of functionality to map commands to their name and trigger them, to avoid 100 if statements in a row

	if command == "ping" {
		commands.Ping(session, message, args)
	}

	if command == "beatmap" {
		commands.Beatmap(session, message, args)
	}
}
