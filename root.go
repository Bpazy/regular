package regular

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/hashicorp/go.net/proxy"
	"github.com/spf13/cobra"
	"log"
	"net"
	"net/http"
)

var (
	// buildVer represents 'regular' build version
	buildVer string

	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "regular",
		Short: "TODO",
		Long: `TODO
`,
		Run: func(cmd *cobra.Command, args []string) {
			Run()
		},
	}

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "版本号",
		Long:  `查看 regular 的版本号`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(buildVer)
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

func Execute() error {
	return rootCmd.Execute()
}

func Run() {
	config := InitConfig()

	bot, err := NewBotApi(config)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("上午涂药", "上午涂药"),
				tgbotapi.NewInlineKeyboardButtonData("晚上涂药", "晚上涂药"),
				tgbotapi.NewInlineKeyboardButtonData("晚上吃药", "晚上吃药"),
			),
		)

		bot.Send(msg)
	}
}

func NewBotApi(config *configuration) (*tgbotapi.BotAPI, error) {
	if config.Proxy.Addr == "" {
		return tgbotapi.NewBotAPI(config.Telegram.Token)
	}

	dialer, err := proxy.SOCKS5("tcp", config.Proxy.Addr, nil, proxy.Direct)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	client.Transport = &http.Transport{DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
		return dialer.Dial(network, addr)
	}}

	return tgbotapi.NewBotAPIWithClient(config.Telegram.Token, client)
}
