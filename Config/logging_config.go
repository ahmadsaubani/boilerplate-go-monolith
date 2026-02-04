package Config

import (
	"fmt"

	"boilerplate-go/Config/DTO/ConfigStructs/Alerts"

	logging "github.com/ahmadsaubani/go-logging-lib"
	"github.com/ahmadsaubani/go-logging-lib/alerts/discord"
	"github.com/ahmadsaubani/go-logging-lib/alerts/email"
	"github.com/ahmadsaubani/go-logging-lib/alerts/slack"
	"github.com/ahmadsaubani/go-logging-lib/alerts/telegram"
)

var AppLogger *logging.Logger

func InitLoggerFromConfig(loggingConfig Alerts.LoggingConfig) {
	filePrefix := loggingConfig.FilePrefix
	if filePrefix == "" {
		filePrefix = "app"
	}

	config := &logging.Config{
		ServiceName:    loggingConfig.ServiceName,
		LogPath:        loggingConfig.LogPath,
		FilePrefix:     filePrefix,
		EnableStdout:   loggingConfig.EnableStdout,
		EnableFile:     loggingConfig.EnableFile,
		EnableLoki:     loggingConfig.EnableLoki,
		EnableRotation: loggingConfig.EnableRotation,
		Alerts:         convertAlertConfig(loggingConfig.Alerts),
	}

	var err error
	AppLogger, err = logging.New(config)
	if err != nil {
		panic(fmt.Sprintf("Failed to initialize logger: %v", err))
	}

	AppLogger.Info("🚀 Logger initialized from YAML configuration")
}

func convertAlertConfig(cfg *Alerts.AlertConfig) *logging.AlertsConfig {
	if cfg == nil || !cfg.Enabled {
		return nil
	}

	alertCfg := &logging.AlertsConfig{
		Enabled:      cfg.Enabled,
		MinLevel:     cfg.MinLevel,
		RateLimitSec: cfg.RateLimitSec,
	}

	if cfg.Discord != nil && cfg.Discord.Enabled {
		alertCfg.Discord = &discord.Config{
			Enabled:    cfg.Discord.Enabled,
			WebhookURL: cfg.Discord.WebhookURL,
			Username:   cfg.Discord.Username,
			AvatarURL:  cfg.Discord.AvatarURL,
		}
	}

	if cfg.Slack != nil && cfg.Slack.Enabled {
		alertCfg.Slack = &slack.Config{
			Enabled:    cfg.Slack.Enabled,
			WebhookURL: cfg.Slack.WebhookURL,
			Channel:    cfg.Slack.Channel,
			Username:   cfg.Slack.Username,
			IconEmoji:  cfg.Slack.IconEmoji,
		}
	}

	if cfg.Telegram != nil && cfg.Telegram.Enabled {
		alertCfg.Telegram = &telegram.Config{
			Enabled:  cfg.Telegram.Enabled,
			BotToken: cfg.Telegram.BotToken,
			ChatID:   cfg.Telegram.ChatID,
		}
	}

	if cfg.Email != nil && cfg.Email.Enabled {
		alertCfg.Email = &email.Config{
			Enabled:    cfg.Email.Enabled,
			SMTPHost:   cfg.Email.SMTPHost,
			SMTPPort:   cfg.Email.SMTPPort,
			Username:   cfg.Email.Username,
			Password:   cfg.Email.Password,
			From:       cfg.Email.From,
			To:         cfg.Email.To,
			UseTLS:     cfg.Email.UseTLS,
			SkipVerify: cfg.Email.SkipVerify,
		}
	}

	return alertCfg
}
