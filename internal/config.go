package internal

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

// TODO
// Add customization option for config file
// Check that default config or given file exists

// teams_url for notify
// teams_file for teams message template file
func GetConfigValue(configName string) (value string) {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	// Look for config in these two locations
	viper.AddConfigPath("$HOME/.config/azurenum")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("fatal error config file: %v", err)
	}

	return viper.Get(configName).(string)
}

func getJsonTemplate() []byte {
	file := GetConfigValue("teams_file")
	json, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("failed opening file... %v", err)
	}
	return json
}
