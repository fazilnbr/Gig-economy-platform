package config

import (
	"github.com/twilio/twilio-go"
	verify "github.com/twilio/twilio-go/rest/verify/v2"
)

type TwilioConfig interface {
	SendOTP(cfg Config, to string) error
	VerifyOTP(cfg Config, phone string, otp string) error
}

type twilioConfig struct{}

// VerifyOTP implements TwilioConfig
func (*twilioConfig) VerifyOTP(cfg Config, phone string, otp string) error {
	return nil
}

func (c *twilioConfig) SendOTP(cfg Config, to string) error {
	accountSid := cfg.TWAccountSID
	serviceSid := cfg.TWVerifyServiseSID
	authToken := cfg.TWAuthTocken
	// fromPhone := cfg.TWFromPhone
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})
	phone := to

	params := &verify.CreateVerificationParams{}
	params.SetTo(phone)
	params.SetChannel("sms")
	_, err := client.VerifyV2.CreateVerification(serviceSid, params)

	return err
}

func NewTwilioConfig() TwilioConfig {
	return &twilioConfig{}
}
