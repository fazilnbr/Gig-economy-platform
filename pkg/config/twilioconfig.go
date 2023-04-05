package config

type TwilioConfig interface {
	SendOTP(cfg Config, to string, message []byte) error
}

type twilioConfig struct{}

func NewTwilioConfig() TwilioConfig {
	return &twilioConfig{}
}

func (c *twilioConfig) SendOTP(cfg Config, to string, message []byte) error {
	return nil
}
