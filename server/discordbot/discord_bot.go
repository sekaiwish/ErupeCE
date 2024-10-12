package discordbot

import (
	"erupe-ce/config"
	"erupe-ce/utils/logger"
	"regexp"

	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

var Commands = []*discordgo.ApplicationCommand{
	{
		Name:        "link",
		Description: "Link your Erupe account to Discord",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "token",
				Description: "The token provided by the Discord command in-game",
				Required:    true,
			},
		},
	},
	{
		Name:        "password",
		Description: "Change your Erupe account password",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "password",
				Description: "Your new password",
				Required:    true,
			},
		},
	},
}

type DiscordBot struct {
	Session      *discordgo.Session
	config       *config.Config
	logger       logger.Logger
	MainGuild    *discordgo.Guild
	RelayChannel *discordgo.Channel
}

func NewDiscordBot() (discordBot *DiscordBot, err error) {
	session, err := discordgo.New("Bot " + config.GetConfig().Discord.BotToken)
	discordLogger := logger.Get().Named("discord")
	if err != nil {
		discordLogger.Fatal("Discord failed", zap.Error(err))
		return nil, err
	}

	var relayChannel *discordgo.Channel

	if config.GetConfig().Discord.RelayChannel.Enabled {
		relayChannel, err = session.Channel(config.GetConfig().Discord.RelayChannel.RelayChannelID)
	}

	if err != nil {
		discordLogger.Fatal("Discord failed to create relayChannel", zap.Error(err))
		return nil, err
	}

	discordBot = &DiscordBot{
		config:       config.GetConfig(),
		logger:       discordLogger,
		Session:      session,
		RelayChannel: relayChannel,
	}

	return
}

func (bot *DiscordBot) Start() (err error) {
	err = bot.Session.Open()

	return
}

// NormalizeDiscordMessage replaces all mentions to real name from the message.
func (bot *DiscordBot) NormalizeDiscordMessage(message string) string {
	userRegex := regexp.MustCompile(`<@!?(\d{17,19})>`)
	emojiRegex := regexp.MustCompile(`(?:<a?)?:(\w+):(?:\d{18}>)?`)

	result := ReplaceTextAll(message, userRegex, func(userId string) string {
		user, err := bot.Session.User(userId)

		if err != nil {
			return "@unknown" // @Unknown
		}

		return "@" + user.Username
	})

	result = ReplaceTextAll(result, emojiRegex, func(emojiName string) string {
		return ":" + emojiName + ":"
	})

	return result
}

func (bot *DiscordBot) RealtimeChannelSend(message string) (err error) {
	if bot.RelayChannel == nil {
		return
	}

	_, err = bot.Session.ChannelMessageSend(bot.RelayChannel.ID, message)

	return
}
func ReplaceTextAll(text string, regex *regexp.Regexp, handler func(input string) string) string {
	result := regex.ReplaceAllFunc([]byte(text), func(s []byte) []byte {
		input := regex.ReplaceAllString(string(s), `$1`)

		return []byte(handler(input))
	})

	return string(result)
}
