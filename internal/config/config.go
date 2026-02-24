package config

const configFileName = "config.json"

type Config struct {
	DBUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c *Config) SetUser(userName string) error {
	return nil
}

func Read() Config {
	var config Config
	return config
}

func getConfigFilePath() (string, error) {
	return "", nil
}

func write(cfg Config) error {
	return nil
}
