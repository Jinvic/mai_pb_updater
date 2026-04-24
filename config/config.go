package config

type Config struct {
	DivingFish DivingFishConfig `mapstructure:"divingfish"`
	Pocketbase PocketbaseConfig `mapstructure:"pocketbase"`
}

type DivingFishConfig struct {
	BaseURL string `mapstructure:"base_url"`
	Etag    string `mapstructure:"etag"`
}

type PocketbaseConfig struct {
	BaseURL        string `mapstructure:"base_url"`
	CollectionName string `mapstructure:"collection_name"`
	Identity       string `mapstructure:"identity"`
	Password       string `mapstructure:"password"`
}