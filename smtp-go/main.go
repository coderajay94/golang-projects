package main

import (
	"log"
	"net/smtp"
)

func main() {
	from := "radx94ra@gmail.com"
	password := "Columbus@2021" // use App Password if using Gmail

	to := []string{"official.ajaykumar94@gmail.com"}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := []byte("Subject: Hello from Go!\r\n" +
		"\r\n" +
		"This is a test email sent using Go.\r\n")

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		log.Fatalf("Failed to send email: %v", err)
	}

	log.Println("✅ Email sent successfully!")
}
