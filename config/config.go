package configs

import "github.com/spf13/viper"

type conf struct {
	WeatherApiKey string `mapstructure:"WEATHER_API_KEY"`
}

var Config conf = conf{}

func LoadConfig(path string) error {
	var cfg *conf
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return err
	}
	Config.WeatherApiKey = cfg.WeatherApiKey
	return nil
}
