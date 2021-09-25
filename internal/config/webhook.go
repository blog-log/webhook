package config

type Webhook struct {
	ProcessorHost  string `envconfig:"PROCESSOR_HOST" required:"true"`
	GithubValidate bool   `envconfig:"GITHUB_VALIDATE" required:"true"`
	GithubSecret   string `envconfig:"GITHUB_WEBHOOK_SECRET" required:"true"`
}
