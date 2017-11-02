package viperfile

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// ReadLocal reads a local config file from the top-level directory
// fileName should be without the extension
// binding is a pointer to the struct that the config will be bound
func ReadLocal(fileName string, binding interface{}) {
	viper.AddConfigPath(".")
	viper.SetConfigName(fileName)
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Cannot find the config file! It should be included in your top-level directory.")
	}
	if viper.Unmarshal(binding) != nil {
		log.Fatalf("Unable to read the config file!")
	}

	log.Printf("Using config %+v\n", binding)

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		if viper.Unmarshal(binding) != nil {
			log.Println("Failed to read the new config file! Using the old config from memory.")
		} else {
			log.Printf("Updated config %+v\n", binding)
		}
	})
}
