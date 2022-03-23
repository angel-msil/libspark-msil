package configs

import (
	"log"
	"github.com/spf13/viper"
	"sync"
)

type providers struct {
	providers map[string]*viper.Viper
	mu        sync.Mutex
}

var baseConfigPath string
var p *providers

// Init is used to initialize the configurations
func Init(path string) {
	baseConfigPath = path
	p = &providers{
		providers: make(map[string]*viper.Viper),
	}
}

// Get is used to get the instance to the config provider for the configuration name
func Get(name string) (*viper.Viper) {
	p.mu.Lock()
	defer p.mu.Unlock()

	// see for an existing provider
	if provider, ok := p.providers[name]; ok {
		// provider already exists
		// use it
		return provider
	}

	// try to get the provider
	provider := viper.New()
	provider.SetConfigName(name)
	provider.AddConfigPath(baseConfigPath)
	err := provider.ReadInConfig()
	if err != nil {
		// config not found
		 log.Println("config %s not found", name)
		 return nil
	}

	// add a watcher for this provider-read an update to a config file while running and not miss a beat
	provider.WatchConfig()

	// successfully found config, store it for future use
	p.providers[name] = provider

	return provider
}
