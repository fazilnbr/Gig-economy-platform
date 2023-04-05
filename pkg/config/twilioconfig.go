package config

type TwilioConfig interface {
	SendOTP(cfg Config, to string, message []byte) error
}

type twilioConfig struct{}

func NewTwilioConfig() MailConfig {
	return &mailConfig{}
}

func (c *mailConfig) SendOTP(cfg Config, to string, message []byte) error {
	return nil
}
