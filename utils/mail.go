package utils

import (
    "fmt"
    "net/smtp"
    "strings"
    "sync"
    "github.com/smtp-http/tiango/datastorage"
)


type MailSender struct {
	MailAddr 	string
	Code 		string
	SmtpServer 	string
	Auth 		smtp.Auth
	User 		string
	Nickname 	string
}

var mailSender *MailSender
var once_sender sync.Once
 
func GetMailSender() *MailSender {
    once_sender.Do(func() {
        mailSender = &MailSender{}
        param := datastorage.GetSysParam()
        mailSender.MailAddr = param.Mail.MailAddr
        mailSender.Code = param.Mail.Code
        mailSender.SmtpServer = param.Mail.SmtpServer
        mailSender.Auth = smtp.PlainAuth("", mailSender.MailAddr, mailSender.Code, mailSender.SmtpServer)
        mailSender.Nickname = "data analysis"
    })
    return mailSender
}

func (m *MailSender) SendMail(to []string,subject string,body string) error {
	content_type := "Content-Type: text/plain; charset=UTF-8"
	user := mailSender.MailAddr

	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + mailSender.Nickname +
        "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)

    err := smtp.SendMail(mailSender.SmtpServer + ":25", mailSender.Auth, user, to, msg)
    if err != nil {
        fmt.Printf("send mail error: %v", err)
        return err
    }

    return nil
}


