package denbot

type Config struct {
	BotToken string `toml:"token"`
}

func NewConfig() *Config {
	return &Config{
		BotToken: "1408091499:AAGrZW8gwvJs33BZZnzoifh4t1BdgFrpbjQ",
	}
}
