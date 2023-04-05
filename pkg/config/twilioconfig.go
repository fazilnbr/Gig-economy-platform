package config

import (
	"fmt"

	"github.com/twilio/twilio-go"
	verify "github.com/twilio/twilio-go/rest/verify/v2"
)

type TwilioConfig interface {
	SendOTP(cfg Config, to string) error
}

type twilioConfig struct{}

func NewTwilioConfig() TwilioConfig {
	return &twilioConfig{}
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
	resp, err := client.VerifyV2.CreateVerification(serviceSid, params)
	fmt.Printf("\n\nsid : %v\n\n",resp.Sid)
	return err
}
