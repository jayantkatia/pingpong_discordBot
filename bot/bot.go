package bot

import (
	"fmt"
	"math/rand"
	"strings"

	"../config"
	"github.com/bwmarrin/discordgo"
)

var BotID string
var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)
	fmt.Println("Test")

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
	flag = false
	err = goBot.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Bot is running")
	<-make(chan struct{})
	return
}

var (
	flag         bool
	numOfPlayers string
	player1      string
	player1Score int = 0
	botScore     int = 0

	player2 string
)

func reset() {
	flag = false
	numOfPlayers = ""
	player1 = ""
	player1Score = 0
	botScore = 0

	player2 = ""
	return
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}

	if strings.HasPrefix(m.Content, config.BotPrefix) {
		if m.Content == "!start" {
			_, _ = s.ChannelMessageSend(m.ChannelID, "Lets play table tennis...\n!1 for 1 player(against bot)\n!2 for player vs player(coming soon)")
			reset()
			flag = true
			player1 = m.Author.Username

		}

		fmt.Println(m.Content)
	}
	if flag && player1 == m.Author.Username {
		if m.Content == "!1" {
			s.ChannelMessageSend(m.ChannelID, "User will type ping or pong")
		} else if m.Content == "ping" || m.Content == "pong" {

			var botPlayChoice int = rand.Intn(3)
			var botPlay string
			if botPlayChoice == 0 {
				botPlay = "ping"
			} else {
				botPlay = "pong"
			}
			fmt.Println(botPlay)

			if m.Content == botPlay {
				botScore++
				s.ChannelMessageSend(m.ChannelID, "Bot goes aggressive...\n")
			} else {
				player1Score++
			}

			player1ScoreString := fmt.Sprintf("%d", player1Score)
			botScoreString := fmt.Sprintf("%d", botScore)
			s.ChannelMessageSend(m.ChannelID, "BotScore -vs- "+player1+" Score\n"+botScoreString+" : "+player1ScoreString)

			if botScore == 5 {
				s.ChannelMessageSend(m.ChannelID, "Bot Wins...\nBot: I knew it from the start...I am gonna win..")
				flag = false
			}
			if player1Score == 5 {
				s.ChannelMessageSend(m.ChannelID, player1+" wins...\nBot: up for a rematch ? this time i am gonna win")
				flag = false
			}
		}

	} else if !flag && m.Content == "!1" || m.Content == "!2" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "please start the game first using !start")
	}
}

// 	if player1 == "" {
// 		player1 = m.Author.Username
// 		_, _ = s.ChannelMessageSend(m.ChannelID, m.Author.Username+" is up for it")
// 	} else if player1 == m.Author.Username {
// 		s.ChannelMessageSend(m.ChannelID, "you are already in the game ...over-excited ehh?")
// 	} else if player2 == "" {
// 		player2 = m.Author.Username
// 		_, _ = s.ChannelMessageSend(m.ChannelID, m.Author.Username+" is up for it \n Lets start the game....\n
// 		The players will write either ping or pong
// 		The bot will randomly choose between these two...player having opposite to the bot gets one point

// 		")

// 	} else {
// 		s.ChannelMessageSend(m.ChannelID, "2 users are already up for it")
// 	}
