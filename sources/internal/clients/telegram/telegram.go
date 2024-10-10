package telegramclient

import (
	"fmt"
	"log"
	"shtem-web/sources/internal/configs"

	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Client struct {
	*tgbotapi.BotAPI

	cfg  *configs.Configs
	mode string
}

func NewTelegamClient(cfg *configs.Configs) (*Client, error) {
	bot, err := tgbotapi.NewBotAPI(cfg.Telegram.Token)

	mode := "\n[Env: Testing]"
	if gin.Mode() == gin.ReleaseMode {
		mode = "\n[Env: Production]"
	}

	return &Client{bot, cfg, mode}, err
}

func (c *Client) NotifyOnEmail(title, text string) {
	msgText := fmt.Sprintf("%s\n%s", title, text)

	msg := tgbotapi.NewMessage(c.cfg.Telegram.SystemChannel, msgText)
	msg.ParseMode = "markdown"

	_, err := c.Send(msg)
	if err != nil {
		log.Println(err)
	}
}

func (c *Client) Notify(text string) {
	msg := tgbotapi.NewMessage(c.cfg.Telegram.SystemChannel, text)
	msg.ParseMode = "markdown"

	_, err := c.Send(msg)
	if err != nil {
		log.Println(err)
	}
}
