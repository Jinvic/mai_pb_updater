package config

import "github.com/spf13/viper"

func InitConfig(configPath string) (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	viper.AddConfigPath(configPath)
	viper.AutomaticEnv()

	viper.SetDefault("divingfish.base_url", "https://www.diving-fish.com/api/maimaidxprober")
	viper.SetDefault("divingfish.etag", "")
	viper.SetDefault("pocketbase.max_batch_size", 50)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	return &config, nil
}

func WriteConfig(key string, value interface{}) error {
	viper.Set(key, value)
	return viper.WriteConfig()
}
