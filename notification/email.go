package notification

import (
	"fmt"
	"net/smtp"
	"strings"
)

func sendEmail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

func SendEmail(toAddress, subject, content string) error {
	user := "niahckcolb@163.com"
	password := `LSHNOOFEAQKDTILS`
	host := "smtp.163.com:25"

	body := `
		<html>
		<body>
		<h3>
	` + content + `
		</h3>
		</body>
		</html>
		`
	err := sendEmail(user, password, host, toAddress, subject, body, "html")
	if err != nil {
		//rateLimiter
		fmt.Println("Send mail error!, err:", err.Error())
		fmt.Println(err)
	} else {
		fmt.Println("Send mail success!")
	}
	return err
}

func MultiSendEmail(toAddresses []string, subject, content string) {
	for _, toAddress := range toAddresses {
		SendEmail(toAddress, subject, content)
	}
}
