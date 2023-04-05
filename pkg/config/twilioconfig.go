package config

type TwilioConfig interface {
	SendOTP(to string) error
}

type twilioConfig struct{}

func NewTwilioConfig() TwilioConfig {
	return &twilioConfig{}
}

func (c *twilioConfig) SendOTP(to string) error {
	
	return nil
}
