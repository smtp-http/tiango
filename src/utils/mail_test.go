package utils 

import (
	"testing"
)


func Test_mail(t *testing.T) {

    to := []string{"Keqiang.Zu@luxshare-ict.com"}
    sender := GetMailSender()
    sender.SendMail(to,"tudou tudou","woshidigua")
}