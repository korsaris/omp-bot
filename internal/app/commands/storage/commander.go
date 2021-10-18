package storage

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/storage/retention"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type StorageCommander struct {
	bot                *tgbotapi.BotAPI
	retentionCommander Commander
}

func NewStorageCommander(
	bot *tgbotapi.BotAPI,
) *StorageCommander {
	return &StorageCommander{
		bot: bot,
		// retentionCommander
		retentionCommander: retention.NewStorageRetentionCommander(bot),
	}
}

func (c *StorageCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Retention {
	case "retention":
		c.retentionCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("StorageCommander.HandleCallback: unknown retention - %s", callbackPath.Retention)
	}
}

func (c *StorageCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Retention {
	case "retention":
		c.retentionCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("StorageCommander.HandleCommand: unknown retention - %s", commandPath.Retention)
	}
}
