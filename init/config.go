package init

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func InitConfig() {
	root, _ := os.Getwd()

	viper.SetDefault("root", root)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(root)
	viper.AddConfigPath("..") // When testing, 'root' is wrong

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// TODO: set default config
			panic(fmt.Errorf("not found config file: %w", err))
		} else {
			panic(err)
		}
	}
}
