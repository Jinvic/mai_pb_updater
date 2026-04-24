package config

type Config struct {
	DivingFish DivingFishConfig `mapstructure:"divingfish"`
}

type DivingFishConfig struct {
	BaseURL string `mapstructure:"base_url"`
	Etag    string `mapstructure:"etag"`
}
