package Alerts

type LoggingConfig struct {
	ServiceName    string       `yaml:"service_name"`
	LogPath        string       `yaml:"log_path"`
	FilePrefix     string       `yaml:"file_prefix"`
	EnableStdout   bool         `yaml:"enable_stdout"`
	EnableFile     bool         `yaml:"enable_file"`
	EnableLoki     bool         `yaml:"enable_loki"`
	EnableRotation bool         `yaml:"enable_rotation"`
	Alerts         *AlertConfig `yaml:"alerts,omitempty"`
}

// AlertConfig for notification settings
type AlertConfig struct {
	Enabled      bool              `yaml:"enabled"`
	MinLevel     string            `yaml:"min_level"`
	RateLimitSec int               `yaml:"rate_limit_sec"`
	Discord      *DiscordConfig    `yaml:"discord,omitempty"`
	Slack        *SlackConfig      `yaml:"slack,omitempty"`
	Telegram     *TelegramConfig   `yaml:"telegram,omitempty"`
	Email        *AlertEmailConfig `yaml:"email,omitempty"`
}

type DiscordConfig struct {
	Enabled    bool   `yaml:"enabled"`
	WebhookURL string `yaml:"webhook_url"`
	Username   string `yaml:"username,omitempty"`
	AvatarURL  string `yaml:"avatar_url,omitempty"`
}

type SlackConfig struct {
	Enabled    bool   `yaml:"enabled"`
	WebhookURL string `yaml:"webhook_url"`
	Channel    string `yaml:"channel,omitempty"`
	Username   string `yaml:"username,omitempty"`
	IconEmoji  string `yaml:"icon_emoji,omitempty"`
}

type TelegramConfig struct {
	Enabled  bool   `yaml:"enabled"`
	BotToken string `yaml:"bot_token"`
	ChatID   string `yaml:"chat_id"`
}

type AlertEmailConfig struct {
	Enabled    bool     `yaml:"enabled"`
	SMTPHost   string   `yaml:"smtp_host"`
	SMTPPort   int      `yaml:"smtp_port"`
	Username   string   `yaml:"username"`
	Password   string   `yaml:"password"`
	From       string   `yaml:"from"`
	To         []string `yaml:"to"`
	UseTLS     bool     `yaml:"use_tls"`
	SkipVerify bool     `yaml:"skip_verify"`
}
