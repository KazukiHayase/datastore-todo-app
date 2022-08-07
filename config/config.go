package config

import "github.com/kelseyhightower/envconfig"

type (
	Config struct {
		Env      string `required:"true" envconfig:"ENV" default:"local"`
		Timezone string `required:"true" envconfig:"TZ" default:"Asia/Tokyo"`

		GCP GCP
	}

	GCP struct {
		ProjectID string `required:"true" envconfig:"GCP_PROJECT_ID"`
		Region    string `required:"true" envconfig:"GCP_REGION"`
	}
)

func Environ() (Config, error) {
	cfg := Config{}
	err := envconfig.Process("", &cfg)

	return cfg, err
}

func (c *Config) IsLocal() bool {
	return c.Env == "local"
}

func (c *Config) IsNotLocal() bool {
	return c.Env != "local"
}
