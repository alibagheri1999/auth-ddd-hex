package utils

import "DDD-HEX/config"

func ConfigSetup() *config.Config {
	config.InitAppConfig()
	if config.Get().IsProduction() {
		config.Init("prod")
	} else if config.Get().IsDeveloping() {
		config.Init("prod")
	} else {
		config.Init("")
	}
	return config.Get()
}
