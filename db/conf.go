package db

import (
	"os"

	"github.com/c8112002/news-api/utils"
	"github.com/spf13/viper"
)

func readDBConf() (*dbconf, error) {
	var c dbconf

	viper.SetConfigName("dbconf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")

	if err := viper.ReadInConfig(); err != nil {
		return &c, err
	}

	var key string
	switch utils.AppEnv {
	case utils.Development:
		key = "development"
	case utils.Production:
		key = "production"
	default:
		key = "production"
	}

	return &dbconf{
		Dialect:    viper.GetString(key + ".dialect"),
		Datasource: os.ExpandEnv(viper.GetString(key + ".datasource")),
		Dir:        viper.GetString(key + ".dir"),
	}, nil
}

type dbconf struct {
	Dialect    string
	Datasource string
	Dir        string
}
