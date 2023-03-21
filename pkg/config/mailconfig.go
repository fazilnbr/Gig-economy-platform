package config

import (
	"net/smtp"
)

type MailConfig interface {
	SendMail(cfg Config, to string, message []byte) error
}

type mailConfig struct{}

func NewMailConfig() MailConfig {
	return &mailConfig{}
}

func (c *mailConfig) SendMail(cfg Config, to string, message []byte) error {
	userName := cfg.SMTPUSERNAME
	password := cfg.SMTPPASSWORD
	smtpHost := cfg.SMTPHOST
	smtpPort := cfg.SMTPPORT

	auth := smtp.PlainAuth("", userName, password, smtpHost)
	return smtp.SendMail(smtpHost+":"+smtpPort, auth, userName, []string{to}, message)
}
