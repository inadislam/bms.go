package utils

import gomail "gopkg.in/gomail.v2"

func ActiveUser(code, email, username string) {
	link := "http://localhost:8080/active-user"
	messBody := "Hello " + username + ", \n Your Activation code is " + code + " \n\n Active your Account By clicking on <a href='" + link + "'>this</a> link"

	mail := gomail.NewMessage()
	mail.SetHeader("From", "support@inadislam.com")
	mail.SetHeader("To", email)
	mail.SetHeader("Subject", "Activate your account")
	mail.SetBody("text/html", messBody)
	dialer := gomail.NewDialer("0.0.0.0", 1025, "", "")
	if err := dialer.DialAndSend(mail); err != nil {
		panic(err)
	}
}
