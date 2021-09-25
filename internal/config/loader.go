package config

import "github.com/kelseyhightower/envconfig"

func Load() (*Webhook, error) {
	var cfg Webhook
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
