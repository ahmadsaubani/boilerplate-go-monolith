package Config

type EnvironmentConfig struct {
	App         AppConfig       `yaml:"app"`
	Logging     LoggingConfig   `yaml:"logging"`
	Databases   DatabasesConfig `yaml:"databases"`
	Email       EmailConfig     `yaml:"email"`
	News        NewsConfig      `yaml:"news_api"`
	MeshNewsApi MeshNewsConfig  `yaml:"mesh_news_api"`
}
