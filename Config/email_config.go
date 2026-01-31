package Config

type EmailConfig struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	User      string `yaml:"user"`
	Password  string `yaml:"password"`
	EmailFrom string `yaml:"email_from"`
}
