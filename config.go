package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/spf13/viper"
)

// MailConfig will configure the IMAP specific settings
type MailConfig struct {
	Secret uuid.UUID

	IMAPServer    string
	IMAPInboxName string
	Username      string
	Password      string
}

type Config struct {
	Mail            MailConfig
	ServerHostname  string
	FeedName        string
	FeedDescription string
}

func GetConfig() Config {
	// Mail defaults
	viper.SetDefault("mail.username", "test@gmail.com")
	viper.SetDefault("mail.password", "abcd123")
	viper.SetDefault("mail.imap_server", "imap.gmail.com:993")
	viper.SetDefault("mail.imap_inbox", "INBOX")

	// Feed defaults
	viper.SetDefault("feed.name", "My Newsletters")
	viper.SetDefault("feed.description", "Local development RSS feed")
	viper.SetDefault("feed.host", "http://localhost")

	viper.SetConfigName("config")         // name of config file (without extension)
	viper.SetConfigType("yaml")           // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/appname/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.appname") // call multiple times to add many search paths
	viper.AddConfigPath(".")              // optionally look for config in the working directory
	err := viper.ReadInConfig()           // Find and read the config file
	if err != nil {                       // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	newUUID, _ := uuid.Parse("9353693a-abfb-456f-9546-8f210ee65f47")

	return Config{
		Mail: MailConfig{
			IMAPServer:    viper.GetString("mail.imap_server"),
			IMAPInboxName: viper.GetString("mail.imap_inbox"),
			Username:      viper.GetString("mail.username"),
			Password:      viper.GetString("mail.password"),

			Secret: newUUID, // For hashing message IDs, this can be regenerated but is not necessary
		},

		ServerHostname:  viper.GetString("feed.host"),
		FeedName:        viper.GetString("feed.name"),
		FeedDescription: viper.GetString("feed.description"),
	}
}
