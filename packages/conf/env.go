package conf

import (
	"strings"
	"sync"
	"github.com/spf13/viper"
)

type(
	Config struct {
		DB Database
		Server  Server
		App Application
		SWAGGER SWAGGER
		JWT JWT
	}

	Server struct {
		Port int
	}

	Application struct {
		NAME string
	}

	Database struct {
		URL 		string
		DBNAME 		string
		COLLECTION 	string
	}

	JWT struct {
		Secret string
		Expired string
		RefreshSecret string
		RefreshExpired string
	}

	SWAGGER struct {
		HOST	string
	}
)

var (
	once 				sync.Once
	configInstance 		*Config
)

func GetConfig() *Config {
	once.Do(
		func() {
			viper.SetConfigName("config");
			viper.SetConfigType("yaml");
			viper.AddConfigPath("./");
			viper.AutomaticEnv();
			viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"));

			if err := viper.ReadInConfig(); err != nil {
				panic(err);
			}

			if err := viper.Unmarshal(&configInstance); err != nil {
				panic(err);
			}
		})

		return configInstance;
}